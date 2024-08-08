---
title: Go语言中如何连接 MySQL，基础必备！
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
abbrlink: 48b3e862
date: 2024-08-09 01:10:47
img:
coverImg:
password:
summary:
---

在现代应用中，数据库操作是必不可少的一部分，而 Go 语言凭借其高效性和并发处理能力，成为了越来越多开发者的选择。

在本教程中，我们将学习如何使用 Go 语言与 MySQL 数据库进行基本的 CRUD（创建、读取、更新、删除）操作。我们将使用 `database/sql` 标准库以及 `github.com/go-sql-driver/mysql` 驱动来实现这些功能。

## 准备工作

### 环境要求

- 创建一个名为 `test_db` 的数据库。
- 创建一个 `users` 表，表结构如下：

```bash
CREATE DATABASE test_db;

USE test_db;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL
);
```

### 安装 MySQL 驱动

使用以下命令安装 MySQL 驱动：

```bash
go get -u github.com/go-sql-driver/mysql
```

## 示例代码

下面是一个完整的 Go 程序，展示如何进行基本的 CRUD 操作：

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // 数据库连接字符串
    dsn := "user:password@tcp(127.0.0.1:3306)/test_db"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    // 防止 db 为 nil，因此需要先判断 err 之后才能调用 Close 方法去释放 db
    defer db.Close()

    // 测试连接
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("成功连接到 MySQL 数据库！")

    // 创建用户
    id := createUser(db, "Alice", 25)
    fmt.Printf("新创建的用户 ID: %d\n", id)

    // 查询所有用户
    getUsers(db)

    // 查询单个用户
    userId := 1
    getUser(db, userId)

    // 更新用户
    affectedRows := updateUser(db, 1, "Alice Smith", 26)
    fmt.Printf("受影响的行数: %d\n", affectedRows)

    // 删除用户
    affectedRows = deleteUser(db, 1)
    fmt.Printf("受影响的行数: %d\n", affectedRows)
}
```

### 创建（Insert）

```go
// 创建用户，返回新创建用户的 ID
func createUser(db *sql.DB, name string, age int) int {
    query := "INSERT INTO users (name, age) VALUES (?, ?)"
    result, err := db.Exec(query, name, age)
    if err != nil {
        log.Fatal(err)
    }
    // 获取最后插入的 ID
    lastInsertId, err := result.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    return int(lastInsertId)
}
```

### 查询（Select）

```go
// 查询所有用户
func getUsers(db *sql.DB) {
    query := "SELECT id, name, age FROM users"
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    // 一定要记得关闭资源
    defer rows.Close()

    fmt.Println("所有用户:")
    for rows.Next() {
        var id int
        var name string
        var age int
        if err := rows.Scan(&id, &name, &age); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }
}

// 查询单个用户
func getUser(db *sql.DB, id int) {
    query := "SELECT id, name, age FROM users WHERE id = ?"
    row := db.QueryRow(query, id)

    var name string
    var age int
    // 调用完了 QueryRow 方法之后，一定要记得调用 Scan 方法，否则持有的数据库连接不会被释放
    if err := row.Scan(&id, &name, &age); err != nil {
        if err == sql.ErrNoRows {
            fmt.Printf("用户 ID %d 不存在\n", id)
        } else {
            log.Fatal(err)
        }
    } else {
        fmt.Printf("用户 ID: %d, Name: %s, Age: %d\n", id, name, age)
    }
}
```

### 更新（Update）

```go
// 更新用户，返回受影响的行数
func updateUser(db *sql.DB, id int, name string, age int) int {
    query := "UPDATE users SET name = ?, age = ? WHERE id = ?"
    result, err := db.Exec(query, name, age, id)
    if err != nil {
        log.Fatal(err)
    }
    // 获取受影响的行数
    affectedRows, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    return int(affectedRows)
}
```

### 删除（Delete）

```go
// 删除用户，返回受影响的行数
func deleteUser(db *sql.DB, id int) int {
    query := "DELETE FROM users WHERE id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        log.Fatal(err)
    }
    // 获取受影响的行数
    affectedRows, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    return int(affectedRows)
}
```

## 代码解析

### 1. 数据库连接

我们使用 `sql.Open` 方法连接到 MySQL 数据库，连接字符串格式为 `user:password@tcp(host:port)/dbname`。在连接后，我们调用 `db.Ping()` 测试数据库连接是否成功。

### 2. 创建用户

`createUser` 函数通过 `INSERT` 语句向 `users` 表中添加新用户，并返回新创建用户的 ID。

### 3. 查询用户

- `getUsers` 函数查询所有用户并打印每个用户的 ID、姓名和年龄。
- `getUser` 函数通过用户 ID 查询单个用户，并处理用户不存在的情况。

### 4. 更新用户

`updateUser` 函数用于更新用户信息，并返回受影响的行数，以确认操作是否成功。

### 5. 删除用户

`deleteUser` 函数用于删除指定 ID 的用户，并同样返回受影响的行数。

## 总结

在本文中，我们展示了如何使用 Go 语言与 MySQL 数据库进行基本的 CRUD 操作。通过本教程，您可以创建、查询、更新和删除用户数据，这为您在开发基于数据库的应用程序时打下了基础。

用这个库还是比较累赘的，代码写起来比较累，并且一般在实际应用中很少会直接拼接 sql 语句，会遇到 SQL 注入的风险，后面会介绍 sqlx 或者 gorm 的使用，这里先了解个基础知识，为后期做做准备。