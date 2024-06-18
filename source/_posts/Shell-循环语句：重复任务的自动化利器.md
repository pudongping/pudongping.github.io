---
title: Shell 循环语句：重复任务的自动化利器
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
abbrlink: 33c4223c
date: 2024-06-18 14:28:06
img:
coverImg:
password:
summary:
---

在日复一日的脚本编程中，循环语句无疑是我们最好的朋友。通过循环，我们可以执行重复的任务，无论是遍历文件列表，处理文本数据，还是简单的数学运算。

今天，我们就来聊聊 shell 脚本中的几种循环语句，它们将如何帮助我们简化编程任务。

## while 循环：当条件满足时循环

`while` 循环非常有用，基本语法是当条件为真（即返回值为 0）时，就执行循环体内的语句。使用它可以执行诸如从 1 加到 100 这样简单但有趣的任务。

**例子：实现 1 到 100 的求和**

```bash
#!/bin/bash

i=1
sum=0

while ((i <= 100))
do
  ((sum += i))
  ((i++))
done

echo "The sum is ===> $sum"
# 输出：The sum is ===> 5050
```

**例子：实现一个简单的加法计算器**

```bash
#!/bin/bash

sum=0

echo '输入数字，进行加法计算（按住 Ctrl + D 组合键获取结果）'

while read n
do
  ((sum += n))
done

echo "The sum is ====> $sum"
```

在这两个例子中，我们可以看到 `while` 循环如何在满足条件的情况下反复执行，直到条件不再成立。特别是在第二个例子中，我们利用了 `read` 命令读取用户输入，这在脚本交互中非常常见。

## for-in 循环：遍历列表元素

`for-in` 循环的用法与 Python 中的非常相似，用于遍历列表中的每个元素。这种方式编写的代码可读性强，易于理解。

**直接给出具体的值作为列表**

```bash
#!/bin/bash

sum=0

for n in 1 2 3 4 5 6
do
  echo $n
  ((sum += n))
done

echo 'The sum is '$sum
# 输出分别为 1 2 3 4 5 6 和 The sum is 21
```

**给出一个取值范围作为列表**

```bash
#!/bin/bash

sum=0
for n in {1..100}
do
  ((sum += n))
done

echo "The result is ${sum}"
# 输出：The result is 5050
```

**使用命令的执行结果作为列表**

```bash
#!/bin/bash

for filename in $(ls *.sh)
do
  echo "当前目录下所有的以 .sh 为后缀的文件为 ==> ${filename}"
done
```

在这些示例中，我们看到 `for-in` 循环如何通过直接列出元素、指定范囹、甚至使用命令的输出作为列表来实现强大的遍历功能。

## select-in 循环：增强脚本交互性

`select-in` 循环是脚本中用于交互的强大工具，它会显示一个带编号的菜单，用户通过输入编号来进行选择，进而执行不同的功能。

```bash
#!/bin/bash

echo '你喜欢哪种运动？'

select sport in '足球' '篮球' '乒乓球' '看电视'
do
  echo "你选择了 $sport"
  break # 加上 break 退出循环
done
```

通过 `select-in` 循环，我们可以轻松地构建用户友好的菜单系统，使得脚本的交互性大大增强。

在编写 shell 脚本时，正确选择循环类型对于提高代码的可读性和可维护性非常重要。

通过这篇文章的介绍，相信你已经对三种不同的循环有了初步的了解，并能够在实际编程中灵活应用它们。接下来，就是在你自己的脚本项目中实践和探索的时候了！