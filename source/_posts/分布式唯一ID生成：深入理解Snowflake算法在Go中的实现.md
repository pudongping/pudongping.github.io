---
title: 分布式唯一ID生成：深入理解Snowflake算法在Go中的实现
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
abbrlink: 821f69
date: 2024-11-17 15:15:59
img:
coverImg:
password:
summary:
---

在分布式系统中，为了确保每个节点生成的 ID 在整个系统中是唯一的，我们需要一种高效且可靠的 ID 生成机制。

## 分布式 ID 的特点

- 全局唯一性：不能出现有重复的 ID 标识，这是基本要求。
- 递增性：确保生成的 ID 对于用户或业务是递增的。
- 高可用性：确保任何时候都能生成正确的 ID。
- 高性能性：在高并发的环境下依然表现良好。

## 分布式 ID 的应用场景

不仅仅是用于用户 ID，实际互联网中有很多场景都需要能够生成类似 MySQL 自增 ID 这样不断增大，同时又不会重复的 ID，以支持业务中的高并发场景。

比较典型的场景有：电商促销时短时间内会有大量的订单涌入到系统，比如每秒 10W+ 在这些业务场景下将数据插入数据库之前，我们需要给这些订单和数据先分配一个唯一 ID，然后再保存到数据库中。

对这个 ID 的要求是希望其中能带有一些时间信息，这样即使我们后端的系统对数据进行了分库分表，也能够以时间顺序对这些数据进行排序。

Snowflake 算法就是这样的一种算法，它最初由 Twitter 开发，并因其高效、稳定、可扩展等优点，被广泛应用于分布式系统中。

## Snowflake 算法（雪花算法）

Twitter 的分布式 ID 生成算法，是一个经过实践考验的算法，它的核心思想是：使用一个 64 位的 long 型的数字作为全局唯一 ID。在这 64 位中，其中 1 位是不用的，然后用其中的 41 位作为毫秒数，用 10 位作为工作机器 id，12 位作为序列号。

![雪花算法](https://upload-images.jianshu.io/upload_images/14623749-3de70a11645db37c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 1 位标识位：最高位是符号位，正数是 0，负数是 1，生成的 ID 一般是正数，所以为 0。
- 时间戳：占用 41bit，单位为毫秒，总共可以容纳约 69 年的时间。当然，我们的时间毫秒计数不会真的从 1970 年开始计，那样我们的系统跑到 2039 年 9 月 7 日 23:47:35 就不能用了，所以这里的时间戳只是相对于某个时间的增量，比如我们的系统上线是 2024-08-20，那么我们的时间戳就是当前时间减去 2024-08-20 的时间戳，得到的偏移量。
- 机器 id：占用 10bit，其中高位 5bit 是数据中心 ID（datacenterId），低位 5bit 是机器 ID（workerId），可以部署在 2^5=32 个机房，每个机房可以部署 2^5=32 台机器，可以容纳 1024 个节点。
- 序列号：占用 12bit，用来记录同毫秒内产生的不同 ID。每个节点每毫秒开始不断累加，最多可以累加到 4095，同一毫秒一共可以产生 4096 个 ID。

## SnowFlake 算法在同一毫秒内最多可以生成多少个全局唯一 ID 呢？

同一毫秒的 ID 数量 = 1024 * 4096 = 4194304，也就是说在同一毫秒内最多可以生成 4194304 个全局唯一 ID。

## 雪花算法的 Go 语言实现

在本文中，我们将通过 Go 语言的两个库——`bwmarrin/snowflake`和`sony/sonyflake`，来详细探讨如何实现基于 Snowflake 算法的分布式唯一 ID 生成器。

### 1. 使用`bwmarrin/snowflake`生成唯一ID

我们首先使用`bwmarrin/snowflake`库来生成唯一ID。

```go
package snow_flake

import (
    "fmt"
    "reflect"
    "time"

    "github.com/bwmarrin/snowflake"
)

func SnowFlake1() {
	var (
		node *snowflake.Node
		st   time.Time
		err  error
	)

	startTime := "2024-08-20" // 初始化一个开始的时间，表示从这个时间开始算起
	machineID := 1            // 机器 ID

	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	// 根据指定的开始时间和机器ID，生成节点实例
	node, err = snowflake.NewNode(int64(machineID))
	if err != nil {
		panic(err)
	}

	// 生成并输出 ID
	id := node.Generate()

	fmt.Printf("Int64  ID: %d type of: %T -> Type %v -> Value %v \n", id, id, reflect.TypeOf(id), reflect.ValueOf(id))
	fmt.Printf("Int64  ID: %d\n", id.Int64()) // 也可以直接调用 Int64() 方法
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())
}
```

**代码解析：**

1. **时间戳与机器ID**：我们首先定义了一个时间戳和机器ID。这里的时间戳用于记录从特定时间开始的毫秒数，而机器ID则用于区分不同节点。
2. **生成节点实例**：`snowflake.NewNode()`函数根据时间戳和机器ID生成一个节点实例。
3. **生成唯一ID**：使用`node.Generate()`方法生成唯一ID，并展示了多种表示形式。

### 2. 使用`sony/sonyflake`生成唯一ID

接下来，我们来看一下`sony/sonyflake`库的实现。

```go
package snow_flake

import (
    "fmt"
    "reflect"
    "time"

    "github.com/sony/sonyflake"
)

func SnowFlake2() {
	var (
		sonyFlake     *sonyflake.Sonyflake
		sonyMachineID uint16
		st            time.Time
		err           error
	)

	startTime := "2024-08-20" // 初始化一个开始的时间，表示从这个时间开始算起
	machineID := 1            // 机器 ID
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}

	sonyMachineID = uint16(machineID)
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) { return sonyMachineID, nil },
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	if sonyFlake == nil {
		panic("sonyflake not created")
	}

	id, err := sonyFlake.NextID()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Int64  ID: %d type of: %T -> Type %v -> Value %v \n", id, id, reflect.TypeOf(id), reflect.ValueOf(id))
}
```

**代码解析：**

1. **StartTime**：类似于`snowflake`，我们通过`StartTime`设置了ID生成的起始时间。
2. **MachineID**：通过`Settings`结构体的`MachineID`字段指定机器ID的获取方式。

## 选择哪个库？

`bwmarrin/snowflake`和`sony/sonyflake`都提供了基于Snowflake算法的分布式唯一ID生成器。选择哪个库取决于你的需求：

- **`bwmarrin/snowflake`**：成熟、广泛应用，如果你需要生成不同进制的ID（如Base2, Base64）或对时间戳的精度要求更高，可以选择这个库。
- **`sony/sonyflake`**：优化了一些性能细节，更适合对性能有更高要求的场景。

## 结论

Snowflake 算法通过简单却有效的方式解决了分布式系统中唯一 ID 生成的问题。无论是 `bwmarrin/snowflake` 还是 `sony/sonyflake`，都提供了强大的工具让我们可以在 Go 语言中轻松实现这一算法。在具体应用中，我们可以根据需求选择适合的库，以确保系统的高效性和稳定性。