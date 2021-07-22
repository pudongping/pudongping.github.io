---
title: 使用 Memcached 实现分布式算法
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Cache
tags:
  - Cache
  - Memcached
  - 缓存
abbrlink: aa516918
date: 2021-07-23 00:35:56
img:
coverImg:
password:
summary:
---

# 使用 Memcached 实现分布式算法

## 分布式算法

![需要通过应用程序来实现分布式算法](https://upload-images.jianshu.io/upload_images/14623749-e26290372bf4ec4b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 余数计算分散法

> 根据 key 来计算 CRC，然后结果对服务器数进行取模得到 memcached 服务器节点。  服务器无法连接的时候，将尝试的连接次数加到 key 后面重新计算。

![余数计算分散法](https://upload-images.jianshu.io/upload_images/14623749-6e626d05ecb0cb98.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 缺点：
添加或移除服务器时，几乎所有缓存要重建，还需要考虑雪崩式崩溃问题。

16 个 Key、5 个服务器

服务器 | key 值 | key 值 | key 值 | key 值
--- | --- | --- | --- | ---
S0 |  | 5 | 10 | 15
S1 | 1 | 6 | 11 | 16
S2 | 2 | 7 | 12 |
S3 | 3 | 8 | 13 |
S4 | 4 | 9 | 14 |


举个例子来说：如果 S3 服务器直接挂掉的话，那么 S3 服务器的流量会全部并给 S4 服务器，也就是说 S4 服务器要承受之前一倍的压力，如果此时 S4 服务器已经无法承受，直接崩掉的话，那么也就是会将原本 S3、S4 的流量全部在并给 S0 服务器

![余数计算分散法](https://upload-images.jianshu.io/upload_images/14623749-b39308b35ba20703.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


服务器 | key 值 | key 值 | key 值 | key 值 | key 值
--- | --- | --- | --- | --- | ---
S0 |  | 4 | 8 | 12 | 16
S1 | 1 | 5 | 9 | 13 |
S2 | 2 | 6 | 10 | 14 |
S3（假设移除了 S3 服务器） | | | | |
S4(3) | 3 | 7 | 11 | 15 |

以上例子说明：当移除一台服务器之后，基本上所有的缓存都需要重建，key 值，除了 S1 服务器的 1 和 S2 服务器的 2，剩下的全部都乱套了

![余数计算分散法](https://upload-images.jianshu.io/upload_images/14623749-d57efed81911e414.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 一致性哈希算法
- 求出服务器节点的哈希值分配到 0~2^32 的圆上
- 求出存储数据键的哈希值映射到圆上
- 从数据映射到的位置开始顺时针查找，将数据保存到找到的第一个服务器上

![一致性哈希算法](https://upload-images.jianshu.io/upload_images/14623749-e59269f6a1f321ef.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![一致性哈希算法](https://upload-images.jianshu.io/upload_images/14623749-27ffaa9971199a33.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


#### 优点：
- 冗余少
- 负载均衡
- 过渡平滑
- 存储均衡

#### 缺点：
依然无法解决雪崩问题
