---
title: Shell 内建命令：Shell 的内在魔力
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Shell
tags:
  - Shell
  - Linux
abbrlink: a915d50b
date: 2024-06-20 22:44:56
img:
coverImg:
password:
summary:
---

# shell 内建命令（内置命令）

今天我们来深入挖掘一下 Shell 的内在魔力——内建命令。

通常来说，内建命令会比外部命令执行得更快，执行外部命令时不但会触发磁盘 I/O，还需要 fork 出一个单独的进程来执行，执行完成后再退出。而执行内建命令相当于调用当前 Shell 进程的一个函数。

## 检查一个命令是否是内建命令

```bash

# cd 是一个内建命令
type cd
# cd is a shell builtin

# 可见 ifconfig 是一个外部文件，它的位置时 /sbin/ifconfig
type ifconfig
# ifconfig is /sbin/ifconfig

```

## Bash Shell 内建命令

命令 | 说明
--- | ---
:	| 扩展参数列表，执行重定向操作
.	| 读取并执行指定文件中的命令（在当前 shell 环境中）
alias | 	为指定命令定义一个别名
bg	| 将作业以后台模式运行
bind	| 将键盘序列绑定到一个 readline 函数或宏
break	| 退出 for、while、select 或 until 循环
builtin	| 执行指定的 shell 内建命令
caller	| 返回活动子函数调用的上下文
cd	| 将当前目录切换为指定的目录
command	| 执行指定的命令，无需进行通常的 shell 查找
compgen	| 为指定单词生成可能的补全匹配
complete	| 显示指定的单词是如何补全的
compopt | 	修改指定单词的补全选项
continue	| 继续执行 for、while、select 或 until 循环的下一次迭代
declare	| 声明一个变量或变量类型。
dirs	| 显示当前存储目录的列表
disown	| 从进程作业表中刪除指定的作业
echo | 	将指定字符串输出到 STDOUT
enable	| 启用或禁用指定的内建shell命令
eval	| 	将指定的参数拼接成一个命令，然后执行该命令
exec	| 	用指定命令替换 shell 进程
exit	| 	强制 shell 以指定的退出状态码退出
export	| 	设置子 shell 进程可用的变量
fc		| 从历史记录中选择命令列表
fg		| 将作业以前台模式运行
getopts		| 分析指定的位置参数
hash		| 查找并记住指定命令的全路径名
help	| 	显示帮助文件
history		| 显示命令历史记录
jobs	| 	列出活动作业
kill	| 	向指定的进程 ID(PID) 发送一个系统信号
let		| 计算一个数学表达式中的每个参数
local	| 	在函数中创建一个作用域受限的变量
logout	| 	退出登录 shell
mapfile		| 从 STDIN 读取数据行，并将其加入索引数组
popd	| 	从目录栈中删除记录
printf	| 	使用格式化字符串显示文本
pushd	| 	向目录栈添加一个目录
pwd		| 显示当前工作目录的路径名
read	| 	从 STDIN 读取一行数据并将其赋给一个变量
readarray	| 	从 STDIN 读取数据行并将其放入索引数组
readonly	| 	从 STDIN 读取一行数据并将其赋给一个不可修改的变量
return	| 	强制函数以某个值退出，这个值可以被调用脚本提取
set		| 设置并显示环境变量的值和 shell 属性
shift	| 	将位置参数依次向下降一个位置
shopt	| 	打开/关闭控制 shell 可选行为的变量值
source		| 读取并执行指定文件中的命令（在当前 shell 环境中）
suspend		| 暂停 Shell 的执行，直到收到一个 SIGCONT 信号
test	| 	基于指定条件返回退出状态码 0 或 1
times	| 	显示累计的用户和系统时间
trap	| 	如果收到了指定的系统信号，执行指定的命令
type	| 	显示指定的单词如果作为命令将会如何被解释
typeset	| 	声明一个变量或变量类型。
ulimit	| 	为系统用户设置指定的资源的上限
umask	| 	为新建的文件和目录设置默认权限
unalias	| 	刪除指定的别名
unset	| 	刪除指定的环境变量或 shell 属性
wait	| 	等待指定的进程完成，并返回退出状态码

## alias 给命令创建别名

### 查看所有别名

```bash
# 不带任何参数，则列出当前 shell 进程中所有别名
alias
```

### 设置别名

```bash
# 为获取当前的 unix 时间戳设置别名 timestamp
alias timestamp='date +%s'
```

### 删除别名

```bash
# 删除 timestamp 别名
unalias timestamp
```

## echo 用于在终端输出字符串

默认在末尾加上了换行符

### 不换行

```bash

#!/bin/bash

name="Alex"
age=26
height=168
weight=62
echo -n "${name} is ${age} years old, "
echo -n "${height}cm in height "
echo "and ${weight}kg in weight."
echo "Thank you!"

# Alex is 26 years old, 168cm in height and 62kg in weight.
# Thank you!
```

### 输出转义字符

```bash

#!/bin/bash


name="Alex"
age=26
height=168
weight=62
echo -e "${name} is ${age} years old,\n "
echo -n "${height}cm in height "
echo "and ${weight}kg in weight."
echo "Thank you!"

# Alex is 26 years old,
#
# 168cm in height and 62kg in weight.
# Thank you!
```

### 强制不换行

