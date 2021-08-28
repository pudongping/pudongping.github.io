---
title: Python 搭建环境与安装 pip 扩展
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
abbrlink: ee47ba79
date: 2021-08-29 01:18:34
img:
coverImg:
password:
summary:
---


# Python 搭建环境与安装 pip 扩展

> [python 在线运行网站](http://www.pythontutor.com/live.html#mode=edit)

## Mac 下安装

### [mac 上安装 python](https://www.python.org/downloads/)
可以直接通过下载官网的[安装器](https://docs.python.org/zh-cn/3/using/mac.html)进行安装。

或者使用

```sh
brew install python3
```

### [mac 下使用 pip](https://pip.pypa.io/en/stable/quickstart/)

可以直接使用 `python3 -m pip ……` 命令

```
# 查看 pip 版本
python3 -m pip --version

# 升级 pip 版本
/Library/Frameworks/Python.framework/Versions/3.8/bin/python3 -m pip install --upgrade pip

# 搜索插件包 query
python -m pip search "query"

# 安装 SomePackage 包
python3 -m pip install SomePackage            # latest version
python3 -m pip install SomePackage==1.0.4     # specific version
python3 -m pip install 'SomePackage>=1.0.4'     # minimum version

# 安装已经通过 [PyPI](https://pypi.org/) 下载好的包（这在没有网的情况下很有用）
python3 -m pip install SomePackage-1.0-py2.py3-none-any.whl

# 显示已安装包的详细信息
python3 -m pip show --files SomePackage

# 列出哪些包已经存在了新版本
python3 -m pip list --outdated

# 升级包版本
python3 -m pip install --upgrade SomePackage

# 卸载包
python3 -m pip uninstall SomePackage

```

如果想直接使用 `pip` 命令时，则需要如下操作，[官网安装 pip 步骤文档](https://pip.pypa.io/en/stable/installing/)

```
# 安装 get-pip.py 脚本
curl https://bootstrap.pypa.io/get-pip.py | python3

# 查看 pip 版本
pip --version 或者 pip3 --version

# 查看已经安装的包
pip list 或者 pip3 list

# 更新 pip 版本
pip install --upgrade pip 或者 pip3 install --upgrade pip


```

### requirements.txt

```python

# 生成 requirements 文件
pip freeze > requirements.txt

# 安装 requirements.txt 依赖命令
pip install -r requirements.txt

# 卸载 requirements.txt 文件中的依赖
pip uninstall -y -r requirements.txt
```


## Windows 下安装背景

系统 windows7 64位，安装的 python 版本为 3.6.7 ，此文档为过程总结。

## Python 与 pip 下载与安装

### 下载
-  进入 `https://www.python.org/` 选择  `downloads`

![官网首页](https://upload-images.jianshu.io/upload_images/14623749-0e913bc27da7cfc2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
- 选择 3.6.7 版本

![选择3.6.7版本](https://upload-images.jianshu.io/upload_images/14623749-da7102d7d5108480.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
- 拉到屏幕底部，选择
    - Windows x86-64 executable installer
      ![选择64位安装包](https://upload-images.jianshu.io/upload_images/14623749-4a7bf34585ea3e31.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### Windows 下安装
- 直接对下载好的安装包双进运行 (Install Python 3.6.7(64-bit) )
    - 选择自定义安装（Customize installation）
    - 勾选 add python3.6 to path（自动添加环境变量）
    - 下一步
- 勾选 (Optional Features)
    - document
    - pip（必须选）
    - .. 其他的多装总比少装强
    - 下一步
- 选择安装路径 （Advanced Options）
    - 安装路径，c:\python\Python3.6（可自行定义 Browse）
    - 安装 （Install）
- 安装完成


### 检查安装效果

- 开始，运行，进入cmd命令行
- 敲命令 python
    - windodws 用 ctrl+z 退出交互界面
- 敲命令 pip
    - 命令存在
- 如果提示命令不存在
    - 很可能是你的 环境变量未设置好

## 安装 pipenv

它是一个项目虚拟环境管理工具。

### 利用 pip 安装

`pip install pipenv`

执行该命令自动进入下载安装过程。

### 创建项目

- 在电脑中创建文件夹，比如 I:\Python\test
- 在命令行切换到此
    - cmd: cd  I:\Python\test
- 执行 pipenv 的初始化命令
    - pipenv --python 3.6

![执行 pipenv 初始化命令之后的画面](https://upload-images.jianshu.io/upload_images/14623749-0742c148428484a3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 进入 pipenv 的虚拟环境
    - pipenv shell

![进入 pipenv 的虚拟环境](https://upload-images.jianshu.io/upload_images/14623749-b40a5bcc9c69f618.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 安装一个库试试
    - pip install requests

![安装 requests 库界面](https://upload-images.jianshu.io/upload_images/14623749-d2ac78cb65d7f691.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 虚拟化环境存放路径
    - 默认 c:\users\administrator\.vi*****

![本次安装虚拟化环境存放路径](https://upload-images.jianshu.io/upload_images/14623749-1cd80d6ed3238a66.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- OK End

## 下载使用 pycharm

直接去官网下载 windows 的 pro版本。

地址 http://www.jetbrains.com/pycharm/download/
![下载 Pro 版本](https://upload-images.jianshu.io/upload_images/14623749-120aa374788beefc.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 创建项目

直接选择我们刚刚的项目路径 `I:\Python\test`

### 让项目使用pipenv的虚拟环境

`file -> setting -> project:test`

![如图所示](https://upload-images.jianshu.io/upload_images/14623749-ac306de03f065e51.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


进一步选择 `interperter` 在下拉选单中选中我们刚刚的虚拟化环境目录即可。
![选择上文中创建的虚拟环境目录](https://upload-images.jianshu.io/upload_images/14623749-175faf09036cc3a6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### interperter 中没有怎么办？

下拉选单中，选择 `show all`
![第一步： interperter 中没有的话，点击下拉选单之后选择 show all](https://upload-images.jianshu.io/upload_images/14623749-6828f88cd9937443.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![第二步：点击加号+](https://upload-images.jianshu.io/upload_images/14623749-6171dcff4886c3bc.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![第三步：选择已经存在的，不要选择新建的](https://upload-images.jianshu.io/upload_images/14623749-e6d11f9351337be9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


然后在 existing environment，在右侧`...`选中虚拟环境的 python.exe 即可。
![第四步：选择已存在的之后点击最右边的 …](https://upload-images.jianshu.io/upload_images/14623749-3b5e2bdbeb07032e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![第五步：选择 python.exe 既可](https://upload-images.jianshu.io/upload_images/14623749-87dab8f9313b439e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
一般情况下，默认路径：

`c:\users\administrator\.virtualenvs\项目名\Scripts\python.exe`

![切记切记！！！](https://upload-images.jianshu.io/upload_images/14623749-61e21223cf87bdf2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后一直点 ok 即可
效果如下
![可以看到会出现两个虚拟环境的选项](https://upload-images.jianshu.io/upload_images/14623749-78490ce15425ed18.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
最终可以看到编辑器中包含了刚刚创建的虚拟环境
![看到这里证明编辑器配置成功了](https://upload-images.jianshu.io/upload_images/14623749-8260285f4c846b36.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 总结

至此，整个项目环境配置完毕。

以后我们所有对项目的命令行操作都需要在项目根目录先进入 `pipenv shell` 然后再执行命令操作，这样就可以使得每一个项目安装独立的安装包。

