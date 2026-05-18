---
title: Python 装饰器很难？那是你没看到这篇文章！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
abbrlink: e9302aa9
date: 2026-05-18 11:00:23
img:
coverImg:
password:
summary:
---

大家好，我是 Alex。很多小伙伴在学习 Python 时，都有一座很多初学者觉得难以翻越的大山，那就是 **装饰器 (Decorator)**。

每当你看到代码里那个神秘的 `@` 符号，是不是总觉得心里没底？面试官问你“什么是闭包，什么是装饰器”时，是不是支支吾吾答不上来？

别担心，今天我们就把高深的概念扔到一边，用最通俗的语言，由浅入深，一层一层剥开装饰器的面纱。看完这一篇，保证你以后写代码时，想不用装饰器都难！

-----

## 01 为啥我们需要装饰器？

在讲技术之前，我们先来看一个实际业务场景。

假设你正在开发一个项目，里面有很多重要的业务函数，比如“处理订单”、“导出报表”等。现在老板提了一个需求：**“给每个业务函数都加上日志记录，我想知道每个函数执行了多久。”**

作为一个老实人，你可能会这样做：

```python
import time

def process_order():
    start_time = time.time()  # 记录开始时间
    print("正在处理订单...")
    time.sleep(1)             # 模拟业务耗时
    end_time = time.time()    # 记录结束时间
    print(f"耗时: {end_time - start_time} 秒")

def export_report():
    start_time = time.time()  # 重复代码...
    print("正在导出报表...")
    time.sleep(2)
    end_time = time.time()    # 重复代码...
    print(f"耗时: {end_time - start_time} 秒")
```

写完后你就会发现，**代码里充斥着大量重复的计时逻辑**。如果老板明天说“耗时要保留两位小数”，你得把几十个函数全改一遍。这简直是维护的噩梦！😱

**然鹅，我们真正想要的是：**

1.  不修改原有的 `process_order` 函数内部代码。
2.  在它执行前后，自动加上计时的功能。

这时候，当当当，**装饰器**就该闪亮登场了。

-----

## 02 一切的基础：函数也是“东西”

在 Python 中，函数是一等公民（First Class Citizen）。这话说得太学术，简单来说就是：**函数和变量一样，可以被传来传去。**

* 函数可以赋值给变量。
* 函数可以作为参数传给另一个函数。
* **函数里可以定义另一个函数，并且把它返回。**  (注意，这是一个重点！)

看个简单的例子：

```python
def outer():
    print("我是外层函数")
    
    def inner():
        print("我是内层函数")
        
    return inner  # 注意：这里返回的是函数对象，没有加括号()

# 调用
func = outer() # 执行 outer，拿到 inner
func()         # 执行 inner
```

理解了这一点，我们就可以大致理解了什么事装饰器了。

-----

## 03 打造第一个装饰器

装饰器的本质，其实就是一个**闭包**。

> 我们可以把它想象成一块“夹心饼干”。

* **原函数**（比如 `process_order`）是饼干的夹心。
* **装饰器**是在外面包的一层酥皮（添加的新功能）。

下面，我们来写一个专门用来计时的装饰器 `timer`：

```python
import time

# 这就是装饰器函数
def timer(func):
    # wrapper 是包装纸，把原函数包起来
    def wrapper():
        print(f"--- 开始执行 {func.__name__} ---")
        start_time = time.time()
        
        func()  # 这里是真正的业务逻辑（夹心）
        
        end_time = time.time()
        print(f"--- 执行完毕，耗时: {end_time - start_time:.2f} 秒 ---")
    
    return wrapper  # 把包装好的新函数送回去
```

现在，我们要用这个 `timer` 来“装饰”我们的业务函数：

```python
def process_order():
    print("正在处理订单...")
    time.sleep(1)

# 手动装饰
process_order = timer(process_order)
# 当然了，如果你觉得这里的 process_order 有歧义的话，你完全可以命名成其他的名字，比如： 
# process_order123 = timer(process_order)
# 那么调用的时候就是 process_order123()

# 再次调用 process_order，其实是在调用 wrapper
process_order()
```

