---
title: Hyperf 框架中将数据导出到 csv 文件中
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - Hyperf
  - Swoole
abbrlink: d0c857b0
date: 2022-11-09 23:15:52
img:
coverImg:
password:
summary:
---

# Hyperf 框架中将数据导出到 csv 文件中

直接看代码吧……

```php

<?php

declare(strict_types=1);

namespace App\Controller;

use Hyperf\HttpServer\Annotation\AutoController;

/**
 * @AutoController
 * Class TestController
 * @package App\Controller
 */
class TestController extends AbstractController
{

    /**
     * 请求测试：
     * curl 'http://127.0.0.1:9511/test/csvExport' >> test.csv
     *
     * @return \Psr\Http\Message\ResponseInterface
     */
    public function csvExport()
    {
        $fileName = 'test.csv';
        $data = [
            ['name', 'age'],
            ['阿猫', '11'],
            ['阿狗', '22'],
            ['阿猪', '33'],
        ];

        $str = '';
        foreach ($data as $v) {
            $str .= mb_convert_encoding(implode(',', $v), 'GBK', 'UTF-8') . "\n";
        }

        return (new \Hyperf\HttpServer\Response())
            ->withHeader('content-type', 'text/csv')
            ->withHeader('content-disposition', "attachment; filename={$fileName}")
            ->withHeader('content-transfer-encoding', 'binary')
            ->withBody(new \Hyperf\HttpMessage\Stream\SwooleStream($str));
    }

}    

```

## 测试

第一种：  
直接使用浏览器访问 `http://127.0.0.1:9511/test/csvExport`

第二种：  
直接通过 `curl` 命令行访问 `curl 'http://127.0.0.1:9511/test/csvExport' >> test.csv`
