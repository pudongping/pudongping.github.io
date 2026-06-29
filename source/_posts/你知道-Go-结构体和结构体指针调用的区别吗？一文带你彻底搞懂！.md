---
title: 你知道 Go 结构体和结构体指针调用的区别吗？一文带你彻底搞懂！
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
abbrlink: 934e573e
date: 2026-06-29 11:07:27
img:
coverImg:
password:
summary:
---

前几天在技术群里看到有小伙伴问了一个经典问题：**"Go 结构体方法调用时，什么时候用值接收器，什么时候用指针接收器？"**

虽然这是个老生常谈的话题，但是，我发现很多同学（包括工作几年的）对此还是一知半解。正好最近在梳理 Go 基础知识，就写下这篇文章。

我在网上搜索时发现，这个问题从 Go 诞生至今一直困扰着开发者，真是"一代传一代"的经典问题。

## 结构体是什么

在 Go 语言中有个基本类型，我们称之为结构体（struct）。是 Go 语言中非常常用的，基本定义如下：

```go
type struct_variable_type struct {
    member definition
    member definition
    ...
    member definition
}
```

结构体（struct）是 Go 中的一种复合数据类型，可以将不同类型的数据组合在一起。

简单示例：

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    u := User{"张三", 25}
    u.Age = 26
    fmt.Println(u.Age)
}
```

输出结果：

```go
26
```

这部分属于基础知识，因此不再过多解释。如果看不懂，建议重新再去看一看 Go 语言语法基础。

## 结构体和指针调用

讲解前置概念后，直接进入本文主题。如下例子：

```go
type User struct {
    Name string
    Age  int
}

func (u User) SetName1(name string) {
    u.Name = name
}

func (u *User) SetName2(name string) {
    u.Name = name
}
```

上面的代码声明了一个 `User` 结构体，其包含两个结构体方法，分别是 `SetName1` 和 `SetName2` 方法，两者之间的差异就是接收器的类型不同。

那么，这两者有什么区别，什么情况下要用哪种，有没有啥注意事项呢？

## 两者区别

从许多小伙伴的反馈来看，这二者之间确实会让人感到困惑，经常会有人纠结要不要使用 "指针"，又担心性能、内存什么的。

实际上，情况并没那么复杂，看看下面的栗子🌰：

```go
func (u User) SetName1(name string)
func (u *User) SetName2(name string)
```

当在一个类型上定义一个方法时，接收器（在上面的例子中是 `u`）的行为就像它是方法的一个参数一样。其相当于：

> 其实上方的代码就是下方代码的一种语法糖而已。

```go
func SetName1(u User, name string) {
    u.Name = name
}

func SetName2(u *User, name string) {
    u.Name = name
}
```

因此结构体方法是要将接收器定义成值，还是指针。这**本质上与函数参数应该是值还是指针是同一个问题**。是不是看到这里，有同学就已经茅塞顿开了？

### 实际效果对比

让我们通过一个完整的例子来看看两者的实际区别：

```go
package main

import "fmt"

type Counter struct {
    Value int
}

// 值接收器
func (c Counter) IncrementByValue() {
    c.Value++
    fmt.Printf("方法内部值：%d\n", c.Value)
}

// 指针接收器
func (c *Counter) IncrementByPointer() {
    c.Value++
    fmt.Printf("方法内部值：%d\n", c.Value)
}

