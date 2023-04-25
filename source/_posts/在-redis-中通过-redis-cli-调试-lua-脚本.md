---
title: 在 redis 中通过 redis-cli 调试 lua 脚本
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Redis
  - 缓存
  - Lua
abbrlink: '81563347'
date: 2023-04-25 18:04:27
img:
coverImg:
password:
summary:
---


# 在 redis 中通过 redis-cli 调试 lua 脚本

## 假设我有如下一段 lua 脚本

```php

<?php

$redis = new \Redis();

$redis->connect('192.168.10.194');

$lua = <<<EOF
local k1 = KEYS[1]
local k2 = KEYS[2]

local v1 = redis.call('get', k1)
local v2 = redis.call('get', k2)

if v1 and v2 then
    local v1ttl = redis.call('ttl', k1)
    local v2ttl = redis.call('ttl', k2)
    return {true, k1, v1, k2, v2, v1ttl, v2ttl, v1..v2}
else
    local now = tonumber(ARGV[1])
    redis.call('hsetnx', 'hash_table', k1, now)
    redis.call('hsetnx', 'hash_table', k2, now)
    local hash_all = redis.call('hgetall', 'hash_table')
    return {false, hash_all}
end

EOF;

$res = $redis->eval($lua, ['name', 'age', time()], 2);

var_dump($res);

```

## 第一种调试方式：使用 --eval 选项直接调试

先创建一个 `.lua` 脚本

```shell

root@a72e6809fb18:/tmp# cat > test.lua <<EOF
 local k1 = KEYS[1]
 local k2 = KEYS[2]

 local v1 = redis.call('get', k1)
 local v2 = redis.call('get', k2)

 if v1 and v2 then
     local v1ttl = redis.call('ttl', k1)
     local v2ttl = redis.call('ttl', k2)
     return {true, k1, v1, k2, v2, v1ttl, v2ttl, v1..v2}
 else
     local now = tonumber(ARGV[1])
     redis.call('hsetnx', 'hash_table', k1, now)
     redis.call('hsetnx', 'hash_table', k2, now)
     local hash_all = redis.call('hgetall', 'hash_table')
     return {false, hash_all}
 end

 EOF

```

如果 k1 和 k2 有值时

**注意：参数和值之间有一个英文逗号，并且英文逗号左右均有一个空格！**

```shell

root@a72e6809fb18:/tmp# redis-cli --eval test.lua name age , 12345
1) (integer) 1
2) "name"
3) "alex"
4) "age"
5) "18"
6) (integer) -1
7) (integer) 70535
8) "alex18"

```

如果 k1 和 k2 不存在时

```shell

root@a72e6809fb18:/tmp# redis-cli --eval test.lua name1 age1 , 12345
1) (nil)
2) 1) "name1"
   2) "1682404163"
   3) "age1"
   4) "1682404163"

```

### 通过增加 `--ldb` 选项启用 Lua 调试器来调试

> 注意：使用 `--ldb` 选项后，执行到最后一步，数据会自动回滚，也就是不会被持久化到 redis 中。你可以观察到，执行到命令最后会有一行 `(Lua debugging session ended -- dataset changes rolled back)` 提示信息。

使用 Redis 内置的 Lua 调试器来调试 Lua 脚本，可以查看变量和命令的结果输出。通过 `--ldb` 选项来启动一个调试会话，可以使用以下命令来控制调试过程：

- step: 执行当前行，并停在下一行。
- continue: 继续执行直到遇到下一个断点或者脚本结束。
- list: 显示当前行周围的源代码。
- break: 设置或删除一个断点。你可以指定一个行号或者不指定任何参数来删除所有断点。
- watch: 查看一个变量的值。你可以指定一个变量名或者不指定任何参数来查看所有变量。
- trace: 查看脚本执行过程中调用的 Redis 命令。
- quit: 退出调试会话。

执行效果如下：

