---
title: 如何在Go语言中实现表单验证？整一个validator吧！
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
abbrlink: '59486883'
date: 2024-11-20 11:19:49
img:
coverImg:
password:
summary:
---

在现代 Web 开发中，表单验证和错误处理是至关重要的环节，尤其是在多语言环境下。

本文将通过一个实际的示例，演示如何使用 Go 语言的 Gin 框架结合 `validator` 包，实现高级的表单验证功能，并且支持国际化（i18n）的错误信息提示。

## 背景与需求

假设我们正在开发一个用户注册功能，需要对用户提交的信息进行严格的验证。例如，用户名不能为空、邮箱格式必须正确、密码和确认密码必须一致、用户年龄应在合理范围内（如 1 到 130 岁），并且日期字段不能早于当前日期。除此之外，系统还需要根据用户的语言偏好提供相应语言的错误提示信息。

## 代码示例

我们将从以下几个方面展开：

1. **表单数据的结构定义**
2. **表单验证器的初始化与自定义**
3. **多语言支持的实现**
4. **处理表单提交与错误返回**

```go
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器
var trans ut.Translator
```

### 表单数据结构定义

首先，我们定义用户提交的表单数据结构 `SignUpParam`。这个结构体中包含了用户注册时所需的各个字段，并通过结构标签（tags）指定了验证规则。

```go
type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Date       string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}
```

- `Age` 字段必须在 1 到 130 岁之间。
- `Name` 字段不能为空。
- `Email` 字段必须是有效的电子邮件地址。
- `Password` 和 `RePassword` 字段必须一致。
- `Date` 字段需要使用自定义校验方法 `checkDate`，确保输入日期晚于当前日期。

### 初始化与自定义表单验证器

在 Gin 框架中，我们可以通过 `binding.Validator.Engine()` 获取到内置的验证器，并对其进行自定义。

在下面的代码中，我们完成了翻译器的初始化，并注册了自定义的标签名称和验证方法。

```go
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册获取 JSON tag 的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 注册结构体级别的验证函数
		v.RegisterStructValidation(SignUpParamStructLevelValidation, SignUpParam{})

		// 注册自定义校验方法
		if err := v.RegisterValidation("checkDate", customFunc); err != nil {
			return err
		}

		// 初始化多语言支持
		zhT := zh.New() 
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

                var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册语言翻译
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			return err
		}

		// 注册自定义翻译
		if err := v.RegisterTranslation(
			"checkDate",
			trans,
			registerTranslator("checkDate", "{0}必须晚于当前日期"),
			translate,
		); err != nil {
			return err
		}
		return
	}
	return
}
```

### 实现自定义校验逻辑

在上面的代码中，我们自定义了两个校验函数：

1. **customFunc**：用于校验日期是否晚于当前日期。
2. **SignUpParamStructLevelValidation**：用于校验两个密码字段是否一致。

```go
func customFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return date.After(time.Now())
}

func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(SignUpParam)
	if su.Password != su.RePassword {
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
```

### 处理多语言错误提示

为了确保错误信息能够根据用户的语言偏好正确返回，我们注册了一个自定义的翻译函数 `registerTranslator`，并在验证失败时使用该函数对错误信息进行翻译。

```go
// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
```

### 主程序逻辑

最后，我们在 Gin 中处理用户的注册请求。当用户提交的数据验证失败时，系统会自动返回翻译后的错误提示信息。

```go
// removeTopStruct 去除字段名中的结构体名称标识
// refer from:https://github.com/go-playground/validator/issues/633#issuecomment-654382345
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func main() {
	// 初始化翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("初始化翻译器失败: %v\n", err)
		return
	}

	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"msg": removeTopStruct(errs.Translate(trans))})
			return
		}

		// 其他的一些业务逻辑操作……

		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("服务器运行失败: %v\n", err)
	}
}
```

## 总结

本文通过一个完整的示例，展示了如何在 Go 语言中使用 Gin 框架实现多语言的表单验证。

我们不仅探讨了基础的验证规则，还介绍了如何自定义验证逻辑以及如何实现国际化的错误提示。这种方式使得我们的应用程序不仅在功能上更加强大，同时也能更好地适应全球化的需求。