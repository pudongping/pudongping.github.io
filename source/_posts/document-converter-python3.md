---
title: 使用 python3 写的文档格式转换小工具
author: Alex
top: false
hide: false
cover: true
toc: true
mathjax: false
categories: Python
tags:
  - Python3
abbrlink: 6cc526bb
date: 2021-06-07 16:45:09
img: https://pudongping.com/medias/featureimages/41.jpg
coverImg: https://pudongping.com/medias/banner/6.jpg
password:
summary: 自己使用 python3 写的文档格式转换工具，目前暂时只支持 docx 和 pdf 文档互转，后面有时间的话，再考虑完善更多的格式互转。
---

# document-converter （文档格式互转工具）

## 项目发起初衷

近期有朋友要我帮忙将 pdf 格式文档数据提取成 docx 格式，平时我自己要么就通过编辑器转换了，要么就通过命令行转换，但这种方式对于一些朋友来说不是很友好，故此有此项目。

目前只支持 docx 和 pdf 互转，后续有时间了，会考虑支持多种常用文档格式互转，如果你对这感兴趣，欢迎提 PR ！

## 项目特性

- 使用 python3 编写的 docx、pdf 互转工具。
- 支持多线程。
- 支持跨平台（macOS、Windows、Linux）。
- 批量 docx 文档转 pdf 格式（效果良好）。
- 批量 pdf 文档转 docx 格式（尝试过多个扩展库，遗憾的是均不支持图片提取）。
- 后续有时间了，会考虑支持多种常用的文档格式互转。

## 效果图

命令行提示

![文档简介](https://upload-images.jianshu.io/upload_images/14623749-155e22f797a6ee6f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

批量 docx 转 pdf

![docx转pdf](https://upload-images.jianshu.io/upload_images/14623749-3df6914674c7e060.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

批量 pdf 转 docx

![pdf转docx](https://upload-images.jianshu.io/upload_images/14623749-800ff9980abdb59a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 使用方法

1. 克隆源代码

```shell
# GitHub
git clone https://github.com/pudongping/document-converter.git

# gitee
git clone https://gitee.com/pudongping/document-converter.git
```

2. 进入项目根目录，并安装相关依赖扩展

```shell
cd document-converter && sudo pip install -r requirements.txt
```

3. 使用文档格式转换

举个例子来说，如果此时我想将 `n` 个 `.docx` 格式的文档转换成 `.pdf` 格式的文档，那么我需要这么操作：

- 将这 `n` 个 `.docx` 文档复制到此项目的 `data/input/docxs` 目录（如果想转 `.pdf` 格式的文档，则需要将文档复制到 `data/input/pdfs` 目录）
- 在项目根目录执行 `python3 main.py --docx-to-pdf` 命令，且等待命令执行结束（你可以去泡杯咖啡，之后静静等待，一般来说会很快，因为支持多线程）
- 之后在 `data/output/pdfs` 目录查看已经转换好的文档。至此，完毕！

4. 命令介绍

命令|说明
---|---
python3 main.py --version | 显示当前应用的版本号
python3 main.py --help 或者 python3 main.py -h | 显示帮助文档
python3 main.py --docx-to-pdf | 执行 docx 文档转成 pdf 格式的文档
python3 main.py --pdf-to-docx | 执行 pdf 文档转成 docx 格式的文档

## 项目目录介绍

```shell

├── LICENSE
├── Pipfile
├── README.md  项目介绍文档
├── app  代码目录
│   ├── __init__.py
│   ├── config  配置文件目录
│   │   ├── __init__.py
│   │   └── app.py  配置相关
│   ├── converter  转换相关代码目录
│   │   ├── __init__.py
│   │   ├── converter_docx.py  *文档转 docx 格式相关代码
│   │   └── converter_pdf.py   *文档转 pdf 格式相关代码
│   └── helper.py  助手函数
├── data  数据相关
│   ├── input  需要转换的文档目录
│   │   ├── docxs  如果需要将 docx 文档转换成其他格式，则默认需要将文档放入此目录下
│   │   │   ├── 文档1.docx
│   │   │   └── 文档2.docx
│   │   └── pdfs  如果需要将 pdf 文档转换成其他格式，则默认需要将文档放入此目录下
│   │       ├── alex1.pdf
│   │       └── sample.pdf
│   └── output  文档转换后存放的文档目录
│       ├── docxs  所有经过转换后的 docx 文档默认会放到此目录下
│       │   ├── alex1.pdf.docx
│       │   └── sample.pdf.docx
│       └── pdfs  所有经过转换后的 pdf 文档默认会放到此目录下
│           ├── 文档1.docx.pdf
│           └── 文档2.docx.pdf
├── main.py  项目入口文件
├── requirements.txt  项目依赖关系清单
└── runtime  项目运行相关
    └── logs  日志目录
        └── 202104
            └── converter-2021-04-24.log  以天为单位记录的操作日志

```

## License

源代码基于 [MIT](https://opensource.org/licenses/MIT) 协议发布。
