---
title: 程序员常犯的坑：别再用 VARCHAR 存 IP 了！用对方式，性能何止提升10倍！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: MySQL
tags:
  - MySQL
  - 技巧
abbrlink: 412dd566
date: 2026-05-18 10:49:56
img:
coverImg:
password:
summary:
---

当我们在使用 MySQL 设计表结构的时候都会遇到一个经典问题：**“如果要存 IP 地址，到底该用什么数据类型？”**

别小看这个问题，如果存储方式选错了，未来不仅浪费空间，还可能导致查询性能低下。今天我就带你一步一步搞懂如何在 MySQL 里优雅地存储 IP 地址。

---

## 为啥不能直接用字符串？

很多小伙伴第一反应：IP 地址不就是 `"192.168.0.1"` 这种字符串嘛？那我直接用 `VARCHAR(15)` 或者 `CHAR(15)` 存不就行了吗？

表面上来看没啥大的问题，但这样做有几个弊端：

1. **空间浪费**：字符串比数字类型占用更多字节，还要额外存储长度信息。
2. **查询效率低**：字符串比较比数字比较要慢，索引效果也差。
3. **不方便计算**：比如要判断 IP 是否落在某个网段里，字符串处理就很麻烦。

所以，直接用字符串存 IP 地址，只能算是**最简单但不优雅的做法**。

那么，到底该怎么做才比较合适呢？

---

## IPv4：用 `INT UNSIGNED` 存整数更高效

IPv4 地址其实就是一个 **32 位的整数**，比如：

```bash
192.168.0.1 → 3232235521
```

在 MySQL 里我们完全可以用 `INT UNSIGNED` 类型来存储，既节省空间（4 字节），又利于索引和范围查询。

* 插入时，用 `INET_ATON()` 转换：

 ```sql
 INSERT INTO user_logs(ip) VALUES (INET_ATON('192.168.0.1'));
```
* 查询时，用 `INET_NTOA()` 再转回字符串：

```sql
SELECT INET_NTOA(ip) FROM user_logs;
```

这样就能同时兼顾 **空间效率** 和 **查询效率**。

---

## IPv6：用 `BINARY(16)` 存二进制

IPv6 地址更长，是 **128 位**，像这样：

```bash
2001:0db8:85a3:0000:0000:8a2e:0370:7334
```

如果直接用字符串存，需要 `CHAR(39)`，非常浪费空间。

更推荐的做法是用 `BINARY(16)` 或 `VARBINARY(16)` 存二进制形式。

MySQL 也有直接现成的函数：

* 插入时，用 `INET6_ATON()`：

```sql
INSERT INTO user_logs(ipv6) VALUES (INET6_ATON('2001:db8::1'));
```
* 查询时，用 `INET6_NTOA()` 转回：

```sql
SELECT INET6_NTOA(ipv6) FROM user_logs;
```

占用空间 **16 字节**，支持索引，性能杠杠滴。

---

## 那么，如果既要支持 IPv4，又要支持 IPv6 怎么办呢？

这就更简单了，直接：

* 字段类型设为 `VARBINARY(16)` 或 `BINARY(16)`
* 统一用 `INET6_ATON()` / `INET6_NTOA()` 存取

MySQL 会自动把 IPv4 地址映射到 IPv6 格式里，做到兼容存储，是不是非常 nice？

---

## 来个实战

以下面的建表语句为例：

```sql
CREATE TABLE user_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ip BINARY(16) NOT NULL COMMENT '存储IPv4/IPv6地址',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入IPv4
INSERT INTO user_logs(ip) VALUES (INET6_ATON('192.168.0.1'));

-- 插入IPv6
INSERT INTO user_logs(ip) VALUES (INET6_ATON('2001:db8::1'));

-- 查询
SELECT INET6_NTOA(ip) AS ip_address FROM user_logs;
```

一套方案，**IPv4 + IPv6 通吃**，相当 OK 👌

---

## 总结

* **IPv4-only** → `INT UNSIGNED`（配合 `INET_ATON` / `INET_NTOA`）
* **IPv6-only** → `BINARY(16)`（配合 `INET6_ATON` / `INET6_NTOA`）
* **通用方案** → `BINARY(16)` 或 `VARBINARY(16)`，一表解决 IPv4 和 IPv6

老铁们，下次设计数据库表的时候，就可以考虑不用 `VARCHAR(15)` 去存 IP 了，用这种方式，性能和优雅度都会大幅提升！🚀 赶紧用起来吧～如果对你有帮助，帮忙点个关注，支持一下呗～
