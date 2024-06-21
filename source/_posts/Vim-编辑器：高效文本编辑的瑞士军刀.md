---
title: Vim 编辑器：高效文本编辑的瑞士军刀
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
  - Vim
abbrlink: 3325b0d1
date: 2024-06-21 23:11:32
img:
coverImg:
password:
summary:
---

Vim，作为编程和系统管理中的强大工具，以其丰富的功能和高度可定制性著称。

在这篇文章中，我们将探索 Vim 的一些高效使用技巧，从打开文件的快捷方法到文本编辑、查找、替换、删除和复制的高级技巧，再到 Vim 配置和插件安装，帮助你提升 Vim 使用技能。

## vim 打开文件的快捷方法

Vim 提供了多种打开文件的选项，让你的工作更加高效：

vim 使用的选项 | 说明
--- | ---
vim -r filename | 恢复上次 vim 打开时崩溃的文件
vim + filename	| 打开文件，并将光标置于最后一行的首部
vim +n filename	| 打开文件，并将光标置于第 n 行的首部
vim +/pattern filename	| 打幵文件，并将光标置于第一个与 pattern 匹配的位置
vim -c command filename	| 在对文件进行编辑前，先执行指定的命令

## 命令行模式下的常用命令

Vim 的命令行模式提供了丰富的快捷命令，以下是一些常用的：

命令 | 说明
---|---
set nu | 显示行号
set nonu | 取消显示行号
set ic | 忽略大小写
set noic | 取消忽略大小写
yy | 复制光标所在行
p  | 粘贴
u  | 撤销
ctrl + r | 反撤销（恢复撤销）
ctrl+d | 翻页 向下翻
ctrl+b | 翻页 向上翻
/关键字 | 查找
dd | 删除当前行
x  | 删除单个字符
o （小写）| 下行插入
O （大写 shift +o）| 上行插入
a  | 字符后插入
A | 行末插入
数字 0 或者 ^ | 光标移动到当前行的行首
$ | 光标移动到当前行的行尾
gg | 光标移动到文件开头
G | 光标移动到文件末尾


### vim 查找文本

要查找的字符串是严格区分大小写的，比如需要查找 `alex` 和 `Alex` 会得到不同的结果

快捷键 | 功能描述
--- | ---
/abc | 从光标所在位置向前查找字符串 abc
/^abc	| 查找以 abc 为行首的行
/abc$	| 查找以 abc 为行尾的行
?abc	| 从光标所在为主向后查找字符串 abc
n	| 向同一方向重复上次的查找指令
N	| 向相反方向重复上次的查找指定

### vim 替换文本

快捷键	| 功能描述
--- | ---
r	| 替换光标所在位置的字符
R	| 从光标所在位置开始替换字符，其输入内容会覆盖掉后面等长的文本内容，按“Esc”可以结束
:s/a1/a2/g	| 将当前光标所在行中的所有 a1 用 a2 替换
:n1,n2s/a1/a2/g	| 将文件中 n1 到 n2 行中所有 a1 都用 a2 替换
:g/a1/a2/g	| 将文件中所有的 a1 都用 a2 替换


### vim 删除文本

快捷键	| 功能描述
--- | ---
x	| 删除光标所在位置的字符
dd	| 删除光标所在行
ndd	| 删除当前行（包括此行）后 n 行文本
dG	| 删除光标所在行一直到文件末尾的所有内容
D	| 删除光标位置到行尾的内容

### 复制文本

快捷键	| 功能描述
--- | ---
p	| 将剪贴板中的内容粘贴到光标后
P（大写）	| 将剪贴板中的内容粘贴到光标前
y	| 复制已选中的文本到剪贴板
yy	| 将光标所在行复制到剪贴板，此命令前可以加数字 n，可复制多行
yw	| 将光标位置的单词复制到剪贴板


## 配置配置文件 `~/.vim/vimrc`

```
# 将 leader 键更改为空格键
let mapleader=" "

# 增加兼容性的常规设置
set nocompatible
filetype on
filetype indent on
filetype plugin on
filetype plugin indent on

# 让 vim 编辑器可以使用鼠标
set mouse=a
# 设置编码格式
# set encoding=utf-8
# 有些终端颜色可能会出问题，加上这一行之后就会好很多
let &t_ut=''

# 按一下 tab 键，缩进为 4 个空格
set expandtab
set tabstop=4
set shiftwidth=4
set softtabstop=4

# 显示高亮
syntax on
# 显示行号
set number
# 显示当前活动行号
set relativenumber
# 当前光标处，显示一条横线
set cursorline
# 当前行不会超出当前窗口，自动换行
set wrap
# 显示按键输出
set showcmd
# 提示
set wildmenu

# 搜索词高亮 high light search
set hlsearch
# 进入命令行模式时，自动取消高亮
exec "nohlsearch"
# 边输入，边高亮
set incsearch
# 搜索时，忽略大小写
set ignorecase
# 智能大小写
set smartcase

# 设置键盘映射
# 将 n 键映射为 h 也就是说当按了 n 键，相当于按了 h 键
noremap n h

# 设置快捷指令
# 删除小写 s 键对应的功能
map s <nop>
# 按了大写 s 键，相当于 `:w 回车` 即为快捷保存文件指令
map S :w<CR>
# 退出
map Q :q<CR>
# 重新加载 vim 配置文件
map R :source $MYVIMRC<CR>

# 向右分屏
map sl :set splitright<CR>:vsplit<CR>
# 向左分屏
map sh :set nosplitright<CR>:vsplit<CR>
# 向上分屏
map sk :set nosplitbelow<CR>:split<CR>
# 向下分屏
map sj :set splitbelow<CR>:split<CR>

# 当前屏，纵向分屏
map sv <C-w>t<C-w>H
# 当前屏，横向分屏
map sb <C-w>t<C-w>K

# 当前配置文件最上方已经将 leader 键更改为空格键，那么这里就是 空格+k 代替了 Ctrl+w
# 分屏之后，光标向上移动
map <LEADER>k <C-w>k
# 分屏之后，光标向下移动
map <LEADER>j <C-w>j
# 分屏之后，光标向左移动
map <LEADER>h <C-w>h
# 分屏之后，光标向右移动
map <LEADER>l <C-w>l

# 横向的分屏往上加 5
map <up> :res +5<CR>
# 横向的分屏往下减 5
map <down> :res -5<CR>
# 纵向的分屏宽度减 5
map <left> :vertical resize-5<CR>
# 纵向的分屏宽度加 5
map <right> :vertical resize+5<CR>

# 新建标签页
map tn :tabe<CR>
# 查看左边的标签页
map tl :-tabnext<CR>
# 查看右边的标签页
map tr :+tabnext<CR>

```

## 安装插件

在 `~/.vim/vimrc` 配置文件中

```

call plug#begin('~/.vim/plugged')

# 安装 vim-airline 插件
Plug 'vim-airline/vim-airline'
# 安装配色
Plug 'connorholyday/vim-snazzy'

call plug#end()

# 详见 https://github.com/connorholyday/vim-snazzy
# 设置配色
color snazzy
# 设置透明背景
let g:SnazzyTransparent = 1

# 然后在命令行模式下输入
# :PlugInstall
```

Vim 是一个功能强大的文本编辑器，通过熟练掌握其快捷键和配置，你可以极大提升编辑效率。希望这篇文章能帮助你更好地使用 Vim，成为文本编辑的高手。