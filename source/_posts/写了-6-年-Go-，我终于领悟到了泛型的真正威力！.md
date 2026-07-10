---
title: 写了 6 年 Go ，我终于领悟到了泛型的真正威力！
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
  - 泛型
abbrlink: 2dbf6c6a
date: 2026-07-10 10:32:08
img:
coverImg:
password:
summary:
---

哈喽，各位小伙伴们！今天来聊一聊 Go 语言中的泛型～哈哈，当然这篇文章的标题的确有点儿夸张哈，这篇文章旨在详细介绍一下 Go 语言中的泛型以及让没有用过 Go 语言泛型的童鞋们快速掌握如何使用泛型。

好了，废话不多说，直接开聊～

在使用 Go 语言进行日常的开发工作中，你是不是经常遇到这样一个让人头疼的场景：为了处理不同类型的数据，你不得不把一段逻辑一模一样的代码复制好几遍，但是归根结蒂其实它们仅仅因为它们的**数据类型不一样**而已。

学习一门语言的新特性，最好的方法就是从“它解决了什么痛点”开始。

## 没有泛型会怎样？

要理解为什么我们需要泛型，首先得知道没有泛型的时候我们得有多麻烦

假设我们现在有一个非常简单的需求，要写一个非常简单的函数：比较两个“数字类型”的大小，返回最大的那一个。

如果是整型（ `int` ），我们就会这样写如下代码：

