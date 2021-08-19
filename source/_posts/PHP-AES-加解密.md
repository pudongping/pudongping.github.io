---
title: PHP AES 加解密
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - AES
  - 加解密
abbrlink: ebe0be72
date: 2021-08-19 15:53:12
img:
coverImg:
password:
summary:
---

# PHP AES 加解密

```php

class Aesmcrypt
{

    /**
     * MCRYPT_ciphername 常量中的一个，或者是字符串值的算法名称。
     */
    const CIPHER = MCRYPT_RIJNDAEL_128;
    /**
     * MCRYPT_MODE_modename 常量中的一个，或以下字符串中的一个：
     * "ecb"，"cbc"，"cfb"，"ofb"，"nofb" 和 "stream"。
     */
    const MODE = MCRYPT_MODE_CBC;

    /**
     * 加密密钥。如果密钥长度不是该算法所能够支持的有效长度，则函数将会发出警告并返回 FALSE
     */
    const KEY = 'bcb04b7e103a0cd8b54763051';


    /**
     * aes 加密
     *
     * @param string $str 需要加密的字符串
     * @param string $key 加密用到的 key
     * @return string 加密后的字符串
     */
    public function encryptAes ($str = '', $key = self::KEY)
    {
        $keyArr = $this->getLegalKey($key);
        $encrypted = mcrypt_encrypt(self::CIPHER, $keyArr['key'], $str, self::MODE, $keyArr['iv']);
        return base64_encode($encrypted);
    }

    /**
     * aes 解密
     *
     * @param string $str 加密后的字符串
     * @param string $key 加密用到的 key
     * @return string 解密后的字符串
     */
    public function decryptAes ($str = '', $key = self::KEY)
    {
        $str = base64_decode($str);
        $keyArr = $this->getLegalKey($key);
        $decrypted = rtrim(mcrypt_decrypt(self::CIPHER, $keyArr['key'], $str, self::MODE, $keyArr['iv']), "\0");
        return $decrypted;
    }


    /**
     * 生成有效的 key
     *
     * @param $key 原始 key
     * @return array
     */
    private function getLegalKey ($key)
    {
        $key = substr(md5($key), 0, 32); // 必须为16，24，32个字符
        $iv = substr(md5($key), 0, 16);
        return array('key' => $key, 'iv' => $iv);
    }


}

```
