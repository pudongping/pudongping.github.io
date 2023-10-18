---
title: 使用 GRPC
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
abbrlink: 8a377bf8
date: 2023-10-18 15:21:01
img:
coverImg:
password:
summary:
---

# 使用 GRPC

## 安装 [protobuf](https://github.com/protocolbuffers/protobuf)

> [官方安装地址](https://grpc.io/docs/protoc-installation/)

- 第一种方式：Mac 下使用 Homebrew 安装

```bash

brew install protobuf

```

- 第二种方式：指定版本号安装

> 这里指定安装 `3.17.0` 版本

```bash

PROTOC_ZIP=protoc-3.17.0-osx-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.17.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

```

或者使用如下方式安装

> 这里示范的是安装 Linux 环境下的 protobuf

```bash

# 下载编译好的二进制包
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.17.0-rc1/protoc-3.17.0-rc-1-linux-x86_64.zip

# 创建文件夹并解压缩到指定文件夹中
mkdir protoc && tar -xzvf protoc-3.17.0-rc-1-linux-x86_64.zip -C protoc

# 将解压后的文件移动到 /usr/local 文件夹下
sudo mv protoc /usr/local/protoc

# 创建软连接，方便在任意目录下使用 protoc 命令
ln -s /usr/local/protoc/bin/protoc /usr/local/bin/protoc

```

- 第三种方式：编译安装

> 这里安装的是 `3.19.1` 版本

```bash

# 下载安装包
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protobuf-all-3.19.1.zip

# 解压缩
unzip protobuf-all-3.19.1.zip && cd protobuf-3.19.1

./configure

make

make install

```

- 第四种方式：直接使用二进制文件

```bash

# 根据你自己的系统下载对应的源码包 （比如我这里使用的是 mac book，我就要下载 osx 压缩包）
cd ~/go-tools && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-osx-x86_64.zip

# 解压缩
unzip protoc-3.19.1-osx-x86_64.zip

# 编辑配置文件
vim ~/.zshrc

# 写入以下配置信息
export PATH="/Users/pudongping/go-tools/protoc-3.19.1-osx-x86_64/bin:$PATH"

source ~/.zshrc
```

- 通过查看版本号，检查是否安装成功

```bash

# protoc 是 Protobuf 的编译器，其主要功能是用于编译 .proto 文件
protoc --version

```

> 如果出现 `protoc: error while loading shared libraries: libprotobuf.so.15: cannot open shared object file: No such file or directory` 那么则需要在命令行中执行 `ldconfig` 命令后，再次运行即可成功。

## 安装 protoc 插件

> 仅安装 protoc 编译器是不够的，针对不同的语言，还需要安装运行时的 protoc 插件，而对应 Go 语言的是 protoc-gen-go 插件。

可以执行以下命令进行安装，但是**不推荐**，因为 protoc-gen-go 是需要与 proto 软件包版本相匹配的，必须要锁定版本号

```bash
# 安装的二进制文件在 `$GOPATH/bin` 目录下
# 安装 protobuf 3.17.0 版本之后直接执行 go get -u github.com/golang/protobuf/protoc-gen-go 命令后，
# 貌似安装的 protoc-gen-go 版本是 v1.27.1，但是指定 protoc-gen-go@v1.27.1 貌似又找不到 v1.27.1 的版本，不知为何，暂且记录下。
# 在项目根目录下执行
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.3
```

**推荐的安装方式如下：**

```bash
GIT_TAG="v1.3.3"
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

然后还可以将其移动到 bin 目录下

```bash
mv $GOPATH/bin/protoc-gen-go /usr/local/go/bin/
```

## 安装 grpc-go 库

grpc-go 包含了 Go 的 grpc 库，我们可以使用如下方式安装

```bash
go get -u google.golang.org/grpc@v1.29.1
```

如果使用以上方式安装，发现无法安装（可能会被墙掉了），我们可以使用如下方式手动安装

```bash
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/
go install google.golang.org/grpc
```

## 生成 proto 文件（生成接口库）

```bash
# 加载 protoc-gen-go 插件 生成 grpc 的 go 服务端/客户端
# 注意：必须要先安装好 protoc-gen-go 插件

# 这里将 .proto 文件写到 proto 目录，执行此命令后将会把对应的 .pb.go 文件也会生成到 proto 目录
protoc --go_out=plugins=grpc:. ./proto/*.proto
```

- `-I` 参数：格式为 `-IPATH` 作用是指定 import 搜索的目录（也就是 Proto 文件中的 import 命令），可指定多个，编译时按照顺序查找，如果不指定则默认当前工作目录，比如：`-I.` 、`-I/usr/local` 、`-I$GOPATH/src`
- `M` 参数：指定导入的 `.proto` 文件路径编译后对应的 golang 包名（不指定默认以 .proto 文件中 import 语句路径），格式为：Mfoo/bar.proto=quux/shme，则在生成、编译 Proto 时将所指定的包名替换为所要求的名字（如：foo/bar.proto 编译后为包名为 quux/shme）
- `--go_out`：设置所生成的 Go 代码输出的目录。该指令会加载 protoc-gen-go 插件，以达到生成 Go 代码的目的。生成的文件以 .pb.go 为文件后缀，这里的 `:` （冒号）有分隔符的作用，后跟命令所需要的参数集，这意味着把生成的 Go 代码输出到指向的 protoc 编译的当前目录。
- `plugins=plugin1+plugin2`：指定要加载的子插件列表。我们定义的 proto 文件是涉及了 RPC 服务的，而默认是不会生成 RPC 代码的，因此需要在 go_out 中给出 plugins 参数，将其传递给 protoc-gen-go 插件，即告诉编译器，请支持 RPC。