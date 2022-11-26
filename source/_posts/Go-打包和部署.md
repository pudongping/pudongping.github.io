---
title: Go 打包和部署
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
abbrlink: 5c09a73d
date: 2022-11-26 22:45:03
img:
coverImg:
password:
summary:
---

# Go 打包和部署


## 打包相关命令

命令 | 含义
--- | ---
go run | 编译并马上运行 go 程序（只接收 main 包下的文件作为参数）
go build | 编译指定的源文件、软件包及其依赖项，但它不会运行编译后的二进制文件。(如果想要指定所生成的二进制文件为其他名称，则可以通过 `-o` 参数进行调整)
go install | 编译并安装源文件、软件包到 `$GOBIN` 目录下。 可以执行 `go install -x` 查看它的编译过程。文件名称为 Go modules 的项目名，而不是目录名

参数 | 含义
--- | ---
-x | 打印编译过程中的所有执行命令，执行生成的二进制文件，比如：`go run -x main.go` 和 `go build -x`
-n | 打印编译过程中的所有执行命令，不执行生成的二进制文件
-a | 强制重新编译所有涉及的依赖，简单来说，就是不利用缓存或已编译好的部分文件，直接所有包都是最新的代码重新编译和关联
-o | 指定生成的二进制文件名称
-p | 指定编译过程中可以并发运行程序的数量，默认值为可用的 CPU 数量（Go 语言默认是支持并发编译的）
-work | 打印临时工作目录的完整路径，在退出时不删除该目录
-race | 启用数据竞争检测，目前仅支持 Linux/amd64、FreeBSD/amd64、Darwin/amd64 和 Windows/amd64 平台
-installsuffix | 在软件包安装的目录中增加后缀标识，以保持输出与默认版本分开

## 跨平台交叉编译

Go 语言支持跨平台交叉编译，也就是说我们可以在 Windows 或 Mac 平台下编写代码，最后将代码编译成能够在 Linux amd64 服务器上运行的程序。

根目录使用以下指令可以静态编译 `Linux` 平台 `amd64` 架构的可执行文件：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o <application-name>

# 比如
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o <application-name> .
```

Windows 依次执行以下四个命令：

```bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o <application-name>
```

参数名 | 含义
--- | ---
CGO_ENABLED | 设置是否在 Go 代码中调用 C 代码。0 为关闭，采用纯静态编译。可通过执行 go env 进行查看
GOOS | 目标操作系统，如 Linux 、Darwin、Windows
GOARCH | 目标操作系统的架构，若不设置，则默认值为程序运行环境的目标计算架构一致
GOHOSTOS | 用于标识程序运行环境的目标操作系统
GOHOSTARCH | 用于标识程序运行环境的目标计算架构

**可以通过如下命令查看 Go 支持 OS 和平台列表**

```bash
go tool dist list
```

系统 | GOOS | GOARCH
--- | --- | ---
Windows 32位 | windows | 386
Windows 64位 | windows | amd64
OS X 32位 | darwin | 386
OS X 64位 | darwin | amd64
Linux 32位 | linux | 386
Linux 64位 | linux | amd64

## 第三方打包工具

以下几个第三方打包工具可以将静态文件（JS、CSS、图片）或者模版文件等非 `.go` 文件打包到一个二进制文件中， [Embed](https://pkg.go.dev/embed) 标准库也可以，但是得使用 go 1.16 版本以后的版本才行。

- [go-bindata/go-bindata](https://github.com/go-bindata/go-bindata)：推荐使用这个
- [gobuffalo/packr](https://github.com/gobuffalo/packr)
- [markbates/pkger](https://github.com/markbates/pkger)
- [rakyll/statik](https://github.com/rakyll/statik)
- [knadh/stuffbin](https://github.com/knadh/stuffbin)
- [github.com/go-bindata/go-bindata](https://github.com/go-bindata/go-bindata)
- [elazarl/go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs)
- [GeertJohan/go.rice](https://github.com/GeertJohan/go.rice)

### 将数据文件转换成 go 代码

> [go-bindata](https://github.com/go-bindata/go-bindata) 库可以将数据文件转换为 Go 代码。例如：常见的配置文件、资源文件（如 swagger ui）
> 因此读取配置信息也可以通过 [go-bindata/go-bindata](https://github.com/go-bindata/go-bindata/) 包提供的方式来读取。使用方式如下：

1. 安装 `go-bindata/go-bindata` 包

```shell
go get -u github.com/go-bindata/go-bindata/...
```

2. 将配置文件生成 go 代码

```shell
# 执行这条命令后，会将 `configs/config.yaml` 文件打包，并通过 `-o` 参数选择指定的路径输出到 `configs/config.go` 文件中
# 再通过 `-pkg` 选项指定生成的包名为 `configs`
go-bindata -o configs/config.go -pkg-configs configs/config.yaml
```

3. 读取文件中的配置信息

```go
data, err := configs.Asset("configs/config.yaml")

if err == nil {
    fmt.Println(string(data))
}
```


## 编译缓存

```bash

# 查看编译缓存所在的目录
go env GOCACHE

# 清理编译缓存
go clean -cache

# 查看程序编译的时间
time go build

