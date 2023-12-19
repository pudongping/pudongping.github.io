---
title: 一文讲懂 HTTP 协议知识
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - HTTP
abbrlink: d8b21c9b
date: 2023-12-19 16:34:20
img:
coverImg:
password:
summary:
---

# HTTP协议

## HTTP是什么？

一种规范，统一了双方的沟通方法。   
浏览器中所显示的内容都是需要从服务器下载到客户端才可以在浏览器中显示。（每刷新一次网页都会产生下载。）

## 两大重点。

- 请求体（request）
    - 发什么？  
      请求报文。
- 响应体（response）
    - 收什么？
      响应报文。

## http 格式

每个 http 请求和响应都遵循相同的格式，一个 http 包含 header 和 body 两部分，其中 body 是可选的。

### request

```bash
// 第一行为请求行
// GET 请求
// 请求的url
// http/1.1 协议版本
GET http://sodevel.com/course/res/95/104 HTTP/1.1

// 以下都是请求头
Host: sodevel.com
Proxy-Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Cookie: XSRF-TOKEN=eyJpdiI6IlphTmxTVXROaGkySTdYdk9xUnc5cUE9PSIsInZhbHVlIjoidWVVRDBiSWxzYVB6ZDk3VmlQSzhuSGF6THI3R2lzOGVodUxXM1I2aWt4WWdmTGw0R1JXbkFvY0NId00xdlhHSSIsIm1hYyI6IjVhMGFhNDY0NjExMTgxZjU4ZGM4NTM0ZGIwYjdiZmY2ZjMzN2YzY2U5YmZjN2NjZmZmZjNjMTg1NjhjNzZiZTQifQ%3D%3D; laravel_session=eyJpdiI6IkJYR3NSUXRtTHVURzFSVXRyZDNHXC93PT0iLCJ2YWx1ZSI6IjhscEdQc2IxMjNEaTBzMjRlNXM5eE1QOUNOeXdENXBTMWpBMFdoMk5oSzc0N1REVGVUVmdWYmhrc05MWEpZSWMiLCJtYWMiOiJmNzcwNjYwYjZmNDFlYzlmOWIwZjQ1MDM0MzA5NzQ2NjYzNTUzMzI5ZjIzMmIzYTAyMjQ4M2UzZGZhNGEzOWZhIn0%3D


// 当遇到连续的两个 \r\n 时，header 部分结束，后面的数据全部是 body
```

