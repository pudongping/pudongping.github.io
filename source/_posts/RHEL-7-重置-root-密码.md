---
title: RHEL 7 重置 root 密码
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: f41190ff
date: 2021-06-12 23:57:29
img:
coverImg:
password:
summary:
---

# RHEL 7 重置 root 密码

- 查看是否为 RHEL7 系统

```bash

# 查看内核版本
cat /etc/redhat-release

# 如果信息为以下内容，则可以采用此方法重置 root 密码
Red Hat Enterprise Linux Server release 7.0 (Maipo)

```

- 重启 Linux 系统主机，并出现引导界面时，按下键盘上的 `e` 键进入内核编辑界面


![图片来源于 www.linuxprobe.com 网站](https://www.linuxprobe.com/wp-content/uploads/2015/02/%E7%AC%AC1%E6%AD%A5%EF%BC%9A%E5%BC%80%E6%9C%BA%E5%90%8E%E5%9C%A8%E5%86%85%E6%A0%B8%E4%B8%8A%E6%95%B2%E5%87%BB%E2%80%9Ce%E2%80%9D.png)


- 在 `linux16` 参数这行的最后面（即`LANG=en_US.UTF-8`尾部）先敲一个 `空格` 然后添加 `\rd.break` 参数，然后按下 `Ctrl + X` 组合键来运行修改过的内核程序


> 这里需要注意的是：
> 1. 页面内容一般情况下是显示不全的，需要按住上下箭头滚动。
> 2. 添加 \rd.break 参数的时候，你也可以不用添加 '\反斜杠' 但是必须要先在 linux16 末尾处敲击一个空格之后写入 rd.break 参数

![图片来源于 www.linuxprobe.com 网站](https://www.linuxprobe.com/wp-content/uploads/2015/02/%E7%AC%AC2%E6%AD%A5%EF%BC%9A%E5%9C%A8linux16%E8%BF%99%E8%A1%8C%E7%9A%84%E5%90%8E%E9%9D%A2%E8%BE%93%E5%85%A5%E2%80%9Crd.break%E2%80%9D%E5%B9%B6%E6%95%B2%E5%87%BB%E2%80%9Cctrl-x%E2%80%9C.png)

- 稍后即可进入系统的紧急救援模式

![图片来源于 www.linuxprobe.com 网站](https://www.linuxprobe.com/wp-content/uploads/2015/02/%E7%AC%AC3%E6%AD%A5%EF%BC%9A%E8%BF%9B%E5%85%A5%E5%88%B0%E4%BA%86%E7%B3%BB%E7%BB%9F%E7%9A%84%E7%B4%A7%E6%80%A5%E6%B1%82%E6%8F%B4%E6%A8%A1%E5%BC%8F.png)

- 依次输入以下命令，等待系统重启操作完毕，然后就可以使用新密码来登录 Linux 系统了。

1. 以可读写的权限重新挂载硬盘上真实系统根目录（ /sysroot ）目录

```bash

# -o remount：将一个已经挂下的档案系统重新用不同的方式挂上。
# 例如原先是唯读的系统，现在用可读写的模式重新挂上。
# -o rw：用可读写模式挂上。
# 可以合并参数写，即 mount -o remount,rw

mount -o remount,rw / /sysroot

# 同样也可以直接写成这样 （直接省略掉根目录 / ）
mount -o remount,rw /sysroot

```

2. 把环境切换到真实系统根目录 /sysroot

```bash
chroot /sysroot
```

3. 修改 root 账户密码

```bash
# 输入 passwd 命令时，是交互式界面。  
# 同样也可以直接输入 echo "redhat" | passwd --stdin root 命令，直接将 root 密码，修改为 redhat

passwd
```

4. 告诉系统下次启动将对文件进行 selinux 上下文重新打标。这也就造成了下次重启的时候时间会很长，autorelabel 是一个隐藏文件，需要注意的是前面有一个点

```bash
touch /.autorelabel
```

5. 退出真实系统根目录环境

```bash
exit
```

6. 重启系统

```bash
reboot
```


![](https://www.linuxprobe.com/wp-content/uploads/2015/02/%E7%AC%AC4%E6%AD%A5%EF%BC%9A%E4%BE%9D%E6%AC%A1%E8%BE%93%E5%85%A5%E4%BB%A5%E4%B8%8B%E5%91%BD%E4%BB%A4.png)

- 修改密码后，首次重启的时间将会比较长，因为系统将对所有文件进行 SeLinux 打标，请耐心等待，整个过程并非死机，请勿在打标过程中手动强制再次重启，否则系统将会永久性损坏导致无法开机。