```shell
root@a72e6809fb18:/tmp# redis-cli --ldb --eval test.lua name1 age1 , 12345
Lua debugging session started, please use:
quit    -- End the session.
restart -- Restart the script in debug mode again.
help    -- Show Lua script debugging commands.

* Stopped at 1, stop reason = step over
-> 1   local k1 = KEYS[1]
lua debugger> step
* Stopped at 2, stop reason = step over
-> 2   local k2 = KEYS[2]
lua debugger> step
* Stopped at 4, stop reason = step over
-> 4   local v1 = redis.call('get', k1)
lua debugger> restart

Lua debugging session started, please use:
quit    -- End the session.
restart -- Restart the script in debug mode again.
help    -- Show Lua script debugging commands.

* Stopped at 1, stop reason = step over
-> 1   local k1 = KEYS[1]
lua debugger> step
* Stopped at 2, stop reason = step over
-> 2   local k2 = KEYS[2]
lua debugger> step
* Stopped at 4, stop reason = step over
-> 4   local v1 = redis.call('get', k1)
lua debugger> step
<redis> get name1
<reply> NULL
* Stopped at 5, stop reason = step over
-> 5   local v2 = redis.call('get', k2)
lua debugger> step
<redis> get age1
<reply> NULL
* Stopped at 7, stop reason = step over
-> 7   if v1 and v2 then
lua debugger> step
* Stopped at 12, stop reason = step over
-> 12      local now = tonumber(ARGV[1])
lua debugger> continue

1) (nil)
2) 1) "name1"
   2) "1682404163"
   3) "age1"
   4) "1682404163"

(Lua debugging session ended -- dataset changes rolled back)

127.0.0.1:6379>
```

#### 再进一步，通过使用 `redis.debug()` 函数打印调试信息

```shell

root@a72e6809fb18:/tmp# cat > redis_debug.lua <<EOF
 local k1 = KEYS[1]
 redis.debug('k1 => ', k1)

 local k2 = KEYS[2]
 redis.debug('k2 => ', k1)

 local v1 = redis.call('get', k1)
 redis.debug('v1 => ', type(v1))

 local v2 = redis.call('get', k2)
 redis.debug('v2 => ', type(v2))

 if v1 and v2 then
      local v1_ttl = redis.call('ttl', k1)
     local v2_ttl = redis.call('ttl', k2)
     redis.debug('k1 => ', v1, type(v1_ttl))
     redis.debug('k2 => ', v2, type(v2_ttl))
     return {true, k1, v1, k2, v2, v1_ttl, v2_ttl, v1..v2}
 else
     redis.debug('both v1 and v2 of not exists', v1, v2)
     local now = tonumber(ARGV[1])
     redis.call('hsetnx', 'hash_table', k1, now)
     redis.call('hsetnx', 'hash_table', k2, now)
     local hash_all = redis.call('hgetall', 'hash_table')
     redis.debug('reset the value is ==> ', hash_all, 'hasl_all type is ==> ', type(hash_all))
     return {false, hash_all}
 end

 EOF

```

调试

