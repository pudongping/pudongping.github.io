---
title: Go语言中json序列化的一个小坑，建议多留意一下
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
abbrlink: 5fc86d1c
date: 2024-11-24 22:27:52
img:
coverImg:
password:
summary:
---

在 Go 语言开发中，JSON（JavaScript Object Notation）因其简洁和广泛的兼容性，通常被用作数据交换的主要序列化格式。然而，当你深入使用 JSON 时，**可能会发现它并不总是最佳选择**。

本文将探讨 JSON 序列化的一些局限性，也算是一个***小坑***吧。并给出一些常用的解决方案。

## JSON 序列化的潜在问题

我们先来看一个使用 JSON 进行序列化和反序列化的示例：

```go
package json_demo

import (
	"encoding/json"
	"fmt"
)

func JsonEnDeDemo() {
	d1 := make(map[string]interface{})
	d2 := make(map[string]interface{})

	var (
		age    int     = 18
		name   string  = "Alex"
		height float32 = 1.75
	)

	d1["name"] = name
	d1["age"] = age
	d1["height"] = height

	ret, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("json.Marshal failed: %v\n", err)
		return
	}
	// json.Marshal: {"age":18,"height":1.75,"name":"Alex"}
	fmt.Printf("json.Marshal: %s\n", string(ret))

	err = json.Unmarshal(ret, &d2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed: %v\n", err)
		return
	}
	// json.Unmarshal: map[age:18 height:1.75 name:Alex]
	fmt.Printf("json.Unmarshal: %v\n", d2)

	// 这里我们可以发现一个问题：Go 语言中的 json 包在序列化 interface{} 类型时，会将数字类型（整型、浮点型等）都序列化为 float64 类型
	for k, v := range d2 {
		// key: age, value: 18, type:float64
		// key: height, value: 1.75, type:float64
		// key: name, value: Alex, type:string
		fmt.Printf("key: %s, value: %v, type:%T \n", k, v, v)
	}
}
```

这段代码展示了如何将一个包含 `name`、`age` 和 `height` 的 Go map 数据结构序列化为 JSON 字符串，然后再反序列化回来。看似一切正常，但请注意反序列化后的数据类型变化。

运行代码后的输出可能会让你感到意外：

```
json.Marshal: {"age":18,"height":1.75,"name":"Alex"}
json.Unmarshal: map[age:18 height:1.75 name:Alex]
key: age, value: 18, type:float64 
key: height, value: 1.75, type:float64 
key: name, value: Alex, type:string 
```

**问题**：我们发现，尽管原始数据中 `age` 是 `int` 类型，`height` 是 `float32` 类型，但经过 JSON 反序列化后，它们全都变成了 `float64` 类型。

Go 语言中的 `encoding/json` 包会将所有数字类型（包括整型、浮点型等）**都**转换为 `float64` ，那么，有没有方式可以不让类型丢失呢？还真有！

## gob 二进制协议，高效且保留类型的 Go 专用序列化

为了避免 JSON 的这一局限性，我们可以使用 Go 语言特有的 GOB 序列化方式。GOB 不仅可以高效地序列化数据，还能够保留原始数据类型。

以下是使用 GOB 进行序列化和反序列化的示例：

```go
package json_demo

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func GobEnDeDemo() {
	d1 := make(map[string]interface{})
	d2 := make(map[string]interface{})

	var (
		age    int     = 18
		name   string  = "Alex"
		height float32 = 1.75
	)

	d1["name"] = name
	d1["age"] = age
	d1["height"] = height

	// encode
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(d1)
	if err != nil {
		fmt.Printf("gob.Encode failed: %v\n", err)
		return
	}
	b := buf.Bytes()
	// gob.Encode:  [13 127 4 1 2 255 128 0 1 12 1 16 0 0 57 255 128 0 3 4 110 97 109 101 6 115 116 114 105 110 103 12 6 0 4 65 108 101 120 3 97 103 101 3 105 110 116 4 2 0 36 6 104 101 105 103 104 116 7 102 108 111 97 116 51 50 8 4 0 254 252 63]
	fmt.Println("gob.Encode: ", b)

	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&d2)
	if err != nil {
		fmt.Printf("gob.Decode failed: %v\n", err)
		return
	}
	// gob.Decode: map[age:18 height:1.75 name:Alex]
	fmt.Printf("gob.Decode: %v\n", d2)

	for k, v := range d2 {
		// key: name, value: Alex, type:string
		// key: age, value: 18, type:int
		// key: height, value: 1.75, type:float32
		fmt.Printf("key: %s, value: %v, type:%T \n", k, v, v)
	}
}
```

从上面的代码中可以看到，GOB 序列化不仅保留了 `age` 的 `int` 类型和 `height` 的 `float32` 类型，还能高效地进行数据编码。这使得 GOB 成为在 Go 程序内部传递数据的理想选择。

## 第三方包 msgpack

`msgpack` 是一种高效的二进制序列化格式，它允许你在多种语言(如JSON)之间交换数据。但它更快更小。

首先需要先下载这个包

```bash
go get -v github.com/vmihailenco/msgpack/v5
```

来看一个使用 msgpack 的示例：

```go
package json_demo

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

func MsgpackEnDeDemo() {
	// msgpack 序列化示例
	d1 := make(map[string]interface{})
	d2 := make(map[string]interface{})

	var (
		age    int     = 18
		name   string  = "Alex"
		height float32 = 1.75
	)

	d1["name"] = name
	d1["age"] = age
	d1["height"] = height

	// encode
	b, err := msgpack.Marshal(d1)
	if err != nil {
		fmt.Printf("msgpack.Marshal failed: %v\n", err)
		return
	}
	// msgpack.Marshal:  [131 164 110 97 109 101 164 65 108 101 120 163 97 103 101 18 166 104 101 105 103 104 116 202 63 224 0 0]
	fmt.Println("msgpack.Marshal: ", b)

	// decode
	err = msgpack.Unmarshal(b, &d2)
	if err != nil {
		fmt.Printf("msgpack.Unmarshal failed: %v\n", err)
		return
	}
	// msgpack.Unmarshal: map[age:18 height:1.75 name:Alex]
	fmt.Printf("msgpack.Unmarshal: %v\n", d2)

	for k, v := range d2 {
		// key: age, value: 18, type:int8
		// key: height, value: 1.75, type:float32
		// key: name, value: Alex, type:string
		fmt.Printf("key: %s, value: %v, type:%T \n", k, v, v)
	}

}
```

**msgpack的优势**：
- **高效紧凑**：数据体积比 JSON 更小，序列化和反序列化速度更快。
- **类型保持**：与 GOB 类似，msgpack 也能保持原始数据类型。

## 总结

- **json**：虽然广泛使用且易于阅读，但在处理数字类型时有潜在的精度问题。
- **gob**：适用于 Go 语言程序内部的数据传输，保留类型且性能优异，但仅适用于 Go。
- **msgpack**：在需要高效、紧凑的跨语言数据交换时非常有用，同时还能保留数据类型。

通过这三种序列化方式的比较，希望你能够根据实际需求选择合适的工具。在需要保证类型和性能的 Go 程序中，gob 和 msgpack 可能是比 json 更好的选择，不过，你也完全可以使用 json 包来反序列化，只不过取值的时候就需要通过类型断言来得到之前的类型。