---
title: MySQL 锁机制
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: MySQL
tags:
  - MySQL
  - 锁
abbrlink: eee5658b
date: 2021-07-29 21:50:02
img:
coverImg:
password:
summary:
---

# MySQL 锁机制

## 锁的概念
- 从对数据操作的类型（读/写）来分
    - 读锁（共享锁）：针对同一份数据，多个读操作可以同时进行而不会互相影响。
    - 写锁（排它锁）：当前写操作没有完成前，它会阻断其他写锁和读锁。

- 锁相关命令

```
# 查看是否有表已经被上锁
show open tables;

# 给 table1 上读锁，给 table2 上写锁
lock table table1 read,table2 write;

# 给所有表解锁
unlock tables;
```

### 表锁（偏向于 MyISAM）

给表加**读**锁的情况：

1. session-1

```
# 对 table1 加读锁
lock table table1 read; 

# 因为读锁是共享锁，因此可以查出数据
select * from table1;

# 但是无法修改和插入新数据
update table1 set name = 'a1' where id = 1;

# 但是不能读取其他的表
select * from table2;
```

2. session-2

```
# 因为读锁是共享锁，因此可以查出数据
select * from table1;

# 也可以读其他的表
select * from table2;

# 当 session-2 想要修改 session-1 已经锁的表 table1 时，已经被阻塞，需要等到 table1 被解锁之后才可以将数据更新
update table1 set name = 'a1' where id = 1;
```

给表加**写**锁的情况：

1. session-1

```
# 给 table1 加写锁
lock table table1 write;

# 能够查询 table1
select * from table1;

# 自己能够更新 table1 的数据
update table1 set name = 'a1' where id = 1;

# 但是不能读取其他的表
select * from table2;
```

2. session-2

```
# 可以查询其他没有上写锁的表
select * from table2;

# 查询 session-1 中已经上写锁的 table1 时，会发生阻塞，需要等到 session-1 将 table1 解锁才会查得出来结果
select * from table1;
```

MyISAM 在执行查询语句 （select） 前，会自动给涉及的所有表加读锁，在执行增删改操作前，会自动给涉及的表加写锁。  
MySQL的表级锁有两种模式：  
表共享读锁（Table Read Lock）  
表独占写锁（Table Write Lock）

锁类型 | 可否兼容 | 读锁 | 写锁
--- | --- | --- | ---
读锁 | 是 | 是 | 否
写锁 | 是 | 否 | 否

**结论**  
结合上表，所有对 MyISAM 表进行操作，会有以下情况：
1. 对 MyISAM 表的读操作（加读锁），不会阻塞其他进程对同一表的读请求，但会阻塞对同一表的写请求。只有当读锁释放后，才会执行其它进程的写操作。
2. 对 MyISAM 表的写操作（加写锁），会阻塞其他进程对同一表的读和写操作，只有当写锁释放后，才会执行其他进程的读写操作。

**简而言之，就是读锁会阻塞写，但是不会阻塞读。而写锁则会把读和写都阻塞**

#### 如何分析表锁 ？

```mysql
show status like 'table%';
```

这里有两个状态变量记录 MySQL 内部表级锁定的情况，两个变量说明如下：
- Table_locks_immediate ：  产生表级锁定的次数，表示可以立即获取锁的查询次数，每立即获取锁值加 1；
- Table_locks_waited ： 出现表级锁定争用而发生等待的次数（不能立即获取锁的次数，每等待一次锁值加 1），此值高则说明存在着较严重的表级锁争用情况。

**此外，MyISAM 的读写锁调度是写优先，这也是 MyISAM 不适合做写为主表的引擎。因为写锁后，其他线程不能做任何操作，大量的更新会使查询很难得到锁，从而造成永远阻塞**


### 行锁 （偏向于 InnoDB）
特点：
1. 开销大，加锁慢；会出现死锁；锁定粒度最小，发生锁冲突的概率最低，并发度也最高。
2. InnoDB 与 MyISAM 的最大不同有两点：一是支持事务（TRANSACTION），二是采用了行级锁

**事务的 4 个特性： ACID**
1. 原子性（Atomicity）：事务是一个原子操作单元，其对数据的修改，要么全部执行，要么全部不执行。
2. 一致性（Consistent）：在事务开始和完成时，数据都必须保持一致状态。这意味着所有相关的数据规则都必须应用于事务的修改，以保持数据的完整性：事务结束时，所有的内部数据结构（如 B树索引或双向链表）也都必须是正确的。
3. 隔离性（Isolation）：数据库系统提供一定的隔离机制，保证事务在不受外部并发操作影响的“独立”环境执行。这意味着事务处理过程中的中间状态对外部是不可见的，反之亦然。
4. 持久性（Durable）：事务完成之后，它对于数据的修改是永久性的，即使出现系统故障也能够保持。

- 间隙锁   
  当我们用**范围条件**而不是相等条件检索数据，并请求共享或排他锁时，InnoDB 会给符合条件的已有数据记录的索引项加锁；（宁可错杀一千，也不可放过一个）对于键值在条件范围内但不存在的记录，叫做“间隙（GAP）”，InnoDB 也会对这个“间隙”加锁，这种锁机制就是所谓的间隙锁（Next-Key锁）。
- 间隙锁的危害  
  因为 query 执行过程中通过范围查找的话，他会锁定整个范围内所有的索引键值，即使这个键值并不存在。间隙锁有一个比较致命的弱点，就是当锁定一个范围值之后，即使某些不存在的键值也会被无辜的锁定，而造成在锁定的时候无法插入锁定键值范围内的任何数据。在某些场景下这可能会对性能造成很大的危害。


- 如何给增加一条行锁 ？

``` 
# 设置一个起点
begin;

# 给指定行上锁
select * from table1 where a = 8 for update;

# 上锁之后，做一些你想要做的操作，比如说需要更新当前数据
update table1 set b = 123 where a = 8;

# 做完你想要做的操作之后，一定要提交
commit;
```

Innodb 存储引擎由于实现了行级锁定，虽然在锁定机制的实现方面所带来的性能损耗可能比表级锁定要更高一些，但是在整体并发处理能力方面要远远优于 MyISAM 的表级锁定的。当系统并发量较高的时候，InnoDB 的整体性能和 MyISAM 相比就会有比较明显的优势了。   
但是，InnoDB 的行级锁同样也有其脆弱的一面，当我们使用不当的时候，可能会让 InnoDB 的整体性能表现不仅不能比 MyISAM 高，甚至可能会更差。

#### 如何分析行锁 ？
通过检查 InnoDB_row_lock 状态变量来分析系统上的行锁的争夺情况

```mysql
show status like 'innodb_row_lock%';
```

对各个状态量的说明如下：
- innodb_row_lock_current_waits ：当前正在等待锁定的数量
- innodb_row_lock_time ： 从系统启动到现在锁定总时间长度
- innodb_row_lock_time_avg ：每次等待所花平均时间
- innodb_row_lock_time_max ：从系统启动到现在等待最常的一次所花的时间
- innodb_row_lock_waits ：系统启动后到现在总共等待的次数

对于这 5 个状态变量，比较重要的主要是 innodb_row_lock_time_avg （等待平均时长）、innodb_row_lock_waits （等待总次数）、innodb_row_lock_time （等待总时长）这三项。尤其是当等待次数很高，而且每次等待时长也不小的时候，我们就需要分析系统中为什么会有如此多的等待，然后根据分析结果着手指定优化计划。