```shell

root@a72e6809fb18:/tmp# redis-cli --ldb --eval redis_debug.lua name age , 12345
Lua debugging session started, please use:
quit    -- End the session.
restart -- Restart the script in debug mode again.
help    -- Show Lua script debugging commands.

* Stopped at 1, stop reason = step over
-> 1   local k1 = KEYS[1]
lua debugger> step
* Stopped at 2, stop reason = step over
-> 2   redis.debug('k1 => ', k1)
lua debugger> step
<debug> line 2: "k1 => ", "name"
* Stopped at 4, stop reason = step over
-> 4   local k2 = KEYS[2]
lua debugger> step
* Stopped at 5, stop reason = step over
-> 5   redis.debug('k2 => ', k1)
lua debugger> step
<debug> line 5: "k2 => ", "name"
* Stopped at 7, stop reason = step over
-> 7   local v1 = redis.call('get', k1)
lua debugger> step
<redis> get name
<reply> "alex"
* Stopped at 8, stop reason = step over
-> 8   redis.debug('v1 => ', type(v1))
lua debugger> step
<debug> line 8: "v1 => ", "string"
* Stopped at 10, stop reason = step over
-> 10  local v2 = redis.call('get', k2)
lua debugger> step
<redis> get age
<reply> "18"
* Stopped at 11, stop reason = step over
-> 11  redis.debug('v2 => ', type(v2))
lua debugger> step
<debug> line 11: "v2 => ", "string"
* Stopped at 13, stop reason = step over
-> 13  if v1 and v2 then
lua debugger> step
* Stopped at 14, stop reason = step over
-> 14      local v1_ttl = redis.call('ttl', k1)
lua debugger> step
<redis> ttl name
<reply> -1
* Stopped at 15, stop reason = step over
-> 15      local v2_ttl = redis.call('ttl', k2)
lua debugger> step
<redis> ttl age
<reply> 68213
* Stopped at 16, stop reason = step over
-> 16      redis.debug('k1 => ', v1, type(v1_ttl))
lua debugger> step
<debug> line 16: "k1 => ", "alex", "number"
* Stopped at 17, stop reason = step over
-> 17      redis.debug('k2 => ', v2, type(v2_ttl))
lua debugger> step
<debug> line 17: "k2 => ", "18", "number"
* Stopped at 18, stop reason = step over
-> 18      return {true, k1, v1, k2, v2, v1_ttl, v2_ttl, v1..v2}
lua debugger> step

1) (integer) 1
2) "name"
3) "alex"
4) "age"
5) "18"
6) (integer) -1
7) (integer) 68213
8) "alex18"

(Lua debugging session ended -- dataset changes rolled back)

127.0.0.1:6379>

```

```shell

root@a72e6809fb18:/tmp# redis-cli --ldb --eval redis_debug.lua name2 age2 , 55333
Lua debugging session started, please use:
quit    -- End the session.
restart -- Restart the script in debug mode again.
help    -- Show Lua script debugging commands.

* Stopped at 1, stop reason = step over
-> 1   local k1 = KEYS[1]
lua debugger> step
* Stopped at 2, stop reason = step over
-> 2   redis.debug('k1 => ', k1)
lua debugger> step
<debug> line 2: "k1 => ", "name2"
* Stopped at 4, stop reason = step over
-> 4   local k2 = KEYS[2]
lua debugger> step
* Stopped at 5, stop reason = step over
-> 5   redis.debug('k2 => ', k1)
lua debugger> step
<debug> line 5: "k2 => ", "name2"
* Stopped at 7, stop reason = step over
-> 7   local v1 = redis.call('get', k1)
lua debugger> step
<redis> get name2
<reply> NULL
* Stopped at 8, stop reason = step over
-> 8   redis.debug('v1 => ', type(v1))
lua debugger> step
<debug> line 8: "v1 => ", "boolean"
* Stopped at 10, stop reason = step over
-> 10  local v2 = redis.call('get', k2)
lua debugger> step
<redis> get age2
<reply> NULL
* Stopped at 11, stop reason = step over
-> 11  redis.debug('v2 => ', type(v2))
lua debugger> step
<debug> line 11: "v2 => ", "boolean"
* Stopped at 13, stop reason = step over
-> 13  if v1 and v2 then
lua debugger> step
* Stopped at 20, stop reason = step over
-> 20      redis.debug('both v1 and v2 of not exists', v1, v2)
lua debugger> step
<debug> line 20: "both v1 and v2 of not exists", false, false
* Stopped at 21, stop reason = step over
-> 21      local now = tonumber(ARGV[1])
lua debugger> step
* Stopped at 22, stop reason = step over
-> 22      redis.call('hsetnx', 'hash_table', k1, now)
lua debugger> step
<redis> hsetnx hash_table name2 55333
<reply> 1
* Stopped at 23, stop reason = step over
-> 23      redis.call('hsetnx', 'hash_table', k2, now)
lua debugger> step
<redis> hsetnx hash_table age2 55333
<reply> 1
* Stopped at 24, stop reason = step over
-> 24      local hash_all = redis.call('hgetall', 'hash_table')
lua debugger> step
<redis> hgetall hash_table
<reply> ["name1","1682404163","age1","1682404163","name2","55333","age2","55333"]
* Stopped at 25, stop reason = step over
-> 25      redis.debug('reset the value is ==> ', hash_all, 'hasl_all type is ==> ', type(hash_all))
lua debugger> step
<debug> line 25: "reset the value is ==> ", {"name1"; "1682404163"; "age1"; "1682404163"; "name2"; "55333"; "age2"; "55333"}, "hasl_all type is ==> ", "table"
* Stopped at 26, stop reason = step over
-> 26      return {false, hash_all}
lua debugger> step

1) (nil)
2) 1) "name1"
   2) "1682404163"
   3) "age1"
   4) "1682404163"
   5) "name2"
   6) "55333"
   7) "age2"
   8) "55333"

(Lua debugging session ended -- dataset changes rolled back)

127.0.0.1:6379>

```

