---
title: Go语言中使用JWT鉴权、Token刷新完整示例，拿去直接用！
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
  - JWT
abbrlink: '9e517530'
date: 2024-11-18 11:44:41
img:
coverImg:
password:
summary:
---

在现代 Web 应用中，JWT（JSON Web Token）已经成为了主流的认证与授权解决方案。它轻量、高效、易于实现，并且非常适合于微服务架构。

在本文中，我们将通过 Go 语言及其流行的 Gin 框架，来深入探讨如何使用 JWT 实现用户认证和安全保护。

## 什么是 JWT？

JSON Web Tokens（JWT）是一种开放标准（RFC 7519），用于在网络应用环境间安全地传递声明。JWT是一个紧凑、URL安全的方式，用于在双方之间传递信息。在认证流程中，JWT被用来验证用户身份，并传递用户状态信息。

其结构主要包括三部分：
- **Header**：包含令牌的类型和签名算法。
- **Payload**：携带用户信息（如用户 ID）和一些标准声明（如签发者、过期时间等）。
- **Signature**：用来验证令牌的真实性，防止被篡改。

JWT 的魅力在于它是自包含的，可以通过令牌直接获取用户信息，而无需在服务器端维护会话状态。

## 使用 Gin 和 JWT 实现用户认证

让我们从实际代码开始，演示如何在 Gin 中集成 JWT 认证。

```go
package main

import (
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)
```

在上述代码中，我们首先导入了必要的包，包括用于处理 JWT 的 `github.com/golang-jwt/jwt/v4` 包和用于错误处理的 `github.com/pkg/errors` 包。

## JWT 结构体定义

```go
var (
	ErrTokenGenFailed         = errors.New("令牌生成失败")
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrTokenNotFound          = errors.New("无法找到令牌")
)

// JWT 定义一个 jwt 对象
type JWT struct {
	Key        []byte // 密钥
	MaxRefresh int64  // 最大刷新时间（分钟）
	ExpireTime int64  // 过期时间（分钟）
	Issuer     string // 签发者
}
```

`JWT` 结构体包含了实现 JWT 所需的关键信息，如密钥、最大刷新时间、过期时间和签发者信息。我们使用这些字段来配置和管理 JWT。

## 生成 JWT

```go
func NewJWT(secret, issuer string, maxRefreshTime, expireTime int64) *JWT {
	if maxRefreshTime <= expireTime {
		log.Fatal("最大刷新时间必须大于 token 的过期时间")
	}

	return &JWT{
		Key:        []byte(secret), // 密钥
		MaxRefresh: maxRefreshTime, // 允许刷新时间
		ExpireTime: expireTime,     // token 过期时间
		Issuer:     issuer,         // token 的签发者
	}
}
```

通过 `NewJWT` 方法，我们可以创建一个 JWT 实例。这个实例将用于生成、解析和刷新 JWT。需要注意的是，最大刷新时间必须大于 token 的过期时间，否则会导致逻辑错误。

## 解析 JWT

```go
func (j *JWT) ParseToken(c *gin.Context, userToken ...string) (*JWTCustomClaims, error) {
	var (
		tokenStr string
		err      error
	)

	if len(userToken) > 0 {
		tokenStr = userToken[0]
	} else {
		tokenStr, err = j.GetToken(c)
		if err != nil {
			return nil, err
		}
	}

	token, err := j.parseTokenString(tokenStr)

	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			switch validationErr.Errors {
			case jwtPkg.ValidationErrorMalformed:
				return nil, ErrTokenMalformed
			case jwtPkg.ValidationErrorExpired:
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
```

`ParseToken` 方法用于解析 JWT 并验证其有效性。如果令牌无效或者过期，会返回相应的错误信息。这个方法是我们在各个需要鉴权的 API 接口中最常用的一个方法。

## 刷新 JWT

```go
func (j *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenStr, err := j.GetToken(c)
	if err != nil {
		return "", err
	}

	token, err := j.parseTokenString(tokenStr)

	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}

	claims := token.Claims.(*JWTCustomClaims)
	maxRefreshTime := time.Duration(j.MaxRefresh) * time.Minute

	if claims.IssuedAt > time.Now().Add(-maxRefreshTime).Unix() {
		claims.StandardClaims.ExpiresAt = j.expireAtTime()
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}
```

`RefreshToken` 方法允许在 token 过期但仍在允许刷新时间内时，重新生成一个新的 token。这对于长时间需要保持登录状态的应用非常有用。

