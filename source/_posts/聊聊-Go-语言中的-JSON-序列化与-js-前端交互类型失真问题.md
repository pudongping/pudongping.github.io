---
title: 聊聊 Go 语言中的 JSON 序列化与 js 前端交互类型失真问题
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
abbrlink: a370ff67
date: 2024-11-21 10:40:21
img:
coverImg:
password:
summary:
---

在 Web 开发中，后端与前端之间的数据交换通常通过 JSON 格式进行。

然而，在处理数字，尤其是大整数时，我们可能会遇到精度丢失的问题。这是因为 JavaScript 中的数字类型只能安全地处理一定范围内的整数。其数字类型是基于 64 位双精度浮点数的 `Number` 类型。这种类型可以安全表示 `-2^53` 到 `2^53` 之间的整数，超过这个范围的整数将无法精确表示，但是我们后端语言的整数范围是超过的，因此就有可能会遇到精度丢失的问题。

本文将通过 Go 语言的 `encoding/json` 包，探讨如何通过 JSON 序列化与反序列化来避免数字精度丢失的问题。

## Go 语言中的 JSON 处理

Go 语言的 `encoding/json` 包提供了强大的 JSON 序列化与反序列化功能。通过合理地使用结构体标签，我们可以控制 JSON 的编码与解码行为。

## 序列化：将大整数转为字符串

在 Go 语言中，如果我们有一个大整数，比如 `math.MaxInt64`，直接序列化为 JSON，那么在 JavaScript 中可能会丢失精度。为了解决这个问题，我们可以**将大整数以字符串的形式**序列化。

因为字符串不存在精度问题，从而从侧边也就解决了数字精度的问题。

```go
type User struct {
	UserID int64  `json:"user_id,string"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

func DigitalDistortionDemo() {
	data := User{
		UserID: math.MaxInt64,
		Name:   "Alex",
		Age:    18,
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	fmt.Printf("r1: %s\n", string(b))
}
```

在上述代码中，我们在 `User` 结构体的 `UserID` 字段上使用了 `json:"user_id,string"` 标签，这告诉 `json.Marshal` 函数将 `UserID` 以字符串的形式序列化。

## 反序列化：将字符串还原为大整数

当从前端接收到的 JSON 字符串中的 `user_id` 为字符串类型时，我们需要确保在反序列化过程中将其正确地转换回大整数。

```go
func DigitalDistortionDemo1() {
	s := `{"user_id":"9223372036854775807","name":"Alex","age":18}`
	var user User
	if err := json.Unmarshal([]byte(s), &user); err != nil {
		log.Fatalf("json unmarshal failed: %v", err)
	}
	fmt.Printf("r2: %+v\n", user)
}
```

在这段代码中，`json.Unmarshal` 函数将 JSON 字符串中的 `user_id` 字段正确地解析为 `User` 结构体中的 `UserID` 字段，即使它是以字符串形式提供的。

这样也就完美解决了，我们后端的数值传给 js 前端，前端丢失精度的问题。

并且因为 js 前端需要字符串类型，而我们后端定义的类型是一个 int64 类型，通过只是加了一个 `string` json tag ，从而就优雅的解决了 js 前端无论接收还是传值都用 string，后端继续使用 int64 类型，不用再做类型转换问题。

## 结论

通过在 Go 语言中合理使用 `encoding/json` 包的结构体标签，我们可以有效地避免在 JSON 序列化与反序列化过程中的数字精度丢失问题。

这种方法对于处理大整数，特别是在与 JavaScript 环境交互时，尤为重要。希望本文能够帮助你更好地理解和应用 JSON 数据交换中的数字精度问题。
