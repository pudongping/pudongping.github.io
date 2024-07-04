---
title: 如何有效提升 API 接口的安全性？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - 安全
abbrlink: 5edede5c
date: 2024-07-04 23:00:08
img:
coverImg:
password:
summary:
---

在如今的互联网应用中，API 接口的安全性已经成为了开发过程中不可忽视的一环。越来越多的应用服务通过 API 进行数据交换，而 API 的安全性直接关系到应用的数据安全与用户隐私。因此，我们需要采取一系列措施来确保 API 的安全。

本文将教你如何通过 API 接口验证规则和接口防刷来提升 API 的安全性。

## API 接口验证规则

API 接口验证是 API 安全的基础。有效的验证机制可以阻止未授权的访问尝试，并确保只有拥有正确凭据的请求才能成功执行。

以下是一种常见且有效的 API 接口验证方法：

1. **请求参数排序**：将所有请求参数按 `ASCII` 码的顺序进行排序。这一步是为了确保发送到服务端的参数在前后端保持一致性，以便于生成可验证的签名(`sign`)。

2. **生成查询字符串**：将排序后的请求参数转换成 `key=value` 的形式，并使用 `&` 连接，形成查询字符串。如 `a=111&b=222`。在此基础上，还可以加上一个与后端开发人员协定好的密钥(`key`)，以增加验证的复杂度。

比如：

```bash
GET /api/data?a=111&b=222&key=secret
```

3. **MD5 加密**：对步骤 2 生成的查询字符串进行 MD5 加密，生成签名(`sign`)。

```php
$sign = md5("a=111&b=222&key=secret");
```

4. **客户端携带签名请求**：客户端在发送请求时，将加密后的签名(`sign`)一同携带发送。

5. **服务端验证**：服务端接收到请求后，按照相同的规则（步骤 1 - 3）对请求进行处理并生成新的 `sign`，然后与客户端发送过来的 `sign` 进行对比。如果两者一致，则验证通过，否则请求无效。

PHP 示例代码如下：

```php
<?php
// 示例代码，假设请求参数是一个关联数组
$params = [
    'b' => 222,
    'a' => 111
];

// 1. 对请求参数进行 ASCII 排序
ksort($params);

// 2. 转成 `a=111&b=222` 这样的结构
$queryString = http_build_query($params);

// 3. 进行 MD5 加密，生成 sign
$secretKey = 'your_secret_key'; // 与后端协定好的 key
$sign = md5($queryString . '&key=' . $secretKey);

// 4. 客户端请求携带参数以及 sign
// 假设这是客户端的请求
$request = [
    'params' => $params,
    'sign' => $sign
];

// 后端验证逻辑示例
function verifyRequest($request, $secretKey) {
    // 获取参数和 sign
    $params = $request['params'];
    $clientSign = $request['sign'];

    // 1. 对请求参数进行 ASCII 排序
    ksort($params);

    // 2. 转成 `a=111&b=222` 这样的结构
    $queryString = http_build_query($params);

    // 3. 进行 MD5 加密，生成新的 sign
    $serverSign = md5($queryString . '&key=' . $secretKey);

    // 4. 比较两个 sign
    return $clientSign === $serverSign;
}

$secretKey = 'your_secret_key'; // 与前端协定好的 key
$isRequestValid = verifyRequest($request, $secretKey);

if ($isRequestValid) {
    echo "请求合法";
} else {
    echo "请求非法";
}
```

## 接口防刷

接口被恶意刷取不仅会消耗服务器资源，还可能导致数据被不当获取。为了防止接口被恶意调用，通常会采用一些防刷策略，比如限制请求频率、使用验证码等。

其中，一种简单有效的防刷策略是利用 Redis 设置请求指纹的过期时间，限制同一签名（`sign`）或同一用户在短时间内的请求频率。

当一个请求被处理后，可以将该请求的签名存入 Redis，并设置一个过期时间，例如 1 小时。如果在 1 小时内再次收到相同的签名请求，则可以认为是重复请求，拒绝处理。~~这里的时间可以按照具体情况设置短一点儿也行。~~

```php
<?php
// 假设已经连接到 Redis 服务器
$redis = new \Redis();
$redis->connect('127.0.0.1', 6379);

// 生成一个唯一的 sign，通常可以使用请求参数的哈希值
$sign = md5(http_build_query($params) . '&key=' . $secretKey);

// 检查请求是否已经存在于 Redis 中
if ($redis->exists($sign)) {
    echo "请求过于频繁，请稍后再试";
} else {
    // 设置 Redis 键值对，带过期时间（秒）
    $expirationTime = 3600; // 1 小时
    $redis->setex($sign, $expirationTime, 'alex');
    // 处理正常请求逻辑
    // ...
}
```

在上述代码中，我们使用 Redis 的 `setex` 方法将请求的 `sign` 作为键，值设置为 `alex`，并设置过期时间为 1 小时。

每次请求时，先检查该 `sign` 是否存在于 Redis 中，如果存在，则认为请求过于频繁，拒绝处理。如果不存在，则将签名存入 Redis，并设置过期时间为 1 小时。这样，即使同一个签名的请求在 1 小时内重复发送，服务器也能正确地拒绝处理。

## 结语

API 接口的安全防护是确保数据安全和服务质量的关键。通过实现严格的验证规则和采用接口防刷措施，我们可以有效地保护我们的 API 免受恶意攻击和滥用。

通过上述介绍，我们学习了如何通过 API 接口验证与接口防刷两种方式来增强我们 API 的安全性。这两种方法在实际开发中非常实用，能有效防止 API 被滥用，保护数据安全。

希望这篇文章能够帮助你在日后的开发工作中更好地保护 API。
