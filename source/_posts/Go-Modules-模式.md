---
title: Go Modules 模式
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
abbrlink: 2f3aad3a
date: 2022-11-14 22:49:58
img:
coverImg:
password:
summary:
---

# Go Modules 模式

## GOPATH 目录

GOPATH 目录下一共包含三个子目录：

- bin：存储所编译生成的二进制文件。
- pkg：存储预编译的目标文件，以加快程序的后续编译速度。
- src：存储所有 `.go` 文件或源代码。在编写 Go 应用程序，程序包和库时，一般会以 `$GOPATH/src/github.com/foo/bar` 的路径进行存放。

> 使用 `go get` 来拉取外部依赖时，会自动下载并安装到 `$GOPATH` 目录下。

## go mod 命令

查看 go mod 都有哪些命令

```bash
# Go 的所有工具都可以使用 `go help` 来查看使用方法
go help mod

# 查看 go mod download 有哪些参数
go help mod download
```

命令 | 作用
:--- | :---:
go mod help | 查看帮助信息
go mod init | 初始化当前文件夹，生成 go.mod 文件
go mod download | 下载 go.mod 文件中指明的所有依赖到本地（默认为 `$GOPATH/pkg/mod` 目录）增加 `-x` 参数 `go mod download -x` 会打印下载信息；`go mod download -json` 用来查看模块下载的 zip 存放位置，以及解压后的位置；
go mod tidy | 整理现有的依赖，执行时会把未使用的 module 移除掉，同时也会增加缺少的包
go mod graph | 查看现有的依赖结构图
go mod edit | 编辑 go.mod 文件，比如修改项目中使用的 go 版本 `go mod edit -go=1.17`
go mod vendor | 导出项目所有的依赖到 vendor 目录（需要执行 go build -mod=vendor 才可以使用 vendor 作为依赖来编译）
go mod verify | 校验一个模块是否被篡改过，校验从 GOPROXY 服务器上下载的 zip 文件与 GOSUMDB 服务器下载下来的哈希值，是否匹配。
go mod why | 查看为什么需要依赖某模块，比如 `go mod why gopkg.in/yaml.v2 gopkg.in/yaml.v3`
go clean -modcache | 可以清空本地下载的 Go Modules 缓存 （会清空 `$GOPATH/pkg/mod` 目录）

## go mod 环境变量

和 `go mod` 比较关联的几个环境变量

```go

go env

GO111MODULE="auto"
GOPROXY="https://goproxy.cn,direct"
GOSUMDB="sum.golang.org"
GONOPROXY=""
GONOSUMDB=""
GOPRIVATE=""

```

### GO111MODULE

Go 语言提供了 GO111MODULE 这个环境变量来作为 Go modules 的开关，其允许设置以下参数：

- auto：只要项目包含了 go.mod 文件的话启用 Go modules，目前在 Go1.11 至 Go 1.14 中仍然是默认值
- on：启用 Go modules，推荐设置，将会是未来版本中的默认值
- off：禁用 Go modules，不推荐设置

#### 设置方式

```shell
go env -w GO111MODULE=on
```

也可以直接在 shell 环境变量中设置，比如我这里使用的是 Mac，且使用的 zsh，则在 `vim ~/.zshrc` 然后添加以下内容，如果是其它 Linux 系列系统，则需要在 `~/.bash_profile` 文件中进行设置。

```shell

export GO111MODULE=on

```

然后要记得 `source ~/.zshrc`

### GOPROXY

这个环境变量主要是用于设置 Go 模块代理（Go module proxy），其作用是用于使 Go 在后续拉取模块版本时直接通过镜像站点来快速拉取

GOPROXY 的默认值是：`https://proxy.golang.org,direct`

```shell

# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct

```

`direct` 是一个特殊指示符，用于指示 Go 在获取源码包时，先尝试在设置 GOPROXY 的地址下抓取，如果遇到 404 或 410 等错误时，再回溯到模块版本的源地址去抓取 （比如 GitHub 等）。


### GOSUMDB

它的值是一个 Go checksum database，用于在拉取模块版本时（无论是从源站拉取还是通过 Go module proxy 拉取）保证拉取到的模块版本数据未经过篡改，若发现不一致，也就是可能存在篡改，将会立即终止。

