---
title: 使用 pyenv 管理多个版本 Python
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
abbrlink: 9fc3a49a
date: 2024-01-15 14:23:50
img:
coverImg:
password:
summary:
---

[pyenv](https://github.com/pyenv/pyenv) 是 Python 的版本管理工具，利用它可以在同一台电脑上安装多个版本的 Python。

## 安装方式

方式一：MacOS 下可以直接通过 Homebrew 来安装

```bash
brew update
# 安装 pyenv
brew install pyenv
# 如果是升级时
brew upgrade pyenv
```

方式二：通过自动程序来安装

```bash
curl https://pyenv.run | bash
```

方式三：通过源码的方式来安装

```bash
git clone https://github.com/pyenv/pyenv.git ~/.pyenv
# 可选的：可以尝试编译一个动态的 bash 扩展来加速 Pyenv ，如果失败了也不用担心，Pyenv 仍然可以正常工作
cd $(pyenv root) && src/configure && make -C src
```

如果是通过源码来安装的话，需要升级 pyenv 时，可通过

```bash
cd $(pyenv root)
git fetch
git tag
git checkout {tag_name}
# eg：git checkout v0.1.0
```

## 添加环境变量

如果你是用的 `Bash` 则需要：

```bash
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bashrc
echo 'command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(pyenv init -)"' >> ~/.bashrc
```

如果你是用的 `Zsh` 则需要：

```bash
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.zshrc
echo '[[ -d $PYENV_ROOT/bin ]] && export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.zshrc
echo 'eval "$(pyenv init -)"' >> ~/.zshrc
```

## pyenv 使用

检查 pyenv 是否安装成功

```bash
pyenv -v
```

查看 pyenv 指令列表

```bash
pyenv commands
```

查看所有可选的 python 版本

```bash
pyenv install -l
```

查看已经安装的所有 python 版本

```bash
# 版本号前面带有 * 号的，则证明当前使用的就是这个版本
pyenv versions
```

显示当前的 Python 版本及其本地路径

```bash
pyenv version
```

下载指定版本的 Python

```bash
# 比如这里下载 python 3.9.6
pyenv install 3.9.6

# 安装完毕之后记得刷新 pyenv shims
pyenv rehash
```

如果安装 python 比较慢时的解决方案：

首先在 pyenv 的根目录创建一个 cache 文件夹

```bash
mkdir -p $(pyenv root)/cache/
```

然后我们先执行一下安装命令，比如 `pyenv install 3.9.6` 它会显示出下载链接，此时，我们直接复制下载链接地址，通过浏览器下载，然后将下载后的 `Python-3.9.6.tar.xz` 文件放到 `$(pyenv root)/cache/` 文件夹中，然后再次执行 `pyenv install 3.9.6` 命令即可，它会自动使用 cache 文件夹中的安装包。

**下载后的 Python 直接在 `$(pyenv root)/versions/` 目录下。**

卸载指定版本的 Python

```bash
# 比如这里卸载 python 3.9.6
pyenv uninstall 3.9.6
```

切换 Python 版本

```bash
# 全局切换（系统全局用系统默认的 Python 比较好，不建议直接对其操作）
pyenv global 3.9.6

# 通过查看 Python 版本来确定是否切换成功
python -V

# 切换回系统版本
pyenv global system

# 用 local 进行指定版本切换，一般开发环境使用（只影响当前所在文件夹）
pyenv local 3.9.6

# 解除 local 设置
pyenv local --unset

# 当前 shell 会话切换，退出 shell 会话后失效
pyenv shell 3.9.6

# 解除 shell 设置
pyenv shell --unset
```

## 卸载 pyenv

```bash
# 如果是使用 homebrew 下载的 pyenv 时，卸载为
brew uninstall pyenv

# 如果是使用 git 拉取源码安装的 pyenv 时，卸载为
rm -rf $(pyenv root)
```

## 结合 pyenv 为每个项目建立自己的虚拟环境

### 如果用的是 pipenv 时

Pipenv 是个包管理工具，它综合了 virtualenv，pip 和 pyenv 三者的功能。可以使用 pipenv 这个工具来安装、卸载、跟踪和记录依赖性，并创建、使用和组织你的虚拟环境。

使用 Homebrew 安装 pipenv

```bash
# 安装 pipenv
brew update && brew install pipenv
# 更新 pipenv
brew update && brew upgrade pipenv
```

也可以使用 pip 来安装和升级 pipenv

```bash
# 安装 pipenv
pip install pipenv
# 更新 pipenv
pip install --upgrade pipenv
```

使用 pipenv 建立虚拟环境

```bash
# 如果系统中没有你想要的 Python 版本 {python_version} 时，pipenv 会调用 pyenv 来安装对应的 Python 版本
cd {your_project_dir} && pipenv --python {python_version}
# eg：pipenv --python 3.9.6
```

我们可以通过 `pipenv --venv` 来查看项目的虚拟环境目录，可以通过 `pipenv --rm` 来删除虚拟环境。


### 如果用的是 [virtualenv](https://github.com/pypa/virtualenv) 时

如果我们安装了 pyenv 时，其实已经自动以 plugin 的形式安装好了 virtualenv 我们只需直接使用就好了。

如果没有安装 virtualenv 时，则可以使用以下命令进行安装

```bash
# 使用 pip 进行安装
pip install virtualenv
# 或者使用 homebrew 安装
brew install pyenv-virtualenv
```

使用 virtualenv 建立虚拟环境

```bash
pyenv virtualenv {python_version} {virtual_env_name}
# eg：pyenv virtualenv 3.9.6 envdemo396

# 查看虚拟环境目录
ls -al $(pyenv root)/versions

# 也可以直接通过 pyenv 命令进行查看
pyenv versions
```

查看当前有哪些虚拟环境

```bash
pyenv virtualenvs
```

激活新创建的虚拟环境

```bash
# 以 envdemo396 虚拟环境为例
pyenv activate envdemo396
```

手动退出虚拟环境

```bash
pyenv deactivate
```

删除虚拟环境

```bash
# 以 envdemo396 虚拟环境为例
rm -rf $(pyenv root)/versions/envdemo396
```