---
title: Redis 数据备份与恢复
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
tags:
  - Redis
  - 缓存
abbrlink: c14e496
date: 2022-12-11 21:41:58
img:
coverImg:
password:
summary:
categories: Redis
---


# Redis 数据备份与恢复

## 1. 命令行执行 save 手动开启 RDB 持久化

使用 RDB 文件做迁移时，需要注意需要先关闭掉目标 redis 的 aof 功能，因为如果二者同时存在的话，会优先于 aof 的方式进行数据恢复。

```
redis-cli -h {target-host} -a {target-password} config set appendonly no
```

### 备份

```shell
# 会阻塞主进程
127.0.0.1:6379> save
OK

# 通过 fork 一个专门 save 的子进程，从而不会阻塞主进程
127.0.0.1:6379> bgsave
Background saving started
```

### 恢复

```
# 查看 redis 默认存放备份文件的目录路径
127.0.0.1:6379> config get dir

# 查看备份 RDB 文件的名称，默认为 `dump.rdb`
127.0.0.1:6379> config get dbfilename
```

将备份之后的 `dump.rdb` 文件放到 `config get dir` 命令得出的目录路径下，然后重启 redis 即可恢复。（建议备份的时候，可以将 redis 暂时关闭）

## 2. 通过命令行手动开启 AOF 持久化

### 备份

```
# 先清空目标 redis 中全部数据
redis-cli -h {target-host} -a {target-password} flushall
# 然后在源 redis 中生成 aof 备份文件
redis-cli -h {source-host} -a {source-password} config set appendonly yes

# 查看生成后的 appendonly.aof 文件所在目录
127.0.0.1:6379> config get dir
# 查看备份的 aof 文件的名称，默认为 `appendonly.aof`
127.0.0.1:6379> config get appendfilename
```

### 恢复

```
# 将 `appendonly.aof` 文件放在当前路径下
redis-cli -h {target-host} -a {target-password} --pipe < appendonly.aof
# 源 redis 关闭 aof 功能
redis-cli -h {source-host} -a {source-password} config set appendonly no
```

将备份之后的 `appendonly.aof` 文件放到 `config get dir` 命令得出的目录路径下，然后重启 redis 也应该可恢复（具体我没有实操，看资料所说如此）。


## 3. 使用 [redis-dump](https://github.com/delano/redis-dump)

redis-dump 是一个用于 redis 数据导入导出的工具（可以以新增的形式导入），是基于 Ruby 实现的，因此需要先安装 Ruby 环境，建议安装 2.6.1 版本以上的 Ruby。

### MAC 上使用 Homebrew 安装 Ruby

```
brew install ruby
```

### 使用 [rvm](https://rvm.io/) （ruby 版本管理器）安装 ruby

#### centos 7 上安装 rvm

```bash

curl -sSL https://rvm.io/mpapis.asc | gpg2 --import -
curl -sSL https://rvm.io/pkuczynski.asc | gpg2 --import -

# 安装成功之后退出终端，然后可以通过 `rvm help` 进行查看
\curl -sSL https://get.rvm.io | bash -s stable

# 如果不想退出终端，可以直接重载配置文件
source /etc/profile.d/rvm.sh

```

##### rvm 常用命令

```bash

# 列出已知的 ruby 版本
rvm list known

# 安装指定版本的 ruby
rvm install 2.3.0

# 更新 rvm
rvm get stable

# 切换到指定 ruby 版本
rvm use 2.2.1

# 设置指定 ruby 版本为默认版本
rvm use 2.2.2 --default

# 查询已经安装的 ruby 版本
rvm list

# 卸载指定的 ruby 版本
rvm remove 1.9.1

```

### 安装 redis-dump

```bash

# 安装 ruby 2.6.1
rvm install 2.6.1

# 使用 2.6.1
rvm use 2.6.1
rvm use 2.6.1 --default

# 查看 ruby 版本
ruby --version

# 替换 gem 镜像地址
gem sources --add https://gems.ruby-china.org/ --remove https://rubygems.org/

# 查看镜像地址是否更换成功
gem sources -l

# 安装 redis-dump
gem install redis-dump -V

```

### redis-dump 备份与恢复

> 以增量的形式恢复

#### 备份

```
# 数据导出
redis-dump -u 127.0.0.1:6379 > data.json

# 导出指定数据库中的数据，比如说 8 号数据库
redis-dump -u 127.0.0.1:6379 -d 8 > data8.json

# 如果 redis 设置了有密码
redis-dump -u {host} -a {password} > data.json
redis-dump -u :{password}@127.0.0.1:6379 > data.json

# 如果需要导出的 redis 是一个 URL 连接地址时，貌似可以这样（没有实操过，具体不清楚）
redis-dump -u :{password}@{domain}:{port}
# eg: redis-dump -u :123456@www.alex.com:9055

```

