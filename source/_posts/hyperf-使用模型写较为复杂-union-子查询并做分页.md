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
	`dfo_account_log`.`log_id`,
	`dfo_account_log`.`change_time`,
	`dfo_account_log`.`user_id`,
	`dfo_account_log`.`pay_points`,
	`dfo_account_log`.`change_type`,
	`dfo_account_log`.`from_user_id`,
	`dfo_u`.`user_id` AS `u_user_id`,
	`dfo_u`.`username`,
	`dfo_u`.`head_pic`,
	`dfo_u`.`vip_time` 
FROM
	`dfo_account_log`
	LEFT JOIN `dfo_users` AS `dfo_u` ON `dfo_account_log`.`user_id` = `dfo_u`.`user_id` 
WHERE
	( `dfo_account_log`.`user_id` = 3649 AND `dfo_account_log`.`pay_points` > 0 ) 
	AND ( `dfo_account_log`.`from_user_id` IS NULL OR `dfo_account_log`.`from_user_id` = 0 ) 
	) UNION
	(
SELECT
	`dfo_account_log`.`log_id`,
	`dfo_account_log`.`change_time`,
	`dfo_account_log`.`user_id`,
	`dfo_account_log`.`pay_points`,
	`dfo_account_log`.`change_type`,
	`dfo_account_log`.`from_user_id`,
	`dfo_u`.`user_id` AS `u_user_id`,
	`dfo_u`.`username`,
	`dfo_u`.`head_pic`,
	`dfo_u`.`vip_time` 
FROM
	`dfo_account_log`
	LEFT JOIN `dfo_users` AS `dfo_u` ON `dfo_account_log`.`from_user_id` = `dfo_u`.`user_id` 
WHERE
	( `dfo_account_log`.`user_id` = 3649 AND `dfo_account_log`.`pay_points` > 0 ) 
	AND ( `dfo_account_log`.`from_user_id` IS NOT NULL OR `dfo_account_log`.`from_user_id` > 0 ) 
	) 
	) AS dfo_al 
ORDER BY
	`log_id` DESC 
	LIMIT 2 OFFSET 0

```

## hyperf 代码为

> 以下代码仅为示例代码，仅提供参考使用，具体实现请根据你自己的业务逻辑实现。

```php

        $tbl = 'account_log';
        $where = [
            [$tbl . '.user_id', '=', auth()->user_id],
            [$tbl . '.pay_points', '>', 0]
        ];
        $fields = [
            $tbl . '.log_id',
            $tbl . '.change_time',
            $tbl . '.user_id',
            $tbl . '.pay_points',
            $tbl . '.change_type',
            $tbl . '.from_user_id',
            'u.user_id AS u_user_id',
            'u.username',
            'u.head_pic',
            'u.vip_time'
        ];

        $table1 = $this->accountLogModel
            ->leftJoin('users as u', $tbl . '.user_id', '=', 'u.user_id')
            ->select($fields)
            ->where($where)
            ->where(function ($query) use ($tbl) {
                $query->whereNull($tbl . '.from_user_id')->orWhere($tbl . '.from_user_id', 0);
            });

        $table2 = $this->accountLogModel
            ->leftJoin('users as u', $tbl . '.from_user_id', '=', 'u.user_id')
            ->select($fields)
            ->where($where)
            ->where(function ($query) use ($tbl) {
                $query->whereNotNull($tbl . '.from_user_id')->orWhere($tbl . '.from_user_id', '>', 0);
            });

        $table = $table1->union($table2);

        $tablePrefix = Db::connection()->getTablePrefix();  // 获取数据表前缀 => 或者使用  $tablePrefix = Db::getConfig('prefix'); 都可以
        $model = $this->accountLogModel
            ->mergeBindings($table->getQuery())
            ->select(['al.*'])
            ->from(Db::raw("({$table->toSql()}) as {$tablePrefix}" . 'al'));

        return $model->orderBy('log_id', 'desc')->paginate(2);

```
