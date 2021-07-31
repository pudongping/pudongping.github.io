---
title: MySQL 性能分析
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: MySQL
tags:
  - MySQL
  - explain
  - 性能优化
abbrlink: 1e4459df
date: 2021-07-31 23:45:17
img:
coverImg:
password:
summary:
---


# MySQL 性能分析

## explain

```mysql

explain select * from users;

-- 本意为显示警告信息。但是和 explain 一块儿使用，就会显示出优化后的 sql。需要注意使用顺序。（只能在 mysql cli 中才会有结果）
show warnings;

```

> 最重要的 5 个字段是：id、type、key、rows、Extra

- **id**： select 查询的序列号，包含一组数字，表示查询中执行 select 子句或操作表的顺序

1. id 相同，执行顺序由上至下
2. id 不同，如果是子查询，id 的序号会递增，id 值越大优先级越高，越先被执行
3. id 有相同也有不同的，id 值越大优先级越高越先被执行，id 值相同的按由上到下执行

- **select_type**：查询的类型，主要是用于区别普通查询，联合查询，子查询等的复杂查询

1. simple ： 简单的 select 查询，查询中不包含子查询或者 union
2. primary ：查询中若包含任何复杂的子部份，最 外层查询则被标记为 primary
3. subquery ： 在 select 或 where 列表中包含子查询
4. derived ： 在 from 列表中包含的子查询被标记为 derived（衍生）MySQL 会递归执行这些子查询，把结果放在临时表里。
5. union ： 若第二个 select 出现在 union 之后，则被标记为 union：若 union 包含在 from 子句的子查询中，外层 select 将被标记为： derived
6. union result ： 从 union 表获取结果的 select

- **table** ： 显示这一行的数据是关于哪张表的
- **type** ： 访问类型排序
    - 常用的从最好到最差依次是，system > const > eq_ref > ref > range > index > ALL
    - 完整的从最好到最差依次是 system > const > eq_ref > ref > fulltext > ref_or_null > index_merge > unique_subquery > index_subquery > range > ALL
    - 一般来说， 得保证查询至少达到 range 级别，最好能达到 ref 级别

1. system ：表只有一行记录（等于系统表），这是 const 类型的特例，平时不会出现，这个也可以忽略不计
2. const ： 表示通过索引一次就找到了，const 用于比较 primary key 或者 unique 索引。因为只匹配一行数据，所以很快。如将主键置于 where 列表中，MySQL 就能将该查询转换为一个常量
3. eq_ref : 唯一性索引扫描，对于每个索引键，表中只有一条记录与之匹配。常见于主键或唯一索引扫描。
4. ref ： 非唯一性索引扫描，返回匹配某个单独值的所有行，本质上也是一种索引访问，它返回所有匹配某个单独值的行，然而，它可能会找到多个符合条件的行，所以他应该属于查找和扫描的混合体。
5. range ：只检索给定范围的行，使用一个索引来选择行，key 列显示使用了哪个索引，一般就是在你的 where 语句中出现了 between 、< 、> 、in 等的查询。这种范围扫描索引扫描比全表扫描要好，因为它只需要开始于索引的某一点，而结束于另一点，不用扫描全部索引
6. index ： Full Index Scan，Index 与 ALL 区别为 index 类型只遍历索引树。这通常比 ALL 快，因为索引文件通常比数据文件小。（也就是说虽然 all 和 index 都是读全表，但 index 是从索引中读取的，而 all 是从硬盘中读的）
7. all ： Full Table Scan，将遍历全表以找到匹配的行

**一般来说，得保证查询至少达到 range 级别，最好能达到 ref**

- **possible_keys** ：显示可能应用在这张表中的索引，一个或多个。查询涉及到的字段上若存在索引，则该索引将被列出，**但不一定被查询实际使用**
- **key** ： 实际使用的索引。如果为 null，则没有使用索引。查询中若使用了覆盖索引，则该索引仅出现在 key 列表中。
- **key_len** : 表示索引中使用的字节数，可通过该列计算查询中使用的索引的长度。在不损失精确性的情况下，长度越短越好。key_len 显示的值为索引字段的最大可能长度，并非实际使用长度，即 key_len 是根据表定义计算而得，不是通过表内检索出的
- **ref** ： 显示索引的哪一列被使用了，如果可能的话，是一个常数。哪些列或常量被用于查找索引列上的值
- **rows** ：根据表统计信息及索引选用情况，大致估算出找到所需的记录所需要读取的行数（行数越小越好）
- **Extra** ：包含不适合在其它列中显示但十分重要的额外信息

