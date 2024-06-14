---
title: Shell 字符串操作：让你的脚本语言更加灵活
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
abbrlink: 68372f77
date: 2024-06-14 18:51:44
img:
coverImg:
password:
summary:
---

字符串在编程中扮演着至关重要的角色，尤其是在 Shell 脚本编程中。无论是处理文件路径、生成动态消息还是执行复杂的文本分析，掌握字符串操作无疑会让你的脚本更加强大且灵活。

今天，我们就来深入了解一些基础而且实用的 Shell 字符串操作技巧，无论你是编程新手还是有一定经验的开发者，掌握字符串操作总能在编写脚本时让你事半功倍。

## 获取字符串长度

有时候，你可能想知道一个字符串有多长，比如在校验用户名或者是切割字符串时。Shell 通过一个简单的表达式就能帮助我们得到答案。

```bash
#!/bin/bash

name=alex
echo ${#name}  # 输出：4
```

通过 `${#变量名}` 的语法，就可以快速获取字符串的长度。这种方式简洁而高效，对于各种字符串操作场景非常实用。

## 字符串连接合合并

在日常的脚本编写中，字符串的连接操作是避不开的。无论是拼接路径，还是生成含有变量的提示信息，字符串连接都扮演着重要的角色。

```bash
#!/bin/bash

name=alex
age=26

str1=$name$begin
str2="$name $age" # 注意，当字符串中包含空格时，最好用双引号包裹起来
str3=$name"=====>$age"
str4="$name =====> $age"
str5="${name}is a web artisan and the age is${age}"

echo $str1  # 输出：alex26
echo $str2  # 输出：alex 26
echo $str3  # 输出：alex=====>26
echo $str4  # 输出：alex =====> 26
echo $str5  # 输出：alexis a web artisan and the age is26
```

连接字符串时，直接使用 `$变量名` 或者是 `${变量名}` 完成。特别是在字符串和变量混合使用的场景下，使用大括号可以帮助明确变量的边界，避免解析上的混乱。

## 字符串截取

字符串截取是另一个非常实用的功能，它允许我们从一个字符串中提取出我们需要的某一部分。这在处理路径、文件名或者是日志分析等场景中特别有用。

```bash
#!/bin/bash

name=alex
age=26

str="hello, my name is ${name}, my age is ${age}. nice to meet you!"

# 从字符串左边开始计数
echo ${str: 4: 6}  # o, my

# 直接截取到字符串的末尾
echo ${str: 4}  # o, my name is alex, my age is 26. nice to meet you!

# 从字符串右边开始计数
echo ${str: 0-11: 8}  # o meet y

# 直接截取到字符串的末尾
echo ${str: 0-11}  # o meet you!

# 使用 # 号截取右边所有字符
echo ${str#*my}  # name is alex, my age is 26. nice to meet you!
echo ${str#*y}  # name is alex, my age is 26. nice to meet you!
# 如果不需要忽略子字符串左边的字符，那么也可以不写 * 号
echo ${str#my}  # hello, my name is alex, my age is 26. nice to meet you!
# 如果希望直到最后一个子字符串才匹配结束，那么可以使用两个 # 号
echo ${str##*my}  # age is 26. nice to meet you!

# 使用 % 号截取左边所有字符
echo ${str%my*}  # hello, my name is alex,
echo ${str%y*}  # hello, my name is alex, my age is 26. nice to meet
echo ${str%}  # hello, my name is alex, my age is 26. nice to meet you!
echo ${str%%my*}  # hello,
```

这些截取方法基于不同的需求，为我们提供了极大的灵活性，但同时也需要我们有足够的练习来熟练掌握。

更多的用法可以对照这张表进行查看：

格式 | 含义
--- | ---
${string: start :length} |	从 string 字符串的左边第 start 个字符开始，向右截取 length 个字符
${string: start} |	从 string 字符串的左边第 start 个字符开始截取，直到最后
${string: 0-start :length}	| 从 string 字符串的右边第 start 个字符开始，向右截取 length 个字符
${string: 0-start}	| 从 string 字符串的右边第 start 个字符开始截取，直到最后
${string#*chars}	| 从 string 字符串第一次出现 *chars 的位置开始，截取 *chars 右边的所有字符
${string##*chars}	| 从 string 字符串最后一次出现 *chars 的位置开始，截取 *chars 右边的所有字符
${string%*chars}	| 从 string 字符串第一次出现 *chars 的位置开始，截取 *chars 左边的所有字符
${string%%*chars}	| 从 string 字符串最后一次出现 *chars 的位置开始，截取 *chars 左边的所有字符

通过以上的介绍，我希望能帮助你了解并掌握 Shell 中的基础字符串操作。

记住，实践是学习的捷径。我鼓励你自行编写脚本，尝试不同的字符串操作，这样你才能更加熟悉并灵活运用它们。