---
title: Python 使用虚拟环境
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
  - Pip
abbrlink: ddef4649
date: 2023-11-07 17:46:01
img:
coverImg:
password:
summary:
---

# Python 使用虚拟环境

Python 需要使用虚拟环境的主要原因包括：

1. 隔离项目依赖：虚拟环境允许您在不同的项目之间隔离依赖关系。这意味着您可以为每个项目创建一个独立的虚拟环境，以确保项目的依赖不会相互干扰。这对于开发多个项目或维护项目的不同版本非常重要，因为它可以防止依赖冲突。
2. 版本管理：虚拟环境允许您在不同的项目中使用不同的 Python 版本。这对于需要支持不同 Python 版本的项目非常有用，因为您可以在不同的虚拟环境中安装和使用特定版本的 Python。
3. 防止全局依赖污染：如果您在全局 Python 环境中安装依赖项，可能会导致全局依赖项的混乱，甚至可能破坏系统依赖项。虚拟环境将项目的依赖项隔离到项目本身的目录中，从而避免了这种情况。
4. 管理依赖项：虚拟环境允许您在项目级别管理依赖项。您可以使用 pip 来安装、升级和卸载依赖项，而不会影响全局 Python 环境。
5. 简化部署：使用虚拟环境，您可以轻松地将项目及其依赖项打包并部署到其他环境中，而不必担心依赖冲突或版本问题。

## [Pipenv](https://github.com/pypa/pipenv)

### 安装

```bash
# 全局安装，如果只想在当前用户模式下安装，可添加参数 `--user`
pip3 install pipenv

# 更新 pipenv
pip3 install --user --upgrade pipenv
```

### 使用

```bash
# 查看 python3 版本
python3 --version
# 假设是 Python 3.11.4

# 指定使用 python 3.11.4 创建虚拟环境
pipenv --python 3.11.4

# 激活虚拟环境
pipenv shell

# 退出当前虚拟环境
exit

# 删除当前虚拟环境
pipenv --rm
```

### 下载依赖包

```bash

# 下载 Pipfile 文件中的所有包
pipenv install

# 安装 requests 插件包并加入到 Pipfile
pipenv install requests
# 安装固定版本的 requests
pipenv install requests==2.22.0

# 只安装开发环境才会使用到的包
pipenv install {package-name} --dev
```

### 更新依赖包

```bash
# 查看所有需要更新的依赖
pipenv update --outdated

# 更新所有包的依赖项
pipenv update

# 更新指定包的依赖项
pipenv update {package-name}
```

### 卸载依赖包

```bash
# 卸载指定模块
pipenv uninstall {package-name}

# 卸载全部包
pipenv uninstall --all

# 卸载全部开发环境所需要依赖的包
pipenv uninstall --all-dev
```

### 其他常用命令

```bash
# 显示目录信息
pipenv --where

# 显示虚拟环境信息
pipenv --venv

# 显示 python 解释器信息
pipenv --py

# 查看当前安装的库及其依赖
pipenv graph

# 检查安全漏洞
pipenv check

# 生成 Pipfile.lock 文件
pipenv lock
```

### requirements.txt

```bash
# 将 Pipfile 和 Pipfile.lock 文件里面的包导出为 requirements.txt 文件
pipenv run pip freeze > requirements.txt
# 或者
pipenv requirements > requirements.txt


# 只使用 `pipenv install` 时会自动检测当前目录下的 requirements.txt 并生成 Pipfile 文件
# 通过 requirements.txt 安装包
pipenv install -r requirements.txt

# 只安装开发环境所需要的包
pipenv install -r --dev requirements.txt
```