GOSUMDB 的默认值为：`sum.golang.org`，在国内也是无法访问的，但是 GOSUMDB 可以被 Go 模块代理所代理，因此我们可以通过设置 `GOPROXY` 来解决，而先前我们所设置的模块代理 `goproxy.cn` （七牛云的 CDN）就能支持代理 `sum.golang.org`，所以这一个问题在设置 GOPROXY 后，可以不需要过度关心。

也可以将其设置为 `off` ，也就是禁止 Go 在后续操作中校验模块版本。但是不建议关闭校验。


### GONOPROXY/GONOSUMDB/GOPRIVATE

- GONOPROXY —— 设置不走 Go Proxy 的 URL 规则；
- GONOSUMDB —— 设置不检查哈希的 URL 规则；
- GOPRIVATE —— 设置私有模块的 URL 规则，会同时设置以上两个变量。

这三个环境变量都是用在当前项目依赖了私有模块，例如自己公司部署的私有 git 仓库或者是 GitHub 中的私有仓库，都是需要进行设置的，否则会拉取失败。

设置 GOPRIVATE 之后，表示该地址为私有仓库，不会从 GOPROXY 所对应的地址上去下载。

**而一般建议直接设置 GOPRIVATE，它的值将作为 GONOPROXY 和 GONOSUMDB 的默认值，所以建议直接设置 GOPRIVATE 即可。**


```shell

# 以下表示 git.example.com 和 github.com/username/package 都是私有仓库，不会进行 GOPROXY 下载和校验
go env -w GOPRIVATE="git.example.com,github.com/username/package"
# 设置后，前缀为 `git.example.com` 和 `github.com/username/package` 的模块都会被认为是私有模块

# 表示所有模块路径为 example.com 的子域名都不进行 GOPROXY 下载和校验
# 需要注意的是不包括 example.com 本身
go env -w GOPRIVATE="*.example.com"

```

## 使用 Go Modules 初始化项目

```shell

# 开启 Go Modules 模块，保证 GO111MODULE=on
go env -w GO111MODULE=on

# 在任意文件夹下创建一个项目（不要求在 $GOPATH/src 目录下创建）
mkdir -p $HOME/modules_test

cd $HOME/modules_test

# 创建 go.mod 文件，同时起当前项目的模块名称
# 如果你是在 `$GOPATH/src` 目录下创建的文件夹可以直接执行 `go mod init` 命令初始化，不需要加模块名称
go mod init github.com/pudongping/moudles_test

# 在该项目下编写源代码，并下载依赖库
# 也可以不加 `-v` 参数
go get -v XXXXXXXX

eg: go get -v github.com/pudongping/moudles_test

```

> 下载的包其实被缓存在 `$GOPATH/pkg/mod` 目录和 `$GOPATH/pkg/sumdb` 目录下

- **`go.mod` 文件中：**

1. module：用于定义当前项目的模块路径。
2. go：用于标识当前模块的 Go 语言版本，值为初始化模块时的版本。
3. require：用于设置一个特定的模块版本。
4. exclude：用于从使用中排除一个特定的模块版本。
5. replace：用于将一个模块版本替换为另外一个模块版本。

```
# v0.0.0 表示版本信息
# 20190718012654 表示所拉取版本的 commit 时间
# fb15b899a751 表示所拉取版本的 commit 哈希值
github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
```

- **`go.sum` 文件中：**

go.sum 文件的作用是：罗列当前项目直接或间接依赖的所有模块版本，保证今后项目依赖的版本不会被篡改。间接依赖的包的哈希值也会被保存。

go.sum 文件中有两种 hash 的形式：

1. `h1:<hash>` 将目标模块版本的 zip 文件开包后，针对所有包内文件依次进行 hash，然后再把它们的 hash 结果按照固定格式和算法组成总的 hash 值，如果不存在，表示可能依赖的库用不上
2. `xxx/go.mod h1.<hash>` 表示 go.mod 文件做的 hash

```shell

# 将目标模块版本的 zip 文件开包后，针对包内所有文件依次进行 hash，然后再将 它们的 hash 结果汇总组成 hash
github.com/gorilla/mux v1.7.4 h1:VuZ8uybHlWmqV03+zRzdwKL4tUnIp1MAQtp1mIFE1bc=

# 针对 go.mod 文件的 hash 值
github.com/gorilla/mux v1.7.4/go.mod h1:DVbg23sWSpFRCP0SfiEN6jmj59UnW/n46BH5rLB71So=

```