## 第二种调试方式：使用 eval redis 命令进行调试

> 使用这种方式数据会被持久化，但是如果 lua 脚本比较复杂，那么命令行写起来就比较麻烦，不太直观。

先进入 redis cli 中

```shell
redis-cli
```

执行命令

```

eval "local k1 = KEYS[1] \local k2 = KEYS[2] \local v1 = redis.call('get', k1) \local v2 = redis.call('get', k2) \if v1 and v2 then \    local v1ttl = redis.call('ttl', k1) \    local v2ttl = redis.call('ttl', k2) \    return {true, k1, v1, k2, v2, v1ttl, v2ttl, v1..v2} \else \    local now = tonumber(ARGV[1]) \    redis.call('hsetnx', 'hash_table', k1, now) \    redis.call('hsetnx', 'hash_table', k2, now) \    local hash_all = redis.call('hgetall', 'hash_table') \    return {false, hash_all} \end" 2 name age 78675

```

执行效果

```shell

127.0.0.1:6379> eval "local k1 = KEYS[1] \local k2 = KEYS[2] \local v1 = redis.call('get', k1) \local v2 = redis.call('get', k2) \if v1 and v2 then \    local v1ttl = redis.call('ttl', k1) \    local v2ttl = redis.call('ttl', k2) \    return {true, k1, v1, k2, v2, v1ttl, v2ttl, v1..v2} \else \    local now = tonumber(ARGV[1]) \    redis.call('hsetnx', 'hash_table', k1, now) \    redis.call('hsetnx', 'hash_table', k2, now) \    local hash_all = redis.call('hgetall', 'hash_table') \    return {false, hash_all} \end" 2 name age 78675
1) (integer) 1
2) "name"
3) "alex"
4) "age"
5) "18"
6) (integer) -1
7) (integer) 65267
8) "alex18"
127.0.0.1:6379>

```

```shell

127.0.0.1:6379> eval "local k1 = KEYS[1] \local k2 = KEYS[2] \local v1 = redis.call('get', k1) \local v2 = redis.call('get', k2) \if v1 and v2 then \    local v1ttl = redis.call('ttl', k1) \    local v2ttl = redis.call('ttl', k2) \    return {true, k1, v1, k2, v2, v1ttl, v2ttl, v1..v2} \else \    local now = tonumber(ARGV[1]) \    redis.call('hsetnx', 'hash_table', k1, now) \    redis.call('hsetnx', 'hash_table', k2, now) \    local hash_all = redis.call('hgetall', 'hash_table') \    return {false, hash_all} \end" 2 name4 age4 78675
1) (nil)
2)  1) "name2"
    2) "55333"
    3) "age2"
    4) "55333"
    5) "name3"
    6) "1682408424"
    7) "age3"
    8) "1682408424"
    9) "name4"
   10) "78675"
   11) "age4"
   12) "78675"
127.0.0.1:6379>

```

