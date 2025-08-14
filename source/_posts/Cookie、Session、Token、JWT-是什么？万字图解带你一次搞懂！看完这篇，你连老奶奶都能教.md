---
title: Cookie、Session、Token、JWT 是什么？万字图解带你一次搞懂！看完这篇，你连老奶奶都能教
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - Cookie
  - Session
  - Token
  - 安全
  - 授权
abbrlink: 65239a79
date: 2025-08-14 12:01:10
img:
coverImg:
password:
summary:
---

在讲这几个专业术语之前，我们先看一下这样的场景：

> 你去银行准备办理业务，柜台工作人员礼貌地问你：“请问您要办理什么业务？”
> 你说：“我要查询我还有多少余额。”
> 等工作人员处理完之后，你紧接着说：“我还想转账 1000 块钱。”
> 这时候，工作人员一脸茫然地看着你：“请问您是谁？需要办理什么业务？”
> 此时你开始懵逼了，不是刚刚已经办理过业务嘛，怎么还要确认我的身份？
> 但是，你作为一位文明人，你还是耐着性子重新提供了你自己的身份证明。然而，当你转账了 1000 块钱后，你发现你还需要再取 500 块钱现金出来，然后紧接着再对工作人员说：“我需要取 500 块钱现金”。
> 没想到，柜台工作人员又一脸茫然地看着你：“请问您是谁？需要办理什么业务？”
> 你心里开始犯嘀咕了，只怕是这柜台人员是金鱼的记忆吧？你都想给 Ta 一大逼兜了……

是不是觉得这事情听起来就很荒谬？

但是，这就是 HTTP 协议的真实写照 —— **无状态**特性。你每发起一次 HTTP 请求，对服务器来说都是一次全新的请求，它完全不会记得你刚刚做过了什么，就像那个“健忘”的银行柜员一样。

## 什么是 HTTP 无状态？

HTTP 协议被设计为**无状态协议**，这意味着每个请求都是独立的，服务器不会保存任何关于客户端的状态信息。就像你每次打电话给客服，都需要重新报上你的姓名和问题一样。

### 无状态带来的现实问题

在实际的 Web 应用中，无状态特性带来了诸多困扰：

- 用户身份无法维持：用户登录后，下一个请求服务器就"忘记"了用户是谁
- 购物车数据丢失：用户添加商品后，换个页面购物车就空了
- 个性化设置失效：用户的主题、语言偏好等设置无法保存
- 用户体验极差：每次操作都要重新验证身份

> 虽然“无状态”这种特性有很多烦恼，但是这也使得 HTTP 协议更加轻量、并发性能更好。

那有没有办法解决这个问题呢？

当然有！可以用 Cookie。

## Cookie

### 什么是 Cookie？

Cookie 本质上是服务器发送给浏览器的一小段数据，浏览器会将这段数据存储起来，并在后续的请求中自动携带。就像银行给你办了一张会员卡，下次来的时候你亮出会员卡，工作人员就知道你是谁了。

我们来看看它的工作原理：

