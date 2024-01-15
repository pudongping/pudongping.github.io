---
title: 搭建 GitLab Git 服务器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
  - Git 服务器
abbrlink: 34b62620
date: 2024-01-15 14:00:58
img:
coverImg:
password:
summary:
---

> [官方安装文档](https://about.gitlab.com/install/?version=ce#centos-7)

GitLab是Ruby开发的自托管的Git项目仓库，可通过Web界面访问公开的或者私人的项目。

> 版本：GitLab 分为社区版（CE）和企业版（EE）  
> 此次安装的 gitlab 版本为：14.10.0

## Ubuntu 系统下

- 更新软件源

```shell
# 推荐使用 apt 工具，其实使用 apt-get 也行
apt update && apt upgrade
```

- 安装依赖包，如果系统已经安装，就不需要再次安装

```shell
apt install -y curl openssh-server ca-certificates tzdata perl
```

- 添加 gitlab 软件包仓库

```shell

# 这里添加的是企业版（EE）地址
curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ee/script.deb.sh | sudo bash

```

- 安装 Postfix 来发送通知邮件

```
apt install -y postfix
```

- 安装 gitlab 企业版（EE）

```shell
# 安装企业版 gitlab
apt install gitlab-ee

# 查看可用的版本
sudo apt-cache madison gitlab-ee
```

- 修改配置文件

```shell
# 配置很简单，一看就懂，就不做介绍
vim /etc/gitlab/gitlab.rb

# 需要注意的是：需要将 `external_url` 后面的值修改成你自己访问 gitlab 的地址
```

- 重新编译配置文件

```shell
gitlab-ctl reconfigure
```

- 重启 gitlab

```shell
gitlab-ctl restart
```

- 查看默认的 root 密码

```shell
cat /etc/gitlab/initial_root_password
```

## CentOS 7 系统下

- 安装依赖包，如果系统已经安装，就不需要再次安装

```shell

# 安装 curl 和 ssh 程序，（一般系统都会自带，因此可以忽略）
# policycoreutils-python 依赖包不影响部署 gitlab 因此也可以不安装
sudo yum install -y curl policycoreutils-python openssh-server

```

- 配置 ssh 服务

```shell

# 启动 ssh 服务
sudo systemctl start sshd

# 设置 ssh 服务为开机自动启动
sudo systemctl enable sshd

```

- 配置 Postfix 来发送通知邮件

```shell

# 安装 Postfix
sudo yum install postfix

# 设置 Postfix 服务为开机自动启动
sudo systemctl enable postfix

```

启动 Postfix 服务

```shell

# 打开 Postfix 的配置文件
vim /etc/postfix/main.cf

# 修改 `inet_interfaces = localhost` 为 `inet_interfaces = all`


# 启动 Postfix 服务
sudo systemctl start postfix
```

- 添加 gitlab 软件包仓库

```shell

# 这里添加的是社区版（CE）地址
curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash

```

- 安装 gitlab 社区版（CE）

```
sudo EXTERNAL_URL="你的 GitLab 服务器的公网 IP 地址或者域名" yum install -y gitlab-ce

# 比如：
sudo EXTERNAL_URL="http://127.0.0.1:8099" yum install -y gitlab-ce
```

- 使用浏览器访问 `你的 GitLab 服务器的公网 IP 地址或者域名` ，账号为 `root` 密码通过以下命令进行查看初始密码

```shell
cat /etc/gitlab/initial_root_password
```

## 如果不想做过多的配置时，直接执行以下两步骤即可：

```shell

curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash

# 默认安装的是最新版
sudo EXTERNAL_URL="你的 GitLab 服务器的公网 IP 地址或者域名" yum install -y gitlab-ce

# 如果需要安装指定版本的话 
yum install -y gitlab-ce-14.10.1-ce.0.el7.x86_64
# 或者
yum install -y gitlab-ce-14.10.1-ce.0.el7

# 查看可用的版本
yum --showduplicates list gitlab-ce
```

出现以下界面，则表示安装成功

![安装成功界面](https://upload-images.jianshu.io/upload_images/14623749-fd85b0c1dd3a1e17.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### gitlab 常用命令

```
# 启动所有 gitlab 组件
gitlab-ctl start

# 停止所有 gitlab 组件
gitlab-ctl stop

# 重启所有 gitlab 组件
gitlab-ctl restart

# 查看服务状态
gitlab-ctl status

# 修改 gitlab 配置文件
vim /etc/gitlab/gitlab.rb

# 重新编译 gitlab 的配置
gitlab-ctl reconfigure 

# 查看 gitlab 配置信息
gitlab-ctl show-config

# 检查 gitlab
gitlab-rake gitlab:check SANITIZE=true --trace

# 查看日志
gitlab-ctl tail
gitlab-ctl tail nginx/gitlab_access.log

# 日志地址 /var/log/gitlab
# 服务地址 /var/opt/gitlab
# nginx 配置路径 /var/opt/gitlab/nginx/conf/nginx.conf

# 查看 gitlab 的版本
cat /opt/gitlab/embedded/service/gitlab-rails/VERSION
```

### 重置账号密码

> [官方文档重置用户密码](https://docs.gitlab.com/ee/security/reset_user_password.html)

#### 使用 Rake 任务

```shell
gitlab-rake "gitlab:password:reset"
```

如下所示：

```shell
[root@iZ447rdnsv9nreZ ~]# gitlab-rake "gitlab:password:reset"
Enter username: root
Enter password: 
Confirm password: 
Password successfully updated for user with username root.
[root@iZ447rdnsv9nreZ ~]# 
```

#### 使用 Rails 控制台

```
# 打开交互模式
gitlab-rails console -e production
# 如果是低版本的 gitlab 可以尝试使用以下命令
gitlab-rails console production

# 定位到 gitlab 数据库中 Users 表中的一个用户
user = User.where(id:1).first
# 或者通过用户名进行查找
user = User.find_by_username 'root'
# 或者通过邮箱地址查找
user = User.find_by(email: 'user@example.com')
# 如果用户 id、用户名、邮箱地址都不记得，那么可以尝试使用以下命令查看所有用户
User.all

# 密码必须不小于 8 个字符，密码后面加不加单引号都可以
# 重置管理员密码为 12345678
user.password=12345678

# 确认管理员的密码为 12345678
user.password_confirmation=12345678

# 保存更改信息，必须要有后面的感叹号
user.save!
```

如下所示：

```shell
[root@iZ447rdnsv9nreZ gitlab]# gitlab-rails console -e production
--------------------------------------------------------------------------------
 Ruby:         ruby 2.7.5p203 (2021-11-24 revision f69aeb8314) [x86_64-linux]
 GitLab:       14.10.0 (88da5554d96) FOSS
 GitLab Shell: 13.25.1
 PostgreSQL:   12.7
-----------------------------------------------------------[ booted in 606.21s ]
Loading production environment (Rails 6.1.4.7)
irb(main):005:0> user = User.where(id:1).first
=> #<User id:1 @root>
irb(main):006:0> user.password=12345678
=> 12345678
irb(main):007:0> user.password_confirmation=12345678
=> 12345678
irb(main):008:0> user.save!
=> true
irb(main):009:0> 
```

或者这样：

```shell
[root@iZ447rdnsv9nreZ ~]# gitlab-rails console -e production
--------------------------------------------------------------------------------
 Ruby:         ruby 2.7.5p203 (2021-11-24 revision f69aeb8314) [x86_64-linux]
 GitLab:       14.10.0 (88da5554d96) FOSS
 GitLab Shell: 13.25.1
 PostgreSQL:   12.7
------------------------------------------------------------[ booted in 41.21s ]
Loading production environment (Rails 6.1.4.7)
irb(main):001:0> User.all
=> #<ActiveRecord::Relation [#<User id:1 @root>]>
irb(main):002:0> user = User.find_by_username 'root'
=> #<User id:1 @root>
irb(main):003:0> user.password='alex123456'
=> "alex123456"
irb(main):004:0> user.password_confirmation='alex123456'
=> "alex123456"
irb(main):005:0> user.save!
=> true
irb(main):006:0> exit
```


### 备份与恢复

需要注意，两台服务器上的 gitlab 必须版本一致，才可以进行恢复。可以使用以下命令进行查看 gitlab 的版本号

```shell
cat /opt/gitlab/embedded/service/gitlab-rails/VERSION
```

#### 备份

```shell
# 执行以下命令生成备份文件
/usr/bin/gitlab-rake gitlab:backup:create
```
生成的备份文件位于 `/var/opt/gitlab/backups` 目录。
可以通过修改配置文件 `vim /etc/gitlab/gitlab.rb` 中的 `gitlab_rails['backup_path']` 项，来调整默认的备份目录。

生成的文件类似于 `1651390965_2022_05_01_14.10.0_gitlab_backup.tar`


#### 恢复

> 账号密码、git 仓库、以及个人设置啥的都会被恢复过来

```shell
# 停止 unicorn 和 sidekiq，保证数据库没有新的连接，不会有写数据情况
sudo gitlab-ctl stop unicorn
sudo gitlab-ctl stop sidekiq

# 进入备份目录进行恢复，1651390965_2022_05_01_14.10.0 为备份文件的时间以及 gitlab 版本号
cd /var/opt/gitlab/backups
gitlab-rake gitlab:backup:restore BACKUP=1651390965_2022_05_01_14.10.0

# 启动 unicorn 和 sidekiq
sudo gitlab-ctl start unicorn
sudo gitlab-ctl start sidekiq
```

### 降级

降级只建议在次要版本和补丁版本之间降级。[官方文档](https://docs.gitlab.com/ee/update/package/downgrade.html) 比如以下将 `14.10.1` 版本降低到 `14.10.0` 版本

1. 停止 gitlab 并删除当前包

```shell

# If running Puma
sudo gitlab-ctl stop puma

# Stop sidekiq
sudo gitlab-ctl stop sidekiq

# If on Ubuntu: remove the current package
sudo dpkg -r gitlab-ce

# If on Centos: remove the current package
sudo yum remove gitlab-ce

```

2. 确定要降级到的 gitlab 版本

```shell

# (Replace with gitlab-ce if you have GitLab FOSS installed)

# Ubuntu
sudo apt-cache madison gitlab-ce

# CentOS:
sudo yum --showduplicates list gitlab-ce

```

3. 将 gitlab 降级到所需版本

```shell

# (Replace with gitlab-ce if you have GitLab FOSS installed)

# Ubuntu
sudo apt install gitlab-ce=14.10.0-ce.0

# CentOS:
sudo yum install gitlab-ce-14.10.0-ce.0.el7

```

4. 重新配置 gitlab

```shell

sudo gitlab-ctl reconfigure

```