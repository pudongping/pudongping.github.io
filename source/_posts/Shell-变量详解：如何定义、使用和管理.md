---
title: Shell 变量详解：如何定义、使用和管理
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
abbrlink: c60e5e2c
date: 2024-06-19 22:38:54
img:
coverImg:
password:
summary:
---

在编写 Shell 脚本时，变量扮演着极为重要的角色。它们使我们能够临时保存数据，以便在脚本的其他部分中使用。

本文将通过简单的语言和清晰的示例，帮助你掌握 Shell 变量的基础知诀，无论你是初学者还是具备一定编程经验的开发者，都能从中获益。

## 变量的基本规则

在 Bash Shell 中，变量的值默认为字符串类型，且在进行变量赋值时，等号`=`两侧不能有空格。

### 变量命名规范

1. 变量名可由数字、字母、下划线构成；
2. 必须以字母或下划线开头；
3. 不能使用 Shell 中的关键字（可通过 `help` 命令查看保留关键字列表）。

### 特殊变量一览

下面的表格简要介绍了几个常用的特殊变量及其含义：

| 变量   | 含义 |
| ------ | ---- |
| `$0`   | 当前脚本的文件名 |
| `$n`   | 脚本或函数的第 n 个参数。注意：当 n≥10 时，应使用 `${n}` 的格式 |
| `$#`   | 传递给脚本或函数的参数个数 |
| `$*`   | 传递给脚本或函数的所有参数（作为一个整体） |
| `$@`   | 传递给脚本或函数的所有参数（作为独立的多个值）|
| `$?`   | 上一个命令的退出状态或函数的返回值 |
| `$$`   | 当前 Shell 进程的 PID |
| `$!`   | 后台运行的最后一个进程的 PID |

## 如何定义变量

定义变量的方式主要有三种：不加引号、单引号和双引号。选择哪种方式取决于你希望如何处理其中的特殊字符和变量。

### 单引号包围

单引号内的内容将完全按字面意义处理，不解析变量或执行命令。

```bash
#!/bin/bash

word='Hello $USER'
echo $word  # 输出 Hello $USER 字符串本身
```

### 双引号包围

双引号内的内容可以解析变量，执行命令。

```bash
#!/bin/bash

word="Hello $USER"
echo $word  # 输出 Hello 后跟当前用户名
```

### 不加引号

不加引号时，如果值中包含空格，需要特别注意，因为 Shell 会将空格后的内容视为另一个命令或参数。

```bash
#!/bin/bash

variable=value
echo ${variable}  # 输出 value
```

## 使用变量

在使用变量时，强烈推荐将变量名包围在花括号`{}`中。这不仅是一个良好的编程习惯，而且有助于明确变量的边界。

```bash
#!/bin/bash

name="alex"
echo "My name is ${name}."  # 输出 My name is alex.
```

## 修改变量的值

变量一旦定义后，我们可以按需更改其值。

```bash
#!/bin/bash

name="alex"
echo ${name}  # 输出 alex
name="harry"
echo ${name}  # 输出 harry
```

## 将命令的输出赋值给变量

可以通过反引号`` ` ``或 `$()` 将命令的输出结果赋值给变量，`$()` 方式具有可嵌套的特性且可读性更强。

```bash
#!/bin/bash

path=$(pwd)
echo ${path}  # 输出当前目录路径
```

示例：计算脚本的运行时间

```bash

#!/bin/bash
begin_time=`date +%s`    #开始时间，使用``替换
sleep 20s                #休眠20秒
finish_time=$(date +%s)  #结束时间，使用$()替换
run_time=$((finish_time - begin_time))  #时间差
echo "begin time: $begin_time"
echo "finish time: $finish_time"
echo "run time: ${run_time}s"

```

## 只读变量

使用 `readonly` 命令可以将变量设置为只读，尝试更改这些变量的值将导致错误。

```bash
#!/bin/bash

name="alex"
readonly name
name="ben"  # 尝试执行将引发错误
```

## 删除变量

使用 `unset` 命令可以删除变量。但需要注意的是，这个命令不能删除只读变量。

```bash
#!/bin/bash

path=$(pwd)
unset path
echo ${path}  # 此时没有任何输出
```

至此，我们对 Shell 变量的定义、使用和管理方法有了基本的了解。通过这篇文章，你应该能够在你的脚本中更灵活地使用变量来存储和修改数据了。

记得实践是学习的最佳方式，所以不妨动手尝试一下吧！