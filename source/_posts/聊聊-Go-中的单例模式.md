---
title: 聊聊 Go 中的单例模式
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
abbrlink: 1b8de6f
date: 2024-07-08 12:10:43
img:
coverImg:
password:
summary:
---

单例模式在软件开发中是一种常见的设计模式，用于确保一个类在任何情况下都仅有一个实例，并提供一个访问它的全局访问点。

在 Go 语言中，实现单例模式通常有两种方式：**饿汉式** 和 **懒汉式**。

今天，我们就来详细了解这两种实现方式，并通过简单易懂的代码示例解释相关概念。

## 饿汉式单例模式

饿汉式单例模式的核心思想是：类加载时就创建实例。由于 Go 语言不同于 Java，没有显式的类概念，我们通常使用结构体来模拟类的行为。下面是一个饿汉式单例模式的实现示例：

```go
// 饿汉式单例模式
package main

type singleton struct {
    count int
}

// 饿汉式单例，程序启动即初始化
var Instance = new(singleton)

// Add 方法用于累加并返回计数值
func (s *singleton) Add() int {
    s.count++
    return s.count
}
```

在这个例子中，我们定义了一个 `singleton` 结构体，并在程序启动时通过 `var` 声明即初始化了 `Instance`。这样就保证了 `Instance` 是全局唯一的，并且在第一次使用前就已经准备好了。

## 懒汉式单例模式

与饿汉式相比，懒汉式单例模式在第一次需要时才创建实例，可以延迟初始化资源。这在某些情况下可以节省资源，但需要考虑并发环境下的线程安全问题。

在 Go 语言中，可以使用[双重检查锁定模式](https://zh.wikipedia.org/wiki/%E5%8F%8C%E9%87%8D%E6%A3%80%E6%9F%A5%E9%94%81%E5%AE%9A%E6%A8%A1%E5%BC%8F) （Double-checked Locking）来解决线程安全问题。

下面是懒汉式单例模式的实现示例：

```go
// 懒汉式单例模式
package main

import (
	"sync"
)

type singleton struct {
	count int
}

var (
	instance *singleton
	mutex    sync.Mutex
)

// New 实例化一个对象
// 这里采用了【双重检查】
// 假设 goroutine X 和 Y 几乎同时调用 New 函数
// 当它们同时进入此函数时，instance 变量值是 nil 因此 goroutine X 和 Y 会同时到达【位置1】
// 假设 goroutine X 会先到达【位置2】，并进入 mutex.Lock() 到达【位置3】，这时，由于 mutex.Lock() 的同步限制
// goroutine Y 无法到达【位置3】只能在【位置2】等候
// goroutine X 执行 instance = new(singleton) 语句，使得 instance 变量得到一个值，此时 goroutine Y 还是只能在【位置2】等候
// goroutine X 释放锁，返回 instance 变量，退出 New 函数
// goroutine Y 进入 mutex.Lock() 到达【位置3】，进而到达【位置4】。由于此时 instance 变量已经不是 nil，因此 goroutine Y 释放锁
// 可见，锁仅用来避免多个 goroutine 同时实例化 singleton
func New() *singleton {
	if instance == nil { // 【位置1】
		// 这里可能有多于一个 goroutine 同时到达 【位置2】
		mutex.Lock()
		// 这里每个时刻只会有一个 goroutine 到达  【位置3】
		if instance == nil { // 【位置4】
			instance = new(singleton)
		}
		mutex.Unlock()
	}

	return instance
}

func (s *singleton) Add() int {
	s.count++
	return s.count
}
```

在这个例子中，我们使用 `mutex` 来保护 `instance` 的创建过程，确保即使在多个 goroutine 同时调用 `New()` 时，实例也只会被创建一次。这种方法称为“双重检查”，因为每次调用 `New()` 时会进行两次 `instance` 是否为 `nil` 的检查：一次在加锁前，一次在加锁后。

## 双重检查锁定模式
双重检查锁定模式是一种优化，它避免了在每次访问实例时都要进行同步操作的开销。这种模式首先检查实例是否已经创建，如果没有，则进行同步。在同步块内部，再次检查实例是否创建，以确保即使多个 goroutine 同时进入同步块，也只有一个能够创建实例。

## 小结

单例模式在需要全局访问点且只希望创建一个实例的场景下非常有用。饿汉式单例模式简单但可能造成资源浪费，而懒汉式单例模式则更加灵活，但需要处理线程安全问题。Go 语言的并发特性使得实现懒汉式单例模式时，双重检查锁定模式成为了一个优雅的解决方案。

通过以上的介绍和代码示例，相信你已经对饿汉式和懒汉式单例模式有了基本的了解和认识。在实际开发中，根据具体情况选用适当的实现方式，是每个 Go 开发者需要考虑的问题。