1. Using filesort ： 说明 mysql 会对数据使用一个外部的索引排序，而不是按照表内的索引顺序进行读取。MySQL 中无法利用索引完成的排序操作称为“文件排序”。**九死一生了，效率比较差**
2. Using temporary ：使用了临时表保存中间结果，MySQL 在对查询结果排序时使用临时表。常用于排序 order by 和分组查询 group by。**十死无生，效率最差！**
3. Using index ： 表示相应的 select 操作中使用了覆盖索引（Covering index），避免访问了表的数据行，效率不错！如果同时出现 using where ，表明索引被用来执行索引键值的查找；如果没有同时出现 using where ，表明索引用来读取数据而非执行查找动作。
4. Using where ：表明使用了 where 过滤
5. using join buffer ：使用了连接缓存
6. impossible where ： where 子句的值总是 false，不能用来获取任何元组
7. select tables optimized away ：在没有 group by 子句的情况下，基于索引优化 min/max 操作或者对于 MyISAM 存储引擎优化 count(*) 操作，不必等到执行阶段再进行计算，查询执行计划生成的阶段即完成优化。
8. distinct ：优化 distinct 操作，在找到第一匹配的元组后即停止找同样值的动作


## 查询截取分析

- 通俗来讲
1. 观察，至少跑 1 天 ，看看生产的慢 sql 情况。
2. 开启慢查询日志，设置阈值，比如超过 5 秒钟的就是慢 sql，并将它抓取出来。
3. explain + 慢 sql 分析
4. show profile
5. 运维经理或者 DBA，进行 sql 数据库服务器的参数调优。

- 学术说法
1. 慢查询的开启并捕获
2. explain + 慢 sql 分析
3. show profile 查询 sql 在 mysql 服务器里面的执行细节和生命周期情况
4. sql 数据库服务器的参数调优

### 提高 order by 的速度
1. order by 时，select * 是一个大忌，只查询需要的字段，这点非常重要。在这里的影响是：
- 当查询的字段大小总和小于 max_length_for_sort_data 而且排序字段不是 text|blob 类型时，会用改进后的算法——单路排序，否则用老算法——多路排序。
- 两种算法的数据都有可能超出 sort_buffer 的容量，超出之后，会创建 tmp 文件进行合并排序，导致多次 I/O，但是用单路排序算法的风险会更大一些，所以要提高 sort_buffer_size
2. 尝试提高 sort_buffer_size 不管用哪种算法，提高这个参数都会提高效率，当然，要根据系统的能力去提高，因为这个参数是针对每个进程的
3. 尝试提高 max_length_for_sort_data 提高这个参数，会增加用改进算法的概率，但是如果设的太高，数据总容量超出 sort_buffer_size 的概率就增大，明显症状是高的磁盘 I/O 活动和低的处理器使用率。


### 为排序使用索引

- MySQL 两种排序方式：文件排序或扫描有序索引排序
- MySQL 能为排序与查询使用相同的索引

KEY a_b_c (a,b,c)

1. order by 能使用索引最左前缀   
   ORDER BY a   
   ORDER BY a,b  
   ORDER BY a,b,c  
   ORDER BY a DESC, b DESC, c DESC
2. 如果 where 使用索引的最左前缀定义为常量，则 order by 能使用索引  
   WHERE a = const ORDER BY b,c   
   WHERE a = const AND b = const ORDER BY c  
   WHERE a = const ORDER BY b,c  
   WHERE a = const AND b > const ORDER BY b,c

## 开启慢查询日志
- 查看是否开启及如何开启慢查询日志

```
# 查看是否开启了慢查询日志
show variables like '%slow_query_log%';

# 开启慢查询日志（临时开启）
set global slow_query_log = 1;
```

如果要永久生效，就必须修改配置文件 my.cnf （其他系统变量也是如此）  
修改 my.cnf 文件，[mysqld] 下增加或修改参数 slow_query_log 和 slow_query_log_file 后，然后重启 MySQL 服务器。

