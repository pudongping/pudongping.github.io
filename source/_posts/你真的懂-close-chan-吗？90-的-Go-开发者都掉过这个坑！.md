---
title: 你真的懂 close(chan) 吗？90% 的 Go 开发者都掉过这个坑！
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
abbrlink: 5e6d50e0
date: 2025-08-08 11:08:45
img:
coverImg:
password:
summary:
---

在日常 Go 并发编程中，我们可能会看到类似以下这样的代码：

```go
// 初始化一个空的 channel，然后在某个位置直接关闭它
// 接收方可以无阻塞地读取到 "完成" 信号
done := make(chan struct{})

close(done)

select {
case <-done:
    fmt.Println("收到完成信号")
default:
    fmt.Println("还没完成")
}
```

这段代码看上去很奇怪：**通道创建了没有任何写入就关闭？和我们常写的 `done <- struct{}{}` 有什么不同？**

这到底是为了什么？这种写法背后有何妙用？

这是个很细节但非常重要的 Go 并发问题，理解它能让你在处理 channel 时避免很多“意料之外”的坑 👀。

---

## ✅ 我们直接先说结论：

> 有些情况下，我们**确实可以**在初始化一个 channel 后**在不发送数据的情况下去 `close(chan)`**，这是**合法且常见的用法**，尤其用于：
>
> * 表示“空数据”或“任务已完成”
> * 用于广播型通知
> * 用于提前返回的通道
> * 用作 `select` 的结束信号

## 🔁 通道基础复习：读、写、关闭

在深入之前，我们先快速复习 Go 通道的几个基本行为：

```go
c := make(chan int, 1)
c <- 42          // 写入数据
v := <-c         // 读取数据（v == 42）
close(c)         // 关闭通道
v, ok := <-c     // ok == false，表示通道已关闭
```

✅ Channel 的三种状态：

| 状态      | 会阻塞吗？ | 返回什么？       | 备注   |
| ------- | ----- | ----------- | ---- |
| 通道有值    | ✅ 不阻塞 | 读到数据        | 正常读写 |
| 通道空未关闭  | ❌ 阻塞  | 无           | 等待写入 |
| 通道空且已关闭 |  ✅ 不阻塞 | 零值，ok=false | 用于通知 |


## 🔍 举个最简单的例子：

```go
done := make(chan struct{})
close(done)

<-done // 可以正常读取，不会阻塞（读到零值）
```

这个通道被关闭了，但你还能从中“读出”零值，这就是 Go 语言通道的一个特殊行为：


### 通道关闭后的行为：

> ❗ 从一个 **已关闭但未被清空的通道** 中读取数据，会：
>
> * 返回已缓冲的数据（如果有）
> * 如果没有数据了，**返回通道元素类型的零值**，且 `ok == false`

举个栗子🌰：

```go
ch := make(chan int, 1)
ch <- 42
close(ch)

v, ok := <-ch
fmt.Println(v, ok) // 输出: 42 true

v2, ok2 := <-ch
fmt.Println(v2, ok2) // 输出: 0 false
```

你应该记住这个核心原则：

> **关闭的通道，后续读取不会阻塞，而是返回通道类型的零值，并携带 `ok=false` 表示通道已关闭。**


## ✍️ 一个对比例子：发送 vs 关闭

我们先用一个例子对比一下“写数据”和“关闭通道”的差异。

### 方式一：通过发送数据通知

```go
done := make(chan struct{})

go func() {
    // 做一些事情
    time.Sleep(time.Second)
    done <- struct{}{}
}()

<-done
fmt.Println("收到通知，继续执行")
```

### 方式二：通过关闭通道通知

```go
done := make(chan struct{})

go func() {
    // 做一些事情
    time.Sleep(time.Second)
    close(done) // 只关闭，不发送
}()

<-done
fmt.Println("收到通知，继续执行")
```

从使用者的角度看，这两段代码好像没有区别，但它们背后的机制完全不同！

---

## 🚨 哪个更安全？

### 🧨 问题：多个 goroutine 能同时读吗？

如果你这样写：

```go
done := make(chan struct{})

for i := 0; i < 10; i++ {
    go func() {
        <-done
        fmt.Println("goroutine 收到退出通知")
    }()
}
```

现在你用哪种方式通知所有 goroutine？

