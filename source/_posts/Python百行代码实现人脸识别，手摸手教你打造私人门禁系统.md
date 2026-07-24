---
title: Python百行代码实现人脸识别，手摸手教你打造私人门禁系统
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
  - face_recognition
  - numpy
  - opencv
abbrlink: fbc2c993
date: 2026-07-24 22:21:35
img:
coverImg:
password:
summary:
---

以前我们提到“人脸识别”，很多人的第一反应是：

> 调用百度AI接口。
> 调用腾讯云人脸识别。
> 买一套商业门禁系统。

但是对于程序员来说，总感觉少了一点参与感。

毕竟：

调用一个接口，只能证明你会调用接口。

真正有意思的是：

> 能不能自己写一个简单的人脸识别系统？

答案是：

可以。

而且并没有想象中那么复杂。

今天我们就使用 Python + OpenCV + face_recognition，实现一个完整的人脸识别 Demo：

它具备两个功能：

1. 录入人脸信息
2. 打开摄像头实时识别人脸

最终效果：

打开摄像头后：

* 如果检测到已经注册的人脸
* 系统会显示用户姓名
* 标记绿色识别框

如果是不认识的人：

* 标记红色框
* 显示 Unknown Person

整个过程：

不需要调用任何云服务。

所有计算：

都发生在你的电脑本地。

## 实现原理简单介绍

在开始写代码之前，我们先了解一下人脸识别到底是怎么工作的。

很多初学者可能认为：

> 人脸识别是不是拿照片直接比较？

实际上不是。

现代人脸识别大概分为几个步骤：

### 第一步：检测人脸

首先需要知道：

图片里面哪里有人脸。

比如：

```
+----------------+
|                |
|     人脸       |
|                |
+----------------+
```

这一步叫：

Face Detection。

### 第二步：提取人脸特征

找到人脸以后：

AI模型会把这张脸转换成一个数字向量。

例如：

```
[
0.124,
-0.532,
0.332,
...
]
```

这个过程叫：

Face Embedding。

简单理解：

就是给每个人生成一个“数字身份证”。

张三：

```
[
0.23,
0.54,
0.12
]
```

李四：

```
[
0.87,
0.21,
0.66
]
```

两个人的数据天然不同。

### 第三步：计算相似度

当摄像头再次拍到人脸：

系统会：

1. 提取当前人脸特征
2. 和数据库里面保存的数据比较

例如：

```
当前人脸:

[0.23,0.54,0.12]


数据库:

张三:
[0.24,0.53,0.11]


距离很近

=> 判断为张三
```

这就是整个识别流程。

## 环境准备

本文代码基于 Python。

建议环境：

* Python 3.9+
* Windows / MacOS / Linux 均可

项目结构：

```python
face-demo

├── face_demo.py

├── face.jpg

├── alex.pkl

└── Pipfile
```

其中：

`.pkl`

文件就是保存的人脸特征数据。

废话不多说，我们直接上干货。

> 文中提供了全部代码，你完全可以将代码拷贝过去，直接运行。

这个 demo 片段代码依赖四个核心库，你可以使用 pip 或者 pipenv 进行安装（当然也可以用最近贼火的 uv）

反正，目的只有一个，把下面的库安装好，就完事儿～

```bash
uv add face_recognition opencv-python numpy dlib
```

**专业避坑：**

这里的 `dlib` 库是真正的幕后大佬，但它在Windows系统下直接 pip install 极大概率会报错。为什么？因为它需要编译C++代码。

Windows用户： 在安装 dlib 之前，你需要先安装 **CMake**，并且在电脑上安装 **Visual Studio C++ Build Tools** 然后再执行安装命令，应该就能丝滑成功。

Mac/Linux用户：基本上直接安装即可，非常友好。

## 源码奉上

准备好环境后，新建一个 Python 文件（比如 `face_door.py`），把下面的代码复制进去。为了方便大家阅读，基本上我都写了比较详细的注释，大家认真看看基本上问题就

