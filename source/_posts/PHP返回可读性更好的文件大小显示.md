---
title: PHP返回可读性更好的文件大小显示
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
abbrlink: 3676e41e
date: 2021-08-16 19:31:56
img:
coverImg:
password:
summary:
---

# 返回可读性更好的文件大小显示

- 方法一

```php

<?php

    /**
     * 返回可读性更好的文件大小
     * 
     * @param $bytes  int 文件大小（字节数）
     * @param int $decimals 保留多少位数
     * @return string 带单位的文件大小字符串
     */
    function human_filesize ($bytes, $decimals = 2)
    {
        $size = ['B', 'kB', 'MB', 'GB', 'TB', 'PB'];
        // 舍去法取整
        $factor = floor((strlen($bytes) - 1) / 3);

        return sprintf("%.{$decimals}f", $bytes / pow(1024, $factor)) . @$size[$factor];
    }

?>

```

### 以上代码做详细解释如下

每相邻的两个存储单位之间以字节数长度的 `3` 倍做间隔，比如看以下表中所示

字节长度 | 字节数 | 换算 | B级 | KB级 | MB级 | GB级
--- | --- | --- | --- | --- | --- | --- |
4 | 1000 | 1000 / 1024 | 0.98 B |  |  |
5 | 10000 | 10000 / 1024 | 9.8 B |  |  |
6 | 10000 0 | 10000 0 / 1024 | 98 B |  |  |
7 | 10000 00 | 10000 00 / 1024 | 980 B | 0.95 KB |  |
8 | 10000 000 | 10000 000 / 1024 | 9800 B | 9.5 KB |  |
9 | 10000 0000 | 10000 0000 / 1024 | 98000 B | 95 KB |  |
10 | 10000 0000 0 | 10000 0000 0 / 1024 | 980000 B | 950 KB | 0.93 MB |
11 | 10000 0000 00 | 10000 0000 00 / 1024 | 980000 0 B | 950 KB | 9.3 MB |
12 | 10000 0000 000 | 10000 0000 000 / 1024 | 980000 00 B | 9500 KB | 93 MB |
13 | 10000 0000 0000 | 10000 0000 0000 / 1024 | 980000B | 95000 KB | 930 MB | 0.91 GB

通过以上的规律，我们可以观察到每两个相邻的量级之间可以通过字节长度的 `3` 倍来划分，减 1 的话类比分页的页码计算，如果字节数长度刚好是 3 的倍数，应该归到前一个量级，比如字节数是 100 bytes，对应长度是 3，显然应该显示为 100B 更好，而不是用 KB 来显示。


- 方法二

```php

<?php

    /**
    * 对文件大小做可读性更好的显示
    *
    * @param integer $bytes    字节数  需要格式转换的字节数
    * @param integer $decimals 保留几位小数
    * @return float  处理好的字节数
    */
    function format_bytes($bytes, $decimals = 2) {
        $units = array('B', 'KB', 'MB', 'GB', 'TB', 'PB');
        for ($i = 0; $bytes >= 1024 && $i < 5; $i++) {
            $size = $bytes /= 1024;
            // 等同于以下代码
            // $size = $bytes = $bytes / 1024;
        }
        return round($size, $decimals) . $units[$i];
    }

>

```

### 调用方式

> 相比而言，方式一会比方式二可读性更加好一些

```php

<?php


    $bytes1 = human_filesize(94875468);  // output: =>  90.48MB
    $bytes2 = format_bytes(94875468);  // output:  =>  90.48MB

    $bytes3 = human_filesize(1024000);  // output:  =>  0.98MB
    $bytes4 = format_bytes(1024000);  // output:  =>  1000KB

    var_dump($bytes1);
    var_dump($bytes2);
    var_dump($bytes3);
    var_dump($bytes4);

>

```
