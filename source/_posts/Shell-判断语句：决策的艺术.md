---
title: Shell 判断语句：决策的艺术
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
abbrlink: 43dbd7ab
date: 2024-06-18 14:18:13
img:
coverImg:
password:
summary:
---

编写 Shell 脚本时，了解如何根据不同条件执行不同的命令是至关重要的。本文旨在以简单易懂的语言，介绍 Shell 脚本中的选择结构——if 语句和 case in 语句，帮助初学者和有经验的开发者深入理解。

## if 语句

if 语句用于基于条件的执行。在 shell 中，if 语句的基础语法如下：

```bash
#!/bin/bash

if condition
then
  statement(s)
fi  
```

在实际应用中，我们通常会遇到需要直接在 if 语句后紧接着写 then 的情况。这时，应该使用分号分隔，否则会导致语法错误。如下所示：

```bash
#!/bin/bash

# 分号是必须的，否则会出现语法错误
if condition; then
  statement(s)
fi
```

### if else 语句

当 if 语句的条件不满足时，else 部分的代码将会被执行：

```bash
#!/bin/bash

if condition
then
   statement1
else
   statement2
fi
```

### if elif else 语句

对于多条件判断，可以使用 if, elif, else 结构来实现：

```bash
#!/bin/bash

if condition1
then
   statement1
elif condition2
then
    statement2
elif condition3
then
    statement3
# 可以有更多的 elif 分支
else
   statementn
fi
```

举一个例子，输入一个整数，输出该整数对应的星期几的英文表示：

```bash
#!/bin/bash

printf "Input integer number: "
read num

if ((num==1)); then
    echo "Monday"
elif ((num==2)); then
    echo "Tuesday"
elif ((num==3)); then
    echo "Wednesday"
elif ((num==4)); then
    echo "Thursday"
elif ((num==5)); then
    echo "Friday"
elif ((num==6)); then
    echo "Saturday"
elif ((num==7)); then
    echo "Sunday"
else
    echo "error"
fi

# Input integer number: 5
# Friday
```

## case in 语句

case in 语句是一种更为灵活的条件分支语句，它能够匹配具体的值或者模式。基本语法如下：

```bash
#!/bin/bash

printf "Input integer number: "
read num

case $num in
  1)
    echo 'Monday'
  ;;
  2)
    echo 'Tuesday'
  ;;
  3)
    echo 'Wednesday'
  ;;
  4)
    echo 'Thursday'
  ;;
  5)
    echo 'Friday'
  ;;
  6)
    echo 'Saturday'
  ;;
  7)
    echo 'Sunday'
  ;;
  *)
    echo 'error'
esac
```

### case in 支持的简单正则表达式

- `*` 表示任意字符串。
- `[abc]` 表示 a、b、c 三个字符中的任意一个。比如，[15ZH] 表示 1、5、Z、H 四个字符中的任意一个。
- `[m-n]` 表示从 m 到 n 的任意一个字符。比如，[0-9] 表示任意一个数字，[0-9a-zA-Z] 表示字母或数字。
- `|` 表示多重选择，类似逻辑运算中的或运算。比如，abc | xyz 表示匹配字符串 "abc" 或者 "xyz"。

一个实际例子如下：

```bash
#!/bin/bash

printf 'Input a character: '
read char

case $char in
  [abc])
    echo 'abc 中任意一个'
  ;;
  [0-9])
    echo '数字'
  ;;
  abc|xyz)
    echo '不是 abc 就是 xyz'
  ;;
  [,.?!])
    echo '特殊符号'
  ;;
  *)
    echo '其他'
esac
```

选择结构是脚本编程中不可或缺的部分，它让我们能够根据不同的条件执行不同的代码。Shell 提供的 if 语句和 case 语句各有千秋，可以根据实际需要选择使用。

通过上面的示例，你应该对 Shell 脚本中的判断语句有了基本的了解。无论你是刚开始学习编程，还是已经有一定的经验，希望本文能帮助你更好地理解和使用 Shell 脚本中的条件判断功能。