```python
import cv2
import face_recognition
import numpy as np
import pickle
import os

# ===================================================
# 一、录入人脸函数
# ===================================================
def register_face(img_path, name):
    """
    从指定图片中检测人脸，并保存人脸特征向量
    :param img_path: 人脸照片路径（尽量纯色背景，五官清晰）
    :param name: 用户姓名（用于文件名保存，建议不要用中文字符）
    """
    # 读取图片
    img = face_recognition.load_image_file(img_path)

    # 提取人脸特征（也就是传说中的 128维 embedding 向量）
    encodings = face_recognition.face_encodings(img)

    if len(encodings) > 0:
        # 现实情况一张图可能有多个脸，这里我们简单粗暴只取第一张人脸
        known_encoding = encodings[0]

        # 使用 pickle 将特征向量序列化保存为本地文件
        data = {"name": name, "encoding": known_encoding}
        with open(f"{name}.pkl", "wb") as f:
            pickle.dump(data, f)

        print(f"[SUCCESS] 人脸注册成功：{name}.pkl 已保存。")
    else:
        print("[ERROR] 未检测到人脸，请使用清晰正面照。")


# ===================================================
# 二、实时人脸识别函数
# ===================================================
def recognize_face():
    """
    打开摄像头实时识别人脸，与已注册人脸进行比对
    """
    # 检查当前目录下是否有注册过的人脸特征文件
    pkl_files = [f for f in os.listdir() if f.endswith(".pkl")]
    if not pkl_files:
        print("[ERROR] 未找到注册人脸文件，请先调用 register_face() 注册人脸。")
        return

    # 将所有注册过的人脸特征加载到内存中
    known_faces = []
    for file in pkl_files:
        with open(file, "rb") as f:
            known_faces.append(pickle.load(f))
    print(f"[INFO] 已加载 {len(known_faces)} 个已注册人脸。")

    # 打开摄像头
    # 【专业提示】：0 通常代表笔记本自带的内置摄像头，1 代表外接的USB摄像头。
    # 如果运行后摄像头没亮或者报错，请把这里的 1 改成 0 试试！
    cap = cv2.VideoCapture(1)
    if not cap.isOpened():
        print("[ERROR] 无法打开摄像头，请检查设备。")
        return

    print("[CAMERA] 摄像头已开启，按 'q' 退出。")

    recognition_status = ""  # 当前识别状态
    
    while True:
        ret, frame = cap.read()
        if not ret:
            break

        # OpenCV 默认读取的格式是 BGR，而 face_recognition 需要 RGB 格式
        # 所以必须做一次色彩空间转换，否则识别率会大打折扣
        rgb_frame = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
        
        # 定位画面中的所有人脸位置
        face_locations = face_recognition.face_locations(rgb_frame)
        # 提取画面中所有人脸的特征向量
        face_encodings = face_recognition.face_encodings(rgb_frame, face_locations)

        recognized = False  # 标志：是否有人脸匹配成功

        # 遍历画面中检测到的每一张人脸
        for (top, right, bottom, left), face_encoding in zip(face_locations, face_encodings):
            name = "Unknown Person" # 默认未知访客
            color = (0, 0, 255)  # 红色画框代表不认识、警告

            # 拿画面中的脸，和我们库里注册过的脸挨个比对
            for known in known_faces:
                # tolerance=0.5 是容错率。数字越小越严格（默认0.6）。
                # 调成0.5可以有效防止把长得像的人误认。
                result = face_recognition.compare_faces([known["encoding"]], face_encoding, tolerance=0.5)
                # 计算人脸距离，距离越小越相似
                distance = face_recognition.face_distance([known["encoding"]], face_encoding)

                if result[0]:
                    # 匹配成功！
                    name = f"{known['name']} Bingo! Similarity: {1 - distance[0]:.2f}"
                    color = (0, 255, 0) # 绿色代表放行
                    recognized = True
                    break

            # 使用 OpenCV 绘制识别框与名字
            # 【专业提示】：cv2.putText 默认不支持中文，强行输入中文会变成问号。
            # 这也是为什么这里的提示语我都用了英文。
            cv2.rectangle(frame, (left, top), (right, bottom), color, 2)
            cv2.putText(frame, name, (left, top - 10),
                        cv2.FONT_HERSHEY_SIMPLEX, 0.9, color, 2)

        # =======================
        # 状态提示区（左上角文字显示）
        # =======================
        if len(face_locations) == 0:
            recognition_status = "no check face"
            color = (0, 255, 255) # 黄色提示
        elif recognized:
            recognition_status = "success"
            color = (0, 255, 0)
            print("[SUCCESS] 人脸识别成功，身份验证通过。")
        else:
            recognition_status = "fail"
            color = (0, 0, 255)
            print("[ERROR] 未匹配到已注册人脸。")

        # 将当前状态渲染在画面左上角
        cv2.putText(frame, recognition_status,
                    (50, 50), cv2.FONT_HERSHEY_SIMPLEX,
                    1.0, color, 2)

        cv2.imshow("Face Recognition Demo", frame)

        # 监听键盘，按下小写字母 q 退出死循环
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break

    # 释放资源
    cap.release()
    cv2.destroyAllWindows()

# ===================================================
# 三、主程序入口
# ===================================================
if __name__ == "__main__":
    # 【第一步】：找一张你自己的清晰正脸照，命名为 face.jpg 放在代码同级目录下。
    # 然后取消下面这行代码的注释，运行一次程序，完成“人脸录入”。
    # 录入成功后，你会看到目录多了一个 zhangsan.pkl 文件，这时候可以把这行代码重新注释掉。
    
    # register_face("face.jpg", "zhangsan")

    # 【第二步】：确保已经生成了 .pkl 文件后，运行下面这个函数。
    # 你的摄像头就会亮起，把脸凑过去，见证奇迹的时刻就到了。
    recognize_face()

    # 【附加功能】：如果你不知道自己的摄像头编号，取消下面这段注释运行一下。
    # 它会帮你测试并打印出哪些摄像头是可用的。
    # for i in range(4):
    #     print(f"测试摄像头 {i} ...")
    #     cap = cv2.VideoCapture(i)
    #     if cap.isOpened():
    #         print(f"[SUCCESS] 摄像头 {i} 可用")
    #         cap.release()
```

好了，上面就是所有的源码了，当然，上面的代码距离工业使用上还有很多地方需要优化的，比如：人脸数据库存储呀，多人脸识别，活体检测什么的，但这对于大家想跑跑 demo ，想了解人脸识别是怎么做的，应该还是足够了。

希望你玩得愉快～
