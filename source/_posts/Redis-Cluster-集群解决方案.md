---
title: Redis Cluster 集群解决方案
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Cache
  - Redis
  - 分布式集群
abbrlink: 663a091e
date: 2021-07-24 23:05:38
img:
coverImg:
password:
summary:
---

# Redis Cluster 集群解决方案

- 多个 Redis 实例协同进行
- 采用 slot （槽）分割数据，是 CRC16 与 16384 取模后分散
- 主从结构和选举算法，保证每个节点的可靠性
- 客户端可以连接任意一个 node 进行操作

![主从协同进行](https://upload-images.jianshu.io/upload_images/14623749-a6ad3bca1885917a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 所有的 redis 节点彼此互联（PING-PONG 机制），内部使用二进制协议优化传输速度和带宽。
- 节点的 fail 是通过集群中超过半数的节点检测失效时才生效。
- 客户端与 redis 节点直连，不需要中间 proxy 层，客户端不需要连接集群所有节点，连接集群中任何一个可用节点即可。
- redis-cluster 把所有的物理节点映射到 [0-16383] slot 上，cluster 负责维护 node <-> slot <-> value

![当有一台服务器出现故障时](https://upload-images.jianshu.io/upload_images/14623749-51f56d899218929f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## Redis Cluster 注意事项

- 不完全支持批量操作：mset、mget
- 事务不能跨节点支持
- 不支持多实例
- key 是最小粒度
- 最少 6 个才能保证组成完整高可用的集群
- 连接的时候只需要连接 1 台服务器即可。
- 如果 1 个主从连接宕机的话，那么集群就宕机了。


## Redis Cluster 配置步骤

***（建议使用官方安装包的方式安装 redis，不要使用 apt-get install 或者 yum 直接安装）***

1. 分别安装 **6 台** 服务器，三个主节点，三个从节点

我这里采用的是虚拟机，相应的 ip 地址分别为：
- 192.168.174.128  （28 号服务器）
- 192.168.174.129  （29 号服务器）
- 192.168.174.130  （30 号服务器）
- 192.168.174.131  （31 号服务器）
- 192.168.174.132  （32 号服务器）
- 192.168.174.133  （33 号服务器）

2. 配置 `redis.conf` 配置文件 ***（在所有的服务器上操作）***

vim /etc/redis/redis.conf

- 第一步：

```

# 默认为本地 ip 地址，需要改成当前服务器的 ip 地址，以便其他服务器可以正常访问
  69 bind 127.0.0.1 ::1
  
# 比如 28 号服务器更改为以下 ip 地址
bind 192.168.174.128

```

- 第二步：

```

# 这个参数的含义是指：禁止公网连接 redis 缓存，这样可以加强 redis 安全性。如果是在线上环境的话，我们不需要更改此参数值，然后需要进行设置账号密码，进行 auth 认证。这里为了测试方便，我们直接改成 no
  89 protected-mode yes
  
# 需要更改为以下
protected-mode no

```

- 第三步：开启集群相关参数

```

# 默认集群是关闭的
 815 # cluster-enabled yes

# 需要更改为以下（去除掉 # 号注释即可）
 815 cluster-enabled yes

```

- 第四步：开启集群配置文件

```
# 默认集群配置文件是关闭的
 823 # cluster-config-file nodes-6379.conf
 
# 需要更改为以下（去除掉 # 号注释即可） 
 823 cluster-config-file nodes-6379.conf 
 
```

- 第五步：开启集群超时时间

```
# 当一个节点出现问题的时候，最大超时连接时间为 15s ，当超过 15s 还没有连接的时候，就会认为该节点出现故障了，就会通过选举算法，将从服务器提升为主服务器。该参数默认是关闭的。
 829 # cluster-node-timeout 15000
 
# 需要更改为以下（去除掉 # 号注释即可）
 829 cluster-node-timeout 15000
 
```

3. 安装 `ruby` 组件。如果不安装这个软件，集群的时候，会报组件错误。你需要在那台服务器上面做集群，你就需要在哪台服务器上安装这个组件，并不是每台服务器上面都安装。这里采用第一台服务器做集群，因此在第一台服务器上安装 `ruby` 组件。***（在 28 号服务器上操作）***

```bash
sudo apt install ruby
```

4. 安装其他组件  ***（在 28 号服务器上操作）***

```bash
sudo gem install redis
```

5. 配置集群  ***（在 28 号服务器上操作）***

因为我是直接采用的 `apt-get install` 的方式安装的 `redis` 因此，服务器上面的 `redis` 工具 `redis-trib.rb` 在 `/usr/share/doc/redis-tools/examples` 目录下。如果你是通过安装包安装的 `redis` 那么请直接到 `redis` 解压目录中执行命令，如： `~/redis-4.0.9/src/redis-trib.rb`

```
/usr/share/doc/redis-tools/examples/redis-trib.rb create --replicas 1 192.168.174.128:6379 192.168.174.129:6379 192.168.174.130:6379 192.168.174.131:6379 192.168.174.132:6379 192.168.174.133:6379
```

![配置集群](https://upload-images.jianshu.io/upload_images/14623749-d756993ee1a78b0e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 如果出现如下报错时

![插槽节点被占用](https://upload-images.jianshu.io/upload_images/14623749-398cfd78326a863f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**问题原因：**
slot 插槽被占用了（这是搭建集群前时，以前 redis 的旧数据和配置信息没有清理干净。）  
**解决方案如下：**
用 redis-cli 登录到每个节点执行  flushall  和 cluster reset  就可以了

![解决方案如图所示](https://upload-images.jianshu.io/upload_images/14623749-460623a9e0c0186b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![再次设置集群，则可以连接成功](https://upload-images.jianshu.io/upload_images/14623749-f2ae4a85d4fd0304.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

6. 连接集群 ***（在任意一台服务器上操作）***

这里我挑选的是 ip 地址为 ： 192.168.174.131 的服务器，特别说明下，当执行 `keys` 命令的时候，只针对于当前服务器

```
# 带 -c 参数表示连接集群
redis-cli -h 192.168.174.131 -c
```

7. 测试

```
alex@alex-virtual-machine:~$ redis-cli -h 192.168.174.131 -c    # 连接的 31 号服务器
192.168.174.131:6379> keys *  # 这里的 keys 也只能查看所有在 31 号服务器上面的 keys
(empty list or set)
192.168.174.131:6379> set aa 111  # 随便设置一个 key
-> Redirected to slot [1180] located at 192.168.174.128:6379  # 数据却在 28 号服务器上被保存
OK
192.168.174.128:6379>   # 并且此时的状态直接跳到了 28 号服务器上面

```

![连接 31 号服务器，数据存到了 28 号服务器上](https://upload-images.jianshu.io/upload_images/14623749-1f41460e9f146d8a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```
alex@alex-virtual-machine:~$ redis-cli -h 192.168.174.129 -c  # 通过 29 号服务器连接集群
192.168.174.129:6379> get aa  # 从 29 号服务器中去取值
-> Redirected to slot [1180] located at 192.168.174.128:6379  # 会直接从 28 号服务器中返回值
"111"
192.168.174.128:6379>   # 并且此时的状态直接跳到了 28 号服务器上面

```

![从 30 号服务器上面取刚刚设置的值，会直接跳到 28 号服务器返回值](https://upload-images.jianshu.io/upload_images/14623749-56e8e278da0250a2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```
alex@alex-virtual-machine:~$ redis-cli -h 192.168.174.130 -c  # 通过 30 号服务器连接集群
192.168.174.130:6379> keys *  # 30 号服务器中并没有设置过 key，如果 30 号服务器可以取出值，证明可以跨服务器取出 keys，但是并没有数据，证明 keys 只能取出当前服务器中的 keys
(empty list or set)
192.168.174.130:6379> 

```

![keys 只能取出当前服务器中的所有 key](https://upload-images.jianshu.io/upload_images/14623749-a8098e86eba5fc31.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
