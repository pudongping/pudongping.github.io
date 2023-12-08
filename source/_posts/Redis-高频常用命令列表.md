---
title: Redis 高频常用命令列表
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Cache
  - Redis
  - 缓存
abbrlink: d7e52b6f
date: 2023-12-09 00:55:25
img:
coverImg:
password:
summary:
---

# Redis 的数据类型

## redis-cli 命令

命令 | 说明 | cli 命令示例
--- | --- | ---
del | 删除 key | del key_name
exists | 检查给定 key 是否存在 | exists key_name
keys | 查找所有符合给定模式 pattern 的 key | keys pattern
type | 返回 key 所存储的值的类型 | type key_name
expire | 设置 key 的过期时间 | expire key_name time_in_seconds
ttl | 返回 key 的剩余过期时间 | ttl key_name
save | RDB 持久化 | save
info | Redis 服务器的各种信息和统计数值 | inro [section]
shutdown | 保存并停止所有客户端 | shutdown [nosave] [save]
flushall | 清空整个 redis 服务器的数据 | flushall
flushdb | 清空当前库中所有的 key | flushdb
select | 切换到指定的数据库 | select db_number


## Redis 字段类型

- 字符串 String
- 散列/哈希 Hash
- 列表 List
- 无序集合 Set
- 可排序集合 Zset

### 字符串 String

> 最大容量为 512M

命令 | 说明 | Cli 命令示例  | PHP 写法
--- | --- | --- | ---
set | 赋值 | set key value | $redis->set('key', 'value');
setex | 赋值并添加过期时间 | setex key expire value | $redis->setex('key', 'expire', 'value');
get | 取值 | get key | $redis->get('key');
incr | 递增数字 | incr key | $redis->incr('int_key');
incrby | 增加指定的数字 | incrby key increment | $redis->incrBy('int_key', number);
decr | 递减数字 | decr key | $redis->decr('key1');
decrby | 减少指定的数字 | decrby key decrement | $redis->decrBy('key1', number);
incrbyfloat | 增加指定浮点数 | incrbyfloat key increment | $redis->incrByFloat('key1', 1.5);
append | 向尾部追加值 | append key value | $redis->append('key', 'value2');
strlen | 获取字符串长度 | strlen key | $redis->strlen('key');
mset | 同时设置多个 key 的值 | mset key1 value1 [key2 value2 ...] | $redis->mSet(array('key0' => 'value0', 'key1' => 'value1'));
mget | 同时获取多个 key 的值 | mget key1 [key2 ...] | $redis->mGet(array('key1', 'key2', 'key3'));


###  散列/哈希 Hash

> 1、与 php 的 array 相似；2、可以保存多个 key-value 对，每个 k-v 都是字符串类型；3、最多 2^32-1 个字段。

命令 | 说明 | Cli 命令示例  | PHP 写法
--- | --- | --- | ---
hset | 赋值 | hset key field value | $redis->hSet(key, field, value);
hmset | 赋值多个字段 | hmset key field1 value1 [field2 values] | $redis->hMset(key, array('field1' => 'value1', 'field2' => 'value2'));
hget | 取值 | hset key field | $redis->hGet(key, field);
hmget | 取多个字段的值 | hmset key field1 [field2] | $redis->hmGet(key, array('field1', 'field2'));
hgetall | 取所有字段的值 | hgetall key | $redis->hGetAll(key);
hlen | 获取字段的数量 | hlen key | $redis->hLen(key);
hexists | 判断字段是否存在 | hexists key field | $redis->hExists(key, field);
hsetnx | 当字段不存在时赋值 | hsetnx key field value | $redis->hSetNx(key, field, value);
hincrby | 增加数字 | hincrby key field increment | $redis->hIncrBy(key, field, num);
hdel | 删除字段 | hdel key field | $redis->hDel(key, field);
hkeys | 获取所有字段名 | hkeys key | $redis->hKeys(key);
hvals | 获取所有字段值 | hvals key | $redis->hVals(key);


###  列表 List

> 1、实现方式为双向链表；2、用于存储一个有序的字符串列表；3、从队列两端添加和弹出元素；4、特别适合于做消息队列。

