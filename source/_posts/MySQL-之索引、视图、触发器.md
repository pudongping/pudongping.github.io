---
title: MySQL 之索引、视图、触发器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: MySQL
tags:
  - MySQL
  - 索引
  - 视图
  - 触发器
abbrlink: b1cb609d
date: 2021-08-01 00:27:16
img:
coverImg:
password:
summary:
---


# MySQL 之索引、视图、触发器

## 索引

- 索引的引入

索引是由数据库表中一列或者多列组合而成，其作用是提高对表中数据的查询速度；类似于图书的目录，方便快速定位，寻找指定的内容。

- 索引的优缺点

优点：提高查询数据的速度  
缺点：创建和维护索引的时间增加了

- 建立索引的建议
1. 一张表建议最多建立 5 个索引
2. 建立复合索引优于单值索引（复合索引占用空间小）

- 建立索引的技巧
1. 如果是左连接则需要在右表关联字段上建立索引，因为左表是查的全部数据。如果是右连接则需要在左表关联字段上建立索引。
2. 尽可能减少 join 语句中的 NestedLoop 的循环总次数。（永远用小结果集驱动大的结果集）
3. 优先优化 NestedLoop 的内层循环。
4. 保证 join 语句中被驱动表上 join 条件字段已经被索引。
5. 当无法保证被驱动表的 join 条件字段被索引且内存资源充足的前提下，不要太吝啬 JoinBuffer 的设置。

- 索引失效的常见原因
1. 查询全部列，不会使用到索引（select *）
2. 不遵守最佳左前缀法则（如果索引了多列，要遵守最左前缀法则。指的是查询从索引的最左前列开始并且**不跳过索引的列**）
3. 在索引列上做任何操作（计算、函数、（自动或手动）类型转换），会导致索引失效而转向全表扫描
4. 存储引擎使用了索引中范围条件右边的列，会导致不会使用到索引
5. 尽量使用覆盖索引（只访问索引的查询（索引列和查询列一致）），减少 select *
6. mysql 在使用不等于 （!= 或 <>）的时候无法使用索引会导致全表扫描
7. is null,is not null 也无法使用索引
8. like 以通配符开头 （'%abc' 或者 '%abc%'）mysql 索引失效会变成全表扫描的操作，当百分号写在右边的时候索引不会失效。解决 '%abc%' 索引失效的方法是，在要模糊查询字段上建立索引，使用覆盖索引的方式查询，则索引则不会失效。
9. varchar 类型 （字符串）不加单引号索引失效（如果是 InnoDB 存储类型，会导致行锁变表锁）
10. 少用 or，用它来连接时索引会失效，即使其中的条件带有索引也不会使用到索引，如果要想使用 or，又想让索引生效，只能将 or 条件中的每一列都加上索引。如果出现 or 的语句中没有一个列加了索引，那么建议使用 union 拼接多个查询语句。
11. not in 和 not exist 不会走索引

- 优化口诀  
  全值匹配我最爱，最左前缀要遵守；    
  带头大哥不能死，中间兄弟不能断；  
  索引列上少计算，范围之后全失效；  
  like百分写最右，覆盖索引不写*；  
  不等空值还有or，索引失效要少用。



- 创建索引的前提

索引的效率取决于**索引列的值是否为散列**，即该列的值如果越互不相同，那么索引效率越高。反过来，如果记录的列存在大量相同的值，例如 gender 列，大约一半的记录值是 M，另一半是 F，因此，对该列创建索引就没有意义。区分度的公式是 `count(distinct col)/count(*)`，表示字段不重复的比例，比例越大我们扫描的记录数越少。

可以对一张表创建多个索引。索引的优点是提高了查询效率，缺点是在插入、更新和删除记录时，需要同时修改索引，因此，索引越多，插入、更新和删除记录的速度就越慢。

***对于主键，关系数据库会自动对其创建主键索引。使用主键索引的效率是最高的，因为主键会保证绝对唯一。***

- 创建表的时候创建索引：

```mysql

-- 创建普通索引
CREATE TABLE t_user1 (
	id INT,
	userName VARCHAR (20),
	PASSWORD VARCHAR (20),
	INDEX (userName)
);

-- 创建唯一性索引并为索引取别名
CREATE TABLE t_user2 (
	id INT,
	userName VARCHAR (20),
	PASSWORD VARCHAR (20),
	UNIQUE INDEX usrn (userName)
);

-- 创建多列索引
CREATE TABLE t_user3 (
	id INT,
	userName VARCHAR (20),
	PASSWORD VARCHAR (20),
	INDEX index_user_pwd (userName,PASSWORD)
);

```

- 在已有表中创建索引

```mysql

-- 在已有表中创建普通索引
CREATE INDEX index_userName ON t_user4(userName);

-- 在已有表中创建唯一性索引
CREATE UNIQUE INDEX index_userName ON t_user4(userName);

-- 在已有表中创建多列索引
CREATE INDEX index_userName_pwd ON t_user4(userName,PASSWORD);
// 或者采用下面的方式
ALTER TABLE students
ADD INDEX idx_name_score (name, score);

-- 使用 ALTER 删除索引
ALTER TABLE t_user5 ADD INDEX index_user(userName)

```