```
slow_query_log = 1 
slow_query_log_file = /var/lib/mysql/slow-query.log
```

慢查询日志是由参数 long_query_time 控制，默认情况下 long_query_time 的值为 10 秒，可以使用以下命令查看

```
show variables like 'long_query_time%';
```

可以使用命令修改，也可以在 my.cnf 参数里面修改。  
假如运行时间正好等于 long_query_time 的情况，并不会被记录下来。也就是说，在 mysql 源码里是 **判断大于 long_query_time，而非大于等于**


设置阙值到 3 秒钟的就是慢 sql

```
set global long_query_time = 3
```

设置慢查询日志阙值后看不出变化？
1. 需要重新连接或新开一个会话才能看到修改值
2. 或者直接使用以下命令也可以看到修改后的结果

```
show global variables like 'long_query_time';
```

如何测试？

```
# 模拟查询超过 4 秒钟
select sleep(4);
```

查看慢查询日志中记录了有多少条慢 sql

```
show global status like '%Slow_queries%';
```

### 慢查询分析工具 mysqldumpslow
可用参数：
- s : 表示按照何种方式排序
- c ：访问次数
- l ：锁定时间
- r ：返回记录
- t ：查询时间
- al ：平均锁定时间
- ar ：平均返回记录数
- at ：平均查询时间
- t ：即返回前面多少条的数据
- g ：后边搭配一个正则匹配模式，大小写不敏感的

```
# 得到返回记录集最多的 10 条 sql
mysqldumpslow -s r -t 10 /var/lib/mysql/slow-query.log

# 得到访问次数最多的 10 条 sql
mysqldumpslow -s c -t 10 /var/lib/mysql/slow-query.log

# 得到按照时间排序的前 10 条里面包含左连接的查询语句
mysqldumpslow -s t -t 10 -g "left join" /var/lib/mysql/slow-query.log

另外建议在使用这些命令时结合 | 和 more 使用，否则有可能出现爆屏情况
mysqldumpslow -s r -t 10 /var/lib/mysql/slow-query.log | more
```

## Show Profile （可以看到每一条 sql 执行的生命周期）

是 mysql 提供可以用来分析当前会话中语句执行的资源消耗情况。可以用于 sql 的调优的测量。  
默认情况下，参数处于关闭状态，并保存最近 15 次的运行结果。

1. 查看是否已经开启

```
# 查看是否已经开启
show VARIABLES like 'profiling';

# 如果没有开启的话，则开启
set profiling = on;
```

2. 运行各种查询语句
3. 使用以下命令，即可看到以上的查询语句

```
show profiles;
```

4. 诊断 sql

```
show profile cpu,block io for query 1;（上一步问题 sql 前面的数字号码）
```

show profile 的可用参数为：
- all ：显示所有的开销信息
- block io ：显示块 io 相关开销
- context switches ：上下文切换相关开销
- cpu ：显示 cpu 相关开销信息
- ipc ：显示发送和接收相关开销信息
- memory ：显示内存相关开销信息
- page faults ： 显示页面错误相关开销信息
- source ：显示和 Source_function，Source_file，Source_line 相关的开销信息
- swaps ：显示交换次数相关开销的信息

### 日常开发需要注意项
1. converting HEAP to MyISAM 查询结果太大，内存都不够用了往磁盘上搬了。
2. Creating tmp table 创建临时表。（拷贝数据到临时表，用完再删除）
3. Copying to tmp table on disk 把内存中临时表复制到磁盘，危险！
4. locked 锁表了


## 全局查询日志
> 永远不要在生产环境中开启这个功能！

- 在配置文件中开启  
  在 mysql 的 my.cnf 中，设置如下：

```
# 开启
general_log = 1
# 记录日志文件的路径
general_log_file = /path/logfile
# 输出格式
log_output = FILE
```

- 直接执行 sql 语句开启

```
# 开启
set global general_log=1;
# 用表格的方式记录查询日志
set global log_output = 'TABLE';
# 此后，你所编写的 sql 语句，将会记录到 mysql 库里的 general_log 表
select * from mysql.general_log;
```