命令 | 说明 | Cli 命令示例  | PHP 写法
--- | --- | --- | ---
lpush | 向列表左端添加元素 | lpush key value | $redis->lPush(key, value);
rpush | 向列表右端添加元素 | rpush key value | $redis->rPush(key, value);
lpop | 从列表左端弹出元素 | lpop key | $redis->lPop(key);
rpop | 从列表右端弹出元素 | rpop key | $redis->rPop(key);
llen | 获取列表中元素个数 | llen key | $redis->lSize(key);
lrange | 获取列表中某一片段的元素 | lrange key start stop | $redis->lRange(key, start, end);
lrem | 删除列表中指定的值 | lrem key count value | $redis->lRem(key, value, count);
lindex | 获取指定索引的元素值 | lindex key index | $redis->lGet(key, index);
lset | 设置指定索引的元素值 | lset key index value | $redis->lSet(key, index, value);
ltrim | 值保留列表指定片段 | ltrim key start stop | $redis->lTrim(key, start, end);
linsert | 向列表中插入元素 | linsert key before/after existing_value value | $redis->lInsert(key, Redis::BEFORE, existing_value, value);


###  无序集合 Set

> 1、集合中每个元素都是不同的；2、元素最多为 2^32-1；3、元素没有顺序

命令 | 说明 | Cli 命令示例  | PHP 写法
--- | --- | --- | ---
sadd | 添加元素 | sadd key value1 [value2 value3 ...] | $redis->sAdd('key1', 'set1');
srem | 删除元素 | srem key value2 [value2 value3 ...] | $redis->sRem('key', 'set2');
smembers | 获得集合中所有元素 | smembers key | $redis->sMembers('key');
sismember | 判断元素是否在集合中 | sismember key value | $redis->slsMember(key, value);
sdiff | 对集合做差集运算 | sdiff key1 key2 [key3 ...] | $redis->sDiff(key1, key2, key3);
sinter | 对集合做交集运算 | sinter key1 key2 [key3 ...] | $redis->sInter('key1', 'key2', 'key3');
sunion | 对集合做并集运算 | sunion key1 key2 [key3 ...] | $redis->sUnion('key1', 'key2', 'key3');
scard | 获得集合中元素的个数 | scard key | $redis->sCard('key1');
sdiffstore | 对集合做差集运算并将结果存储 | sdiffstore destination key1 key2 [key3 ...] | $redis->sDiffStore('output', key1, key2, key3);
sinterstore | 对集合做交集运算并将结果存储 | sinterstore destination key1 key2 [key3 ...] | $redis->sInterStore('output', 'key1', 'key2', 'key3');
sunionstore | 对集合做并集运算并将结果存储 | sunionstore destination key1 key2 [key3 ...] | $redis->sUnionStore('output', 'key1', 'key2', 'key3');
srandmember | 随机获取集合中的元素 | srandmember key [count] | $redis->sRandMember('key1', 2);
spop | 随机弹出一个元素 | spop key | $redis->sPop('key1');


###  可排序集合 Zset

> 1、集合是有序的；2、支持插入，删除，判断元素是否存在；3、可以获取分数最高/最低的前 N 个元素。

命令 | 说明 | Cli 命令示例  | PHP 写法
--- | --- | --- | ---
zadd | 添加元素 | zadd key score1 value1 [score2 value2 score3 value3 ...] | $redis->zAdd('key', 1, 'val1');
zscore | 获取元素的分数 | zscore key value | $redis->zScore(key, val2);
zrange | 获取正序排名在某索引区间范围的元素 | zrange key start stop [withscore] | $redis->zRange('key1', 0, -1);
zrevrange | 获取倒序排名在某索引区间范围的元素 | zrevrange key start stop [withscore] | $redis->zRevRange('key1', 0, -1);
zrangebyscore | 获取指定分数范围内的元素 | zrangebyscore key min max | $redis->zRangeByScore(key, start, end, array(withscores, limit));
zincrby | 增加某个元素的分数 | zcard key | $redis->zSize('key');
zcount | 获取指定分数范围内的元素个数 | zcount key min max | $redis->zCount(key, start, end);
zrem | 删除一个或多个元素 | zrem key value1 [value2 ...] | $redis->zDelete('key', 'val2');
zremrangebyrank | 按照排名索引区间范围删除元素 | zremrangebyrank key start stop | $redis->zRemRangeByRank('key', 0, 1);
zremrangebyscore | 按照分数范围删除元素 | zremrangebyscore key min max | $redis->zRemRangeByScore('key', 0, 3);
zrank | 获取正序排序的元素的排名 | zrank key value | $redis->zRank(key, value);
zrevrank | 获取逆序排序的元素的排名 | zrevrank key value | $redis->zRevRank(key, value);