```

## 压缩编译后的二进制文件

> **如非必要情况，不建议压缩**

- 方式一：去掉 DWARF 调试信息和符号表信息

```bash
go build -ldflags="-w -s"
```

参数名 | 含义 | 副作用
--- | --- | ---
-w | 去除 DWARF 调试信息 | 会导致异常 panic 抛出时，调用堆栈信息没有文件名、行号信息
-s | 去除符号表信息 | 无法使用 gdb 调试

- 方式二：使用 [upx](https://github.com/upx/upx)、 [upx 教程](https://niconiconi.fun/2019/01/10/upx-tutorial/) 工具对可执行文件进行压缩

## 编译信息写入（使用 ldflags 设置编译信息）

> [go version 信息注入](https://ms2008.github.io/2018/10/08/golang-build-version/)

- 先在 `main.go` 文件中写入代码

```go

var (
	buildTime    string // 二进制文件编译时间
	buildVersion string // 二进制文件编译版本
	goVersion    string // 打包二进制文件时的 go 版本信息
	gitCommitID  string // 二进制文件编译时的 git 提交版本号
)

func main() {

	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Build Time: %s \n", buildTime)
		fmt.Printf("Build Version: %s \n", buildVersion)
		fmt.Printf("Build Go Version: %s \n", goVersion)
		fmt.Printf("Build Git Commit Hash ID: %s \n", gitCommitID)
		return
	}

}

```

- 编译代码时执行如下命令

```bash
go build -o app-service -ldflags \
"-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X 'main.goVersion=$(go version)' -X main.gitCommitID=`git rev-parse HEAD`"

```

> 在上述命令中，通过 `-ldflags` 命令的 `-X` 参数可以在链接时将信息写入变量中，其格式为：`package_name.variable_name=value`

- 查看编译后的二进制文件和版本信息

```bash
./app-service -version

# output is:
# Build Time: 2022-03-27,13:50:26
# Build Version: 1.0.0
# Build Go Version: go version go1.16.3 darwin/amd64
# Build Git Commit Hash ID: b3473e9cc98148f5c94b53c1cada7de133143462
```

## 部署

### 关于守护进程

- `nohup` —— linux 系统内置，直接使用 `nohup ./<executable-file> &` 即可
- `systemctl` —— linux 系统内置，需要配置 systemctl 管理配置文件
- `supervisor`
- [go-supervisor](https://github.com/ochinchina/supervisord) —— Go 语言实现但 supervisor，好处是不需要安装 Python 环境

#### 1. 使用 supervisor 部署：

```
# 假设关闭一个监听端口为 3000 的服务
kill -9 $(lsof -ti:3000)
```

创建 supervisor 相关配置信息

vim /etc/supervisor/conf.d/alex-blog.conf

```bash
# 程序名称，stop start 等管理时使用
[program:alex-blog]
# 进入该目录运行命令，确保了和项目相关的一些配置文件可以得到加载，比如 .env 文件
directory=/data/www/blog
# 以绝对路径的方式执行 alex-blog 二进制文件
command=/data/www/blog/alex-blog
# 重启时发送的信号，确保端口正常关闭
stopsignal=TERM
# 是否自启动
autostart=true
# 是否自动重启
autorestart=true
# 执行程序的用户
user=www-data
# 输出日志位置
stdout_logfile=/data/log/supervisor/blog/stdout.log
# 错误输出日志
stderr_logfile=/data/log/supervisor/blog/stderr.log
```

重载 Supervisor 配置文件

```bash
# 重载 supervisor 配置文件
supervisorctl reload

# 查看程序名称为 alex-blog 的程序状态
# 如果直接执行 `supervisorctl status` 命令的话，则是查看所有任务的状态
supervisorctl status alex-blog

# 关闭所有任务
supervisorctl shutdown

# 启动任务
supervisorctl start <程序名>

# 关闭任务
supervisorctl stop <程序名>
```

#### 2. 使用 docker 部署：

##### 编译后的二进制文件放在 docker 中运行

先要在宿主机中项目根目录下进行编译

```shell
# 交叉编译生成 Linux 平台的可执行文件
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-world
```

编写 Dockerfile 文件

```dockerfile
# 基础镜像
FROM alpine:3.12
# 或者使用 Scratch 镜像
# FROM scratch

# 维护者
MAINTAINER alex

# docker build 时执行命令
RUN mkdir -p /go-project/demo \
&& ln -sf /dev/stdout /go-project/demo/storage/service.log

# 工作目录
WORKDIR /go-project/demo

# 拷贝
COPY hello-world /go-project/demo/hello-world
# 或者直接将当前目录下所有的文件拷贝到容器中
# COPY . /go-project/demo

# 这里暴露端口与否都行
# EXPOSE 8501

# docker run 时执行的命令
ENTRYPOINT ["./hello-world"]
```

构建镜像

```shell
docker build -t go-project:v1.0.0 .
```

运行容器

```shell
# 宿主机的 9501 端口映射到了容器的 8501 端口
# -d 用于该程序在后台运行
# -p 用于映射端口
docker run -d -p 9501:8501 go-project:v1.0.0
```

#### nginx 配置信息

```nginx

upstream go-project {
    # go-project HTTP Server 的 IP 及 端口
    server 127.0.0.1:9501;
}

server
{
    listen 80;
    server_name goblog.com;
    
    access_log   /data/log/nginx/goblog/access.log;
    error_log    /data/log/nginx/goblog/error.log;
    
    location /static/ {
      alias /www/wwwroot/gitlab/go-project/public/uploads/; #静态资源路径
    }
    
    location / {
        # 将客户端的 Host 和 IP 信息一并转发到对应节点
        proxy_redirect             off;
        proxy_set_header           Host             $http_host;
        proxy_set_header           X-Real-IP        $remote_addr;
        proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        
        # 转发Cookie，设置 SameSite
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        # 执行代理访问真实服务器
        proxy_pass                 http://go-project;
    }
    
}

```
