---
title: Go语言中的加解密利器：go-crypto库全解析
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
abbrlink: 4f904664
date: 2024-11-25 10:31:50
img:
coverImg:
password:
summary:
---

在软件开发中，数据安全和隐私保护越来越受到重视。Go 语言以其简洁高效的特性，成为了许多开发者的首选。然而，在实际项目中使用加解密时，还是需要在标准库的基础上做一些封装。`go-crypto` 库应运而生，它是一个专为 Golang 设计的加密解密工具库，提供了 AES 和 RSA 等多种加密算法的支持。

本文将从安装、特性、基本与高级功能，以及实际应用场景等多个角度，全面介绍这个库。

## go-crypto 库简介

`go-crypto` 是一个为 Golang 设计的加密解密工具库，它实现了多种常用的加密算法，包括 AES 和 RSA 等。通过这个库，开发者可以轻松地在 Go 语言项目中实现数据的加密和解密，保障数据传输和存储的安全性。

## 安装

要在你的 Go 项目中使用 `go-crypto`，首先需要通过 `go get` 命令安装：

```shell
go get -u github.com/pudongping/go-crypto
```

## 特性

`go-crypto` 库提供了以下特性：

1. **AES加解密方法**：支持电码本模式（ECB）、密码分组链接模式（CBC）、计算器模式（CTR）、密码反馈模式（CFB）和输出反馈模式（OFB）。
2. **RSA加解密方法**：支持 RSA 加密和解密。

接下来，我就分别以 Go 和 PHP 加解密分别来演示其用法。

## AES 加解密

### CBC 模式

CBC 模式是密码分组链接模式，它通过将前一个块的加密结果与当前块的明文进行 XOR 操作，增加了加密数据的安全性。以下是使用 `go-crypto` 库进行 AES-CBC 加密和解密的示例：

#### Go加密，PHP解密（AES-128-CBC）

go 加密

```go
package main

import (
	"fmt"
	"github.com/pudongping/go-crypto"
)

func main() {
	plaintext := "hello world! My name is Alex Pu"
	key := "1234567890123456" // 密钥字节长度必须为16个字节

	ciphertext, err := go_crypto.AESCBCEncrypt(plaintext, key)
	if err != nil {
		fmt.Println("出错啦！", err)
	}
	fmt.Println(ciphertext)
}
```

PHP 解密

```php
<?php
$key = '1234567890123456';
$iv = mb_substr($key, 0, 16);
$s = 'BRK08I0OYOoFwhgIBT1qjFywFkLADdeLQfVZM7CPKJ8=';

$str = base64_decode($s);
$decrypted = openssl_decrypt($str, 'AES-128-CBC', $key, OPENSSL_RAW_DATA, $iv);
if (!$decrypted) {
    echo '解密失败' . PHP_EOL;
} else {
    echo($decrypted) . PHP_EOL;
}
?>
```

#### php 加密，go 解密（AES-128-CBC）

PHP 加密

```php

$string = 'hello world! alex';
$key = '1234567890123456';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-128-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: YAZkDJYi7e9O09TRNvUf+6sFMlI8AQvZ5GVU+xJIuOc=
echo $s . PHP_EOL;

```

Go 解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "YAZkDJYi7e9O09TRNvUf+6sFMlI8AQvZ5GVU+xJIuOc="
    key := "1234567890123456"
    
    plaintext, err := go_crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

### ECB 模式

ECB 模式是电码本模式，它是最简单的加密模式，但安全性较低，通常不推荐使用。以下是使用 `go-crypto` 库进行 AES-ECB 加密和解密的示例：

#### Go加密，PHP解密（AES-128-ECB）

go 加密

```go
package main

import (
	"fmt"
	"github.com/pudongping/go-crypto"
)

func main() {
	plaintext := "hello world! My name is Alex Pu"
	key := "1234567890123456" // 密钥字节长度必须为16个字节

	ciphertext, err := go_crypto.AESECBEncrypt(plaintext, key)
	if err != nil {
		fmt.Println("出错啦！", err)
	}
	fmt.Println(ciphertext)
}
```

php 解密

```php
<?php
$key = '1234567890123456';
$s = 'sRFeHhndretZFZE9/7WdGuGw1QYl8l/IlI1XEtpVzxI=';

$str = base64_decode($s);
$decrypted = openssl_decrypt($str, 'AES-128-ECB', $key, OPENSSL_RAW_DATA);
if (!$decrypted) {
    echo '解密失败' . PHP_EOL;
} else {
    echo($decrypted) . PHP_EOL;
}
?>
```

## RSA 加解密

`go-crypto` 库还提供了 RSA 加密和解密的功能。以下是使用 `go-crypto` 库进行 RSA 加密和解密的示例：

```go
package main

import (
	"fmt"
	"github.com/pudongping/go-crypto"
)

func main() {
	privateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
...
-----END RSA PRIVATE KEY-----`)

	publicKey := []byte(`-----BEGIN PUBLIC KEY-----
...
-----END PUBLIC KEY-----`)

	plaintext := "hello world"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := go_crypto.RSAEncrypt(publicKey, []byte(plaintext))
	if err != nil {
		fmt.Println(err)
		return
	}

	plaintext1, err := go_crypto.RSADecrypt(privateKey, ciphertext)
	fmt.Println("解密 ==> ", string(plaintext1))
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

## 应用场景

假设你正在开发一个需要安全通信的分布式系统，`go-crypto` 库可以用于加密敏感数据，如用户信息、支付信息等，确保数据在传输过程中的安全性。通过使用 AES 加密，你可以保护数据不被未授权访问，而 RSA 加密则可以用于安全地传输密钥。

## 结语

`go-crypto` 库为 Go 语言开发者提供了一个强大而灵活的加密解密工具。通过本文的详细介绍，希望你能深入理解并掌握 `go-crypto` 的使用方法，为你的项目增加一层安全保障。在实际开发中，合理利用加密技术，可以显著提高系统的安全性和可靠性。
