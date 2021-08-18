---
title: PHP 使用 CASE WHEN 进行批量更新 （当前基于 laravel 编写）
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Laravel
abbrlink: eaf6377d
date: 2021-08-18 19:50:18
img:
coverImg:
password:
summary:
---

# 使用 CASE WHEN 进行批量更新 （当前基于 laravel 编写）

## 最终执行的 sql 语句为

```mysql

UPDATE articles SET 
view_count = CASE
	WHEN id = 183 AND user_id = 1 THEN 6
	WHEN id = 182 AND user_id = 11 THEN 4
	WHEN id = 181 AND user_id = 15 THEN 4
	WHEN id = 180 AND user_id = 5 THEN 1
	ELSE view_count END,
updated_at = CASE 
	WHEN id = 183 AND user_id = 1 THEN '2020-11-06 06:44:58'
	WHEN id = 182 AND user_id = 11 THEN '2020-11-06 06:44:58'
	WHEN id = 181 AND user_id = 15 THEN '2020-11-06 06:44:58'
	WHEN id = 180 AND user_id = 5 THEN '2020-11-06 06:44:58'
	ELSE updated_at END

```

## 封装成方法为



```php

if (! function_exists('batch_update')) {
    /**
     * $where = [ 'id' => [180, 181, 182, 183], 'user_id' => [5, 15, 11, 1]];
     * $needUpdateFields = [ 'view_count' => [11, 22, 33, 44], 'updated_at' => ['2019-11-06 06:44:58', '2019-11-30 19:59:34', '2019-11-05 11:58:41', '2019-12-13 01:27:59']];
     *
     * 最终执行的 sql 语句如下所示
     *
     * UPDATE articles SET
     * view_count = CASE
     * WHEN id = 183 AND user_id = 1 THEN 44
     * WHEN id = 182 AND user_id = 11 THEN 33
     * WHEN id = 181 AND user_id = 15 THEN 22
     * WHEN id = 180 AND user_id = 5 THEN 11
     * ELSE view_count END,
     * updated_at = CASE
     * WHEN id = 183 AND user_id = 1 THEN '2019-12-13 01:27:59'
     * WHEN id = 182 AND user_id = 11 THEN '2019-11-05 11:58:41'
     * WHEN id = 181 AND user_id = 15 THEN '2019-11-30 19:59:34'
     * WHEN id = 180 AND user_id = 5 THEN '2019-11-06 06:44:58'
     * ELSE updated_at END
     *
     *
     * 批量更新数据
     *
     * @param string $tableName  需要更新的表名称
     * @param array $where  需要更新的条件
     * @param array $needUpdateFields  需要更新的字段
     * @return bool|int  更新数据的条数
     */
    function batch_update(string $tableName, array $where, array $needUpdateFields)
    {

        if (empty($where) || empty($needUpdateFields)) return false;
        // 第一个条件数组的值
        $firstWhere = $where[array_key_first($where)];
        // 第一个条件数组的值的总数量
        $whereFirstValCount = count($firstWhere);
        // 需要更新的第一个字段的值的总数量
        $needUpdateFieldsValCount = count($needUpdateFields[array_key_first($needUpdateFields)]);
        if ($whereFirstValCount !== $needUpdateFieldsValCount) return false;
        // 所有的条件字段数组
        $whereKeys = array_keys($where);

        // 绑定参数
        $building = [];

//        $whereArr = [
//          0 => "id = 180 AND ",
//          1 => "user_id = 5 AND ",
//          2 => "id = 181 AND ",
//          3 => "user_id = 15 AND ",
//          4 => "id = 182 AND ",
//          5 => "user_id = 11 AND ",
//          6 => "id = 183 AND ",
//          7 => "user_id = 1 AND ",
//        ]
        $whereArr = [];
        $whereBuilding = [];
        foreach ($firstWhere as $k => $v) {
            foreach ($whereKeys as $whereKey) {
//                $whereArr[] = "{$whereKey} = {$where[$whereKey][$k]} AND ";
                $whereArr[] = "{$whereKey} = ? AND ";
                $whereBuilding[] = $where[$whereKey][$k];
            }
        }

//        $whereArray = [
//            0 => "id = 180 AND user_id = 5",
//            1 => "id = 181 AND user_id = 15",
//            2 => "id = 182 AND user_id = 11",
//            3 => "id = 183 AND user_id = 1",
//        ]
        $whereArrChunck = array_chunk($whereArr, count($whereKeys));
        $whereBuildingChunck = array_chunk($whereBuilding, count($whereKeys));

        $whereArray = [];
        foreach ($whereArrChunck as $val) {
            $valStr = '';
            foreach ($val as $vv) {
                $valStr .= $vv;
            }
            // 去除掉后面的 AND 字符及空格
            $whereArray[] = rtrim($valStr, "AND ");
        }

        // 需要更新的字段数组
        $needUpdateFieldsKeys = array_keys($needUpdateFields);

        // 拼接 sql 语句
        $sqlStr = '';
        foreach ($needUpdateFieldsKeys as $needUpdateFieldsKey) {
            $str = '';
            foreach ($whereArray as $kk => $vv) {
//                $str .= ' WHEN ' . $vv . ' THEN ' . $needUpdateFields[$needUpdateFieldsKey][$kk];
                $str .= ' WHEN ' . $vv . ' THEN ? ';
                // 合并需要绑定的参数
                $building[] = array_merge($whereBuildingChunck[$kk], [$needUpdateFields[$needUpdateFieldsKey][$kk]]);
            }
            $sqlStr .= $needUpdateFieldsKey . ' = CASE ' . $str . ' ELSE ' . $needUpdateFieldsKey . ' END, ';
        }

        // 去除掉后面的逗号及空格
        $sqlStr = rtrim($sqlStr, ', ');

        $tblSql = 'UPDATE ' . $tableName . ' SET ';

        $tblSql = $tblSql . $sqlStr;

        $building = array_reduce($building,"array_merge",array());
//        return [$tblSql, $building];
        return \DB::update($tblSql, $building);
    }
}

```
