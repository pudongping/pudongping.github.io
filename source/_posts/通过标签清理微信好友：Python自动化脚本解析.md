---
title: 通过标签清理微信好友：Python自动化脚本解析
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python
  - 自动化
abbrlink: 5e38f4b6
date: 2024-12-02 10:08:03
img:
coverImg:
password:
summary:
---

微信已经成为我们日常生活中不可或缺的社交工具。随着使用时间的增长，我们的微信好友列表可能会变得越来越臃肿。

在上一篇文章中，我写了一个自动根据用户状态（好友将我们删除了还是拉黑了）将用户打上不同标签的工具。那么，已经将用户打好了标签之后，我们就可以**根据指定的标签名来直接删除好友了**。

在本文中，我将逐步分析这个用 Python 编写的自动化脚本，它可以通过标签批量清理微信好友。该脚本使用了 `uiautomator2` 库，这是一个强大的Android UI 自动化工具，广泛用于 Android 设备的自动化操作。本脚本就是通过模拟用户的点击、滑动等操作，实现自动化清理指定标签下的微信好友。

## 环境准备

1. 安装 Python 3.x
2. 安装 uiautomator2 库：`pip install uiautomator2`
3. 准备一台 Android 设备，并开启开发者模式
4. 确保设备与电脑在同一网络环境下

## 使用步骤

1. 首先给要清理的好友打上统一的标签（这点也可以直接运行上文我分享的脚本，可以自动化打标签）
2. 然后在脚本中设定脚本名称，执行脚本。

以下是对代码的逐行分析，希望帮助大家更好地理解脚本的实现过程。

## 代码分析

```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
通过标签清理微信好友
"""
import time
import uiautomator2 as u2
```

### 说明：

- `#!/usr/bin/env python3` 是脚本的设定，告诉操作系统用 Python3 来执行该脚本。
- `# -*- coding: utf-8 -*-` 用于指定文件的编码格式为UTF-8，避免中文字符出现乱码。
- 导入了 `time` 模块，用于实现延时操作。
- `uiautomator2` 是一个用于控制 Android 设备 UI 的 Python 库，本文脚本依赖它来执行自动化任务。

## 类定义

```python
class WXClearFriendByTag:
    """
    通过标签清理微信好友
    """
    def __init__(self, tag_name: str, max_times_tag: int = 5, max_times_friend: int = 5):
        self.d = u2.connect()  # 连接设备
        self.clear_tag_name = tag_name  # 需要被清理好友的标签名
        self.max_times_tag = max_times_tag  # 查找指定标签的最大循环次数
        self.max_times_friend = max_times_friend  # 清理好友的最大循环次数
```

### 说明：

- `WXClearFriendByTag` 类的作用是通过标签清理微信中的好友。初始化时，我们传入标签名(`tag_name`)，以及查找标签和清理好友的最大次数。
- `self.d = u2.connect()` 用于连接设备，`u2.connect()`会自动选择连接到的设备。如果有多个设备，默认连接第一个设备。
- `max_times_tag` 和 `max_times_friend` 分别指定了查找标签和清理好友的最大次数。

## 启用调试模式

```python
def enable_debug(self):
    """
    开启调试模式
    :return:
    """
    print(f"设备信息 ==> {self.d.info}")
    print(f"设备IP ==> {self.d.wlan_ip} 设备号 ==> {self.d.serial}")
    self.d.implicitly_wait(30)  # 设置查找元素等待时间，单位秒
```

### 说明：
- `enable_debug` 方法用于启用调试模式，打印设备信息（如设备IP、设备号）以及设置UI元素查找的默认超时时间。
- `self.d.implicitly_wait(30)` 设置了 UI 元素查找的最大超时时间为 30 秒，意味着在操作前，系统会等待最多 30 秒来查找元素。

## 进入标签页面

```python
def go_to_tag_page(self):
    """
    进入标签页面
    :return:
    """
    print("打开【通讯录】")
    self.d(text='通讯录', className='android.widget.TextView').click_exists(timeout=3)
    print("打开【标签】")
    self.d(text='标签', className='android.widget.TextView').click_exists(timeout=3)
```

### 说明：
- `go_to_tag_page` 方法模拟点击进入微信的“通讯录”页面，再点击进入“标签”页面。`click_exists(timeout=3)` 用来判断元素是否存在，如果存在则点击。

## 查找标签并点击