这只是刷新 token 的一种思路，还有一种思路也可以刷新 token，但是就需要用到两个 token，一个 access_token 和 refresh_token ，这里我直接将代码贴进来，大家可以参考参考。

```go
package main

import (
	"log"
	"time"

	jwtPkg "github.com/golang-jwt/jwt/v4"
)

type ARJWT struct {
	// 密钥，用以加密 JWT
	Key []byte

	// 定义 access token 过期时间（单位：分钟）即当颁发 access token 后，多少分钟后 access token 过期
	AccessExpireTime int64

	// 定义 refresh token 过期时间（单位：分钟）即当颁发 refresh token 后，多少分钟后 refresh token 过期
	// 一般来说，refresh token 的过期时间会比 access token 的过期时间长
	RefreshExpireTime int64

	// token 的签发者
	Issuer string
}

func NewARJWT(secret, issuer string, accessExpireTime, refreshExpireTime int64) *ARJWT {
	if refreshExpireTime <= accessExpireTime {
		log.Fatal("refresh token 过期时间必须大于 access token 过期时间")
	}
	return &ARJWT{
		Key:               []byte(secret),    // 密钥
		AccessExpireTime:  accessExpireTime,  // access token 过期时间
		RefreshExpireTime: refreshExpireTime, // refresh token 过期时间
		Issuer:            issuer,            // token 的签发者
	}
}

// GenerateToken 生成 access token 和 refresh token
func (arj *ARJWT) GenerateToken(userId string) (accessToken, refreshToken string, err error) {
	// 生成 access token 在 access token 中需要包含我们自定义的字段，比如用户 ID
	mc := JWTCustomClaims{
		UserID: userId,
		StandardClaims: jwtPkg.StandardClaims{
			// ExpiresAt 是一个时间戳，代表 access token 的过期时间
			ExpiresAt: time.Now().Add(time.Duration(arj.AccessExpireTime) * time.Minute).Unix(),
			// 签发人
			Issuer: arj.Issuer,
		},
	}

	// 生成 access token
	accessToken, err = jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, mc).SignedString(arj.Key)
	if err != nil {
		log.Printf("generate access token failed: %v \n", err)
		return "", "", err
	}

	// 生成 refresh token
	// refresh token 只需要包含标准的声明，不需要包含自定义的声明
	refreshToken, err = jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, jwtPkg.StandardClaims{
		// ExpiresAt 是一个时间戳，代表 refresh token 的过期时间
		ExpiresAt: time.Now().Add(time.Duration(arj.RefreshExpireTime) * time.Minute).Unix(),
		// 签发人
		Issuer: arj.Issuer,
	}).SignedString(arj.Key)

	return
}

func (arj *ARJWT) ParseAccessToken(tokenString string) (*JWTCustomClaims, error) {
	claims := new(JWTCustomClaims)

	token, err := jwtPkg.ParseWithClaims(tokenString, claims, func(token *jwtPkg.Token) (interface{}, error) {
		return arj.Key, nil
	})

	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			switch validationErr.Errors {
			case jwtPkg.ValidationErrorMalformed:
				return nil, ErrTokenMalformed
			case jwtPkg.ValidationErrorExpired:
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	if _, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

func (arj *ARJWT) RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// 先判断 refresh token 是否有效
	if _, err = jwtPkg.Parse(refreshToken, func(token *jwtPkg.Token) (interface{}, error) {
		return arj.Key, nil
	}); err != nil {
		return
	}

	// 从旧的 access token 中解析出 JWTCustomClaims 数据出来
	var claims JWTCustomClaims
	_, err = jwtPkg.ParseWithClaims(accessToken, &claims, func(token *jwtPkg.Token) (interface{}, error) {
		return arj.Key, nil
	})
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		// 当 access token 是过期错误，并且 refresh token 没有过期时就创建一个新的 access token 和 refresh token
		if ok && validationErr.Errors == jwtPkg.ValidationErrorExpired {
			// 重新生成新的 access token 和 refresh token
			return arj.GenerateToken(claims.UserID)
		}
	}

	return
}

```

关于这两种刷新 token 的方式对比，可以直接参考阅读我这里的文章，`https://github.com/pudongping/golang-tutorial/tree/main/project/jwt_demo` 有比较详细的说明。

## 结语

通过本文，我们探索了如何在 Go 中使用 Gin 框架实现 JWT 鉴权，包括 token 的生成、解析、刷新等功能。这套方案不仅高效而且易于扩展，可以满足大多数 Web 应用的安全需求。

完整的代码在这里：`https://github.com/pudongping/golang-tutorial/blob/main/project/jwt_demo/jwt.go`