## lua 脚本自身错误，如何处理？

那，假设我写的 lua 脚本自身就是有问题的呢？那应该如何排查？

> 因为不管是 lua 脚本自身语法有错误，还是因为 lua 脚本根本就没有返回值时，均会返回 `false` 所以有时候你根本搞不清楚 lua 脚本到底有没有执行成功，当然你完全可以通过返回一个固定值来进行判断，也是可行的。

比如我将代码调整成以下：

```php

<?php

$redis = new \Redis();

$redis->connect('192.168.10.194');

$lua = <<<EOF
local k1 = KEYS[1]
local k2 = KEYS[2]

-- 故意将 get 命令写成 get1 以便引发一个 lua 脚本错误
local v1 = redis.call('get1', k1)
local v2 = redis.call('get', k2)

if v1 and v2 then
    local v1_ttl = redis.call('ttl', k1)
    local v2_ttl = redis.call('ttl', k2)
    return {true, k1, v1, k2, v2, v1_ttl, v2_ttl, v1..v2}
else
    local now = tonumber(ARGV[1])
    redis.call('hsetnx', 'hash_table', k1, now)
    redis.call('hsetnx', 'hash_table', k2, now)
    local hash_all = redis.call('hgetall', 'hash_table')
    return {false, hash_all}
end

EOF;

$res = $redis->eval($lua, ['name', 'age', time()], 2);

var_dump($res);

```

执行情况

> 其实这里返回 `false` 是因为 lua 脚本执行失败，并不是因为 lua 脚本执行成功，但是没有返回值导致。

```
bash-5.0# php demo.php 
bool(false)
bash-5.0# 
```

### 使用 lua 脚本中的 `pcall` 函数处理错误与异常

> [lua 错误处理和异常](https://www.lua.org/pil/8.4.html)

代码更改为如下：

```php

<?php

$redis = new \Redis();

$redis->connect('192.168.10.194');

$lua = <<<EOF
local ok, res = pcall(function()

local k1 = KEYS[1]
local k2 = KEYS[2]

-- 故意将 get 命令写成 get1 以便引发一个 lua 脚本错误
local v1 = redis.call('get1', k1)
local v2 = redis.call('get', k2)

if v1 and v2 then
    local v1_ttl = redis.call('ttl', k1)
    local v2_ttl = redis.call('ttl', k2)
    return {true, k1, v1, k2, v2, v1_ttl, v2_ttl, v1..v2}
else
    local now = tonumber(ARGV[1])
    redis.call('hsetnx', 'hash_table', k1, now)
    redis.call('hsetnx', 'hash_table', k2, now)
    local hash_all = redis.call('hgetall', 'hash_table')
    return {false, hash_all}
end

end)

return {ok, res}

EOF;

$res = $redis->eval($lua, ['name', 'age', time()], 2);

var_dump($res);

```

返回结果

```shell

# 可以通过返回值中的第一个值来进行判断 lua 脚本是否执行成功，如果为 false ，第二个值即为错误提示信息

bash-5.0# php demo.php 
array(2) {
  [0]=>
  bool(false)
  [1]=>
  string(61) "@user_script: 6: Unknown Redis command called from Lua script"
}
bash-5.0# 


# lua 脚本无异常时，第一个值为 1，第二个值即为匿名函数中返回的值

bash-5.0# php demo.php 
array(2) {
  [0]=>
  int(1)
  [1]=>
  array(8) {
    [0]=>
    int(1)
    [1]=>
    string(4) "name"
    [2]=>
    string(4) "alex"
    [3]=>
    string(3) "age"
    [4]=>
    string(2) "18"
    [5]=>
    int(-1)
    [6]=>
    int(62010)
    [7]=>
    string(6) "alex18"
  }
}

```
