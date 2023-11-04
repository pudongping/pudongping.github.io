---
title: 认识下 JWT
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - PHP
  - JWT
abbrlink: '22722346'
date: 2023-11-04 13:26:50
img:
coverImg:
password:
summary:
---

# 认识下 JWT

[JWT](https://jwt.io/) 是 JSON Web Token 的缩写，是一个非常轻巧的规范，这个规范允许我们使用 JWT 在用户和服务器之间传递安全可靠的信息。
JWT 由头部（header）、载荷（payload）与签名（signature）组成，一个 JWT 类似下面这样：

```bash

{
    "typ":"JWT",
    "alg":"HS256"
}
{
    "iss":"http://localhost",
    "iat":1587000625,
    "exp":1618536625,
    "nbf":1587000625,
    "jti":"iCxsfo97UVUijjjP",
    "sub":1,
    "prv":"13e8d028b391f3b7b63f21933dbad458ff21072e"
}
signature

```


- 头部声明了加密算法；
- 载荷中有两个比较重要的数据，`exp` 是过期时间，`sub` 是 JWT 的主体，这里就是用户的 id；
    - aud（Audience）：受众，也就是接受 JWT 的一方。
    - exp（ExpiresAt）：所签发的 JWT 过期时间，过期时间必须大于签发时间。
    - jti（JWT Id）：JWT 的唯一标识。
    - iat（IssuedAt）：签发时间
    - iss（Issuer）：JWT 的签发者。
    - nbf（Not Before）：JWT 的生效时间，如果未到这个时间则为不可用。
    - sub（Subject）：主题
- 最后的 signature 是由服务器进行的签名，保证了 token 不被篡改。

signature 签名的生成公式示例如下：

```bash
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

> JWT 最后是通过 Base64 编码的，也就是说，它可以被翻译回原来的样子来的。所以不要在 JWT 中存放一些敏感信息。

用户 id，过期时间等数据都保存在 Token 中了，所以并不需要将 Token 保存在服务器中，客户端请求的时候在 Header 中携带 Token，服务器获取 Token 后，进行 `base64_decode`  即可获取数据进行校验，由于已经有了签名，所以不用担心数据被篡改。

### Token 验证

有了 token 之后该如何验证 token 的有效性，并得到 token 对应的用户呢？其实原理很简单，Laravel 为我们准备好了 `auth` 这个中间件

1. 获取客户端提交的 token
2. 检测 token 中的签名 signature 是否正确
3. 判断 payload 数据中的 exp，是否已经过期
4. 根据 payload 数据中的 sub，取数据库中验证用户是否存在
5. 上述检测不正确，则抛出相应异常

### 安装 jwt-auth

[jwt-auth](https://github.com/tymondesigns/jwt-auth) 是 Laravel 和 lumen 的 JWT 组件，首先来安装一下。

```bash

composer require tymon/jwt-auth

```

安装完成后，我们需要设置一下 JWT 的 secret，这个 secret 很重要，用于最后的签名，更换这个 secret 会导致之前生成的所有 token 无效。

```bash

php artisan jwt:secret

```

可以看到在 .env 文件中，增加了一行 `JWT_SECRET`。
修改 config/auth.php，将 `api guard` 的 `driver` 改为 `jwt`。

config/auth.php

```php

'guards' => [
    'web' => [
        'driver' => 'session',
        'provider' => 'users',
    ],

    'api' => [
        'driver' => 'jwt',
        'provider' => 'users',
    ],
],

```

user 模型需要继承 `Tymon\JWTAuth\Contracts\JWTSubject` 接口，并实现接口的两个方法 `getJWTIdentifier()` 和 `getJWTCustomClaims()`。

app\Models\User.php


```php

<?php

namespace App\Models;

use Auth;
use Spatie\Permission\Traits\HasRoles;
use Tymon\JWTAuth\Contracts\JWTSubject;
use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Auth\MustVerifyEmail as MustVerifyEmailTrait;
use Illuminate\Contracts\Auth\MustVerifyEmail as MustVerifyEmailContract;

class User extends Authenticatable implements MustVerifyEmailContract, JWTSubject


    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    public function getJWTCustomClaims()
    {
        return [];
    }
}

```

`getJWTIdentifier` 返回了 User 的 id，`getJWTCustomClaims` 是我们需要额外在 JWT 载荷中增加的自定义内容，这里返回空数组。打开 tinker，执行如下代码，尝试生成一个 token。

```php

$user = User::first();
Auth::guard('api')->login($user);

```


`jwt-auth` 有两个重要的参数，可以在 .env 中进行设置

- `JWT_TTL` 生成的 token 在多少分钟后过期，默认 60 分钟
- `JWT_REFRESH_TTL`   生成的 token，在多少分钟内，可以刷新获取一个新 token，默认 20160 分钟，14 天。


这里需要理解一下 JWT 的过期和刷新机制，过期很好理解，超过了这个时间，token 就无效了。刷新时间一般比过期时间长，只要在这个刷新时间内，即使 token 过期了， 依然可以换取一个新的 token，以达到应用长期可用，不需要重新登录的目的。