#### 恢复

```
# 导入命令
cat data.json | redis-load
# 或者
< data.json redis-load

# 导入数据到 8 号数据库
cat data8.json | redis-load -u 127.0.0.1:6379 -a 123456 -d 8
# 或者
< data8.json redis-load -u 127.0.0.1:6379 -a 123456 -d 8
# 如果以上命令是因为 utf-8 格式报错时，可以加上 `-n` 参数
cat data8.json | redis-load -n -u 127.0.0.1:6379 -a 123456 -d 8

# 如果 redis 设置了有密码
cat data.json | redis-load -u :password@127.0.0.1:6379
```

## 4. 通过脚本实现迁移

```bash

#!/bin/bash
# 通过这个脚本执行备份有两个缺点，一是使用了 `keys *` ，二是那就是会将 source_host 中所有的 key
# 同步到 target_host 时，会自动改成永不过期
# 详见 `restore` 命令
# 后期有时间了，再考虑优化的事情吧
# document link: https://www.redis.com.cn/commands/restore.html

source_host=127.0.0.1
source_port=6379
source_password=''
source_db=1
target_host=127.0.0.1
target_port=6379
target_password=''
target_db=2


if [[ -z $source_password ]] && [[ -z $target_password ]]
then
  redis-cli -h ${source_host} -p ${source_port} -n ${source_db} keys '*' | while read key
  do
      redis-cli -h ${source_host} -p ${source_port} -n ${source_db} --raw dump $key | perl -pe 'chomp if eof' | redis-cli -h ${target_host} -p ${target_port} -n ${target_db} -x restore $key 0
      echo "migrate key $key"
  done
elif [[ -z $source_password ]] && [[ -n $target_password ]]
then
  redis-cli -h ${source_host} -p ${source_port} -n ${source_db} keys '*' | while read key
  do
      redis-cli -h ${source_host} -p ${source_port} -n ${source_db} --raw dump $key | perl -pe 'chomp if eof' | redis-cli -h ${target_host} -p ${target_port} -a ${target_password} -n ${target_db} -x restore $key 0
      echo "migrate key $key"
  done
elif [[ -n $source_password ]] && [[ -z $target_password ]]
then
  redis-cli -h ${source_host} -p ${source_port} -a ${source_password} -n ${source_db} keys '*' | while read key
  do
      redis-cli -h ${source_host} -p ${source_port} -a ${source_password} -n ${source_db} --raw dump $key | perl -pe 'chomp if eof' | redis-cli -h ${target_host} -p ${target_port} -n ${target_db} -x restore $key 0
      echo "migrate key $key"
  done
else
  redis-cli -h ${source_host} -p ${source_port} -a ${source_password} -n ${source_db} keys '*' | while read key
  do
      redis-cli -h ${source_host} -p ${source_port} -a ${source_password} -n ${source_db} --raw dump $key | perl -pe 'chomp if eof' | redis-cli -h ${target_host} -p ${target_port} -a ${target_password} -n ${target_db} -x restore $key 0
      echo "migrate key $key"
  done
fi

# 其实就是利用了 redis 的 dump 和 restore 命令
# eg：
# 127.0.0.1:6379[1]> set hello "hello, dumping world!"
# OK
# 127.0.0.1:6379[1]> dump hello
# "\x00\x15hello, dumping world!\t\x00\x03\xbfc\xcey\xa1\x9e\xfc"
# 127.0.0.1:6379[1]> restore hello1 0 "\x00\x15hello, dumping world!\t\x00\x03\xbfc\xcey\xa1\x9e\xfc"
# OK
# 127.0.0.1:6379[1]> get hello1
# "hello, dumping world!"
# 127.0.0.1:6379[1]>

```

## 5. redis 使用 migrate 命令迁移数据脚本

```bash

#!/bin/bash

src_redis="10.8.163.1"
src_port="6379"
dest_redis="10.8.132.13"
dest_port="6379"

for y in $(redis-cli -h ${src_redis} -p ${src_port} keys "*"); do
   redis-cli -h ${src_redis} -p ${src_port} migrate ${dest_redis} ${dest_port} ${y} 0 1000 copy replace
   echo "$(date +%F\ %T) Copy key ${y} to new redis...."
done

```
