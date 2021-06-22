---
title: MacOS 安装 homebrew
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: macOS
tags:
  - macOS
  - homebrew
abbrlink: 6fa24dc0
date: 2021-06-22 19:21:53
img:
coverImg:
password:
summary:
---

# MacOS 安装 homebrew

命令 | 说明
--- | ---
brew update | 更新 Homebrew 自身
brew outdated | 查看哪些安装包需要更新
brew upgrade | 更新所有的包
brew upgrade $FORMULA | 更新指定的包
brew cleanup | 清理所有包的旧版本
brew cleanup $FORMULA | 清理指定包的旧版本
brew cleanup -n | 查看可清理的旧版本包，不执行实际操作
brew pin $FORMULA | 锁定某个包
brew unpin $FORMULA | 取消锁定
brew info $FORMULA | 显示某个包的信息
brew info | 显示安装了包数量，文件数量，和总占用空间
brew deps --installed --tree | 查看已安装的包的依赖，树形显示
brew list | 列出已安装的包
brew rm $FORMULA | 删除某个包
brew uninstall --force $FORMULA | 删除所有版本

## 参考文献

> [Homebrew 官网地址](https://brew.sh/)  
[Homebrew GitHub 地址](https://github.com/homebrew/install#uninstall-homebrew)  
[官网源码下载](https://github.com/Homebrew/brew/releases/tag/2.7.3)  
[官网 GitHub 地址](https://github.com/homebrew/install#uninstall-homebrew)  
[清华大学开源软件镜像站 - 安装 homebrew](https://mirrors.tuna.tsinghua.edu.cn/help/homebrew/)  
[使用国内镜像安装 homebrew](https://www.jianshu.com/p/5822d24a651e)  


> 如果使用官网推荐的方式下载时提示以下错误信息时，则表示网络超时，建议直接使用源码包的形式安装

错误信息如下所示：

```bash
curl: (7) Failed to connect to raw.githubusercontent.com port 443: Connection refused
```

## 直接使用包安装
1. 进入 [Homebrew 的 GitHub 仓库 tag 地址](https://github.com/Homebrew/brew/tags) 下载最新的 tag
2. 根据系统选择下载：mac、windows 可以下载 zip 文件， linux 可以下载 tar.gz 的文件。
3. 下载后解压。
4. 进入本地 `Homebrew` 的存放路径，如果找不到的话，可以直接在 `terminal` 下输入以下命令

```bash

# 进入  usr/local 目录
cd /usr/local

# 使用访达(finder) 打开当前目录
open .

```

5. 查看 `/usr/local` 目录下是否有 `Homebrew` 文件夹（注意大小写），如果你发现没有 `Homebrew` 文件夹，则执行以下命令创建 `Homebrew` 文件夹

```bash

mkdir -p /user/local/Homebrew

```

6. 将第三步中解压后的内容全部复制到 `/usr/local/Homebrew` 目录

7. 重启命令行窗口，输入 `brew` 命令，出现 brew 相关的 help 页面，即表示已经安装成功

```bash

cd /user/local/Homebrew && brew

```

敲击 `brew` 命令时，出现以下内容时，表示已经安装 Homebrew 成功

```bash
Example usage:
  brew search [TEXT|/REGEX/]
  brew info [FORMULA...]
  brew install FORMULA...
  brew update
  brew upgrade [FORMULA...]
  brew uninstall FORMULA...
  brew list [FORMULA...]

Troubleshooting:
  brew config
  brew doctor
  brew install --verbose --debug FORMULA

Contributing:
  brew create [URL [--no-fetch]]
  brew edit [FORMULA...]

Further help:
  brew commands
  brew help [COMMAND]
  man brew
  https://docs.brew.sh
```

8. 安装完成后，一定要下载一个软件测试下，比如下载 `wget`

```bash
brew install wget
```

9. 更换 brew 的下载源

- [Homebrew 源](http://mirrors.ustc.edu.cn/help/brew.git.html)
```bash
# 替换 USTC 镜像
cd "$(brew --repo)"
git remote set-url origin https://mirrors.ustc.edu.cn/brew.git

# 重置为官方地址
cd "$(brew --repo)"
git remote set-url origin https://github.com/Homebrew/brew.git
```

- [Homebrew Bottles 源](http://mirrors.ustc.edu.cn/help/homebrew-bottles.html)
```bash
# 请在运行 brew 前设置环境变量 HOMEBREW_BOTTLE_DOMAIN ，值为 https://mirrors.ustc.edu.cn/homebrew-bottles

# 对于 bash 用户
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles' >> ~/.bash_profile

source ~/.bash_profile


# 对于 zsh 用户
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles' >> ~/.zshrc

source ~/.zshrc
```

- [Homebrew Core 源](http://mirrors.ustc.edu.cn/help/homebrew-core.git.html)
```bash
# 替换 USTC 镜像
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-core.git

# 重置为官方地址
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://github.com/Homebrew/homebrew-core
```

- [Homebrew Cask 源](http://mirrors.ustc.edu.cn/help/homebrew-cask.git.html)
```bash
# 替换为 USTC 镜像
cd "$(brew --repo)"/Library/Taps/homebrew/homebrew-cask
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-cask.git

# 重置为官方地址
cd "$(brew --repo)"/Library/Taps/homebrew/homebrew-cask
git remote set-url origin https://github.com/Homebrew/homebrew-cask
```

- [Homebrew Cask Versions 源](http://mirrors.ustc.edu.cn/help/homebrew-cask-versions.git.html)
```bash
# 替换为 USTC 镜像
cd "$(brew --repo)"/Library/Taps/homebrew/homebrew-cask-versions
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-cask-versions.git

# 重置为官方地址
cd "$(brew --repo)"/Library/Taps/homebrew/homebrew-cask-versions
git remote set-url origin https://github.com/Homebrew/homebrew-cask-versions.git
```
