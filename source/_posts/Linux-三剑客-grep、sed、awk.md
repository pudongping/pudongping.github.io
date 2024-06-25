---
title: Linux 三剑客 grep、sed、awk
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: 5ffd9f1c
date: 2024-06-25 11:04:58
img:
coverImg:
password:
summary:
---


在 Linux 的命令行世界里，有三个强大的文本处理工具：`grep`、`sed` 和 `awk`。它们被统称为 "Linux 三剑客"，它们各自拥有独特的功能，可以帮助我们高效地进行各种文本处理任务。让我们一一了解它们。

## grep

`grep` 的全称为 "global regular expression print"，意味着它可以通过正则表达式来搜索文本，并把匹配的行打印出来。

### grep 命令常用选项及含义

选项 | 含义
--- | ---
-c | 仅列出文件中包含匹配模式的行数
-i | 忽略模式中的字母大小写
-l | 列出包含匹配行的文件名
-n | 在每一行的最前面列出行号
-v | 列出不匹配模式的行
-w | 仅匹配整个单词，忽略部分匹配的行

### 示例

```shell
# 查找 demo.txt 文件中含有 "alex" 字符串的行
grep "alex" demo.txt

# 查找 demo.txt 文件中有多少行出现了 "alex" 字符串
grep -c "alex" demo.txt
# 假设输出是 50
```

## sed

`sed`（stream editor）是一种强大的流式文本编辑器。它可以根据脚本命令来处理文本文件中的数据。这些命令可以直接在命令行中输入，也可以存储在一个脚本文件中。

### sed 使用方式

1. 每次仅读取一行内容；
2. 根据提供的规则命令匹配并修改数据。注意，`sed` 默认不会直接修改源文件数据，而是会将数据复制到缓冲区中，修改也仅在缓冲区中进行；
3. 输出执行结果。

当一行数据匹配并处理完成后，`sed` 会继续读取下一行数据，并重复这个过程，直到将文件中所有数据处理完毕。

### 示例

#### 清空掉文件中所有的内容

```shell
# 不会直接修改源文件 demo.txt，而是在终端显示修改后的结果，即不显示任何内容，因为所有行都被删除了
sed 'd' demo.txt
```

## awk

`awk` 是一个用于文本分析的编程语言和工具。它非常擅长列出数据和报表，而且它的语法十分灵活，功能强大。

### awk 的基本使用

`awk` 通过对数据进行模式扫描和处理来达到文本处理的目的。它默认以空格为字段分隔符，将一行划分为多个字段。

### 示例

#### 输出文件的每一行的第二个字段

```shell
# 假设 demo.txt 是以空格分隔的字段的文本文件
awk '{print $2}' demo.txt
```

#### 分析日志文件并汇总信息

假设有一个日志文件 `access.log`，记录了网站的访问信息。我们想要统计出现次数最多的 IP 地址。

```shell
awk '{print $1}' access.log | sort | uniq -c | sort -nr | head -n 10
```

解释：
- `awk '{print $1}' access.log`：使用 awk 打印出日志中每行的第一个字段（一般是 IP 地址）。
- `sort`：对 IP 地址进行排序。
- `uniq -c`：压缩连续重复的行并计数。
- `sort -nr`：根据次数逆序排序。
- `head -n 10`：展示前 10 行。

grep、sed 和 awk 是 Linux 系统中文本处理的三大法宝。grep 用于搜索文本，sed 用于编辑文本，而 awk 则用于更复杂的文本分析和处理。

通过组合使用 `grep`、`sed` 和 `awk`，我们可以轻松地处理复杂的文本数据，有效提高我们的工作效率。

希望本文能帮助你入门并熟练掌握这些工具。