## go get 拉取命令

> 下载的模块会被放置于 `$GOPATH/pkg/mod` 目录中

命令 | 作用
--- | ---
go get	| 拉取依赖，会进行指定性拉取（更新），并不会更新所依赖的其它模块。
go get -u	| 更新现有的依赖，会强制更新它所依赖的其它全部模块，不包括自身。
go get -u -t ./…	| 更新所有直接依赖和间接依赖的模块版本，包括单元测试中用到的。
go get golang.org/x/text@latest	| 拉取最新的版本，若存在 tag，则优先使用。
go get golang.org/x/text@master	| 拉取 master 分支的最新 commit。
go get golang.org/x/text@v0.3.2	| 拉取 tag 为 v0.3.2 的 commit。
go get golang.org/x/text@342b2e	| 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2。

## 修改项目模块的版本依赖关系

```shell

go mod edit -replace=<老版本>=<需要替换的版本>

# 比如
go mod edit -replace=demo-package@v1.0.0=demo-package@v2.0.0

```

## go list 命令以及参数

参数 | 作用
--- | ---
-f | 用于查看对应依赖结构体中的指定的字段，其默认值就是 `{{.ImportPath}}`，也就是导入路径，因此我们一般不需要进行调整
-json | 显示的格式，若不指定该选项，则会一行行输出。
-u | 显示能够升级的模块信息
-m | 显示当前项目所依赖的全部模块

比如查看 `gin` 框架的版本

```bash
go list -m -versions -json github.com/gin-gonic/gin
```

输出如下：

```bash
{
	"Path": "github.com/gin-gonic/gin",
	"Version": "v1.7.7",
	"Versions": [
		"v1.1.1",
		"v1.1.2",
		"v1.1.3",
		"v1.1.4",
		"v1.3.0",
		"v1.4.0",
		"v1.5.0",
		"v1.6.0",
		"v1.6.1",
		"v1.6.2",
		"v1.6.3",
		"v1.7.0",
		"v1.7.1",
		"v1.7.2",
		"v1.7.3",
		"v1.7.4",
		"v1.7.6",
		"v1.7.7"
	],
	"Time": "2021-11-24T13:54:13Z",
	"Dir": "/Users/pudongping/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7",
	"GoMod": "/Users/pudongping/go/pkg/mod/cache/download/github.com/gin-gonic/gin/@v/v1.7.7.mod",
	"GoVersion": "1.13"
}
```

## 私有库使用 Go Modules 时

- 需要将 `GOPRIVATE` 环境变量设置成你私有库的域名

```shell
# GO111MODULE 设置成 on 或者 auto 都行
GO111MODULE="auto"
# GOPROXY 最好设置成国内镜像地址
GOPROXY="https://goproxy.cn,direct"
# GOPRIVATE 一定要设置成你的私有库域名，比如
GOPRIVATE="gitlab.xxx.com"
```

- 然后执行以下命令即可。**注意：前提是你能够通过 ssh 公钥拉取代码**

```shell
# 以下假设我私有 git 仓库地址为 gitlab.xxx.com:2222
# 那么则需要调整为

cat << EOF >> ~/.gitconfig
[url "ssh://git@gitlab.xxx.com:2222"]
        insteadOf = https://gitlab.xxx.com
EOF

# 或者执行（效果都是一样的）
git config --global url."ssh://git@gitlab.xxx.com:2222".insteadof "https://gitlab.xxx.com"
```

- 测试一下下载一个包

```shell
# 仅仅作为示范，此地址根本就不存在
go get -v gitlab.xxx.com/utils/arrayx
```

## 参考

- [Go Modules 终极入门](https://segmentfault.com/a/1190000021854441)
- [Go Modules 私有不合规库怎么解决引用问题](https://mp.weixin.qq.com/s/sWlTylbW2f1llbz232P2Fw)
- [go modules 使用本地库、合规库、私有库](https://studygolang.com/articles/35234)
- [私有化仓库的 GO 模块使用实践](https://studygolang.com/articles/35235)
