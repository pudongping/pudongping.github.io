---
title: Redis 的持久化
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Redis
  - Cache
  - 缓存
abbrlink: 6d929828
date: 2021-07-25 15:04:56
img:
coverImg:
password:
summary:
---

# Redis 的持久化

## 两种持久化方式
- RDB 指定的时间间隔内保存数据快照
- AOF 先把命令追加到操作日志的尾部，保存所有历史操作命令

### RDB 模式
- **优点**
1. 适合用于进行备份
2. fork 出子进程进行备份，主进程没有任何 IO 操作
3. 恢复大数据集时的速度快

- **缺点**
1. 特定条件下进行一次持久化，易丢失数据
2. 庞大数据时，保存时会出现性能问题

设置方式

配置文件路径： sudo vim /etc/redis/redis.conf

```
# 备份的频次
save 900 1    # 900 秒内，有 1 次更新操作，就将数据同步到数据文件
save 300 10
save 60 10000

# 备份的文件名
 253 dbfilename dump.rdb

# 备份的目录路径
 263 dir /var/lib/redis

```

默认的备份文件为： vim /var/lib/redis/dump.rdb

### AOF 模式
- **优点**
1. 数据非常完整，故障恢复丢失数据少
2. 可对历史操作进行处理
- **缺点**
1. 文件的体积大
2. 速度低于 RDB 且故障恢复速度慢

设置方式

配置文件路径： sudo vim /etc/redis/redis.conf

```
# 当 appendonly 参数为 yes 时，则开启 AOF 模式
 672 appendonly yes

# 备份的文件名
 676 appendfilename "appendonly.aof"

# 同步的方式
 701 # appendfsync always  // 同步持久化，每次数据变更都会立刻保存到磁盘上，需要实时记录，因此效率不高，但是数据十分完整
 702 appendfsync everysec  // 异步持久化，每隔 1s 记录一次
 703 # appendfsync no  // 不同步

```

默认的备份文件为： vim /var/lib/redis/appendonly.aof

> 两种模式可以同时开启，同时开启的时候会优先执行 AOF 模式的备份文件，进行 AOF 模式恢复，同时开启的时候，需要注意在 redis 使用之初就要先开启 AOF 模式，以免 AOF 模式，只会记录部分命令，导致恢复数据不完整。


## 合理地使用 Redis

- 防止内存占满：
1. 设置超时时间
2. 不存放过大文件（最好不要超过 500 字节）
3. 不存放不常用数据

- 提高使用效率
1. 合理使用不同的数据结构类型
2. 慎用正则处理或者批量操作 Hash、Set 等。（因为 redis 是单线程，如果正则匹配 key 的话，可能会影响其他命令的使用）