![请求流程](https://upload-images.jianshu.io/upload_images/14623749-505d729ac85418fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### reponse

```bash
// 响应行
// 协议版本
// 请求状态
HTTP/1.1 200 OK

// 响应头开始
Server: nginx
Content-Type: text/html; charset=UTF-8
Transfer-Encoding: chunked
Connection: close
Vary: Accept-Encoding
X-Powered-By: PHP/7.2.6
Cache-Control: no-cache, private
Date: Mon, 31 Dec 2018 06:54:59 GMT
Set-Cookie: XSRF-TOKEN=eyJpdiI6IjFUVHVvOVwvMURRT1BTbmFzczhNTWhnPT0iLCJ2YWx1ZSI6IkhNdmtmcTcyWGNlSUNRTjVDU0hUcU1QenY3K1pDcHAzdTVHUGFXMVNhMlp4clg3bVNidzArMzhhaE1wSHJVVXciLCJtYWMiOiIxZmNhZmY2OGYwYmEzNzFhZDhlOGViYjE4YzRjNmExMTE2MjhiODI5M2Q2NWFmZjM2NWM2MDAwYzg0NzJkOTRlIn0%3D; expires=Mon, 31-Dec-2018 08:54:59 GMT; Max-Age=7200; path=/
Set-Cookie: laravel_session=eyJpdiI6ImtqdG9sa0lIZ3hNRWhkNGpFM08rSXc9PSIsInZhbHVlIjoiSkZjeWlOcCtVdks4VlFHRVB5KzRoUzZQZFUyMWd1YzRDNmdmUHhUU1IwVEhyaUZ5eXdcL0NWREk5THFEQndzNEkiLCJtYWMiOiI1ZmViNDU0ZDVkNDFlYzU0ZjZjMGRjNGYyYWFjZGFjYjg2OGMyZWFlODI3ZDY1NzBhOWE5MjAyMmU5ZWVkNjQ1In0%3D; expires=Mon, 31-Dec-2018 08:54:59 GMT; Max-Age=7200; path=/; httponly
Content-Encoding: gzip
Proxy-Connection: keep-alive

// 响应头结束
// 当遇到连续的两个 \r\n 时，响应头部分结束，后面的数据全部是 body
```

![响应流程](https://upload-images.jianshu.io/upload_images/14623749-911fe20dbfb82de8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**这里需要注意的是：**

`响应头和响应正文（标准的网页html代码或者图片）中间有一个空行，只不过Chrome浏览器将中间的空行`

```php
<?php
$fp = fsockopen("www.example.com", 80, $errno, $errstr, 30);
if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {

// 一个 \r\n 代表换行
    $out = "GET / HTTP/1.1\r\n";
    $out .= "Host: www.example.com\r\n";
    $out .= "Connection: Close\r\n\r\n";
    fwrite($fp, $out);
    while (!feof($fp)) {
        echo fgets($fp, 128);
    }
    fclose($fp);
}
?> 
```

### 状态码：

> 1-- 已经接受请求，处理中  
2-- 请求成功  
3-- 涉及跳转网址  
4-- 客户端错误，比如请求了不存在的网址  
5-- 服务器错误，服务器端可能出现问题

- 100~199：表示服务端成功接收客户端请求，要求客户端继续提交下一次请求才能完成整个处理过程。
- 200~299：表示服务端成功接收请求并已完成整个处理过程。最常用就是：200
- 300~399：为完成请求，客户端需进一步细化请求。比较常用的如：客户端请求的资源已经移动一个新地址使用 302 表示将资源重定向，客户端请求的资源未发生改变，使用 304，告诉客户端从本地缓存中获取。
- 400~499：客户端的请求有错误，如：404 表示你请求的资源在 web 服务器中找不到，403 表示服务器拒绝客户端的访问，一般是权限不够。
- 500~599：服务器端出现错误，最常用的是：500

### 注意点：

```bash
setcookie(); 和 header(); 前不能有任何输出。

原因是：这两个函数是用于设置头信息的，http协议必须是先设置头信息后才会有输出内容。
```

![http协议](https://upload-images.jianshu.io/upload_images/14623749-69402ab4b1507951.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![http协议](https://upload-images.jianshu.io/upload_images/14623749-d334afce5c97f48b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

1. 浏览器通过http协议发起请求到服务器（Apache或者nginx）
2. 服务器（Apache或者nginx）根据不同的需求交给php文件处理。
3. php可能会需要请求MySQL
4. 如果需要请求mysql，mysql将结果返回给php文件
5. php文件将结果返还给服务器（Apache或者nginx）
6. 服务器（Apache或者nginx）将最终结果通过http协议发送给浏览器。


### http协议特点

1. 无状态：不能记住本次请求产生的数据，不同请求之间，数据不能共享。
2. 有会话：session和cookie。
3. 缓存：通过Cache-Control标记，可通知浏览器如何缓存该页面。
4. 同源：浏览器为了防止跨域攻击，多数要求同源策略，HTTP可以通过标记来开放限制。
5. 认证：通过Authenticate标记，可实现基于http的认证。
6. 代理

- ### 资料文档
> https://www.cnblogs.com/ranyonsue/p/5984001.html  
或者
https://developer.mozilla.org/zh-CN/docs/Web/HTTP

- ### http 动词
动词 |	描述 |	是否幂等
--- | --- | ---
GET	 | 获取资源，单个或多个	 | 是
POST |	创建资源 |	否
PUT	| 更新资源，客户端提供完整的资源数据 |	是
[PATCH](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/PATCH) |	更新资源，客户端提供部分的资源数据 |	否
DELETE |	删除资源 |	是

`幂等性`，指一次和多次请求某一个资源应该具有同样的副作用，也就是一次访问与多次访问，对这个资源带来的变化是相同的。

> 为什么 PUT 是幂等的而 PATCH 是非幂等的，因为 PUT 是根据客户端提供了完整的资源数据，客户端提交什么就替换为什么，而 PATCH 有可能是根据客户端提供的参数，动态的计算出某个值，例如每次请求后资源的某个参数减1，所以多次调用，资源会有不同的变化。

### http 状态码

> http 状态码详细讲解：https://httpstatuses.com/  
json.api 格式规范化：http://jsonapi.org.cn/format/

状态码 | 状态 | 说明
--- | --- | ---
200  | OK  | 对成功的 GET、PUT、PATCH 或 DELETE 操作进行响应。也可以被用在不创建新资源的 POST 操作上
201 | Created | 对创建新资源的 POST 操作进行响应。应该带着指向新资源地址的 Location 头
202 | Accepted | 服务器接受了请求，但是还未处理，响应中应该包含相应的指示信息，告诉客户端该去哪里查询关于本次请求的信息
204 | No Content | 对不会返回响应体的成功请求进行响应（比如 DELETE 请求）
304 | Not Modified | HTTP缓存header生效的时候用
400 | Bad Request | 请求异常，比如请求中的body无法解析
401 | Unauthorized | 没有进行认证或者认证非法
403 | Forbidden | 服务器已经理解请求，但是拒绝执行它
404 | Not Found | 请求一个不存在的资源
405 | Method Not Allowed | 所请求的 HTTP 方法不允许当前认证用户访问
410 | Gone | 表示当前请求的资源不再可用。当调用老版本 API 的时候很有用
415 | Unsupported Media Type | 如果请求中的内容类型是错误的
422 | Unprocessable Entity | 用来表示校验错误
429 | Too Many Requests | 由于请求频次达到上限而被拒绝访问

### http 提交数据常用的有两种方式

1. application/x-www-form-urlencoded(默认值)
2. multipart/form-data （form 表单提交文件的时候，需要增加 enctype="multipart/form-data"）
