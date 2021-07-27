---
title: MySQL 主从复制
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: MySQL
tags:
  - MySQL
  - 主从复制
abbrlink: 6b30c31b
date: 2021-07-27 23:08:10
img:
coverImg:
password:
summary:
---


## MySQL 主从复制

- 主从复制原理

![主从复制原理](https://upload-images.jianshu.io/upload_images/14623749-e70a71423894062b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 主从复制的基本原则
1.  每个 slave 只有一个 master
2.  每个 slave 只能有一个唯一的服务器 ID
3.  每个 master 可以有多个 salve

### 一主一从常见配置
- mysql 版本一致且后台以服务运行
- 主从都配置在 [mysqld] 节点下，都是小写

#### 主数据库配置，修改 /etc/my.cnf 配置文件
1. 主服务器唯一 ID **必须**

```
server-id=1
```

2. 启用二进制日志 **必须**

```
# mysqlbin 为官方推荐的文件名
log-bin=自己的本地路径/mysqlbin

log-bin=/var/local/mysql-server5.7/data/mysqlbin
```

3. 启用错误日志 **可选**

```
# mysqlerr 为官方推荐的文件名
log-err=自己的本地路径/mysqlerr

log-err=/var/local/mysql-server5.7/data/mysqlerr
```

4. 根目录 **可选**

```
basedir="自己的本地路径"

basedir="/var/local/mysql-server5.7/"
```

5. 临时目录 **可选**

```
tmpdir="自己的本地路径"

tmpdir="/var/local/mysql-server5.7/"
```

6. 数据目录 **可选**

```
datadir="自己的本地路径"

datadir="/var/local/mysql-server5.7/"
```

7. 主机，读写都可以 **可选**

```
read-only=0
```

8. 设置不要复制的数据库 **可选**

```
binlog-ignore-db=mysql
```

9. 设置需要复制的数据库 **可选**

```
binlog-do-db=需要复制的主数据库名字
```



完整的配置为：
```
[mysqld]

port=3306

server-id=1
log-bin=/var/local/mysql-server5.7/data/mysqlbin
log-err=/var/local/mysql-server5.7/data/mysqlerr
basedir="/var/local/mysql-server5.7/"
tmpdir="/var/local/mysql-server5.7/"
datadir="/var/local/mysql-server5.7/"
read-only=0
binlog-ignore-db=mysql
```


#### 从数据库配置，修改 /etc/my.cnf 配置文件

1. 从服务器唯一 ID **必须**

```
# 默认配置文件中将此行注释掉了，可以直接取消注释即可，如果没有找到，也可以自己添加
server-id=2
```

2. 启用二进制日志 **必须**

```
# 默认配置文件中将此行注释掉了，可以直接取消注释即可，如果没有找到，也可以自己添加

log-bin=mysql-bin
```

- 主从机器都关闭掉防火墙
- 在主数据库上建立账户并授权 slave

```
# 从数据库用于拷贝的账号为：alex
# 从数据库用于拷贝的密码为：123456
# 从数据库所在服务器的 ip 地址为：127.0.0.12

grant replication slave on *.* to 'alex'@'127.0.0.12' identified by '123456';

# 刷新权限
flush privileges;

# 查询 master 的状态，并记录下 File 和 Position 的值
show master status;
```

- 在从数据库上配置

```
# 主数据库所在服务器的 ip 地址为：127.0.0.1
# 在主数据库建立用于从数据库拷贝的账号为：alex
# 在主数据库建立用于从数据库拷贝的密码为：123456
# 主数据库中查询出的 File 值为：mysqlbin.000035
# 主数据库中查询出的 Position 的值为：341

change master to master_host='127.0.0.1', master_user='alex',master_password='123456',master_log_file='mysqlbin.000035',master_log_pos=341;
 
# 启动从服务器复制功能
start slave;

# 执行以下命令，当 Slave_IO_Running: Yes 和 Slave_SQL_Running: Yes 同时为 yes 时，表示主从复制已经打通了
show slave status\G
```

- 测试主从复制是否成功？
    - 主数据库新建一个库、新建表、插入一条记录，从数据库去查询是否含有以上数据即可

- 如何停止从数据库复制功能？

```
# 从数据库上
stop slave;
```

- 当 Slave_IO_Running 和 Slave_SQL_Running 参数不同时为 yes 时？

```
# 先停止从数据库复制功能（从数据库上执行）
stop slave;

# 查询 master 的状态，并记录下 File 和 Position 的值 （主数据库上执行，这次会有新的 File 值和 Position 值）
show master status;

change master to master_host='主数据库所在服务器的 ip 地址', master_user='alex',master_password='123456',master_log_file='新的 File 值',master_log_pos=新的Position 值;
```
