---
title: NVM 包管理器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: NPM
tags:
  - NVM
  - NPM
abbrlink: 63668a37
date: 2021-09-01 20:50:42
img:
coverImg:
password:
summary:
---

# nvm 包管理器

- [官网](https://github.com/nvm-sh/nvm)
- [Windows 直接选择 nvm-setup.zip 安装](https://github.com/coreybutler/nvm-windows/releases)
- [查看 nodejs 的版本号序列](https://nodejs.org/en/blog/)

## npm 更改下载源为淘宝镜像

```sh

# 单次使用
npm install --registry=https://registry.npm.taobao.org

# 永久使用
npm config set registry=https://registry.npm.taobao.org

# 检测是否成功
npm config get registry 或者 npm info express

# 还原 npm 的仓库地址
npm config set registry https://registry.npmjs.org/

```

## mac 下安装 nvm

```sh

cd ~

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash

# 添加环境变量
vim ~/.zshrc
# 添加如下内容，如果是安装过 oh-my-zsh 
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm

# 如果没有安装 oh-my-zsh 时，直接按照如下方式添加也可以，就算是已经安装了 oh-my-zsh 也可以这么操作，不冲突
vim ~/.bashrc
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

```

## nvm 相关命令

```sh

# 查看 nvm 版本号
nvm -v

# windows 下查看可安装的 node.js 版本
nvm list available

# 安装指定版本的 node.js
nvm install 10.15.0

# 使用特定版本的 node.js
nvm use 10.15.0

# 版本切换
nvm install 8.11.2  # 比如先安装 8.11.2 的版本，然后切换到 8.11.2
nvm use 8.11.2

# 查看当前已安装的 node.js
nvm list

# 删除指定版本的 node.js
nvm uninstall 8.11.2

```