**运行结果：**

```text
--- 开始执行 process_order ---
正在处理订单...
--- 执行完毕，耗时: 1.00 秒 ---
```

看！我们没有改动 `process_order` 的任何一行代码，却给它加上了计时功能。这就是装饰器的雏形。

-----

## 04 那个神奇的 @ 符号

虽然上面的 `process_order = timer(process_order)` 也能用，但写起来不够优雅。有的童鞋也就有疑问了，在 Python 的一些框架中，也没见这么用哇，都是直接一个 @ 符号就搞定了，比如，定义路由的时候。

是的，Python 贴心地为我们提供了**语法糖**（Syntactic Sugar），也就是那个 `@` 符号。

下面的代码，和上面的手动赋值是完全等价的：

```python
# 下面的 @timer 就相当于执行了: process_order = timer(process_order)
@timer
def process_order():
    print("正在处理订单...")
    time.sleep(1)

process_order()
```

是不是瞬间清爽了许多？

-----

## 05 进阶：如果原函数有参数怎么办？

上面的装饰器有个致命弱点：`wrapper` 函数没有定义参数。如果原本的 `process_order` 需要传参，程序就会报错。

为了做一个**万能的装饰器**，我们需要用到 `*args` 和 `**kwargs`。这两个兄弟可以代表“任意数量、任意类型的参数”。

**升级版装饰器 `v2.0`：**

```python
def timer_pro(func):
    # 接收任意参数
    def wrapper(*args, **kwargs):
        print(">>> 计时开始")
        start = time.time()
        
        # 把接收到的参数，原封不动地传给原函数
        # 并且接收原函数的返回值
        result = func(*args, **kwargs) 
        
        end = time.time()
        print(f">>> 计时结束，耗时 {end - start:.4f} 秒")
        
        # 千万别忘了把原函数的返回值送回去！
        return result
    
    return wrapper

@timer_pro
def add(a, b):
    time.sleep(0.5)
    return a + b

# 测试
sum_val = add(10, 20)
print(f"计算结果: {sum_val}")
```

> **💡 敲黑板：**
> 在写装饰器时，记得要在 `wrapper` 内部返回 `result`。否则，你的 `add` 函数虽然执行了，但外面拿到的结果会是 `None`（因为被装饰器“吃掉”了）。

-----

## 06 完美主义：保留函数的特征

到这里，装饰器已经能处理 99% 的情况了。但是还有一个小瑕疵。

如果你打印 `add.__name__`（查看函数名），你会发现它竟然变成了 `wrapper`！

```python
print(add.__name__) 
# 输出: wrapper
# 期望: add
```

这在调试或者生成文档时会很麻烦，相当于原函数的“身份证”被篡改了。为了解决这个问题，我们需要引入 Python 标准库中的 `functools.wraps`。

**最终完美版装饰器 `v3.0`：**

```python
from functools import wraps
import time

def timer_final(func):
    @wraps(func)  # <--- 这一行是关键！它负责把原函数的信息复制给 wrapper
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        end = time.time()
        print(f"{func.__name__} 耗时: {end - start:.4f}s")
        return result
    return wrapper

@timer_final
def sub(a, b):
    """这是一个减法函数"""
    return a - b

print(sub.__name__)  # 输出: sub
print(sub.__doc__)   # 输出: 这是一个减法函数
```

现在，这个装饰器不仅功能强大，而且“无痕”，完全保留了原函数的元数据。这才是专业 Python 开发者该有的写法。😎

-----

## 总结一下

Python 装饰器其实并不神秘，它就是一种**设计模式**：

1.  **不改变原函数代码**：符合“开闭原则”（对扩展开放，对修改关闭）。
2.  **闭包**：函数嵌套函数，内层函数引用外层变量。
3.  **应用场景**：日志记录、性能测试、权限校验（比如 Flask 中的 `@login_required`）、缓存数据等。

希望这篇文章能帮你彻底拿下装饰器！

如果你觉得有用，记得**点赞**、**收藏**、**分享**哦，笔芯 ❤️