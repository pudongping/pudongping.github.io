---
title: 一起来看看如何在Go中使用Swagger？
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
abbrlink: 53c8a188
date: 2024-05-09 00:57:07
img:
coverImg:
password:
summary:
---

# Swagger go 的使用

## http 服务

### 安装

```bash

go get -u github.com/swaggo/swag/cmd/swag@v1.7.8
go get -u github.com/swaggo/gin-swagger@v1.3.3
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template

```

### 验证是否安装成功

```bash

swag -v

# swag version v1.7.8

```

### 写入注解

注解 |	描述
--- | ---
@Summary	| 摘要
@Produce	| API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	| 参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	| 响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Failure	| 响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
@Router	| 路由，从左到右分别为：路由地址，HTTP 方法

```go

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

```

### 针对项目写注解

```go

// @title gin-blog-service 博客系统
// @version 1.0
// @description gin-blog-service 学习 gin 写的一个博客系统
// @termsOfService https://github.com/pudongping/gin-blog-service
func main() {
    
}

```

### 生成文档

```bash

swag init

```

### 添加 swagger 访问路由

> 这里以 `gin` 框架中使用为例

```go

import (

    "github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

    // 初始化 docs 包，内含有 swagger 生成的文档
	_ "github.com/pudongping/gin-blog-service/docs"
)


func NewRouter() *gin.Engine {

	r := gin.New()

    // 如果有额外需求，可以手动指定访问地址
	// swaggerUrl := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}

```


## grpc 服务

### 安装 protoc 插件 protoc-gen-swagger

> protoc-gen-swagger 的作用是通过 proto 文件来生成 swagger 定义（.swagger.json）

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

### 下载 Swagger UI 文件

Swagger 提供可视化的接口管理平台，也就是 Swagger UI，我们首先需要到 https://github.com/swagger-api/swagger-ui 上将其源码压缩包下载下来，接着在项目的 third_party 目录下新建 swagger-ui 目录，将其 dist 目录下的所有资源文件拷贝到我们项目的 third_party/swagger-ui 目录中去。

### 静态资源转换

```bash
# 将静态文件转换为 go 代码
go get -u github.com/go-bindata/go-bindata/...
```

在项目的 `pkg` 目录下新建 `swagger` 目录，并在项目根目录下执行以下转换命令

```bash
# 在执行完毕后，应当在项目的 pkg/swagger 目录下创建了 data.go 文件
go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...
```

### Swagger UI 处理和访问

为了让刚刚转换的静态资源代码能够让外部访问到，我们需要安装 go-bindata-assetfs 库，它能够结合 net/http 标准库和 go-bindata 所生成 Swagger UI 的 Go 代码两者来供外部访问

```bash
go get -u github.com/elazarl/go-bindata-assetfs/...
```

引入并使用

```go

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/pudongping/go-grpc-service/pkg/swagger"
)

func runHttpServer() *http.ServeMux {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

    // 代码大致调整这里 start
	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	httpMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	
	httpMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)

		http.ServeFile(w, r, p)
	})	
    // 代码大致调整这里 end	

	return httpMux
}

```

重新运行服务，通过浏览器访问 `http://127.0.0.1:8004/swagger-ui/` 即可访问 swagger 面板

访问自己的 swagger 文档接口 `http://127.0.0.1:8004/swagger/tag.swagger.json`
