---
title: samba 实现 windows 和 centos7 文件共享
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: CentOS
tags:
  - CentOS
  - Samba
  - 文件共享
abbrlink: 7c9f7082
date: 2021-06-25 10:37:52
img:
coverImg:
password:
summary:
---

# samba 实现 windows 和 centos7 文件共享

- 安装 samba

```sh
yum -y install samba
```

![安装 samba 服务](https://upload-images.jianshu.io/upload_images/14623749-edbeb855b86ab130.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 查看 samba 主配置文件

```sh
cat /etc/samba/smb.conf
```

查看主配置文件感觉内容太多时，可以执行以下命令过滤掉注释信息

```sh
# 先将主配置文件做一个备份
mv /etc/samba/smb.conf /etc/samba/smb.conf.bak

# 过滤掉以井号 「#」、分号 「;」 、空行，并覆盖主配置文件
cat /etc/samba/smb.conf.bak | grep -v "#" | grep -v ";" | grep -v "^$" > /etc/samba/smb.conf

# 再次查看主配置文件
cat /etc/samba/smb.conf
```

![查看 samba 主配置文件](https://upload-images.jianshu.io/upload_images/14623749-eb2ca8538712d808.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 配置共享资源

1. 创建用于访问共享资源的账户信息

```sh
# 先创建一个系统用户 （用户名可以随意写，只要自己记得就好，这里为 sambauser）
# Samba服务程序的数据库要求账户必须在当前系统中已经存在，否则日后创建文件时将导致文件的权限属性混乱不堪，由此引发错误。
useradd sambauser

# 创建 samba 服务程序的用户 （执行以下命令后会提示输入密码，这里的密码用于创建共享时验证，因此需要牢记）
pdbedit -a -u sambauser
```

![创建用于共享资源的账户信息](https://upload-images.jianshu.io/upload_images/14623749-c69f38676209f8fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

2. 创建用于共享资源的文件目录

```sh
# 创建共享资源的文件目录，这里创建的文件名为 smb-share，文件名可以自由创建
mkdir -p /mnt/smb-share

# 设置文件夹权限
chown -Rf sambauser:sambauser /mnt/smb-share

# 设置该目录的 SELinux 安全上下文
semanage fcontext -a -t samba_share_t /mnt/smb-share

# 执行 restorecon 命令，让应用于目录的新 SELinux 安全上下文立即生效。
restorecon -Rv /mnt/smb-share

# 设置 SELinux 服务与策略
# 先筛选出所有与Samba服务程序相关的SELinux域策略
getsebool -a | grep samba

# 根据策略的名称（和经验）选择出正确的策略条目进行开启即可：
# eg：如果你的共享目录在家目录 /home 下，那么就需要开启
setsebool -P samba_enable_home_dirs on
# 这里我使用的共享目录在 /mnt 目录下，因此暂时不需要设置
```

![创建用于共享资源的文件目录](https://upload-images.jianshu.io/upload_images/14623749-fa8568d3e863ecbc.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 在 Samba 服务程序的主配置文件中，写共享配置信息

```sh
vim /etc/samba/smb.conf
```

写入以下内容


```sh
# 共享名称为 smb-share
[smb-share]
# 共享信息
comment = smb-share
# 共享目录为 /mnt/smb-share
path = /mnt/smb-share
# 关闭“所有人可见”
public = no
# 允许写入操作
writable = yes
```

![写共享配置信息](https://upload-images.jianshu.io/upload_images/14623749-a6be19d4ad79f97c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

4. 重启 samba 服务

```sh
# 重启 samba 服务
systemctl restart smb

# 将 samba 服务加入到开启启动项中
systemctl enable smb

# 清空 iptables 防火墙
iptables -F

# 保存 iptables 防火墙设置信息
service iptables save
```

至此，samba 服务器已经配置完毕

- Windows 挂载共享

1. 打开 `运行`，并输入 samba 服务器的 ip 地址

```
# 这里我的虚拟机的 ip 地址为 192.168.127.3，请换成你自己的 samba 服务器的 ip 地址
\\192.168.127.3
```

![打开运行，并访问 samba 服务器 ip](https://upload-images.jianshu.io/upload_images/14623749-dd44e0b89a8bf325.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

2. 输入完毕之后，直接点击 `确定`，就会跳出需要你输入账号和密码，这里的账号需要填写你 samba 服务程序的账号和密码，我这里已经刚刚输入一次了，因此第二次打开的时候，就没有提示要我输入账号和密码。我这里应该输入的账号为：sambauser

3. 测试

![账号和密码通过之后即可看到 centos 上面的共享文件夹](https://upload-images.jianshu.io/upload_images/14623749-42493846e7cac55b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![windows 中随便写入一个文件](https://upload-images.jianshu.io/upload_images/14623749-e38dea0acf233162.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

进入 `/mnt/smb-share` 目录查看文件

```sh
more /mnt/smb-share/fad.txt
```

![centos 查看文件](https://upload-images.jianshu.io/upload_images/14623749-e5f789ab12d4f1a7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

还可以将共享文件夹映射为网络驱动器，这样下次就可以直接在 `我的电脑` 中打开了。  
映射网络驱动器的方法为，直接对着共享文件夹，右键 => 映射网络驱动器 => 直接确定 即可

![映射网络驱动](https://upload-images.jianshu.io/upload_images/14623749-b88434639aae9732.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
