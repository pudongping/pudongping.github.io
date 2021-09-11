---
title: Linux 频繁使用命令
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: 6f6880af
date: 2021-09-12 02:47:08
img:
coverImg:
password:
summary:
---

# Linux 频繁使用命令

<!-- TOC -->

- [date](#date)
- [timedatectl](#timedatectl)
- [reboot](#reboot)
- [poweroff](#poweroff)
- [wget](#wget)
- [ps](#ps)
- [pstree](#pstree)
- [nice](#nice)
- [pidof](#pidof)
- [kill](#kill)
- [killall](#killall)
- [uname](#uname)
- [uptime](#uptime)
- [free](#free)
- [who](#who)
- [last](#last)
- [ping](#ping)
- [tracepath](#tracepath)
- [netstat](#netstat)
- [history](#history)
- [sosreport](#sosreport)
- [find](#find)
- [locate](#locate)
- [whereis](#whereis)
- [which](#which)
- [cat](#cat)
- [more](#more)
- [less](#less)
- [head](#head)
- [tail](#tail)
- [tr](#tr)
- [wc](#wc)
- [stat](#stat)
- [grep](#grep)
- [cut](#cut)
- [diff](#diff)
- [uniq](#uniq)
- [sort](#sort)
- [touch](#touch)
- [cp](#cp)
- [rm](#rm)
- [dd](#dd)
- [file](#file)

<!-- /TOC -->

## date

参数 |	作用
--- | ---
%S	| 秒（00～59）
%M	| 分钟（00～59）
%H	| 小时（00～23）
%I	| 小时（00～12）
%m	| 月份（1~12）
%p	| 显示出AM或PM
%a	| 缩写的工作日名称（例如：Sun）
%A	| 完整的工作日名称（例如：Sunday）
%b	| 缩写的月份名称（例如：Jan）
%B	| 完整的月份名称（例如：January）
%q	| 季度（1~4）
%y	| 简写年份（例如：20）
%Y	| 完整年份（例如：2020）
%d	| 本月中的第几天
%j	| 今年中的第几天
%n	| 换行符（相当于按下回车键）
%t	| 跳格（相当于按下Tab键）

```shell

# 查看当前系统时间
date
# Sat Sep 5 09:13:45 CST 2020

# 按照“年-月-日 小时:分钟:秒”的格式查看当前系统时间的 date
date "+%Y-%m-%d %H:%M:%S"
# 2020-09-05 09:14:35

# 将系统的当前时间设置为 2020 年 11 月 1 日 8 点 30 分
date -s "20201101 8:30:00"
# Sun Nov 1 08:30:00 CST 2020
date
# Sun Nov 1 08:30:08 CST 2020

# date 命令中的参数 %j 可用来查看今天是当年中的第几天。这个参数能够很好地区分备份时间的早晚，即数字越大，越靠近当前时间。
date "+%j"
# 305

```

## <span id="timedatectl">timedatectl</span>

> time date control

timedatectl 命令用于设置系统的时间，发现电脑时间跟实际时间不符？如果只差几分钟的话，我们可以直接调整。但是，如果差几个小时，那么除了调整当前的时间，还有必要检查一下时区了。

参数 |	作用
--- | ---
status	| 显示状态信息
list-timezones	| 列出已知时区
set-time	| 设置系统时间
set-timezone	| 设置生效时区

```shell

# 修改时区
timedatectl set-timezone Asia/Shanghai

# 修改系统日期
timedatectl set-time 2021-05-18

# 修改系统时间
timedatectl set-time 9:30

```

## reboot

重启系统

## poweroff

关闭系统

## wget

> web get

wget 命令用于在终端命令行中下载网络文件

参数 |	作用
--- | ---
-b	| 后台下载模式
-P	| 下载到指定目录
-t	| 最大尝试次数
-c	| 断点续传
-p	| 下载页面内所有资源，包括图片、视频等
-r	| 递归下载

```shell

# 递归下载 www.linuxprobe.com 网站内的所有页面数据以及文件，下载完后会自动保存到当前路径下一个名为 www.linuxprobe.com 的目录中
wget -r -p https://www.linuxprobe.com

```

## ps

> processes

用于查看系统中的进程状态

参数 |	作用
--- | ---
-a	| 显示所有进程（包括其他用户的进程）
-u	| 用户以及其他详细信息
-x	| 显示没有控制终端的进程

## pstree

> process tree

用于以树状图的形式展示进程之间的关系

## nice

用于调整进程的优先级，语法格式为“nice优先级数字 服务名称”

在 top 命令输出的结果中，PR 和 NI 值代表的是进程的优先级，数字越低（取值范围是-20～19），优先级越高。在日常的生产工作中，可以将一些不重要进程的优先级调低，让紧迫的服务更多地利用 CPU 和内存资源，以达到合理分配系统资源的目的。

```shell

# 将 bash 服务的优先级调整到最高
nice -n -20 bash

```

## pidof

用于查询某个指定服务进程的 PID 号码值，语法格式为 “pidof [参数] 服务名称”

```shell

# 每个进程的进程号码值（PID）是唯一的，可以用于区分不同的进程。例如，执行如下命令来查询本机上 sshd 服务程序的 PID：
pidof sshd
# 2156

```

## kill

```shell

# 有时系统会提示进程无法被终止，此时可以加参数 -9，表示最高级别地强制杀死进程
kill -9 2156

```

## killall

用于终止某个指定名称的服务所对应的全部进程，语法格式为 “killall [参数] 服务名称”

```shell

pidof httpd
# 13581 13580 13579 13578 13577 13576

killall httpd

# 再次查看
pidof httpd

```

## uname

uname 命令用于查看系统内核版本与系统架构等信息，英文全称为 “unix name”，语法格式为 “uname [-a]” 。在使用 uname 命令时，一般要固定搭配上 -a 参数来完整地查看当前系统的内核名称、主机名、内核发行版本、节点名、压制时间、硬件名称、硬件平台、处理器类型以及操作系统名称等信息。

```shell

uname -a

```

如果要查看当前系统版本的详细信息，则需要查看 redhat-release 文件

```shell

cat /etc/redhat-release

```

## uptime

> 建议负载值保持在 1 左右，在生产环境中不要超过 5 就好。

uptime 命令用于查看系统的负载信息。它可以显示当前系统时间、系统已运行时间、启用终端数量以及平均负载值等信息。平均负载值指的是系统在最近 1分钟、5分钟、15分钟内的压力情况，负载值越低越好。

```shell

uptime
#  0:16  up 5 days, 14:47, 8 users, load averages: 1.41 1.44 1.56

```

## free

free 命令用于显示当前系统中内存的使用量信息

```shell

free -h

```

## who

who 命令用于查看当前登入主机的用户终端信息。

```shell

who

# 登录的用户名 终端设备 登录到系统的时间
# pudongping console  Sep  6 09:29

```

## last

last 命令用于调取主机的被访记录，Linux 系统会将每次的登录信息都记录到日志文件中，如果哪天想翻阅了，直接执行这条命令就行。

## ping

ping 命令用于测试主机之间的网络连通性。

参数 |	作用
--- | ---
-c	| 总共发送次数
-l	| 指定网卡名称
-i	| 每次间隔时间（秒）
-W	| 最长等待时间（秒）

```shell

ping -c 4 www.baidu.com

```

## tracepath

tracepath 命令用于显示数据包到达目的主机时途中经过的所有路由信息。当两台主机之间无法正常 ping 通时，要考虑两台主机之间是否有错误的路由信息，导致数据被某一台设备错误地丢弃。这时便可以使用 tracepath 命令追踪数据包到达目的主机时途中的所有路由信息，以分析是哪台设备出了问题。

```shell

tracepath www.baidu.com

```

## netstat

netstat 命令用于显示如网络连接、路由表、接口状态等的网络相关信息，英文全称为 “network status”

参数 | 作用
--- | ---
-a	| 显示所有连接中的Socket
-p	| 显示正在使用的Socket信息
-t	| 显示TCP协议的连接状态
-u	| 显示UDP协议的连接状态
-n	| 使用IP地址，不使用域名
-l	| 仅列出正在监听的服务状态
-i	| 现在网卡列表信息
-r	| 显示路由表信息


## history

默认会显示出用户在本地计算机中执行过的最近 1000 条命令记录，可以通过修改 `/etc/profile` 文件中的 `HISTSIZE` 变量值来改变默认数量。历史命令都会被记录到 `~/.bash_history` 文件中。

```shell

# 清空命令历史记录
history -c

```

## sosreport

sosreport 命令用于收集系统配置及架构信息并输出诊断文档。

## find

参数 |	作用
--- | ---
-name	| 匹配名称
-perm	| 匹配权限（mode为完全匹配，-mode为包含即可）
-user	| 匹配所有者
-group	| 匹配所有组
-mtime -n +n	| 匹配修改内容的时间（-n指n天以内，+n指n天以前）
-atime -n +n	| 匹配访问文件的时间（-n指n天以内，+n指n天以前）
-ctime -n +n	| 匹配修改文件权限的时间（-n指n天以内，+n指n天以前）
-nouser	| 匹配无所有者的文件
-nogroup	| 匹配无所有组的文件
-newer f1 !f2	| 匹配比文件f1新但比f2旧的文件
--type b/d/c/p/l/f	| 匹配文件类型（后面的字幕字母依次表示块设备、目录、字符设备、管道、链接文件、文本文件）
-size	| 匹配文件的大小（+50KB为查找超过50KB的文件，而-50KB为查找小于50KB的文件）
-prune	| 忽略某个目录
-exec …… {}\;	| 后面可跟用于进一步处理搜索结果的命令

```shell

# 获取 /etc 目录中所有以 host 开头的文件列表
find /etc -name "host*" -print

# 搜索整个系统中包含 SUID 权限的所有文件
find / -perm -4000 -print

# 在整个系统中找出所有归属于 alex 用户的文件，并复制到 /home/alex/findresults 目录中
find / -user alex -exec cp -a {} /home/alex/findresults/\;
# 其中的 `{}` 表示 find 命令搜索出的每一个文件，并且命令的结尾必须是 `\;`

```

## locate

locate 命令用于按照名称快速搜索文件所对应的位置。使用 find 命令进行全盘搜索虽然更准确，但是效率有点低。如果仅仅是想找一些常见的且又知道大概名称的文件，可以使用 locate 命令。在使用 locate 命令时，先使用 updatedb 命令生成一个索引库文件，这个库文件的名字是 `/var/lib/mlocate/mlocate.db` ，后续在使用 locate 命令搜索文件时就是在该库中进行查找操作，速度会快很多。

```shell

updatedb
ls -l /var/lib/mlocate/mlocate.db

# 使用 locate 命令搜索出所有包含 alex 名称的文件所在的位置
locate alex

```

## whereis

whereis 命令用于按照名称快速搜索二进制程序（命令）、源代码以及帮助文件所对应的位置，语法格式为 “whereis 命令名称”。

简单来说，whereis 命令也是基于 updatedb 命令所生成的索引库文件进行搜索，它与 locate 命令的区别是不关心那些相同名称的文件，仅仅是快速找到对应的命令文件及其帮助文件所在的位置。

## which

which 命令用于按照指定名称快速搜索二进制程序（命令）所对应的位置，语法格式为 “which 命令名称”。

which 命令是在 PATH 变量所指定的路径中，按照指定条件搜索命令所在的路径。也就是说，如果我们既不关心同名文件（find 与 locate），也不关心命令所对应的源代码和帮助文件（whereis），仅仅是想找到命令本身所在的路径，那么这个 which 命令就太合适了。

## cat

> cat 是 concatenate（连接、连续）的简写

- 连接合并文件

```shell

echo 11 > a.txt
echo 22 > b.txt

# 合并 a.txt 和 b.txt 文件中所有的内容到 c.txt 文件中
cat a.txt b.txt > c.txt

# 查看 c.txt 文件中的内容
cat c.txt  # 11 22

# -n 参数可以显示行号
cat -n c.txt
# 1 11
# 2 22

```

## more

## less

选项 | 含义
--- | ---
-N | 显示每行的行号
-S | 行过长时将超出部分舍弃
-m | 显示类似 more 命令的百分比


less 进入交互界面时的指令及功能

交互指令 | 功能
--- | ---
/字符串 |	向下搜索“字符串”的功能
?字符串	| 向上搜索“字符串”的功能
n	| 重复*前一个搜索
N	| 反向重复前一个搜索
k	| 向上移动一行
j	| 向下移动一行
u	| 向上移动半页
d	| 向下移动半页
b	| 向上移动一页
空格键	| 向下移动一页
g	| 移动到第一行
G	| 移动至最后一行
h 或 H	| 显示帮助界面
q 或 Q 或 ZZ	| 退出 less 命令
v	| 使用配置的编辑器编辑当前文件

## head

```shell

# 只想查看文本中前 10 行的内容
head -n 10 test.log

```

## tail

```shell

# 只想查看文本内容的最后 10 行
tail -n 10 test.log

# 实时查看文本
tail -f test.log

```

## tr

tr 命令用于替换文本内容中的字符，英文全称为 “translate”，语法格式为 “tr [原始字符] [目标字符]”

```shell

# 把文本内容中的英文全部替换成大写
cat test.log | tr [a-z] [A-Z]

```

## wc

wc 命令用于统计指定文本文件的行数、字数或字节数，英文全称为 “word counts”，语法格式为 “wc [参数] 文件名称”

参数 |	作用
--- | ---
-l	| 只显示行数
-w	| 只显示单词数
-c	| 只显示字节数

```

# 统计当前系统中有多少个用户
wc -l /etc/passwd

```

## stat

stat 命令用于查看文件的具体存储细节和时间等信息，英文全称为 “status”，语法格式为 “stat 文件名称”。
除了修改时间之外，Linux 系统中的文件包含 3 种时间状态，分别是 Access Time（内容最后一次被访问的时间，简称为 **Atime**），Modify Time（内容最后一次被修改的时间，简称为 **Mtime**）以及Change Time（文件属性最后一次被修改的时间，简称为 **Ctime**）。

```shell

stat test.log

```

## grep

grep 命令用于按行提取文本内容，语法格式为 “grep [参数] 文件名称”

参数|	作用
--- | ---
-b	| 将可执行文件(binary)当作文本文件（text）来搜索
-c	| 仅显示找到的行数
-i	| 忽略大小写
-n	| 显示行号
-v	| 反向选择——仅列出没有“关键词”的行。

```shell

# 找出当前系统中不允许登录系统的所有用户信息
grep /sbin/nologin /etc/passwd

```

## cut

cut 命令用于按 **“列”** 提取文本内容，语法格式为 “cut [参数] 文件名称”，按“列”搜索，不仅要使用 `-f` 参数设置需要查看的列数，还需要使用 `-d` 参数来设置间隔符号。

```shell

# 提取 /etc/passwd 文件中的用户名信息，即提取以冒号为间隔符号的第一列内容
cut -d : -f 1 /etc/passwd
# 或者（只提取前 5 个用户信息）
cat /etc/passwd | grep -v '#' | head -n 5 | cut -d : -f 1

```

## diff

diff 命令用于比较多个文件之间内容的差异，英文全称为 “different”，语法格式为 “diff [参数] 文件名称A 文件名称B”。

在使用 diff 命令时，不仅可以使用 --brief 参数来确认两个文件是否相同，还可以使用 -c 参数来详细比较出多个文件的差异之处。

```shell

# 判断两个文件是否相同
diff --brief a.txt b.txt

# 查看文件内容的具体不同
diff -c a.txt b.txt

```

## uniq

uniq 命令用于去除文本中连续的重复行，英文全称为 “unique”，语法格式为“uniq [参数] 文件名称”。 **中间不能夹杂其他文本行（非相邻的默认不会去重）**

```shell

uniq alex.txt

```

## sort

sort 命令用于对文本内容进行再排序，语法格式为 “sort [参数] 文件名称”。

参数	| 作用
--- | ---
-f	| 忽略大小写
-b	| 忽略缩进与空格
-n	| 以数值型排序
-r	| 反向排序
-u	| 去除重复行
-t	| 指定间隔符
-k	| 设置字段范围

```shell

# 将文本中的内容按照字母顺序进行排序
sort alex.txt

# 将文本内容进行去重操作
sort -u alex.txt

# 以第 3 个字段中的数字作为排序依据，以冒号 : 指定分隔符，以数字规则进行排序
sort -t : -k 3 -n alex.txt

```

## touch

touch 命令用于创建空白文件或设置文件的时间，语法格式为 “touch [参数] 文件名称”。

参数	| 作用
--- | ---
-a	| 仅修改“读取时间”（atime）
-m	| 仅修改“修改时间”（mtime）
-d	| 同时修改 atime 与 mtime


```shell

# 修改文件的读取时间和修改时间
touch -d "2021-09-12 15:44" alex.txt

```

## cp

cp 命令用于复制文件或目录，英文全称为 “copy”，语法格式为 “cp [参数] 源文件名称 目标文件名称”。

参数 |	作用
--- | ---
-p	| 保留原始文件的属性
-d	| 若对象为“链接文件”，则保留该“链接文件”的属性
-r	| 递归持续复制（用于目录）
-i	| 若目标文件存在则询问是否覆盖
-a	| 相当于 -pdr（p、d、r为上述参数）

## rm

参数	| 作用
--- | ---
-f	| 强制执行
-i	| 删除前询问
-r	| 删除目录
-v	| 显示过程

## dd

dd 命令用于按照指定大小和个数的数据块来复制文件或转换文件，语法格式为 “dd if=参数值of=参数值count=参数值bs=参数值”。

参数	| 作用
--- | ---
if	| 输入的文件名称
of	| 输出的文件名称
bs	| 设置每个“块”的大小
count	| 设置要复制“块”的个数

```shell

# 从 /dev/zero 设备文件中取出 1 个大小为 560M 的数据块，保存成名为 560_file 的文件
dd if=/dev/zero of=560_file count=1 bs=560M

# 将光驱设备中的光盘制作成 iso 格式的镜像文件
dd if=/dev/cdrom of=RHEL.iso

```

## file

file 命令用于查看文件的类型，语法格式为 “file 文件名称”。
