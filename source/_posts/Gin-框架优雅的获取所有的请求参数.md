---
title: Gin 框架优雅的获取所有的请求参数
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
  - Gin
abbrlink: ccb188ca
date: 2022-05-14 15:41:49
img:
coverImg:
password:
summary:
---

# Gin 框架优雅的获取所有的请求参数

嗯，话不多说，直接看代码 ^_^

将以下代码写入某个 `.go` 文件中，比如，我这里写入 `gin_demo.go` 文件中，然后执行 `go run gin_demo.go`

## 主要代码

```go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/foo", func(c *gin.Context) {
		inputs, err := RequestInputs(c)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
		} else {
			data := gin.H{
				"method": c.Request.Method,
				"params": inputs,
			}
			fmt.Printf("返回值为 ====> %#v \n", data)
			c.JSON(200, data)
		}
	})

	r.Run(":9501")
}

// RequestInputs 获取所有参数
func RequestInputs(c *gin.Context) (map[string]interface{}, error) {

	const defaultMemory = 32 << 20
	contentType := c.ContentType()

	var (
		dataMap  = make(map[string]interface{})
		queryMap = make(map[string]interface{})
		postMap  = make(map[string]interface{})
	)

	// @see gin@v1.7.7/binding/query.go ==> func (queryBinding) Bind(req *http.Request, obj interface{})
	for k := range c.Request.URL.Query() {
		queryMap[k] = c.Query(k)
	}

	if "application/json" == contentType {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// @see gin@v1.7.7/binding/json.go ==> func (jsonBinding) Bind(req *http.Request, obj interface{})
		if c.Request != nil && c.Request.Body != nil {
			if err := json.NewDecoder(c.Request.Body).Decode(&postMap); err != nil {
				return nil, err
			}
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	} else if "multipart/form-data" == contentType {
		// @see gin@v1.7.7/binding/form.go ==> func (formMultipartBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	} else {
		// ParseForm 解析 URL 中的查询字符串，并将解析结果更新到 r.Form 字段
		// 对于 POST 或 PUT 请求，ParseForm 还会将 body 当作表单解析，
		// 并将结果既更新到 r.PostForm 也更新到 r.Form。解析结果中，
		// POST 或 PUT 请求主体要优先于 URL 查询字符串（同名变量，主体的值在查询字符串的值前面）
		// @see gin@v1.7.7/binding/form.go ==> func (formBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseForm(); err != nil {
			return nil, err
		}
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			if err != http.ErrNotMultipart {
				return nil, err
			}
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	}

	var mu sync.RWMutex
	for k, v := range queryMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}
	for k, v := range postMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}

	return dataMap, nil
}

```

## 测试

- get 请求，只有 query 参数时

```shell
curl --location --request GET '127.0.0.1:9501/foo?field1=value1&field2=false&field3=12&field4=4.78'

# 返回值为 ====> gin.H{"method":"GET", "params":map[string]interface {}{"field1":"value1", "field2":"false", "field3":"12", "field4":"4.78"}}
```

- get 请求，有 query 参数，且有 `multipart/form-data` 参数时

```shell
curl --location --request GET '127.0.0.1:9501/foo?field1=value1&field2=false&field3=12&field4=4.78' \
--form 'f1="vv1"' \
--form 'f2="vv2"' \
--form 'f3="false"' \
--form 'f4="7.08"' \
--form 'f5="张三"' \
--form 'f6="88"' \
--form 'f7[]="hello"' \
--form 'f7[]="world"' \
--form 'field1="覆盖了 field1"'

# 返回值为 ====> gin.H{"method":"GET", "params":map[string]interface {}{"f1":"vv1", "f2":"vv2", "f3":"false", "f4":"7.08", "f5":"张三", "f6":"88", "f7[]":[]string{"hello", "world"}, "field1":"覆盖了 field1", "field2":"false", "field3":"12", "field4":"4.78"}}

#{
#   "method": "GET",
#   "params": {
#       "f1": "vv1",
#       "f2": "vv2",
#       "f3": "false",
#       "f4": "7.08",
#       "f5": "张三",
#       "f6": "88",
#       "f7[]": [
#           "hello",
#           "world"
#       ],
#       "field1": "覆盖了 field1",
#       "field2": "false",
#       "field3": "12",
#       "field4": "4.78"
#   }
#}

```