```go
// 比较两个 int 类型的大小
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

这确实没啥问题，但是如果我们遇到另外一个场景，如果我们需要比较浮点数（ `float64` ）的大小呢？

那么，你就只能含泪再写一个：

```go
// 比较两个 float64 类型的大小
func MaxFloat64(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
```

淦！那么问题来了，如果还有 `float32`、`int32`、`int64` 呢？可能你就得复制粘贴数次，然后更改一下数据类型。

有些童鞋就说了，复制粘贴一下又不麻烦，是啊，确实复制粘贴一下不麻烦，但是你顶不住同类型的方法多啊，假设又有了另外一个需求：打印出一个切片中所有的元素。

那么，你的代码就会类似这样写：

```go
package main

import "fmt"

func PrintInts(s []int) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

如果需要打印其他类型的切片，你又得反反复复的复制粘贴……

那么，有没有解决方案呢？

既然聊到了，拿肯定是有的。

## Go1.18 版本之前的权宜之计：空接口 `interface{}`

在 Go 1.18 之前，有经验的开发童鞋们通常会用空接口 `interface{}` 和反射机制来解决这个问题。

```go
// 使用空接口处理多种类型
func MaxInterface(a, b interface{}) interface{} {
    // 这里需要大量的类型断言
    // 并且如果传入了不能比较的类型（比如 map ），程序在运行时（ Runtime ）会直接 Panic 崩溃！
    // 代码需要写的太长且不安全，这里由于篇幅原因就不完整演示了
    return nil
}
```

**空接口的致命缺点：**

1. 丧失了类型安全：编译器在编译阶段无法帮你检查错误，如果传错了类型，程序运行时才会报错崩溃
2. 性能损耗：反射和类型断言会带来额外的性能开销（但是绝对不能说因为性能问题，就不能使用）
3. 代码极其臃肿：大量的 `switch x.(type)` 让人看得眼花缭乱

## 有了泛型之后

那么，什么是泛型咧？

泛型，简单来说就是 **“类型的参数化”**，将数据类型也作为参数传递进去。

其实我们观察文章开头写的代码，你会发现一个规律：它们除了数据类型不一样以外，代码逻辑都是一样的！

那么，现在有了泛型之后，就可以将这类方法进行整合成一个方法就完事了。

以前我们写函数，参数是具体的变量（比如 `a = 10` ）；现在用泛型，**数据类型本身也变成了一个参数**（比如告诉函数“我现在要用 `int` 类型来执行你”）。

## 一步步教你写出第一个泛型函数

说再多文字内容，都不如写点儿代码来得实在。还是以上面我们讨论的比较两个数字大小的需求来说

现在，我们用泛型来重写前面的 `Max` 函数

先看代码，然后我们逐一拆解：

```go
package main

import "fmt"

// 自定义一个类型约束，表示这几种数字类型都可以用
type Number interface {
    int | int32 | int64 | float32 | float64
}

// 这就是一个泛型函数！
// [T Number] 是类型参数列表，表示 T 可以是 Number 里的任何一种类型
func Max[T Number](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    // 传入 int 类型
    fmt.Println("最大的 int 是:", Max(10, 20))         // 输出: 20
    
    // 传入 float64 类型
    fmt.Println("最大的 float64 是:", Max(3.14, 2.71)) // 输出: 3.14
}
```

核心语法拆解（敲黑板，重点来了！）

我们仔细看上面的这段代码： `func Max[T Number](a, b T) T`

1. `[T Number]` ：这叫做 **类型参数列表**。它紧跟在函数名后面，用方括号 `[]` 括起来。

* `T` 是我们给类型起的“代号”（你叫 `A` 、 `MyType` 都可以，你可以就把它理解成是一个占位符就好了，但约定俗成用 `T` 代表 Type ）。
* `Number` 叫做**类型约束**，它告诉编译器：“这个 `T` 不能是阿猫阿狗，它必须是 `Number` 里定义的那几种数字类型之一”。

2. `(a, b T)` ：普通的函数参数列表。表示形参 `a` 和 `b` 的类型都是那个 `T` 数据类型 。
3. 最后的 `T` ：表示函数的返回值类型也是 `T` 。

当你调用 `Max(10, 20)` 的时候， Go 编译器会自动推断你传入的是 `int` ，于是它会在底层自动把 `T` 替换成 `int` ，并执行对比逻辑。

好了，通过上面的简单了解，是不是对泛型已经有了初步认知了？

## 最常用的内置约束： `any` 和 `comparable`

其实在实际开发中，你可能不需要每次都像上面那样自定义一个 `Number` 。 Go 语言已经内置了一些常用的类型约束。

### 1. `any` 约束：任何数据类型都能接收

`any` 其实就是 `interface{}` 的别名。

```go
// 打印任何类型的切片（ Slice ）
// [T any] 表示 T 可以是任意类型
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Printf("%v ", v)
    }
    fmt.Println()
}

func main() {
    PrintSlice([]int{1, 2, 3})          // 输出: 1 2 3
    PrintSlice([]string{"Go", "泛型"})  // 输出: Go 泛型
}
```

### 2. `comparable` 约束：只接受能用 `==` 和 `!=` 比较的类型

如果你需要判断两个值是否相等（比如在 Map 中查找对应的 Key ），你就必须用 `comparable` 。

它支持数字、字符串、布尔值等，但**不支持切片（ Slice ）和字典（ Map ）**， 因为它们不能用 `==` 直接比较。

```go
// 查找元素在切片中的索引，找不到返回 -1
// [T comparable] 表示 T 必须是能够使用 == 比较的类型
func FindIndex[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target { // 因为有了 comparable 约束，这里才能用 ==
            return i
        }
    }
    return -1
}
```

如果你能够看懂上面的代码，恭喜你，你已经学会了 `slices` 包中的 `slices.Index()` 方法。

## 不止是函数，还有泛型类型！

泛型不仅能用在函数上，也能用在 `struct`（结构体）上。这在定义通用的数据容器（比如栈、队列、链表）时非常有用。

比如现在我们要实现一个通用的“栈（ Stack ）”，它可以存放各种类型的数据，就可以如下这么写代码：

```go
// 定义一个泛型结构体
// [T any] 表示这个栈可以存放任何类型的数据
type Stack[T any] struct {
    elements []T
}

// 给泛型栈添加 Push 方法 (入栈)
// 注意接收者要写成 *Stack[T]
func (s *Stack[T]) Push(value T) {
    s.elements = append(s.elements, value)
}

// 给泛型栈添加 Pop 方法 (出栈)
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T // 声明一个 T 类型的零值
        return zero, false
    }
    // 获取最后一个元素
    lastIndex := len(s.elements) - 1
    value := s.elements[lastIndex]
    // 缩容切片
    s.elements = s.elements[:lastIndex]
    return value, true
}

func main() {
    // 实例化一个专门存放 int 的栈
    var intStack Stack[int]
    intStack.Push(100)
    
    // 实例化一个专门存放 string 的栈
    var stringStack Stack[string]
    stringStack.Push("Hello")
}
```

通过上面这段代码，我们就用一份逻辑，轻松创建了不同类型的栈数据结构，是不是现在代码看起来就灰常的nice了？

## 总结

1. **为什么需要泛型**： 为了减少重复代码，提高类型安全，告别 `interface{}` 带来的运行时风险
2. **基本语法**： 在函数名或类型名后加上 `[T 约束类型]`
3. **约束类型**： 可以使用 `any` （任意类型，其实就是 interface{}）、 `comparable` （可判等类型），或者用 `interface` 用 `|` 符号自定义类型集合。

在需要处理通用逻辑、编写底层工具库或通用数据结构时，使用泛型就可以写出更加优雅的代码。

现在你应该对泛型有一定认知了吧？如果这篇文章对你有帮助，别忘了点个**赞**和**在看**哈～

最后抛出一个问题供大家讨论：

```go
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
       ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
       ~float32 | ~float64
}
```

为什么有时候定义一个泛型约束可以定义成这样？` ~int` 前面的 `~` 表示什么含义？为什么有时候又不需要呢？

这里我就抛砖引玉一下，欢迎各位风流倜傥、玉树临风的靓仔们一起聊一聊～
