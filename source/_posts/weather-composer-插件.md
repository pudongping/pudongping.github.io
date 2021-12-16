---
title: weather composer 插件
author: Alex
top: true
hide: false
cover: true
toc: true
mathjax: false
coverImg: https://pudongping.com/medias/banner/2.jpg
summary: 本人基于高德开放平台编写的 PHP 天气信息 composer 组件，并且能够友好的支持 laravel 框架。
categories: 开源
tags:
  - Composer
  - GitHub
  - Laravel
  - PHP
abbrlink: d2a0cef0
date: 2021-06-07 10:10:36
img:
password:
---


<h1 align="center">Weather</h1>

<p align="center">:rainbow: 基于高德开放平台的 PHP 天气信息组件。</p>

## 安装

```sh
$ composer require pudongping/weather -vvv
```

## 配置

在使用本扩展之前，你需要去 [高德开放平台](https://lbs.amap.com/dev/id/newuser) 注册账号，然后创建应用，获取应用的 API Key。


## 使用

```php
use Pudongping\Weather\Weather;

$key = 'xxxxxxxxxxxxxxxxxxxxxxxxxxx';

$weather = new Weather($key);
```

###  获取实时天气

```php
$response = $weather->getLiveWeather('上海');
```

返回值示例：

```json
{
    "status": "1",
    "count": "1",
    "info": "OK",
    "infocode": "10000",
    "lives": [
        {
            "province": "上海",
            "city": "上海市",
            "adcode": "310000",
            "weather": "阴",
            "temperature": "10",
            "winddirection": "东",
            "windpower": "≤3",
            "humidity": "73",
            "reporttime": "2021-03-17 09:31:44"
        }
    ]
}
```

### 获取近期天气预报

```
$response = $weather->getForecastsWeather('上海');
```

返回值示例：

```json
{
    "status": "1",
    "count": "1",
    "info": "OK",
    "infocode": "10000",
    "forecasts": [
        {
            "city": "上海市",
            "adcode": "310000",
            "province": "上海",
            "reporttime": "2021-03-17 10:03:15",
            "casts": [
                {
                    "date": "2021-03-17",
                    "week": "3",
                    "dayweather": "小雨",
                    "nightweather": "小雨",
                    "daytemp": "12",
                    "nighttemp": "9",
                    "daywind": "东北",
                    "nightwind": "东北",
                    "daypower": "4",
                    "nightpower": "4"
                },
                {
                    "date": "2021-03-18",
                    "week": "4",
                    "dayweather": "小雨",
                    "nightweather": "小雨",
                    "daytemp": "13",
                    "nighttemp": "10",
                    "daywind": "东",
                    "nightwind": "东",
                    "daypower": "4",
                    "nightpower": "4"
                },
                {
                    "date": "2021-03-19",
                    "week": "5",
                    "dayweather": "小雨",
                    "nightweather": "小雨",
                    "daytemp": "14",
                    "nighttemp": "9",
                    "daywind": "北",
                    "nightwind": "北",
                    "daypower": "4",
                    "nightpower": "4"
                },
                {
                    "date": "2021-03-20",
                    "week": "6",
                    "dayweather": "阴",
                    "nightweather": "阴",
                    "daytemp": "15",
                    "nighttemp": "8",
                    "daywind": "西北",
                    "nightwind": "西北",
                    "daypower": "4",
                    "nightpower": "4"
                }
            ]
        }
    ]
}
```

### 获取 XML 格式返回值

以上两个方法第二个参数为返回值类型，可选 `json` 与 `xml`，默认 `json`：

```php
$response = $weather->getLiveWeather('上海', 'xml');
```

返回值示例：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<response>
    <status>1</status>
    <count>1</count>
    <info>OK</info>
    <infocode>10000</infocode>
    <lives type="list">
        <live>
            <province>上海</province>
            <city>上海市</city>
            <adcode>310000</adcode>
            <weather>小雨</weather>
            <temperature>10</temperature>
            <winddirection>东北</winddirection>
            <windpower>≤3</windpower>
            <humidity>73</humidity>
            <reporttime>2021-03-17 10:03:15</reporttime>
        </live>
    </lives>
</response>
```

### 参数说明

```
array | string   getLiveWeather(string $city, string $format = 'json')
array | string   getForecastsWeather(string $city, string $format = 'json')
```

> - `$city` - 城市名/[高德地址位置 adcode](https://lbs.amap.com/api/webservice/guide/api/district)，比如：“上海” 或者（adcode：310000）；
> - `$format`  - 输出的数据格式，默认为 json 格式，当 format 设置为 “`xml`” 时，输出的为 XML 格式的数据。


### 在 Laravel 中使用

> laravel version >= 8.x 或者直接使用以上通用方法，以上方法适用于任意 php 框架

执行 `artisan` 命令，发布配置文件

```sh
php artisan vendor:publish --provider="Pudongping\Weather\ServiceProvider"
```

在 Laravel 中使用也是同样的安装方式，配置写在 `config/weather.php` 中，或者自己在 `config` 目录下新建 `weather.php` 文件，填写以下内容亦可。

```php
<?php

return [
    'key' => env('WEATHER_API_KEY'),
];

```

然后在 `.env` 中配置 `WEATHER_API_KEY` ：

```env
WEATHER_API_KEY=xxxxxxxxxxxxxxxxxxxxx
```

可以用两种方式来获取 `Pudongping\Weather\Weather` 实例：

#### 方法参数注入

```php

public function showWeather(Weather $weather) 
{
    $response = $weather->getLiveWeather('上海');
    
    // 或者直接使用城市的 adcode 传参
    // $response = $weather->getLiveWeather('310000');

}

```

#### 服务名访问

```php

public function showWeather()
{
    $response = app('weather')->getLiveWeather('上海');
    
    // 或者直接使用城市的 adcode 传参
    // $response = app('weather')->getLiveWeather('310000');
    
}

```

## 参考

- [高德开放平台天气接口](https://lbs.amap.com/api/webservice/guide/api/weatherinfo/)

## License

MIT
