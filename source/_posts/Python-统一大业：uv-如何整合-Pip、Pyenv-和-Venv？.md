---
title: Python 统一大业：uv 如何整合 Pip、Pyenv 和 Venv？
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
  - Pyenv
  - Venv
  - uv
abbrlink: c387be6c
date: 2026-05-18 11:02:18
img:
coverImg:
password:
summary:
---

Python 的包管理一直是个让开发者“又爱又恨”的话题。从 `pip` 到 `virtualenv`，再到 `poetry`、`conda`、`pdm`，工具层出不穷，但似乎总觉得“差点意思”——要么慢，要么依赖冲突让人头秃。

最近，Python 圈子杀出了一匹黑马—— **`uv`**。它是 `Ruff`（那个快到离谱的代码检查工具）团队的新作，同样用 **Rust** 编写。这玩意的特点就是一个字：**快**！而且它是奔着“大一统”去的，**能管包、管环境、甚至管 Python 版本。**

今天咱们就来盘一盘这个被称为“Python 基础设施未来”的工具。

## 01. 为什么要用 uv？

简单说，`uv` 想把原本分散的活儿全干了：

- **速度极快**：比 `pip` 和 `pip-tools` 快 10-100 倍（不夸张，Rust 写的）。
- **All-in-One**：你不再需要 `pyenv` 装 Python，不需要 `venv` 建环境，不需要 `pip` 装包，一个 `uv` 全搞定。

> 很多同学习惯用 `pipenv`，因为它解决了 `pip` 无法锁定依赖版本的问题。但 `pipenv` 最大的痛点是**解析依赖速度慢**，甚至有时锁定一个环境要喝杯咖啡的时间。`uv` 同样支持锁定文件（`uv.lock`），但它是瞬时完成的。如果你喜欢 `pipenv` 的自动化感，`uv` 会给你同样的体验，但快得像闪电。

## 02. 安装 uv：开启极速之旅

安装非常简单，甚至不需要你先安装 Python。

### macOS

推荐使用 Homebrew 安装

```bash
brew install uv
```

或者使用官方安装脚本来安装

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### Linux

官方推荐使用 `curl` 安装：

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### Windows

可以使用 Winget 来安装

```bash
winget install uv
```

也可以使用 `powershell` 进行安装：

```powershell
powershell -c "irm https://astral.sh/uv/install.ps1 | iex"
```

### 如果你已经有了 pip

当然，你也可以像安装普通 Python 包一样安装它（不推荐作为首选，但很方便）：

```bash
pip install uv
```

*注：安装完成后，记得重启一下终端，或者按照提示配置环境变量。*

### 验证

安装完成之后，验证安装是否成功：

```bash
uv --version
```

## 03. 管理 Python 版本：忘掉 pyenv 吧

以前我们要安装不同版本的 Python，通常会用 `pyenv` 或者去官网下载安装包。现在 `uv` 直接接管了。

**查看可用的 Python 版本：**

```bash
uv python list
```

**安装特定的 Python 版本：**

```bash
# 安装 Python 3.12
uv python install 3.12

# 安装 Python 3.9
uv python install 3.9
```

`uv` 会自动下载并管理这些版本，你甚至不需要手动配置系统路径，它会在项目运行时自动查找合适的版本。

## 04. 项目初始化与虚拟环境

以前我们习惯先 `python -m venv .venv`，现在 `uv` 把这套流程变得更符合现代工程习惯。

### 初始化一个新项目

```bash
# 创建一个名为 my-project 的新项目
uv init my-project

# 进入目录
cd my-project
```

`uv init` 会帮你生成 `pyproject.toml`（现代 Python 项目的标准配置文件）和一个 `.python-version` 文件。

### 自动创建虚拟环境

当你在这个项目里添加依赖或者运行代码时，`uv` 会**自动**为你创建和管理虚拟环境。

如果你非要手动创建一个传统的虚拟环境，也可以：