```bash

#!/bin/bash


name="Alex"
age=26
height=168
weight=62
echo -e "${name} is ${age} years old,\c "
echo -e "${height}cm in height "
echo "and ${weight}kg in weight."
echo "Thank you!"

# Alex is 26 years old,168cm in height
# and 62kg in weight.
# Thank you!

```

## printf

格式替代符 | 含义
--- | ---
%s | 输出一个字符串
%d | 输出一个整型
%c | 输出一个字符
%f | 输出一个小数
%-10s | 指一个宽度为 10 个字符（-表示左对齐，没有则表示右对齐），任何字符都会被显示在 10 个字符宽的字符内，如果不足则自动以空格填充，超过也会将内容全部显示出来。
%-4.2f | 格式化为小数，其中 `.2` 指保留 2 位小数


```bash

#!/bin/bash

printf "%-10s %-8s %-4s\n" 姓名 性别 体重kg
printf "%-10s %-8s %-4.2f\n" alex 男 62.3452

# 姓名     性别   体重kg
# alex       男      62.35

```

## read 用来从标准输入中读取数据并赋值给变量

> 如果没有进行重定向，默认就是从键盘读取用户输入的数据；如果进行了重定向，那么可以从文件中读取数据。

read 命令的用法为：

```bash
# options 表示选项
# variables 表示用来存储数据的变量，可以有一个，也可以有多个
read [-options] [variables]
```

options 支持的选项有：

选项 | 说明
--- | ---
-a new_array  |	把读取的数据赋值给数组 new_array，从下标 0 开始。
-d delimiter |		用字符串 delimiter 指定读取结束的位置，而不是一个换行符（读取到的数据不包括 delimiter）。
-e	 |	在获取用户输入的时候，对功能键进行编码转换，不会直接显式功能键对应的字符。
-n num	 |	读取 num 个字符，而不是整行字符。
-p prompt |		显示提示信息，提示内容为 prompt。
-r	 |	原样读取（Raw mode），不把反斜杠字符解释为转义字符。
-s	 |	静默模式（Silent mode），不会在屏幕上显示输入的字符。当输入密码和其它确认信息的时候，这是很有必要的。
-t seconds	 |	 设置超时时间，单位为秒。如果用户没有在指定时间内输入完成，那么 read 将会返回一个非 0 的退出状态，表示读取失败。
-u fd	 |	使用文件描述符 fd 作为输入源，而不是标准输入，类似于重定向。

```bash

#!/bin/bash
# 使用 read 命令给多个变量赋值

read -p "Enter your name, age and city ===> " name age city

echo "你的名字为：${name}"
echo "你的年龄为：${age}"
echo "你所在的城市为：${city}"


# Enter your name, age and city ===> alex 26 Shanghai
# 你的名字为：alex
# 你的年龄为：26
# 你所在的城市为：Shanghai

#########################################################

#!/bin/bash
# 在指定时间内输入密码
if
    read -t 20 -sp "Enter password in 20 seconds(once) ====> " pass1 && printf "\n" && # 第一次输入密码
    read -t 20 -sp "Please confirm your password again in 20 seconds ====> " pass2 && printf "\n" && # 确认密码
    [ $pass1 == $pass2 ]  # 判断两次输入的密码是否相等
then
    echo "your password is ok!"
else
    echo "Invalid password"
fi

```

## exit 用来退出当前 shell 进程，并返回一个退出状态

- 可以使用 `$?` 接收这个退出状态
- 可以接受一个整数值作为参数，代表退出状态，如果不指定，默认状态值是 0
- 退出状态为 0 表示成功，退出状态非 0 表示执行出错或失败
- 退出状态只能是一个介于 `0~255` 之间的整数，其中只有 0 表示成功，其他值都表示失败

```bash

#!/bin/bash
echo "before exit"  # 只会输出 before exit
exit 1
echo "after exit"  # 不会输出

```

## declare 和 typeset 用来设置变量属性

> typeset 已经被废弃，建议使用 declare

declare 的用法为：

```bash
# - 表示设置属性
# + 表示取消属性
# aAfFgilprtux 表示具体的选项
declare [+/-] [aAfFgilprtux] [变量名=变量值]
```

aAfFgilprtux 支持的选项有：

选项 | 说明
--- | ---
-f [name]	| 列出之前由用户在脚本中定义的函数名称和函数体。
-F [name]		| 仅列出自定义函数名称。
-g name		| 在 Shell 函数内部创建全局变量。
-p [name]		| 显示指定变量的属性和值。
-a name		| 声明变量为普通数组。
-A name		| 声明变量为关联数组（支持索引下标为字符串）。
-i name 	| 	将变量定义为整数型。
-r name[=value] 		| 将变量定义为只读（不可修改和删除），等价于 readonly name。
-x name[=value]		| 将变量设置为环境变量，等价于 export name[=value]。

```bash

#!/bin/bash
# 将变量声明为整数并进行计算

declare -i x y ret
x=11
y=22
ret=$x+$y
echo $ret  # 33

```

内建命令是 Shell 的核心功能，它们提供了快速且强大的工具来处理日常任务。掌握这些内建命令，可以帮助你更高效地编写 Shell 脚本和命令行程序。

希望这篇文章能够帮助你更好地理解和使用 Shell 内建命令。