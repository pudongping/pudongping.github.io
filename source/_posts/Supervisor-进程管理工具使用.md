---
title: Supervisor 进程管理工具使用
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Supervisor
  - 进程管理工具
abbrlink: '21919295'
date: 2021-06-19 17:09:12
img:
coverImg:
password:
summary:
---

# Supervisor 进程管理工具使用

- [pypi 插件链接地址](https://pypi.org/project/supervisor/)
- [官方文档地址](http://supervisord.org/)
- [supervisor 从安装到使用](https://www.jianshu.com/p/3658c963d28b)

> php artisan horizon 和 php artisan queue:work 命令一样，都可以正常处理异步任务  
php artisan horizon:terminate => Horizon 进程会等待当前正在执行的任务执行完毕，然后再退出进程。

## 在阿里云 CentOS 7.6 上

- 安装 supervisor

```sh
yum install -y supervisor
```

- 查看主配置信息 `supervisord.conf` 留意 `include` 选项

```sh
cat /etc/supervisord.conf

# 这里决定了你所需要写的进程配置文件格式，这里是 ini ，因此我们也必须写成 ini 后缀的文件
[include]
files = supervisord.d/*.ini

```

- 添加自定义进程配置信息

vim /etc/supervisord.d/larablog.ini

```sh

[program:larablog-horizon]
process_name=%(program_name)s
command=php /www/wwwroot/larablog/artisan horizon
autostart=true
autorestart=true
user=www
redirect_stderr=true
stdout_logfile=/www/wwwroot/larablog/storage/logs/worker.log

```

- program:larablog-horizon 代表这个配置的名称是 larablog-horizon；
- process_name= 代表这个进程在 Supervisor 内部的命名；
- command= 代表要执行的命令；
- autostart=true 代表这个进程跟随 Supervisor，只要 Supervisor 启动了，就启动这个进程；
- autorestart=true 代表要求 Supervisor 监听进程状态，假如异常退出就再次启动，重启次数默认有 3 次限制；
- user=www 代表以 www 身份启动进程；
- redirect_stderr=true 代表输出错误信息；
- stdout_logfile= 代表将进程的输出保存到日志文件中。


- 更新配置

```sh
sudo supervisorctl update
```

如果遇到报错 `error: <class 'socket.error'>, [Errno 2] No such file or directory: file: /usr/lib64/python2.7/socket.py line: 224` 则执行以下命令

```sh
sudo supervisord -c /etc/supervisord.conf

# 再次尝试执行重载配置命令
sudo supervisorctl update

sudo supervisorctl -c /etc/supervisord.conf
```

- 查看进程状态

```sh
sudo supervisorctl status
```

## 在 ubuntu 上

- 安装 supervisor

```sh
sudo apt-get install supervisor
```

- 添加自定义进程配置信息

vim /etc/supervisor/conf.d/larablog.conf

```sh

[program:larablog-horizon]
process_name=%(program_name)s
command=php /www/wwwroot/larablog/artisan horizon
autostart=true
autorestart=true
user=www
redirect_stderr=true
stdout_logfile=/www/wwwroot/larablog/storage/logs/worker.log

```

- 更新配置

```sh
sudo supervisorctl update
```

- 检查是否正常运行

```sh
sudo supervisorctl status
```

- 单独启动一个指定名称的进程

```sh
sudo supervisorctl start <process-name>

# 比如启动名称为 larablog-horizon 的进程
sudo supervisorctl start larablog-horizon
```