```bash
# 在当前目录创建 .venv
uv venv

# 激活环境（macOS/Linux）
source .venv/bin/activate

# 激活环境（Windows）
.venv\Scripts\activate
```

然后在项目中指定 python 的版本

```bash
# 为当前项目固定 Python 3.11
uv python pin 3.11
```

## 05. 包管理与依赖添加

这是大家最常用的功能。`uv` 的命令设计非常直观

安装包：

```bash
# 安装最新版本
uv pip install requests

# 安装特定版本
uv pip install requests==2.31.0

# 从 requirements.txt 安装
uv pip install -r requirements.txt
```

安装包到开发环境

```bash
uv pip install --dev pytest
```

执行后，`uv` 会做三件事：

1. 下载包（如果缓存里没有）。
2. 安装到虚拟环境。
3. 自动更新 `pyproject.toml` 和 `uv.lock` 文件（锁定版本，确保团队协作一致性）。

升级包

```bash
uv pip upgrade requests
```

卸载包

```bash
uv pip uninstall requests
```


### 同步依赖 (Sync)

当你拉取了同事的代码，或者切换了分支，需要把环境同步到最新状态：

```bash
# 根据 lock 文件同步环境
uv sync
```

## 06. 运行代码

以前我们需要先 `source .venv/bin/activate` 激活环境，然后再 `python app.py`。
使用 `uv`，你可以直接运行：

```bash
# 在项目的虚拟环境中运行 app.py
uv run app.py

# 运行安装在环境中的工具，比如 pytest
uv run pytest
```

`uv run` 会自动检测并使用当前项目的虚拟环境，非常省心。

## 07. 旧项目迁移：如何从 pip/poetry/pipenv 迁移到 uv？

如果你手头有一个老项目，想享受 `uv` 的速度，迁移其实很平滑。

### 场景一：只有 requirements.txt

这是最常见的场景。

```bash
# 1. 初始化 uv 项目
uv init

# 2. 直接安装 requirements.txt 里的依赖
uv pip install -r requirements.txt
```

### 场景二：直接替代 pip 使用

如果你不想改变项目结构，只想单纯加速 `pip install` 的过程，你可以把 `uv` 当作 `pip` 的替身：

```bash
# 创建环境
uv venv

# 激活环境 (Windows: .venv\Scripts\activate)
source .venv/bin/activate 

# 用 uv 的 pip 接口安装
uv pip install requests
```

### 场景三：从 Poetry 迁移

`uv` 原生支持解析 `pyproject.toml`。虽然它有自己的格式，但它能很好地兼容标准。你可以直接在 Poetry 项目目录下尝试 `uv sync`，或者索性删除 `poetry.lock`，使用 `uv add` 重新构建依赖树。

### 场景四：从 Pipenv 迁移

如果你的项目里有 `Pipfile` 和 `Pipfile.lock`，迁移到 `uv` 非常直接：

1. **生成标准依赖**：由于 `uv` 目前主要基于 `pyproject.toml` 标准，你可以先用 `pipenv` 导出依赖：

```bash
pipenv requirements > requirements.txt
```

2. **让 uv 接管**：

```bash
# 初始化 uv
uv init
# 将导出的依赖添加到 uv 记录中
cat requirements.txt | xargs uv add
```

3. **清理旧物**：迁移完成后，你可以放心地删掉 `Pipfile` 和 `Pipfile.lock`，拥抱 `uv.lock`。

## 总结

`uv` 不仅仅是一个包安装器，它正在演变成 Python 开发的**标准工作流工具**。

* **对于新手**：不用再纠结什么是 venv，什么是系统 Python，`uv init` 之后直接写代码。
* **对于老手**：极速的安装体验和符合直觉的 `Workspaces` 管理，能大幅提升开发效率。

天下武功，唯快不破。建议大家赶紧在自己的开发机上试一试，大概率你会回不去！

