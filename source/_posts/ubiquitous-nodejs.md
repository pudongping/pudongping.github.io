---
title: ubiquitous-nodejs
author: Alex
top: true
hide: false
cover: true
toc: true
mathjax: false
coverImg: https://pudongping.com/medias/banner/12.jpg
summary: "⚡ ubiquity-nodejs 是一个本人基于 node.js 开发的 web 脚手架。\U0001F618 支持模版渲染、Restful API、ORM 等特性，遵循 MVC 架构。"
categories: 开源
tags:
  - GitHub
  - Node.js
abbrlink: 6b5490ac
date: 2021-06-07 10:27:15
img:
password:
---

<p align="center">
    <h1 align="center"><a href="https://pudodngping.com">ubiquitous-nodejs</a></h1>
    <p align="center">⚡ ubiquity-nodejs 是一个基于 node.js 的 web 脚手架。😘 支持模版渲染、Restful API、ORM 等特性，遵循 MVC 架构。</p>
</p>

## 如何部署项目？

仓库地址：
- [GitHub](https://github.com/pudongping/ubiquitous-nodejs.git)
- [Gitee](https://gitee.com/pudongping/ubiquitous-nodejs.git)

本地部署
1. 直接使用 git 拉取项目源码，并进入项目根目录

```shell
git clone https://gitee.com/pudongping/ubiquitous-nodejs.git ubiquitous && cd ./ubiquitous
```

2. 安装项目依赖包

```shell
npm install
```

3. 填写配置信息

```shell
# 复制开发环境配置文件
cp ./config/development.js.example ./config/development.js

# 复制生产环境配置文件
cp ./config/production.js.example ./config/production.js

# 复制测试环境配置文件
cp ./config/test.js.example ./config/test.js
```

4. 导入初始数据表
> 这里只是为了演示 restful api 的 CRUD 功能，故此创建了一个初始数据库。

```shell
# 连接数据库
mysql -h <your-mysql-host> -u <your-mysql-account> -p <your-mysql-password>;

# 导入初始数据表
source ./database/ubiquitous.sql;
```

5. 启动项目
> 以下方式执行任意一条命令即可

```shell
# 直接使用 node 执行项目入口文件的方式
node --use_strict app.js

# 使用 npm 执行
# 运行开发环境
npm start
# 或者
npm run dev

# 运行生产环境
npm run prod

# 运行测试环境
npm run test
```

6. 搭建完毕！enjoy it! ✌️
> 可以打开浏览器访问 `localhost:9500` 或者 `127.0.0.1:9500` 进行访问，即可看到首页欢迎语。

## 关于项目目录

```shell

├── LICENSE  License
├── README.md  项目说明文档
├── app.js  项目入口文件
├── bootstrap  项目启动目录
│   ├── boot.js  启动文件
│   ├── db.js  数据库封装
│   ├── rest.js  restful api 方法封装
│   ├── static-files.js  静态文件处理方法封装
│   └── templating.js  模版处理方法封装
├── config  配置文件目录
│   ├── development.js  开发环境配置文件
│   ├── development.js.example  开发环境配置文件模版
│   ├── production.js.example  生产环境配置文件模版
│   └── test.js.example  测试环境配置文件模板
├── constants  常量文件
│   └── ErrorCode.js  api 错误码常量
├── controllers  控制器目录
│   ├── auth  auth 模块目录
│   │   └── user_controller.js  用户控制器（这里作为演示 restful api 写的 demo）
│   └── home_controller.js  首页控制器 （这里作为演示模版调用写的 demo）
├── database  数据文件目录
│   └── ubiquitous.sql  初始化数据库文件
├── lib  工具目录
│   ├── api_error.js  自定义错误异常类
│   └── helper.js  助手函数
├── loader.js  项目加载文件（这里定义项目全局变量）
├── models  模型目录
│   ├── WebSite.js  站点模型文件
│   ├── auth  auth 模块模型目录
│   │   └── User.js  用户模型文件
│   └── model.js  自动化扫描加载所有的模型
├── package-lock.json  插件包描述锁文件
├── package.json  插件包描述文件
├── routes  路由目录
│   ├── api.js  restful api 路由目录
│   └── web.js  web 路由目录
├── services  服务层目录
│   └── auth  auth 模块服务层目录
│       └── user_service.js  用户服务层文件
├── static  静态文件目录
│   ├── css  样式文件目录
│   │   ├── googleapis-fonts.css
│   │   └── iview.css
│   ├── fonts  字体文件目录
│   └── js  js 文件目录
│       ├── iview.min.js
│       └── vue.min.js
└── views  视图层文件目录
    ├── base.html  视图基础文件
    ├── home  home 模块视图文件目录
    │   └── hello.html
    └── home.html  首页视图文件

```

## 项目所使用依赖包

插件包 | 作用
--- | ---
[koa](https://koa.bootcss.com/) | 使用 Koa2 作为 web 框架
[koa-router](https://github.com/koajs/router) | 处理 url
[koa-bodyparser](https://github.com/koajs/bodyparser) | 解析原始 request 的 body
[nunjucks](https://mozilla.github.io/nunjucks/cn/templating.html) | 模版引擎
[cross-env](https://github.com/kentcdodds/cross-env#readme) | 环境脚本的跨平台设置
[mz](https://github.com/normalize/mz#readme) | 支持 Promise 的 fs 模块
[mime](https://github.com/broofa/mime#readme) | 读取文件的 mime
[mysql2](https://github.com/sidorares/node-mysql2#readme) | Node.js 的 mysql 驱动程序
[sequelize](https://www.sequelize.com.cn/) | Node.js 的 ORM 框架
[moment](http://momentjs.cn/) | 日期处理

## 使用演示
- 如果需要使用模版引擎的方式，请查看 `controllers/home_controller.js` 文件。
- 如果需要使用 api 的方式，请查看 `controllers/auth/user_controller.js` 文件。

## 感谢支持

如果你觉得本项目对你有所帮助,请帮忙给个 `Star`。  
如果你想贡献一份力量,欢迎提交 `Pull Request`。