#### ❌ 错误方式：

```go
done <- struct{}{} // 只能唤醒一个 goroutine，其余会阻塞
```

#### ✅ 正确方式：

```go
close(done) // 所有 goroutine 同时唤醒
```

### ✅ 结论：

> 如果你需要**广播通知多个 goroutine**，请使用 `close(chan)`，不要使用 `<- chan` 单播写法。

## ⚖️ 两种方式对比

| 区别点       | `close(done)` ✅推荐方式     | `done <- struct{}{}` ❌不推荐替代 |
| --------- | ----------------------- | --------------------------- |
| 类型        | 关闭信号广播                  | 数据发送                        |
| 是否阻塞      | 永远非阻塞（任何 goroutine 都能读） | 写后只有一个 goroutine 能读，且只能读一次  |
| 多个监听者     | 支持多个 goroutine 并发读取     | 只能有一个 goroutine 成功读取        |
| 多次 select | 永远能读取（返回零值）             | 第一次读后通道就空了，再读就阻塞            |
| panic 风险  | 安全（可多次读取）               | 如果没人读或者写多次就会 panic（死锁或写已满）  |
| 是否语义清晰    | ✅ 代表「完成 / 停止」语义         | ❌ 容易误导为发送具体数据               |

---

## 📍 经典应用场景：优雅关闭 goroutine

假设你有一个 worker，在后台监听数据通道并处理，我们希望在主进程结束时通知它退出。

```go
func worker(stop <-chan struct{}) {
    for {
        select {
        case <-stop:
            fmt.Println("worker 接收到停止信号，退出")
            return
        default:
            fmt.Println("worker 正在工作...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
```

主函数中：

```go
func main() {
    stop := make(chan struct{})
    go worker(stop)

    time.Sleep(2 * time.Second)
    close(stop) // 通知 worker 退出

    time.Sleep(time.Second)
}
```

**关键点：**

* 我们没有发送任何数据到 `stop`，只用了 `close(stop)` 作为通知机制。
* goroutine 在 `<-stop` 中非阻塞地检测到通道已关闭，从而退出。

---

## 🔬 为什么 close(chan) 能用来“广播”？

这是因为在 Go 的底层实现中，所有阻塞在 `<-chan` 的 goroutine 都会在通道关闭时被唤醒，而且不会 panic。这使得 `close(chan)` 成为**低成本的事件广播机制**。

> **一句话总结：**
>
> `chan <- value` 是“发消息”，`close(chan)` 是“发通知”。

---

## 🎯 struct{} 为什么是通用信号类型？

你会发现我们一直使用的是 `chan struct{}`，而不是 `chan bool` 或 `chan int`。

这是因为：

1. `struct{}` 是 Go 中占用空间最小的类型（0 字节）
2. 作为信号通道，它表示“有没有信号”，不需要实际数据
3. 在工具和库中（如 `context.Context.Done()`），也都采用了 `chan struct{}` 作为通知手段



## 🔁 所以，什么时候会主动 close 一个刚初始化的通道？

### 1. ✅ 用于返回“空结果”的通道

```go
func getEmptyChannel() <-chan int {
	ch := make(chan int)
	close(ch) // 表示：这个通道里不会有任何值
	return ch
}
```

这种常用于提前终止 / 空数据返回的并发场景。

---

### 2. ✅ 用作广播通知（常见于 `select`）

```go
done := make(chan struct{})
close(done)

select {
case <-done:
    fmt.Println("收到完成信号")
default:
    fmt.Println("还没完成")
}
```

这是非阻塞通知的经典写法（关闭通道可被所有 goroutine 同时读到，不用写数据）。

---

### 3. ✅ 作为提前取消信号

```go
cancel := make(chan struct{})
close(cancel) // 提前取消任务

go func() {
    select {
    case <-cancel:
        fmt.Println("任务被取消了")
    case <-time.After(2 * time.Second):
        fmt.Println("任务完成")
    }
}()
```

`close(chan)` 是 Go 中的一种 **轻量级广播机制**，相比发送数据更高效也更干净。

## 🏁 写在最后

`close(chan)` 虽然简单，但却是 Go 并发设计中最优雅、最高效的一种“广播机制”。
理解它不仅能让你的 goroutine 更加优雅退出，还能帮你构建稳定的并发系统。



