---
title: PHP RSA 加解密
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - RSA
abbrlink: f7924170
date: 2021-08-27 20:27:18
img:
coverImg:
password:
summary:
---


# PHP RSA 加解密

最近在做支付宝的跨境支付，自己写了一个 composer 包，里面涉及到 RSA 加密，以及验签，故分享之。也方便以后自己随时拿过来用 😃

```php

<?php
/**
 * RSA 加解密
 *
 * 可以线上生成一对公私钥的网站： http://www.metools.info/code/c80.html
 */

/**
 * RSA 签名
 *
 * @param string $data 待签名数据
 * @param string $privateKeyPath 私钥文件路径
 * @return string
 */
function rsaSign(string $data, string $privateKeyPath): string
{
    $priKey = file_get_contents($privateKeyPath);
    $res = openssl_get_privatekey($priKey);
    openssl_sign($data, $sign, $res);
    openssl_free_key($res);
    $sign = base64_encode($sign);
    return $sign;
}

/**
 * RSA 验签
 *
 * @param string $data 待签名数据
 * @param string $aliPublicKeyPath 公钥文件路径
 * @param string $sign 要校对的签名数据
 * @return bool
 */
function rsaVerify(string $data, string $aliPublicKeyPath, string $sign): bool
{
    $pubKey = file_get_contents($aliPublicKeyPath);
    $res = openssl_get_publickey($pubKey);
    $result = (bool)openssl_verify($data, base64_decode($sign), $res);
    openssl_free_key($res);
    return $result;
}

/**
 * RSA 解密
 *
 * @param string $content 需要解密的密文数据
 * @param string $privateKeyPath  私钥文件路径
 * @return string  解密后的明文内容
 */
function rsaDecrypt(string $content, string $privateKeyPath): string
{
    $priKey = file_get_contents($privateKeyPath);
    $res = openssl_get_privatekey($priKey);
    // 用 base64 将内容还原成二进制
    $content = base64_decode($content);
    // 把需要解密的内容，按 128 位拆开解密
    $result = '';
    for ($i = 0; $i < strlen($content) / 128; $i++) {
        $data = substr($content, $i * 128, 128);
        openssl_private_decrypt($data, $decrypt, $res);
        $result .= $decrypt;
    }
    openssl_free_key($res);
    return $result;
}

// 假设需要对如下这段字符串加解密
$str = 'aa=123&bb=456&cc=789&dd=1212';

// osrx0siWqK0+5C6ANNk/2pqEYoWa74UzFsUPFfv5FnhrOU9abyup+h2AY/4LqlSnvH9ztBcx//EpdJI9yI/xnfB14LdiDrPH1bJUJ5oJMafAo4QxL47eAPKT8ZKbufKg+YTf8kx7xnJ5kSyxcHzzhZyVvth4pGstFTUeL/5RpGRxlOQj/viHLkocYM2h1hzunqvcHWKzQqTdmi8g9atxDLPrASdTdolsD0TYbFj4Bn1S/0ziomcesz3IFi0CO6UsSM2N1jBOdtmhrecxv6WcUAgy3y9B7o4vF3hGG7HuhD437bO0XVWMrJ2NcRHhiQolMi6zmeX50ZLUSI63ve7ucA==
$mySign = rsaSign($str, './private_key.pem');

$is = rsaVerify($str, './public_key.pem', $mySign);

$decryptStr = rsaDecrypt($mySign, './private_key.pem');

var_dump($mySign);
echo PHP_EOL;
var_dump($is);
echo PHP_EOL;
var_dump($decryptStr);

```
