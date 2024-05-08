---
title: 推荐几个好用的GRPC调试工具
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: RPC
tags:
  - RPC
  - GRPC
abbrlink: 47f9e31c
date: 2024-05-09 01:41:44
img:
coverImg:
password:
summary:
---

# GRPC 调试工具

## 1. [grpcui](https://github.com/fullstorydev/grpcui)

- 下载

```bash
# 这时，会在 $GOPATH/bin 目录下，生成一个 grpcui 可执行文件
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

- 测试安装是否成功

```bash

grpcui -help

```

- 调试

```bash

# 这里的地址需要设置为：grpc 项目域名或者 ip + 端口号
# -plaintext 参数忽略 tls 证书的验证过程
grpcui -plaintext localhost:12345

```

## 2. [grpcurl](https://github.com/fullstorydev/grpcurl)

- 下载

```bash
# 这时，会在 $GOPATH/bin 目录下，生成一个 grpcurl 可执行文件
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

- 测试安装是否成功

```bash
grpcurl -help
```

- 调试

```bash
# 查看服务列表
# grpcurl 工具默认使用 TLS 认证（可通过 -cert 和 -key 参数设置公钥和密钥）可以通过指定 plaintext 选项来忽略 TLS 认证
grpcurl -plaintext localhost:5200 list
# 比如输出以下内容：
# grpc.reflection.v1alpha.ServerReflection
# proto.TagService

# 查看服务的方法列表
grpcurl -plaintext localhost:5200 list proto.TagService
# 比如输出以下内容：
# proto.TagService.GetTagList

# 查看方法的细节
grpcurl -plaintext localhost:5200 describe proto.TagService
# 比如输出以下内容：
# proto.TagService is a service:
# service TagService {
#  rpc GetTagList ( .proto.GetTagListRequest ) returns ( .proto.GetTagListResponse );
# }

# 调用方法
# 参数必须为 json 格式的字符串
# `proto.TagService/GetTagList` 是所需要调用的接口
grpcurl -plaintext -d '{"name":"go"}' localhost:5200 proto.TagService/GetTagList
# 或者
grpcurl -plaintext -d '{"name":"go"}' localhost:5200 proto.TagService.GetTagList
```

以上两种方式都需要启动反射服务才可以使用，代码类似如下

```go
package main

import (
	"log"
	"net"

	pb "github.com/pudongping/go-grpc-service/proto"
	"github.com/pudongping/go-grpc-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	pb.RegisterTagServiceServer(s, server.NewTagServer())

	listen, err := net.Listen("tcp", ":5200")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 注册反射服务，方便让 grpcurl 或者 grpcui 用作调试
	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
```

## 3. [bloomrpc](https://github.com/bloomrpc/bloomrpc)

> 使用方式非常简单，GitHub 仓库上有一张动图演示，只需要上传 proto 文件即可，不需要在代码中开启反射。

**如果发现 bloomrpc 包使用 int64 时失真，那么可以考虑使用 [crossoverJie/ptg](https://github.com/crossoverJie/ptg) 包（看他的文档说已经解决掉 bloomrpc int64 失真的问题，但是我自己没有去验证）**

- 安装

```
# macOS/Homebrew
brew install --cask bloomrpc

# 或者直接根据系统下载对应的包文件，下载地址如下：
https://github.com/bloomrpc/bloomrpc/releases
# 比如要是 macOS 系统下
wget https://github.com/bloomrpc/bloomrpc/releases/download/1.5.3/BloomRPC-1.5.3.dmg
```

## 4. [kreya](https://kreya.app/)
