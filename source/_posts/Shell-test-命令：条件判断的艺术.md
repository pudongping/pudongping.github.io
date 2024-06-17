---
title: 'Shell test [] 命令：条件判断的艺术'
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
abbrlink: 8deddbf7
date: 2024-06-17 14:23:46
img:
coverImg:
password:
summary:
---

在编程世界里，`shell` 脚本是一种强大而又灵活的方式，用于处理文件、执行命令以及进行自动化操作。

今天，我们将深入探讨 Shell 脚本中的 `test` 命令，这是一种用来检测不同条件（如数值比较、字符串比较和文件存在性检测）是否成立的命令。

## 简介

`test` 命令用以判断一系列表达式是否成立，根据条件的成立与否，它会返回相应的退出状态码。一般来说，如果条件成立，退出状态码为 0；如果条件不成立，退出状态为非 0 值。

语法总览：

```bash
# 标准语法，判断 expression 成立时，退出状态为 0，否则为非 0 值
test expression

# 简写形式，推荐使用，因为更加直观
[ expression ]
```

## 数值比较

首先，我们来看看如何使用 `test` 进行数值比较。

### 示例：判断两个数是否相等

```bash
#!/bin/bash
# 读入两个数
read a b

# 使用 test 命令判断是否相等
if test $a -eq $b; then
  echo "相等"
else
  echo "不相等"
fi
```

## 注意事项

1. 使用 `==`、`>`、`<` 这些符号时，它们只能用于字符串比较，不能用于数字。对于数值，应使用 `-eq`、`-gt`、`-lt` 这样的操作符。
2. 尽管 Shell 支持 `-gt` 和 `-lt` 之类的数值比较操作符，但不支持 `>=` 和 `<=`。
3. 建议在使用变量时，尤其是在字符串比较中，将变量用双引号包围以防止空值或带有空格的值带来问题。
4. 对于整型数值的比较，更推荐使用 `(())` 来进行。

## 文件检测

`test` 命令还可以用于文件的检测，包括文件类型、权限和比较等。

### 文件类型判断

选项 | 作用
:-- | ---
-b filename	 | 判断文件是否存在，并且是否为块设备文件。
-c filename  | 判断文件是否存在，并且是否为字符设备文件。
-d filename	 | 判断文件是否存在，并且是否为目录文件。
-e filename	 | 判断文件是否存在。
-f filename	 | 判断文件是否存在，井且是否为普通文件。
-L filename	 | 判断文件是否存在，并且是否为符号链接文件。
-p filename	 | 判断文件是否存在，并且是否为管道文件。
-s filename	 | 判断文件是否存在，并且是否为非空。
-S filename	 | 判断该文件是否存在，并且是否为套接字文件。

### 文件权限判断

选项 | 作用
:-- | ---
-r filename | 判断文件是否存在，并且是否拥有读权限。
-w filename | 判断文件是否存在，并且是否拥有写权限。
-x filename | 判断文件是否存在，并且是否拥有执行权限。
-u filename | 判断文件是否存在，并且是否拥有 SUID 权限。
-g filename | 判断文件是否存在，并且是否拥有 SGID 权限。
-k filename | 判断该文件是否存在，并且是否拥有 SBIT 权限。

### 文件比较

选项 | 作用
:-- | ---
filename1 -nt filename2 | 判断 filename1 的修改时间是否比 filename2 的新。
filename -ot filename2	 | 判断 filename1 的修改时间是否比 filename2 的旧。
filename1 -ef filename2	 | 判断 filename1 是否和 filename2 的 inode 号一致，可以理解为两个文件是否为同一个文件。这个判断用于判断硬链接是很好的方法

### 示例：检测文件是否存在并可写

```bash
#!/bin/bash

# 读入文件名和内容
read filename
read content

# 检测文件是否可写且内容非空
if test -w "$filename" && test -n "$content"; then
  echo "$content" > "$filename"
  echo "内容写入文件成功"
else
  echo "内容写入失败"
fi
```

## 与数值比较相关的 test 选项

> test 只能用来比较整数，小数相关的比较还是得用 `bc` 命令

选项 | 作用
:-- | ---
num1 -eq num2 |	判断 num1 是否和 num2 相等。
num1 -ne num2 |	判断 num1 是否和 num2 不相等。
num1 -gt num2 |	判断 num1 是否大于 num2 。
num1 -lt num2 |	判断 num1 是否小于 num2。
num1 -ge num2 |	判断 num1 是否大于等于 num2。
num1 -le num2 |	判断 num1 是否小于等于 num2。

## 与字符串判断相关的 test 选项

选项 | 作用
:-- | ---
-z str	| 判断字符串 str 是否为空。
-n str	| 判断宇符串 str 是否为非空。
str1 = str2 <br/> str1 == str2	|  =和==是等价的，都用来判断 str1 是否和 str2 相等。
str1 != str2	| 判断 str1 是否和 str2 不相等。
str1 \\> str2	| 判断 str1 是否大于 str2。\\>是>的转义字符，这样写是为了防止>被误认为成重定向运算符。
str1 \\< str2	| 判断 str1 是否小于 str2。同样，\\<也是转义字符。

### 字符串判断

接下来，让我们学习如何利用 `test` 命令进行字符串相关的判断。

#### 示例：检测两个字符串是否相等

```bash
#!/bin/bash

# 读入两个字符串
read str1
read str2

# 检测字符串是否为空
# 防止 $str1 和 $str2 是空字符串时出现错误，因此需要用双引号括
if [ -z "$str1" ] || [ -z "$str2" ]; then
  echo "字符串不能为空"
  exit 1
elif [ "$str1" != "$str2" ]; then
  echo "两个字符串不相等"
  exit 2
else
  echo "字符串相等"
fi
```

### 逻辑运算

最后，`test` 命令还支持逻辑运算，这包括 `逻辑与`、`逻辑或` 和 `逻辑非`。

选项 | 作用
:-- | ---
expression1 -a expression	| 逻辑与，表达式 expression1 和 expression2 都成立，最终的结果才是成立的。
expression1 -o expression2	| 逻辑或，表达式 expression1 和 expression2 有一个成立，最终的结果就成立。
!expression	| 逻辑非，对 expression 进行取反。

#### 示例：利用逻辑或进行字符串空判断

```bash
#!/bin/bash

# 读入两个字符串
read str1
read str2

# 使用逻辑或检测字符串是否为空
# 使用 -o 选项取代上面的 ||
if [ -z "$str1" -o -z "$str2" ]; then
  echo "字符串不能为空"
  exit 1
elif [ "$str1" != "$str2" ]; then
  echo "两个字符串不相等"
  exit 2
else
  echo "字符串相等"
fi
```

通过上述介绍和示例，相信你已经对 `test` 命令有了更深入的了解。`test` 命令的灵活性使得它成为 Shell 脚本中不可或缺的工具，希望你能在实践中灵活运用它。