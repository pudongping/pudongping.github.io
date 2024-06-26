---
title: Go 基础数据类型
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
abbrlink: 1308b152
date: 2024-05-09 01:25:27
img:
coverImg:
password:
summary:
---

# Go 基础数据类型

对于浮点类型需要被自动推导的变量，其类型将被自动设置为 float64，而不管赋值给它的数字是否是用 32 位长度表示的  
在实际开发中，应该尽可能地使用 float64 类型，因为 math 包中所有有关数学运算的函数都会要求接收这个类型。


## Go 支持的数据类型

### 基本数据类型

- 布尔类型：bool
- 整型： int8、byte、int16、int、uint、uintptr 等
- 浮点类型：float32（单精度浮点数，可以精确到小数点后 7 位）、float64（双精度浮点数，可以精确到小数点后 15 位）
- 复数类型：complex64、complex128
- 字符串：string
- 字符类型：rune
- 错误类型：error

### 复合类型

- 指针：pointer
- 数组：array
- 切片：slice [切片动态操作图片示例](https://ueokande.github.io/go-slice-tricks/)
- 字典：map
- 通道：chan
- 结构体：struct
- 接口：interface

#### 整型

类型	| 长度（单位：字节）	| 说明 | 值范围 | 默认值
--- | --- | --- | --- | ---
int8 |  1	 |  带符号8位整型	 |  -128~127	 |  0
uint8	 |  1	 |  无符号8位整型，与 byte 类型等价	 |  0~255	 |  0
int16	 |  2	 |  带符号16位整型 |  	-32768~32767	 |  0
uint16	 |  2	 |  无符号16位整型	 |  0~65535 |  	0
int32	 |  4	 |  带符号32位整型，与 rune 类型等价	 |  -2147483648~2147483647 |  	0
uint32	 |  4	 |  无符号32位整型	 |  0~4294967295 |  	0
int64	 |  8	 |  带符号64位整型 |  	-9223372036854775808~9223372036854775807	 |  0
uint64	 |  8	 |  无符号64位整型	 |  0~18446744073709551615	 |  0
int	 |  32位或64位	 |  根据宿主机的机器字长决定（int 和 uint 是一样的大小）	 |  32 位的机器就是 int32，64 位就是 int64	 |  0
uint	 |  32位或64位 |  	根据宿主机的机器字长决定（int 和 uint 是一样的大小）	 |  32 位的机器就是 int32，64 位就是 int64 |  	0
uintptr |  	与对应指针相同 |  	无符号整型，足以存储指针值的未解释位	 |  32位平台下为4字节，64位平台下为8字节 |  0