![cookie流程图](https://upload-images.jianshu.io/upload_images/14623749-bde04e4890b205a9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 实战操作

#### 后端（PHP）设置 Cookie

```php
<?php
// 登录成功后设置Cookie
setcookie("user_id", "123", time() + 3600, "/"); // 有效期1小时
echo "Cookie已设置";
?>
```

#### 前端（JS）读取 Cookie

```js
// 获取指定Cookie值
function getCookie(name) {
  const value = "; " + document.cookie;
  const parts = value.split("; " + name + "=");
  if (parts.length === 2) return parts.pop().split(";").shift();
}
console.log(getCookie("user_id")); // 输出: 123
```

我们可以看到 Cookie 的出现完美解决了 HTTP 无状态的问题，那么，Cookie 有什么弊端吗？

### Cookie 的弊端

#### 1. 安全风险：

- Cookie 存储在客户端，用户可以查看和修改 Cookie 中的所有内容，数据容易被篡改。比如，原本我们设置到 cookie 的 user_id 为 123，我们也可以直接改成 456，这样就可以访问到 user_id 为 456 的用户数据了。
- 敏感信息暴露在客户端不安全
- 容易受到 XSS 和 CSRF 攻击

#### 2. 存储限制：

- 大小限制（通常 4KB）
- 数量限制（每个域名最多几十个）

既然 Cookie 是把敏感数据放到客户端极其不安全，那我们能不能把用户数据（已经授权的数据）放在服务端呢？当然，这也是可以的。这就要说到 Session 了。

## Session

### 什么是 Session？

Session 是直接将用户状态信息存储在服务器端，然后颁发一个 Session ID 给到客户端浏览器（相当于颁发了一把钥匙给到了浏览器），浏览器再通过 Cookie 携带这个 Session ID 去访问服务端数据。

我们来看看 Session 的工作原理：

![Session 工作原理流程图](https://upload-images.jianshu.io/upload_images/14623749-f25d527f1fa94487.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 操作实战

#### 后端（PHP）Session 用法

```php
<?php
session_start(); // 注意，一定要先启动 Session

// 登录成功后
$_SESSION['user_id'] = 123;
echo "Session已设置";
?>
```

#### 前端（JS）无感知，Session ID 由浏览器自动携带

```js
// 前端无需手动处理，浏览器自动携带PHPSESSID
fetch('/user/profile')
  .then(res => res.json())
  .then(data => console.log(data));
```

Session 可以直接将用户数据存储到服务器的任何一个地方，比如：内存、文件、Redis、数据库等……并且成熟的 Session 扩展库还会维护 Session 的生命周期，如：用户不活动（如 30 分钟）超时或主动注销后自动销毁。

## Cookie 和 Session 的对比

对比项 | Session | Cookie
--- | --- | ---
数据存储位置 |	服务器 |	用户浏览器
安全性	| 高（敏感数据不传输）	| 低（可能被窃取）
存储数据类型	| 任意复杂对象	| 仅文本（≤4KB）
生命周期控制	| 由服务器主动管理	| 由浏览器过期时间决定

Session 的出现已经可以解决大部分 Web 端的授权认证了。但是，客户端 APP 呢？也可以使用 Cookie 和 Session 吗？答案，当然是：**可以！**

虽然一些 HTTP 库可以处理 Cookie 和 Session，但是却始终没有在浏览器环境中使用起来那么丝滑的，毕竟在浏览器环境中 Cookie 和 Session 都可以被自动处理掉。并且，如果后端是服务器集群呢？那么我们还得考虑 Session 的存储问题。

那么，有没有一种方式既可以不用存储又能够安全授权认证呢？

既然说到这里，那么肯定是有解决方案的。这就是：Token。

## Token

### 什么是 Token？

Token 是服务器签发的数字令牌，包含用户身份和权限信息，客户端（浏览器/APP）每次请求时自主出示令牌，无需服务器存储会话状态。

我们来看看 Token 的工作原理流程图：

![token工作原理流程图](https://upload-images.jianshu.io/upload_images/14623749-19ab7c6b15c6c634.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

虽然服务端确实是可以不用存储用户信息了，但是客户端还是需要保存一下 Token 的。一般会使用 LocalStorage / APP 内存进行保存。

### JWT 是什么？

JWT 全称为 JSON Web Token，是 Token 的一种规范，也是现在用得最流行的一种 Token 范式。它是一种自包含的令牌格式，将用户信息（Payload）和防伪签名打包成字符串，实现无状态身份验证。

JWT 由三部分组成：

![JWT 结构](https://upload-images.jianshu.io/upload_images/14623749-a765bf3ec348c601.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

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

用户 id，过期时间等数据都保存在 Token 中了，所以并**不需要将 Token 保存在服务器中**，客户端请求的时候在 Header 中携带 Token，服务器获取 Token 后，进行 `base64_decode`  即可获取数据进行校验，由于已经有了签名，所以不用担心数据被篡改。

有了 token 之后该如何验证 token 的有效性，并得到 token 对应的用户呢？我们的常规流程是：

1. 获取客户端提交的 token
2. 检测 token 中的签名 signature 是否正确
3. 判断 payload 数据中的 exp，是否已经过期
4. 根据 payload 数据中的 sub，取数据库中验证用户是否存在
5. 上述检测不正确，则抛出相应异常

为了方便你可以更好的理解这个过程，我们一起来看看下面这个流程图：

![JWT 核心工作流程](https://upload-images.jianshu.io/upload_images/14623749-046c4059f0c8594b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

是不是现在就清晰了许多？

## Cookie vs Session vs Token

| 特性                 | Cookie                          | Session                         | Token (以 JWT 为例)              |
| :------------------- | :------------------------------ | :------------------------------ | :------------------------------ |
| **存储位置**         | **浏览器**                      | **服务器** (Session 存储)       | **客户端** (LocalStorage/Cookie) |
| **存储内容**         | Key-Value 数据                  | 用户会话数据对象                | 自包含的签名令牌 (包含声明)      |
| **状态管理**         | 客户端存储状态片段              |  服务器端管理状态                |  无状态                      |
| **传输方式**         | 自动通过 `Cookie` 请求头        | 通过 Cookie/URL 传递 Session ID | 手动添加 (常放 `Authorization` Header 中) |
| **安全性**           | 较低 (易被窃取/篡改)            | 较高 (敏感数据在服务器)         | 较高 (签名防篡改)               |
| **扩展性 (分布式)**  | 好 (无状态)                     | **差** (需 Session 共享方案)    | **极好** (天然无状态)           |
| **跨域支持**         | 受限 (同源策略/CORS 配置)       | 受限 (依赖 Session ID 传输)     | 好 (配合 CORS)                  |
| **移动端/ Web端**  | 一般浏览器环境为主           |  一般浏览器环境为主       |  浏览器、APP端都可以使用                       |
| **典型应用场景**     | 跟踪、简单配置、Session ID 载体 | 传统服务器渲染 Web 应用         | 现代 SPA、移动 App、API 服务、微服务 |

现在你应该对 Cookie、Session、Token 不再陌生了吧？

创作不易，如果对你有所帮助，还望帮忙点个赞，关注一下，谢谢。