func main() {
    counter := Counter{Value: 10}
    
    fmt.Printf("初始值：%d\n", counter.Value)
    
    // 使用值接收器
    counter.IncrementByValue()
    fmt.Printf("调用值接收器后：%d\n", counter.Value)
    
    // 使用指针接收器
    counter.IncrementByPointer()
    fmt.Printf("调用指针接收器后：%d\n", counter.Value)
}
```

输出结果：

```
初始值：10
方法内部值：11
调用值接收器后：10
方法内部值：11
调用指针接收器后：11
```

看到了吗？这就是核心区别所在。值类型不可修改，指针类型可直接修改。

## 如何选择？

整体有以下几个考虑因素，按重要程度顺序排列：

**1. 在使用上的考虑：方法是否需要修改接收器？**  如果需要，接收器必须是一个指针。
**2. 在效率上的考虑：**  如果接收器很大，比如：一个大的结构体，使用指针接收器会好很多。
**3. 在一致性上的考虑：**  如果类型的某些方法必须有指针接收器，那么其余的方法也应该有指针接收器，所以无论类型如何使用，方法集都是一致的。

回到上面的例子中，从功能使用角度来看：

- 如果 `IncrementByPointer` 方法修改了 `c` 的字段，调用者是可以看到这些字段值变更的，因为其是指针引用，本质上是同一份。
- 相对 `IncrementByValue` 方法来讲，该方法是用调用者参数的副本来调用的，本质上是值传递，它所做的任何字段变更对调用者来说是看不见的。

另外对于基本类型、切片和小结构等类型，值接收器是非常廉价的。

因此除非方法的语义需要指针，那么值接收器是最高效和清晰的。在内存管理方面，也不需要过度关注。出现问题时再优化就好了。

## Go 的语法糖

可能有小伙伴会疑问："我明明定义的是指针接收器，为什么可以用结构体变量直接调用？"

```go
func main() {
    user := User{Name: "张三"}
    
    // 这样调用为什么不报错？
    user.SetName2("李四")  // SetName2 接收器是 *User
}
```

这是因为 Go 提供了便利的语法糖：

- 当你用结构体变量调用指针接收器方法时，Go 会自动转换为 `(&user).SetName2("李四")`
- 当你用结构体指针调用值接收器方法时，Go 会自动转换为 `(*userPtr).SetName1("王五")`

也就是下面的代码，完全不会报错。

```go
package main

type User struct {
	Name string
	Age  int
}

func (u User) SetName1(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	u := User{Name: "Alice", Age: 30}
	u.SetName1("Bob")
	println(u.Name) // 输出: Alice

	u.SetName2("Charlie")
	println(u.Name) // 输出: Charlie

	u1 := &User{Name: "David", Age: 25}
	u1.SetName1("Eve")
	println(u1.Name) // 输出: David

	u1.SetName2("Frank")
	println(u1.Name) // 输出: Frank
}
```

这个语法糖让代码更加简洁，但最好还是理解一下底层机制。

## 接口实现的坑

这里有一个很重要的细节，容易踩坑：

```go
package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

type Document struct {
	Content string
}

func (d *Document) Print() {
	fmt.Println(d.Content)
}

func main() {
	var p Printer
	doc := Document{Content: "Hello"}

	// 这样会编译错误！
	// p = doc //  Document does not implement Printer (Print method has pointer receiver)

	// 正确的做法
	p = &doc // *Document implements Printer
	p.Print()
}
```

记住：**如果方法是指针接收器，那么只有指针类型才实现了该接口。**

## 性能考量

对于大结构体，性能差异是显著的：

```go
type LargeStruct struct {
    Data [10000]int  // 40KB 的数据
}

func (ls LargeStruct) ProcessByValue() {
    // 每次调用都复制 40KB
}

func (ls *LargeStruct) ProcessByPointer() {
    // 只传递 8 字节指针（64 位系统）
}
```

显然，对于大结构体，指针接收器是更好的选择。

## 总结

在本文中，我们针对 Go 结构体和结构体指针调用有什么区别，这个问题进行了深入浅出的分析和说明。

1. **值接收器**：操作副本，无法修改原始数据，适合小结构体和不需修改的场景
2. **指针接收器**：操作原始数据，可以修改，适合大结构体和需要修改的场景
3. **选择原则**：优先考虑是否需要修改，其次考虑性能，最后考虑一致性

而在本文中所介绍的部分内容，实际上在 Go 官方文档中都有相应说明。这确实是一个被问了无数次的经典问题。

**谁再疑惑这个问题，转发这篇文章，学就完了。**

如果对你有所帮助，那么，就帮忙点个赞支持一下吧～