- post 请求，既有 query 参数，也有 `application/x-www-form-urlencoded` 参数时

```shell
curl --location --request POST '127.0.0.1:9501/foo?field1=value1&field2=false&field3=12&field4=4.78' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'f1=vv1' \
--data-urlencode 'f2=vv2' \
--data-urlencode 'f3=false' \
--data-urlencode 'f4=7.08' \
--data-urlencode 'f5=张三' \
--data-urlencode 'f6=88' \
--data-urlencode 'f7[]=hello' \
--data-urlencode 'f7[]=world' \
--data-urlencode 'field1=覆盖了 field1'

# 返回值为 ====> gin.H{"method":"POST", "params":map[string]interface {}{"f1":"vv1", "f2":"vv2", "f3":"false", "f4":"7.08", "f5":"张三", "f6":"88", "f7[]":[]string{"hello", "world"}, "field1":"覆盖了 field1", "field2":"false", "field3":"12", "field4":"4.78"}}

#{
#   "method": "POST",
#   "params": {
#       "f1": "vv1",
#       "f2": "vv2",
#       "f3": "false",
#       "f4": "7.08",
#       "f5": "张三",
#       "f6": "88",
#       "f7[]": [
#           "hello",
#           "world"
#       ],
#       "field1": "覆盖了 field1",
#       "field2": "false",
#       "field3": "12",
#       "field4": "4.78"
#   }
#}

```

- post 请求，既有 query 参数，也有 `application/json` 参数时

```shell

curl --location --request POST '127.0.0.1:9501/foo?field1=value1&field2=false&field3=12&field4=4.78' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "张t",
    "account": "",
    "introduction": null,
    "avatar": "/avatar/2022/05/11/ruQvUHOF-6512bd43d9caa6e02c990b0a82652dca-20220511111053.jpg",
    "favorite": [
        "football",
        "basketball",
        "pingpong-ball"
    ],
    "num": [
        1,
        5,
        8
    ],
    "books": {
        "name": "bookA",
        "price": 18.98
    },
    "students": [
        {
            "name": "张三",
            "age": 18,
            "height": 180.00
        },
        {
            "name": "李四",
            "age": 24.5,
            "height": 165.55
        }
    ],
    "field1": "覆盖了 field1"
}'

# 返回值为 ====> gin.H{"method":"POST", "params":map[string]interface {}{"account":"", "avatar":"/avatar/2022/05/11/ruQvUHOF-6512bd43d9caa6e02c990b0a82652dca-20220511111053.jpg", "books":map[string]interface {}{"name":"bookA", "price":18.98}, "favorite":[]interface {ootball", "basketball", "pingpong-ball"}, "field1":"覆盖了 field1", "field2":"false", "field3":"12", "field4":"4.78", "introduction":interface {}(nil), "name":"张t", "num":[]interface {}{1, 5, 8}, "students":[]interface {}{map[string]interface {}{"age":18, "height", "name":"张三"}, map[string]interface {}{"age":24.5, "height":165.55, "name":"李四"}}}}

#{
#    "method": "POST",
#    "params": {
#        "account": "",
#        "avatar": "/avatar/2022/05/11/ruQvUHOF-6512bd43d9caa6e02c990b0a82652dca-20220511111053.jpg",
#        "books": {
#            "name": "bookA",
#            "price": 18.98
#        },
#        "favorite": [
#            "football",
#            "basketball",
#            "pingpong-ball"
#        ],
#        "field1": "覆盖了 field1",
#        "field2": "false",
#        "field3": "12",
#        "field4": "4.78",
#        "introduction": null,
#        "name": "张t",
#        "num": [
#            1,
#            5,
#            8
#        ],
#        "students": [
#            {
#                "age": 18,
#                "height": 180,
#                "name": "张三"
#            },
#            {
#                "age": 24.5,
#                "height": 165.55,
#                "name": "李四"
#            }
#        ]
#    }
#}

```

以上代码通过阅读 `gin` 框架 `ShouldBind` 方法绑定参数逻辑编写。经过粗略测试，均可以获取所有的参数，并将结果绑定到 `map[string]interface{}` 中。