```python
def find_tag_click(self):
    i = 1
    while i <= self.max_times_tag:
        print(f"第{i}次查询")
        status = self.find_tag_by_name()
        if status == 'done':
            break
        else:
            self.d.swipe(100, 1000, 100, 200)
            i += 1
    if i > self.max_times_tag:
        print(f"没有找到标签：{self.clear_tag_name}")
        exit()

    tag = self.d(text=self.clear_tag_name)
    if tag.exists() is False:
        print(f"没有找到标签：{self.clear_tag_name}")
        exit()

    print(f"找了{i}次，找到了标签：{self.clear_tag_name} 点击")
    tag.click_exists(timeout=5)
```

### 说明：
- `find_tag_click` 方法的核心是查找指定标签并点击。首先通过 `find_tag_by_name()` 查找标签，若没有找到，就会通过滑动屏幕继续查找，最多重复 `max_times_tag` 次。
- 滑动操作 `self.d.swipe(100, 1000, 100, 200)` 会从屏幕下方滑动至上方，模拟手动滑动以加载更多标签。
- 如果找到了目标标签，会点击该标签进入标签页面。

## 通过标签名查找标签

```python
def find_tag_by_name(self):
    run_status = 'doing'
    elems = self.d(resourceId="com.tencent.mm:id/hs8")
    for elem in elems:
        tag_name = elem.get_text(timeout=10)
        friend_count = elem.sibling(resourceId="com.tencent.mm:id/hs7").get_text(timeout=10)
        print(f"标签是：{tag_name}，好友数：{friend_count}")
        if tag_name == self.clear_tag_name:
            run_status = 'done'
            break
    return run_status
```

### 说明：
- `find_tag_by_name` 方法通过标签的resourceId查找所有标签，获取每个标签的名称和好友数。
- 如果找到指定的标签名，会返回 `'done'`，否则返回 `'doing'`，供 `find_tag_click` 方法判断是否继续查找。

## 清理每个好友

```python
def clear_every_friend(self):
    """
    清理每个好友
    :return:
    """
    time.sleep(3)
    elems = self.d(resourceId='com.tencent.mm:id/kbq')
    if elems.exists() is False:
        print("没有找到好友")
        exit()

    for elem in elems:
        time.sleep(1)
        friend_nickname = elem.get_text(timeout=10)
        print(f'进入好友详情页面 --> {friend_nickname}')

        elem.click(timeout=5)

        self.d(resourceId='com.tencent.mm:id/coy').click(timeout=5)
        time.sleep(1)
        self.d(text='删除').click(timeout=5)
        time.sleep(3)
        self.d(text='删除', resourceId='com.tencent.mm:id/mm_alert_ok_btn').click(timeout=5)
        time.sleep(1)
```

### 说明：
- `clear_every_friend` 方法负责清理每个好友。首先，通过`resourceId='com.tencent.mm:id/kbq'` 获取所有好友元素。
- 点击每个好友，进入好友详情页面后，点击右上角的菜单，选择删除该好友。
- 删除操作是通过点击【删除】按钮实现的，并最终点击确认框中的【删除】。

## 清理标签中的所有好友

```python
def clear_friends_in_tag(self):
    """
    在标签中清理好友
    :return:
    """
    i = 1
    while i <= self.max_times_friend:
        print(f"第{i}次清理好友")
        self.clear_every_friend()
        i += 1

    print(f"清理了{i}次好友")
```

### 说明：
- `clear_friends_in_tag` 方法用于在标签中清理所有好友。它会循环调用 `clear_every_friend` 方法，直到达到 `max_times_friend` 次。

## 脚本入口

```python
if __name__ == '__main__':
    tag_name = "清粉-账号问题"
    wx = WXClearFriendByTag(tag_name)
    wx.enable_debug()
    wx.find_tag_click()
    wx.clear_friends_in_tag()
```

### 说明：
- `if __name__ == '__main__':` 是Python脚本的入口，确保脚本从此处开始执行。
- 创建 `WXClearFriendByTag` 类的实例，传入标签名“清粉-账号问题”，并依次调用 `enable_debug`、`find_tag_click` 和 `clear_friends_in_tag` 方法，完成清理操作。

## 总结

这段代码展示了如何使用 `uiautomator2` 库进行微信好友管理自动化。通过标签筛选好友并删除，非常适合需要批量清理好友的场景。

对于初学者来说，理解此脚本能帮助你掌握如何使用 `uiautomator2` 控制 Android 设备，执行 UI 操作。

希望本文的解析能帮助你更好地理解代码实现。如果你有更多问题，欢迎随时提问！