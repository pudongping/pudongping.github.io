---
title: 探索 Go 语言中的内存对齐：为什么结构体大小会有所不同？
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
abbrlink: 8f8634cd
date: 2024-11-22 11:00:10
img:
coverImg:
password:
summary:
---

在 Go 语言中，内存对齐是一个经常被忽略但非常重要的概念。理解内存对齐不仅可以帮助我们写出更高效的代码，还能避免一些潜在的性能陷阱。

在这篇文章中，我们将通过一个简单的例子来探讨 Go 语言中的内存对齐机制，以及为什么相似的结构体在内存中会占用不同的大小。

## 示例代码

我们先来看一段代码：

```go
package memory_alignment

import (
	"fmt"
	"unsafe"
)

type A struct {
	a int8
	b int8
	c int32
	d string
	e string
}

type B struct {
	a int8
	e string
	c int32
	b int8
	d string
}

func Run() {
	var a A
	var b B
	fmt.Printf("a size: %v \n", unsafe.Sizeof(a))
	fmt.Printf("b size: %v \n", unsafe.Sizeof(b))
	// a size: 40
	// b size: 48
}
```

在这个例子中，我们定义了两个结构体 `A` 和 `B`。它们的字段基本相同，只是排列顺序不同。然后，我们使用 `unsafe.Sizeof` 来查看这两个结构体在内存中的大小。

结果却令人惊讶：结构体 `A` 的大小是 40 字节，而结构体 `B` 的大小是 48 字节。为什么会出现这样的差异呢？这就是我们今天要讨论的内存对齐的作用。

## 内存对齐概念

内存对齐是指编译器为了优化内存访问速度，而对数据在内存中的位置进行调整的一种策略。不同类型的数据在内存中的对齐要求不同，例如：

- `int8` 类型的变量通常对齐到 1 字节边界。
- `int32` 类型的变量通常对齐到 4 字节边界。
- 指针（如 `string`）通常对齐到 8 字节边界。

为了满足这些对齐要求，编译器可能会在结构体的字段之间插入一些“填充”字节，从而确保每个字段都能正确对齐。

## 结构体内存布局解析

让我们深入分析一下 `A` 和 `B` 两个结构体的内存布局，看看编译器是如何为它们分配内存的。

### 结构体 A 的内存布局

```plaintext
| a (int8) | b (int8) | padding (2 bytes) | c (int32) | d (string, 8 bytes) | e (string, 8 bytes) |
```

- `a` 和 `b` 是 `int8` 类型，各占 1 字节。
- `c` 是 `int32` 类型，需要 4 字节对齐，`b` 后面会有 2 个填充字节。
- `d` 和 `e` 是 `string` 类型，各占 8 字节。

总大小为：1 + 1 + 2 + 4 + 8 + 8 = 24 字节。

### 结构体 B 的内存布局

```plaintext
| a (int8) | padding (7 bytes) | e (string, 8 bytes) | c (int32) | padding (4 bytes) | b (int8) | padding (3 bytes) | d (string, 8 bytes) |
```

- `a` 是 `int8` 类型，占 1 字节，后面有 7 个填充字节，以便 `e` 能够对齐到 8 字节边界。
- `c` 是 `int32` 类型，需要 4 字节对齐，因此在 `c` 后面没有填充。
- `b` 是 `int8` 类型，需要填充 3 个字节来对齐到 `d` 的 8 字节边界。

总大小为：1 + 7 + 8 + 4 + 4 + 1 + 3 + 8 = 36 字节。

**请注意**，Go 编译器可能会将 `d` 和 `e` 视为 8 字节对齐类型（取决于系统和编译器的实现），因此总大小可能是 48 字节。

## 如何优化结构体内存布局

为了减少结构体的内存占用，我们可以按照字段的对齐要求来重新排列字段。例如：

- 先声明大的字段（如 `string` 和 `int32`），然后是小的字段（如 `int8`），可以减少内存中的填充字节。

我们可以将 `B` 结构体改成以下形式：

```go
type OptimizedB struct {
    e string
    d string
    c int32
    a int8
    b int8
}
```

这样可以减少内存填充，从而优化内存占用。

## 总结

内存对齐是编译器优化内存访问速度的一个重要策略。虽然它对大多数应用程序的影响可能较小，但在高性能场景或内存受限的环境中，理解并优化内存对齐可能会带来显著的性能提升。

在 Go 语言中，了解结构体的内存对齐规则，合理排列结构体字段顺序，不仅可以提高程序的性能，还能减少内存的浪费。这是一种简单而有效的优化手段，希望大家在以后的编程实践中能够灵活运用。