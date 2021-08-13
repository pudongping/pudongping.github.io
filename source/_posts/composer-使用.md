---
title: composer 使用
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Composer
abbrlink: 9b3c7295
date: 2021-08-13 23:38:21
img:
coverImg:
password:
summary:
---

# composer 使用

## [语义化版本](http://semver.org)
### 版本号的组成
major.minor.patch => 主版本号.次版本号.修订号
- major：通常会发生 api 变更，不向后兼容
- minor：新增功能，但是向后兼容
- patch：补丁，向后兼容，修复 bug

### 版本符号
- `~`：指定向后兼容的最小版本
    - `~1.2` 等于 >=1.2 && <2.0.0
    - `~1.2.3` 等于 >=1.2.3 && <1.3.0
- `^`：允许大版本前的所有版本
    - `^1.2` 等于 >=1.2 && <2.0.0
    - `^1.2.3` 等于 >=1.2.3 && <2.0 （区别在这里）
### 使用版本号
- 确切版本：1.0.2
- 范围：>=1.0、>=1.0 <2.0、>=1.0 <1.1 || >=1.2
- 连字符范围：1.0 - 2.0
- 通配符：1.0.*
- 波浪运算符：~1.2.3
- ^运算符：^1.2.3

> [packagist 包管理官方网站](https://packagist.org/)

## 安装

### linux 下安装 composer

```shell

# 直接从阿里云镜像中下载 composer.phar
wget https://mirrors.aliyun.com/composer/composer.phar

# 或者直接从官网下载 composer.phar 二进制执行文件
wget https://getcomposer.org/download/2.0.11/composer.phar

# 或者使用 composer 中国镜像安装器安装
curl -sS http://install.phpcomposer.com/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer

# 添加执行权限
chmod u+x composer.phar

# 将 composer.phar 移动到系统环境变量中，以方便在任意位置都可以直接执行 composer 命令
mv composer.phar /usr/local/bin/composer
# 如果没有 /usr/local/bin 目录，则需要自己创建
mkdir -p /usr/local/bin

# 如果你只想为你的用户安装 composer 并避免需要 root 权限，那么你可以直接执行以下命令，前提是你得有 ~/local/bin 目录
mv composer.phar ~/local/bin

# 添加阿里云镜像加速
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer
# 或者添加中国镜像
composer config -g repo.packagist composer https://packagist.phpcomposer.com

# 解除镜像
composer config -g --unset repos.packagist

```

### windows 下安装 composer

```
# 在 D 盘中新建一个 composer 目录

# 直接从官网下载 composer.phar 二进制执行文件
wget https://getcomposer.org/download/2.0.11/composer.phar

# 使用 php.exe 执行 composer.phar 测试是否正常
I:\xampp\php\php.exe composer.phar -V

# 简化命令
echo @I:\xampp\php\php.exe "d:\composer\composer.phar" %*>composer.bat

# 将 D:/composer 目录添加到环境变量中去

# 查看 composer 版本号，确定是否已经安装成功
composer -V （大写的 V）
```


### composer.json 和 composer.lock
当项目中存在 composer.lock 文件时，使用 `composer install` 命令安装依赖时，`composer.lock` 都会解析并安装你在 `composer.json` 中所列出来的依赖，但是 `composer` 会严格使用 `composer.lock` 文件所列出来的版本以确保项目中得所有成员所安装的版本都是一致的。**需要将 composer.lock** 文件提交到代码管理中。


### 自动加载

执行 `composer install` 命令时，会生成 `vendor/autoload.php`  文件，只需 include 这个文件就可以使用这些包所提供的类

```php

require __DIR__ . '/vendor/autoload.php';

$log = new Monolog\Logger('name');
$log->pushHandler(new Monolog\Handler\StreamHandler('app.log', Monolog\Logger::WARNING));
$log->addWarning('Foo');

```

你也可以直接在 `composer.json` 文件中添加一个 `autoload` 指令来添加自己的自动加载声明

```php

{
    "autoload": {
        "psr-4": {"Acme\\": "src/"}
    }
}

```

composer 会为 Acme 命名空间注册一个 `PSR-4` 的自动加载。
你定义一个命名空间指向目录的映射。 在 vendor 目录同级的 src 目录将成为你项目的根目录。一个案例，文件名 src/Foo.php 需包含 Acme\Foo 类。
添加 `autoload` 指令之后，你必需重新运行 `composer dump-autoload`来重新生成 vendor/autoload.php 文件。

### [composer 命令的使用](https://learnku.com/docs/composer/2018/03-cli/2084)

- 查看 composer 全局配置

```
composer config -l -g
```

- 初始化 composer.json

```
# 创建扩展包时，初始化 composer.json
composer init
```

- 安装所有扩展包

```
# 会读取当前目录的 composer.json 文件，解决依赖关系，并将他们安装到 vendor 文件夹中
composer install

# 制定 composer 略过 require-dev 选项里的扩展包，只加载需要的扩展包
composer install --no-dev
```

- 安装指定的扩展包

```
composer require <package-name>
# 比如安装 vendor/package:2.* 扩展包
composer require vendor/package:2.*
```

- 更新扩展包（更新命令执行完毕后，同时也会更新 composer.lock 文件，不要乱用 composer update ，如果需要更新包，建议直接使用 composer require 命令）

```
# 更新所有的依赖包
composer update

# 更新指定的包
composer update <package-name>
# 比如更新 vendor/package 包
composer update vendor/package
# 一次性更新多个包
composer update example/example1 example/example2
# 更新指定包到指定版本
composer require example/example3:1.2.3
```

- 移除扩展包

```
composer remove <package-name>
# 比如移除 vendor/package:2.*  扩展包
composer remove vendor/package:2.* 
```

- 搜索扩展包

```
composer search <package-name>
# 比如搜索 monolog 扩展包
composer search monolog
```

- 列出所有的可用的包

```
# 查看所有的包信息
composer show

# 查看包详细信息
composer show <package-name>
# 比如查看 monolog/monolog 包详细信息
composer show monolog/monolog 1.0.2
```

- 验证 composer.json 文件是否合法

```
composer validate
```

- 更新 composer

```
# 更新 composer 到最新版本
composer self-update

# 更新 composer 到指定版本
composer self-update 1.0.0

# 回滚到上一个安装版本
composer self-update -r
```

