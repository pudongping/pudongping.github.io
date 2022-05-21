---
title: Golang 使用 CASE WHEN 进行批量更新
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
  - 批量更新
abbrlink: 84fd0244
date: 2022-05-22 02:34:25
img:
coverImg:
password:
summary:
---

# Golang 使用 CASE WHEN 进行批量更新

> 这是使用 `Go` 语言写的 `CASE WHEN` 拼接语句，如果需要 `PHP` 版本的，可以参考我的这篇文章 [PHP 使用 CASE WHEN 进行批量更新 （当前基于 laravel 编写）](https://pudongping.com/posts/eaf6377d.html)

## 以下代码最终返回的 sql 语句为

```mysql

UPDATE articles SET 
view_count = CASE
		WHEN id = 180 AND user_id = 5  THEN 11 
		WHEN id = 181 AND user_id = 15 THEN 22 
		WHEN id = 182 AND user_id = 11 THEN 33 
		WHEN id = 183 AND user_id = 1  THEN 44 
		ELSE view_count END,
updated_at = CASE
		WHEN id = 180 AND user_id = 5  THEN 1653147405 
		WHEN id = 181 AND user_id = 15 THEN 1653147406 
		WHEN id = 182 AND user_id = 11 THEN 1653147407 
		WHEN id = 183 AND user_id = 1  THEN 1653147408 
		ELSE updated_at 
END

```

## 封装成方法为

```go

// BatchUpdate 使用 case when 进行批量更新
// 最终执行的 sql 语句为：
// UPDATE articles SET
// view_count = CASE
//		WHEN id = 180 AND user_id = 5  THEN 11
//		WHEN id = 181 AND user_id = 15 THEN 22
//		WHEN id = 182 AND user_id = 11 THEN 33
//		WHEN id = 183 AND user_id = 1  THEN 44
//		ELSE view_count END,
// updated_at = CASE
//		WHEN id = 180 AND user_id = 5  THEN 1653147405
//		WHEN id = 181 AND user_id = 15 THEN 1653147406
//		WHEN id = 182 AND user_id = 11 THEN 1653147407
//		WHEN id = 183 AND user_id = 1  THEN 1653147408
//		ELSE updated_at
// END
//
//
// tableName := "articles"
// where := make(map[string][]int)
// where["id"] = []int{180, 181, 182, 183}
// where["user_id"] = []int{5, 15, 11, 1}
// needUpdateFields := make(map[string][]int)
// needUpdateFields["view_count"] = []int{11, 22, 33, 44}
// needUpdateFields["updated_at"] = []int{1653147405, 1653147405, 1653147405, 1653147405}
func BatchUpdate(tableName string, where, needUpdateFields map[string][]int) string {
	if len(where) == 0 || len(needUpdateFields) == 0 {
		return ""
	}

	// 所有的条件字段数组
	var whereKeys []string
	for k := range where {
		whereKeys = append(whereKeys, k)
	}
	// 第一个 where 条件所有的值
	firstWhere := where[whereKeys[0]]

	// 所有需要更新的字段数组
	var needUpdateFieldsKeys []string
	for k := range needUpdateFields {
		needUpdateFieldsKeys = append(needUpdateFieldsKeys, k)
	}

	if len(firstWhere) != len(needUpdateFields[needUpdateFieldsKeys[0]]) {
		// 更新的条件与更新的字段值数量不相等
		return ""
	}

	var s1 []string
	for k := range firstWhere {
		for _, vv := range whereKeys {
			s1 = append(s1, fmt.Sprintf("%s = %v AND ", vv, where[vv][k]))
		}
	}

	// 按照 where 条件字段数量做切割
	whereSize := len(whereKeys)
	batches := make([][]string, 0, (len(s1)+whereSize-1)/whereSize)
	for whereSize < len(s1) {
		s1, batches = s1[whereSize:], append(batches, s1[0:whereSize:whereSize])
	}
	batches = append(batches, s1)

	var whereArr []string
	for _, v := range batches {
		whereArr = append(whereArr, strings.TrimSuffix(strings.Join(v, " "), "AND "))
	}

	// 拼接 sql 语句
	sqlStr := ""
	for _, v := range needUpdateFieldsKeys {
		str := ""
		for kk, vv := range whereArr {
			str += fmt.Sprintf(" WHEN %v THEN %v ", vv, needUpdateFields[v][kk])
		}
		sqlStr += fmt.Sprintf("%s = CASE %s ELSE %s END, ", v, str, v)
	}

	// 去除掉最后面的逗号及空格
	sqlStr = strings.TrimSuffix(sqlStr, ", ")

	caseWhenSql := fmt.Sprintf("UPDATE %s SET %s", tableName, sqlStr)

	return caseWhenSql
}

```

## 使用方式

```go

func main() {

	tableName := "articles"
	where := make(map[string][]int)
	where["id"] = []int{180, 181, 182, 183}
	where["user_id"] = []int{5, 15, 11, 1}

	needUpdateFields := make(map[string][]int)
	needUpdateFields["view_count"] = []int{11, 22, 33, 44}
	needUpdateFields["updated_at"] = []int{1653147405, 1653147406, 1653147407, 1653147408}

	sql := BatchUpdate(tableName, where, needUpdateFields)
	fmt.Println(sql)

}

```

一般 orm 都会支持原生 sql 执行，在这里我只返回了拼接好的 sql 语句，最后只需要将 sql 语句拿去执行即可。当然，如果字段值类型需要调整的，则需要按需调整下。
