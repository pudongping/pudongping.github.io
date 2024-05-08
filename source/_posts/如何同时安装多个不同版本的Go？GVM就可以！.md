---
title: 如何同时安装多个不同版本的Go？GVM就可以！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
abbrlink: c10fde57
date: 2024-05-09 00:39:39
img:
coverImg:
password:
summary:
---

# [GVM](https://github.com/moovweb/gvm)

> 安装过程中，如果遇到了错误，还是直接先访问仓库地址看看文档介绍，可能解决得更快。

## 安装并使用

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# 也可以先把脚本先下载下来，然后执行
curl -o gvm-installer https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer
chmod +x gvm-installer
./gvm-installer
```

查看可用的 go 版本

```bash
gvm listall
```

安装特定版本的 go

```bash
# 这里的版本号由 `gvm listall` 命令返回的版本号之一
gvm install go1.16.5
# 下载后的 go 位于 `~/.gvm/gos` 目录下
```

查看当前正在使用的 go 版本

```bash
gvm list

# gvm gos (installed)
#
#   go1.17
#   go1.17.13
#   go1.20.8
#=> go1.21.0
#   system
```

使用特定版本的 go

```bash
gvm use go1.16.5
# 设置为默认版本的 go
gvm use go1.16.5 --default
```

卸载特定版本的 go

```bash
gvm uninstall go1.16.5
```

**慎重！！！** 完全卸载掉 GVM 和所有安装的 Go 版本

```bash
gvm implode

# 如果以上的命令不奏效的话，也可以直接
rm -rf $GVM_ROOT
# 然后去除掉 ~/.zshrc 文件中的类似以下内容即可（如果使用的是 bash 则在 ~/.bashrc 文件中）
[[ -s "/Users/pudongping/.gvm/scripts/gvm" ]] && source "/Users/pudongping/.gvm/scripts/gvm"
```

## 使用 GVM pkgset

> 使用 pkgset 可以将你每一个项目的依赖包都分隔开来。

```bash
# 查看系统环境变量
echo $PATH
# 可见已经将 `~/.gvm/bin` 加入了环境变量中
# /Users/pudongping/.gvm/bin
```

先使用特定的版本，比如 `go1.21.0`

```bash
gvm use go1.21.0

# Now using version go1.21.0

gvm list

# gvm gos (installed)
#
#   go1.17
#   go1.17.13
#   go1.20.8
#=> go1.21.0
#   system

echo $PATH
# /Users/pudongping/.gvm/pkgsets/go1.21.0/global/bin:/Users/pudongping/.gvm/gos/go1.21.0/bin:/Users/pudongping/.gvm/pkgsets/go1.21.0/global/overlay/bin:/Users/pudongping/.gvm/bin
# 此时可见多添加了
# ~/.gvm/pkgsets/{GO_VERSION}/global/bin
# ~/.gvm/gos/{GO_VERSION}/bin
# ~/.gvm/pkgsets/{GO_VERSION}/global/overlay/bin

echo $GOPATH
# /Users/pudongping/.gvm/pkgsets/go1.21.0/global
```

查看 **global** pkgset

```bash
gvm pkgset list

#gvm go package sets (go1.21.0)
#
#=>  global
```

创建特定的 pkgset

```bash
# 创建一个名为 `ppGvm` 的 pkgset
gvm pkgset create ppGvm

gvm pkgset list
#
#gvm go package sets (go1.21.0)
#
#=>  global
#    ppGvm
 
# 此时查看系统环境变量，还不会将 `ppGvm` 加入    
echo $PATH
# /Users/pudongping/.gvm/pkgsets/go1.21.0/global/bin:/Users/pudongping/.gvm/gos/go1.21.0/bin:/Users/pudongping/.gvm/pkgsets/go1.21.0/global/overlay/bin:/Users/pudongping/.gvm/bin
echo $GOPATH
# /Users/pudongping/.gvm/pkgsets/go1.21.0/global

# 使用 `ppGvm` pkgset
gvm pkgset use ppGvm
# Now using pkgset go1.21.0@ppGvm

gvm pkgset list
#
#gvm go package sets (go1.21.0)
#
#    global
#=>  ppGvm

# 此时可以看到，已经将 `ppGvm` 加入到环境变量中
echo $PATH
# /Users/pudongping/.gvm/pkgsets/go1.21.0/ppGvm/bin:/Users/pudongping/.gvm/pkgsets/go1.21.0/ppGvm/overlay/bin:/Users/pudongping/.gvm/pkgsets/go1.21.0/global/bin:/Users/pudongping/.gvm/gos/go1.21.0/bin:/Users/pudongping/.gvm/pkgsets/go1.21.0/global/overlay/bin:/Users/pudongping/.gvm/bin

echo $GOPATH
# /Users/pudongping/.gvm/pkgsets/go1.21.0/ppGvm:/Users/pudongping/.gvm/pkgsets/go1.21.0/global
```

如果此时你通过 `go get` 下载 go 包的话，你可以看到新的包被添加到了以下目录下

```bash
cd $( awk -F':' '{print $1}' <<< $GOPATH )
# 或者直接 cd /Users/pudongping/.gvm/pkgsets/go1.21.0/ppGvm

pwd
#/Users/pudongping/.gvm/pkgsets/go1.21.0/ppGvm
```
