---
title: 如何在Redis中快速推算两地之间的距离？——Geo篇
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
  - 缓存
abbrlink: 9c96a9d2
date: 2024-07-18 18:22:22
img:
coverImg:
password:
summary:
---

处理地理位置数据已成为许多应用程序的核心需求。无论是推送附近的餐馆还是对全国范围内的服务点进行分析，快速而准确地处理和检索地理位置信息都至关重要。Redis，作为一种高性能的内存数据库，为我们提供了这样的解决方案。

> Redis 在 3.2 推出 Geo 类型，该功能可以推算出地理位置信息，两地之间的距离。有效的经度从 -180 度到 180 度。有效的纬度从 -85.05112878 度到 85.05112878 度，如果超过此范围，redis 会直接报错。

通过本文，我们将一步步探索 Redis 如何帮助我们处理地理位置数据，不仅适合初学者，也能让有经验的开发者有所收获。

## 添加地理位置数据
首先，我们需要向 Redis 中添加一些中国城市的地理位置数据：

> 你可以通过这个网站 `http://www.jsons.cn/lngcode/` 来查询一下一些城市的经纬度。

```bash
127.0.0.1:6379> geoadd china:city 116.40 39.90 beijing
(integer) 1
127.0.0.1:6379> geoadd china:city 121.47 31.23 shanghai
(integer) 1
127.0.0.1:6379> geoadd china:city 106.50 29.53 chongqing 114.05 22.52 shengzhen
(integer) 2
127.0.0.1:6379> geoadd china:city 120.16 30.24 hangzhou 108.96 34.26 xian
(integer) 2
```

这里，`geoadd` 命令用于向指定的 key（这里是 china:city）中添加地理空间位置信息。每条记录包括经度、纬度以及位置的名称。

你是否会好奇 geo 是通过什么类型在 Redis 中存储的？那我们不妨可以使用 type 命令来进行查看一下：

```bash
127.0.0.1:6379> type china:city
zset
```

我们可以看到，其实 **geo 底层实现原理就是 zset ** ，这表明 Redis 内部使用了有序集合存储地理空间信息，每个地点的名称是有序集合的成员，而其经纬度则用于计算分数，以确保成员的有序性。

## 查询地理空间信息

当信息存入数据库后，我们可以执行各种查询：

### 查询单个城市位置

```bash
127.0.0.1:6379> geopos china:city shanghai
1) 1) "121.47000163793563843"
   2) "31.22999903975783553"
```

geopos 命令用于获取一个或多个成员的地理位置信息（经度和纬度），这个命令返回上海的经纬度。

### 计算两城市间距离

单位：

- `m` 表示单位为米，也是默认单位。
- `km` 表示单位为千米。
- `mi` 表示单位为英里。
- `ft` 表示单位为英尺。

```bash
127.0.0.1:6379> geodist china:city shanghai chongqing
"1447673.6920"
```

geodist 命令用于计算两个位置之间的距离，默认单位是米。上面返回的是上海到重庆的距离。

### 查询指定范围内的城市

```bash
127.0.0.1:6379> georadius china:city 100 20 1000 km
(empty array)
127.0.0.1:6379> georadius china:city 110 30 1000 km
1) "chongqing"
2) "xian"
3) "shengzhen"
4) "hangzhou"
```
查看以（110,30）为中心，半径 1000 公里的范围内，有哪些城市。

## 带有选项的地理位置查询

Redis 地理空间查询还支持多种选项，例如，返回搜索结果的坐标和距离，或者限制返回结果的数量：

### 查询并返回坐标

```bash
127.0.0.1:6379> georadius china:city 110 30 1000 km withcoord
1) 1) "chongqing"
   2) 1) "106.49999767541885376"
      2) "29.52999957900659211"
```

通过添加 withcoord 选项，georadius 命令可以返回位置元素的名称和它们的地理位置信息。

### 查询并返回距离

```bash
127.0.0.1:6379> georadius china:city 110 30 1000 km withdist
1) 1) "chongqing"
   2) "341.9374"
2) 1) "xian"
   2) "483.8340"
3) 1) "shengzhen"
   2) "924.6408"
4) 1) "hangzhou"
   2) "977.5143"
```

返回每个城市到查询中心的距离。

### 查询并限制结果数量

```bash
127.0.0.1:6379> georadius china:city 110 30 1000 km withdist withcoord count 2
1) 1) "chongqing"
   2) "341.9374"
   3) 1) "106.49999767541885376"
      2) "29.52999957900659211"
2) 1) "xian"
   2) "483.8340"
   3) 1) "108.96000176668167114"
      2) "34.25999964418929977"
```

仅返回两个最近的城市及其坐标和距离。

## 查询某个距离范围内的元素

```bash
127.0.0.1:6379> georadiusbymember china:city chongqing 1000000 m
1) "chongqing"
2) "xian"
```

georadiusbymember 命令根据指定成员的位置和给定的距离，返回范围内的位置元素。

上文中，我们已经知道 geo 的底层是通过 zset 来实现的，那么也就意味着我们也可以通过 zset 命令来操作 geo。

```bash
127.0.0.1:6379> zrange china:city 0 -1
1) "chongqing"
2) "xian"
3) "shengzhen"
4) "hangzhou"
5) "shanghai"
6) "beijing"
127.0.0.1:6379> zrem china:city xian
(integer) 1
127.0.0.1:6379> zrange china:city 0 -1
1) "chongqing"
2) "shengzhen"
3) "hangzhou"
4) "shanghai"
5) "beijing"
```

## 实践意义

在实际开发中，你可以使用 Redis 的地理空间功能来实现各种基于位置的服务，如商家定位、配送范围估算、最近服务点查询等。通过上述例子，我们可以看到，Redis 提供的地理空间功能既强大又易于使用，能够帮助开发者在构建地理空间数据相关应用时，提高开发效率和应用性能。

## 结语

Redis 的地理空间数据处理模块为处理和查询地理信息提供了强大而高效的方法。无论你是在处理简单的位置数据查询还是构建复杂的地理信息系统（GIS），Redis 都能为你提供必要的支持。