- 查看索引

```mysql

show index from table_name\G;

```

- 删除索引

```mysql

-- 删除索引
-- DROP INDEX 索引名 ON 表名;
DROP INDEX index_user ON t_user5;

```

- 索引检索原理

![索引检索原理](https://upload-images.jianshu.io/upload_images/14623749-8dbf9172d1163963.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![索引的分析](https://upload-images.jianshu.io/upload_images/14623749-7b8ef85580ae6d04.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 哪些情况下应该建立索引 ？
1. 主键自动建立唯一索引
2. 频繁作为查询条件的字段应该创建索引
3. 查询中与其它表关联的字段，外键关系建立索引
4. 频繁更新的字段不适合创建索引（因为每次更新不单单是更新了记录还会更新索引）
5. where 条件里用不到的字段不创建索引
6. 单键索引还是组合索引的选择问题？（在高并发下倾向创建组合索引）
7. 查询中排序的字段，排序字段若通过索引去访问将大大提高排序速度
8. 查询中统计或者分组字段

- 哪些情况下不要建立索引？
1. 表记录太少
2. 经常增删改的表
3. 数据重复且分布平均的表字段，因此应该只为最经常查询和最经常排序的数据列建立索引。注意：如果某个数据列包含许多重复的内容，为它建立索引就没有太大的实际效果。


## 视图

- 视图的引入
1. 试图是一种虚拟的表，是从数据库中一个或者多个表中导出来的表。
2. 数据库中只存放了视图的定义，而并没有存放视图中的数据，这些数据存放在原来的表中。
3. 使用视图查询数据时，数据库系统会从原来的表中取出对应的数据。

- 视图的作用
1. 使操作简便化；eg：如果一张表中有 100 个字段，需求只需要 20 个字段，那么可以定义一个视图只取出 20 个字段。
2. 增加数据的安全性；eg：如果写代码的时候不想要别人知道某些字段，那么可以定义视图，只取出安全系数低的字段
3. 提高表的逻辑独立性；

- 创建视图

```mysql

CREATE [ALGORITHM = {UNDEFIEND | MERGE | TEMPTABLE}]
    VIEW 视图名 [(属性清单)]
    AS SELECT 语句
    [WITH [CASCADED | LOCAL] CHECK OPTION]
    
-- 创建单表视图
CREATE VIEW v1 AS SELECT userName,password FROM t_user4

-- 创建单表视图并给视图字段取别名
CREATE VIEW v1(u,p) AS SELECT userName,password FROM t_user4

-- 查询视图结果
SELECT * FROM v1

-- 在多表上创建视图
CREATE VIEW v2 AS SELECT bookName.bookTypeName FROM t_book,t_booktype WHERE t_book.bookTypeId = t_booktype.id    
    
```

## 触发器

- 触发器的引入

触发器（trigger）是由事件来触发某个操作，这些事件包括 insert 语句、 update 语句和 delete 语句。当数据库系统执行这些事件时，就会激活触发器执行相应的操作。

- 创建与使用触发器

1. 创建只有一个执行语句的触发器

```mysql

CREATE TRIGGER 触发器名 BEFORE | AFTER 触发事件 ON 表名 FOR EACH ROW 执行语句

-- 其中 new 和 old 为过渡变量， new 代表新的数据，old 代表旧的数据
eg：

CREATE TRIGGER trig_book AFTER INSERT
	ON t_book FOR EACH ROW
		UPDATE t_booktype SET bookNum=bookNum+1 WHERE new.bookTypeId=t_booktype.id

-- 执行以下语句之后将会触发触发器		
INSERT INTO t_book VALUES(NULL,'php学习',100,'ke',1)		

```

2. 创建有多个执行语句的触发器

```mysql

CREATE TRIGGER 触发器名 BEFORE | AFTER 触发事件
    ON 表名 FOR EACH ROW
    BEGIN
        执行语句列表
    END  

-- 其中，因为 mysql 遇到分号（;）之后会认为语句终止（分号前面的语句为执行语句），
因此需要使用 delimiter 来手动定义在 | 符号中间的语句才为执行语句   
eg：

delimiter |
CREATE TRIGGER trig_book2 AFTER DELETE
	ON t_book FOR EACH ROW
	BEGIN 
		UPDATE t_booktype SET bookNum=bookNum-1 WHERE old.bookTypeId=t_booktype.id;
		INSERT INTO t_log VALUES (NULL,NULL,'在book表里删除了一条数据');
		DELETE FROM t_test WHERE old.bookTypeId=t_test.id;
	END
|
delimiter;

-- 执行以下语句之后将会触发触发器
DELETE FROM t_book WHERE id=3;
    
```

- 查看触发器

1. 直接执行 sql 语句

```mysql

SHOW TRIGGERS;

```

2. 在系统数据库中 information_schema 库中查看 TRIGGERS 表

- 删除触发器

```mysql

DROP TRIGGER [触发器名];

eg:
DROP TRIGGER trig_book2;

```
