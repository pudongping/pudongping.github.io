---
title: hyperf 使用模型写较为复杂 union 子查询并做分页
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - Hyperf
  - PHP
  - Swoole
abbrlink: 7dec67c2
date: 2021-06-11 08:59:03
img:
coverImg:
password:
summary:
---

# hyperf 使用模型写 union 子查询并做分页

## 最终需要实现的 sql 语句为如下所示：

```sql

SELECT
	`dfo_al`.* 
FROM
	(
	(
SELECT
	`a`.`log_id`,
	`a`.`change_time`,
	`a`.`user_id`,
	`a`.`pay_points`,
	`a`.`change_type`,
	`a`.`from_user_id`,
	`u`.`user_id` AS `u_user_id`,
	`u`.`username`,
	`u`.`head_pic`,
	`u`.`vip_time` 
FROM
	`a`
	LEFT JOIN `dfo_users` AS `u` ON `a`.`user_id` = `u`.`user_id` 
WHERE
	( `a`.`user_id` = 3649 AND `a`.`pay_points` > 0 ) 
	AND ( `a`.`from_user_id` IS NULL OR `a`.`from_user_id` = 0 ) 
	) UNION
	(
SELECT
	`a`.`log_id`,
	`a`.`change_time`,
	`a`.`user_id`,
	`a`.`pay_points`,
	`a`.`change_type`,
	`a`.`from_user_id`,
	`u`.`user_id` AS `u_user_id`,
	`u`.`username`,
	`u`.`head_pic`,
	`u`.`vip_time` 
FROM
	`a`
	LEFT JOIN `dfo_users` AS `u` ON `a`.`from_user_id` = `u`.`user_id` 
WHERE
	( `a`.`user_id` = 3649 AND `a`.`pay_points` > 0 ) 
	AND ( `a`.`from_user_id` IS NOT NULL OR `a`.`from_user_id` > 0 ) 
	) 
	) AS dfo_al 
ORDER BY
	`log_id` DESC 
	LIMIT 2 OFFSET 0

```

## hyperf 代码为

> 以下代码仅为示例代码，仅提供参考使用，具体实现请根据你自己的业务逻辑实现。

```php

        $where = [
            ['a.user_id', '=', auth()->user_id],
            ['a.pay_points', '>', 0]
        ];
        $fields = [
            'a.log_id',
            'a.change_time',
            'a.user_id',
            'a.pay_points',
            'a.change_type',
            'a.from_user_id',
            'u.user_id AS u_user_id',
            'u.username',
            'u.head_pic',
            'u.vip_time'
        ];

        $table1 = AccountLog::query()
            ->from('account_log as a')
            ->leftJoin('users as u', 'a.user_id', '=', 'u.user_id')
            ->select($fields)
            ->where($where)
            ->where(function ($query) {
                $query->whereNull('a.from_user_id')->orWhere('a.from_user_id', 0);
            });

        $table2 = AccountLog::query()
            ->from('account_log as a')
            ->leftJoin('users as u', 'a.from_user_id', '=', 'u.user_id')
            ->select($fields)
            ->where($where)
            ->where(function ($query) {
                $query->whereNotNull('a.from_user_id')->orWhere('a.from_user_id', '>', 0);
            });

        $table = $table1->union($table2);

        $tablePrefix = Db::connection()->getTablePrefix();  // 获取数据表前缀 => 或者使用  $tablePrefix = Db::getConfig('prefix'); 都可以
        $model = AccountLog::query()
            ->mergeBindings($table->getQuery())
            ->select(['al.*'])
            ->from(Db::raw("({$table->toSql()}) as {$tablePrefix}" . 'al'));

        return $model->orderBy('log_id', 'desc')->paginate(2);

```
