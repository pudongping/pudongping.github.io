---
title: Go语言中init 函数、变量、常量的执行顺序
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
abbrlink: 40b0fc5c
date: 2024-05-09 01:19:47
img:
coverImg:
password:
summary:
---

# init 函数、变量、常量的执行顺序

![img](https://cdn.learnku.com/uploads/images/202011/11/1/ipbtiTb7Fd.png!large)

- 假如 main 引入了 pkg1 最终依赖于 pkg3，pkg3 中的 init() 方法会优先被执行；
- 同一个包里，单文件的情况，init() 优先于其他方法执行，包括 main()；
- 同一个包里的常量和变量声明会优先于 init() 方法执行；
- 同一个文件里允许多个 init() 存在，会按照自上而下的顺序执行；
- 同一个包，多个文件里存在 init() 的情况，执行顺序是按文件名的字母排序执行。


## 包导入路径优先级

### 如果使用 `govendor` 时

1. 先从项目根目录的 `vendor` 目录下查找
2. 然后从 $GOROOT/src 目录下查找
3. 然后从 $GOPATH/src 目录下查找
4. 都找不到时，报错

### 如果使用 `go modules` 时

1. 如果导入的包有域名，则都会在 $GOPATH/pkg/mod 下查找，找不到就去域名对应的网站下寻找，找不到或者找到的不是一个包，则报错
2. 如果导入的包没有域名，比如 `fmt` 包，则就去 $GOROOT 下找
3. 如果项目根目录下存在 `vendor` 目录，则不管导入的包有没有域名，都只会在 `vendor` 目录中查找

> 通常 `vendor` 目录是通过执行 `go mod vendor` 命令生成的。
