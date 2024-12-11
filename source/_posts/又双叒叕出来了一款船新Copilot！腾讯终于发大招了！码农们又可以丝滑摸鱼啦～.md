---
title: 又双叒叕出来了一款船新Copilot！腾讯终于发大招了！码农们又可以丝滑摸鱼啦～
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: AI
tags:
  - AI
  - Copilot
  - 编程助手
abbrlink: bd2219d
date: 2024-12-12 00:00:58
img:
coverImg:
password:
summary:
---

这几天在腾讯云开发者社区写文章时，老是给我下面这个弹窗提示，要我去体验一下这款 **AI 代码助手**。

![腾讯云天天提示我](https://upload-images.jianshu.io/upload_images/14623749-246d69a60d716959.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

说实话，自从去年年底 GitHub Copilot 腾空出世之后，陆陆续续出现了不少好用的 AI 代码助手，像 **CodeGeeX 智能编程助手**、JetBrains 官方出品的 **JetBrains AI Assistant**、字节跳动旗下豆包的 **MarsCode AI**、再到 **Cursor AI 编辑器**…… 在 AI 代码助手这方面，真可谓百花齐放，其中完全免费的，有体验期限的，直接需要付费的……一大把，数不胜数。

## GitHub Copilot 初体验感受

我记得去年 GitHub Copilot 刚出来那会儿，官方提供了为期 30 天（具体是 30 天，还是 60 天，有点儿记不太清了）的体验期，于是，我就赶紧在 IDE 中装了 GitHub Copilot 插件来体验，可能是初次体验这种 AI 代码助手吧，一时就有点儿像刘姥姥进大观园一般，哈哈😂。不禁觉得这玩意儿真是太神奇了！比如，我想要写一个函数方法，我只需要将注释信息写好，然后写一个 `function` 开头，后面的方法名、方法逻辑体完全可以一路敲击 `Tab` 出来！完全不需要自己写一丁点儿业务逻辑代码！

有时候，我想借鉴一下一些开源代码段的时候，如果是以前的编码方式，我会一个窗口打开编辑器，另外一个窗口打开浏览器，然后一边看浏览器中的代码是否符合自己的实际需要，再往编辑器中写，然而，现在我发现似乎我身边站了一位隐身的人，他时时刻刻盯着我的屏幕，并且就跟我肚里的蛔虫一样，知道我大部份想法，我唯一需要做的就是敲几次 `Tab` 键！

后来，体验期过了之后，我就没有续费，然而，我发现没使用 GitHub Copilot 之后，一时半会儿，我竟然还有些许不太习惯了。当然，最后我还是用上了 GitHub Copilot，并且到现在为止依旧用的好好的。在今年，受邀也体验过豆包的 MarsCode，体验下来，感觉和 GitHub Copilot 相差不大，但是，在细节处理方面可能 GitHub Copilot 会更加细致一点。~~之前也写过一篇文章，做了一下简单的对比，感兴趣的童鞋可以去翻一翻历史文章，或者去**我的公主耗「左诗右码」**上找找~~。可是，不得不说，毕竟 MarsCode 免费呀！嗯……真香！

到现在为止也体验过不少 AI 代码助手了，但是，发现功能都大同小异，并没有太多惊喜。可是，既然腾讯也出了一款 AI 代码助手，再怎么说也是有大厂金主爸爸背书的，想必应该多多少少还是会有一些独特之处的吧？说不多说，直接开始体验……

## 体验「腾讯云 AI 代码助手」

打开首页，就直接被**个人免费使用**这 6 个字吸引了，这简直是我这种穷13的福音啊！

![首页](https://upload-images.jianshu.io/upload_images/14623749-a0be8eba0f7892b2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 主要功能

作为一款 AI 代码助手，比较“通用”的几个功能都还是有的。

1. 补全行内代码
2. 根据上文补充下文
3. 函数代码块内补全代码
4. 智能对话

并且腾讯云 AI 代码助手也提供了和 MarsCode AI 类似的 **在线 IDE** 的功能，个人一直都比较喜欢这个功能，因为在有些场景下，可能我只需要简单的跑一下某段示例代码，看下运行效果，但是我本地又不想搭建环境时，这个**在线 IDE**的功能就能够很好的帮我解决问题。

当然了，这个功能也不是一个非常新奇的功能，隔壁的 GitHub 在早几年就已经有了在线 IDE 的功能，但是一直以来用的就不是那么频繁。一方面是因为一些“众所周知”的国情原因，另外一方面是有些编程环境还是得自己搭建（不知道现在是否已经有默认的编程语言环境，我已经很久没有体验过了，如果我说错了，就请各位大神在评论区指正一下🤝）。

## 支持多种编程语言和编辑器

主流的编程语言和编辑器都支持。~~要是主流的某门编程语言或者编辑器不支持，估计开发这款 AI 代码助手的程序员就要被拉去祭天了吧……~~

![支持的语言和IDE](https://upload-images.jianshu.io/upload_images/14623749-6a833c547f03a16c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

那么，如何使用呢？

## 使用方式

### Visual Studio Code

直接在插件市场搜索框中输入 **Tencent Cloud AI Code Assistant**，然后直接点击 `Install` 即可安装。

![就是这个](https://upload-images.jianshu.io/upload_images/14623749-5c0fa24f51dd1b1b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

具体的一些快捷键，可以参考如下：

![VS Code 中快捷键](https://upload-images.jianshu.io/upload_images/14623749-4d0f8440b27791e6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### JetBrains IDEs

你可以去你用到的 IDE 中找到**插件**，比如，这里我用的是 PyCharm，然后你可以直接搜索 **腾讯云** 出现下面截图中的图标，然后直接点安装即可。

![就是这个](https://upload-images.jianshu.io/upload_images/14623749-6d07d911c5b96882.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

具体的一些快捷键，可以参考如下：

![IDE 中快捷键](https://upload-images.jianshu.io/upload_images/14623749-ddc85f6070c74585.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

如果你之前使用过 GitHub Copilot 的话，你会发现二者的快捷键不能说几乎一样，简直可以用**完全**来形容！当然，这样设计也是有好处的，至少对于我们使用者来说，可以非常丝滑的在各个 AI 代码助手工具直接来回切换。

## 实际体验

上文中，我也说到了，我现在还在使用 GitHub Copilot，并且我也想对比一下二者之间的差异，再者我也非常想体验一下腾讯云的在线 IDE。之前我在玩豆包的 MarsCode AI 时，也是直接用的在线 IDE，因此下面我就以腾讯云的在线 IDE 为基础来体验一下。

想要体验**在线 IDE**，可以直接在首页点击**IDE在线体验**即可跳转到编辑器界面

![IDE在线体验](https://upload-images.jianshu.io/upload_images/14623749-7b41a25c599d207b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后会慢慢的等着启动，看这启动界面，还有点儿酷呢。

![启动界面](https://upload-images.jianshu.io/upload_images/14623749-6e676650acdc3f85.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

在启动的过程中，需要我们登录，我个人建议的是：**直接用自己的微信授权登录**，因为腾讯的很多产品都是可以共用授权的，直接用微信授权登录，一方面可以少记一次账号密码，另外一方面还可以让自己的账号在多个产品间数据共享。当然，具体情况具体分析了，根据自己的实际情况来定就行。

然而，当我们怀着期待，打开编辑器界面时

![在线编辑器界面](https://upload-images.jianshu.io/upload_images/14623749-3c35b39372a3a556.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

惊不惊喜！？意不意外！？是不是跟 VS Code 就是孪生兄弟？

算了，今天，我们不是来讨论编辑器界面的，还是来体验一下代码提示究竟如何吧！是骡子是马，总还是要拿出来溜一溜的。

### 线上 IDE 编辑代码

![界面](https://upload-images.jianshu.io/upload_images/14623749-35590abca2f323fa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

默认为我们准备了 5 个代码示例，现在我用 go 代码作为示例，让它帮我写一段快速排序。

![有点儿翻车了](https://upload-images.jianshu.io/upload_images/14623749-b115ddb5a4db3db2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然而，他竟然一下子给了我这么多莫名其妙的注释？什么鬼！？

可能是我提示的不太对，或许是这个问题太简单了？AI 不屑一顾？后来，我就再按了一次 `Enter` 键然后有提示时，我就持续敲 `Tab` 键，坚决不做一点儿人为干预

![依旧还是有点儿翻车](https://upload-images.jianshu.io/upload_images/14623749-ce4a3b47aaf192db.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然而，前面回答的好好的，后面又开始**稀里糊涂的胡说八道**了。

这不对劲儿呀，应该没有那么水吧。我一度认为，肯定是我的操作问题，才导致他没有正确提示的。于是，我删除了原始文件中的所有代码，然后，重新在文件首行加了注释信息，之后再一路 `Tab` ，最终的结果就如下：

![最终结果](https://upload-images.jianshu.io/upload_images/14623749-8a5235179b3a4680.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

终于，终于，没有辜负我的期望，打印结果也正确！

代码如下：

```go
// 帮我写一段快速排序的代码

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("before sort:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("after sort:", arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	pivot := arr[(left+right)/2]
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	quickSort(arr, left, j)
	quickSort(arr, i, right)
}
```

我们从图中也可以看到，该**线上 IDE**默认是安装了 `go` 语言环境的，这为我们调试代码提供了不少方便。当然，除了 `go` 语言编程环境，`python` 、`java`、`node` 也都有默认的环境。

### 体验 AI 对话

当然，在线上 IDE 中，我们也可以非常方便的和 AI 代码助手进行对话，直接点击左侧栏中的这个按钮即可。

![AI 对话](https://upload-images.jianshu.io/upload_images/14623749-67f844ad6904c837.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

你可以直接在对话框中问任何问题，同时，也可以**就某个代码文件中的某个片段进行提问**，相当的 nice！

![对话](https://upload-images.jianshu.io/upload_images/14623749-5e548240fa21358f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 体验感受

初步体验下来，感觉**腾讯云 AI 代码助手**在功能上和 `Github Copilot` 以及豆包的 `MarsCode` 没有太大的差别，该有的功能都有了。可能腾讯云 AI 代码助手还有其他新奇的功能，我还没有体验到，希望能够在日后工作中慢慢发掘吧。

当前，我仅仅从我日常工作的角度去体验了一下这款工具，不一定十分客观，但是我能够肯定的说，**腾讯云 AI 代码助手这款工具，确实很不错！** 有了这款工具的加持，确实能够有效的提高我们的工作效率。**非常值得用起来！** 现在，再也不是 GitHub Copilot 一家独大了，目前我体验过的 `MarsCode` 和 **腾讯云 AI 代码助手** 都是非常不错的平替！

## 众多 AI 编程助手简单对比（来自各个网友评价）

另外，我**简单汇总了一下网友们对各种 AI 编程助手的评价**，希望对你有所参考。同时，也希望你能够在评论区发表你独特的见解。

| AI编程助手             | 优点                                                         | 缺点                                                         |
|----------------------|------------------------------------------------------------|------------------------------------------------------------|
| GitHub Copilot        | - 提供跨多种编程语言的AI驱动代码补全和生成<br>- 基于OpenAI的Codex模型，能够根据上下文生成高质量代码建议<br>- 由GitHub和微软推出<br>- 机器学习技术提供实时代码建议<br>- 支持多种编程语言<br>- 减少人为失误，提高工作效率<br>- 快速提供示例代码，减少查阅文档时间 | - 部分生成的代码可能不符合项目最佳实践，需谨慎审查<br>- 对隐私敏感项目不太友好，因为需要依赖云端处理<br>- 需要付费<br>- 网络延迟问题经常发生                             |
| 豆包 MarsCode         | - 由字节跳动公司推出<br>- 提供智能化的代码补全、生成、优化等功能<br>- 支持多种编程语言和主流IDE | - 目前市场推广有限，用户数量较少<br>- 需要与更成熟的工具竞争，部分功能可能尚待完善 |
| Cursor               | - 注重隐私和自然语言编程<br>- 提供智能且快速的代码补全<br>- 全面的代码建议<br>- 多文件编辑<br>- 集成文档<br>- 上下文感知聊天 | - 作为新兴工具，可能在某些方面不如一些成熟的AI编程助手 |
| bolt.new             | - 全栈Web应用开发沙盒<br>- 支持多种编程语言和技术栈<br>- 提供即时交互式编程环境<br>- 允许快速部署应用<br>- 集成版本控制和协作功能 | - 可能在某些高级功能上不如一些专业的AI编程助手 |
| 通义灵码             | - 阿里巴巴推出的AI编程助手<br>- 基于通义大模型<br>- 支持多种编程语言和开发环境<br>- 特别适用于企业级项目的代码生成和优化<br>- 阿里云推出的基于通义大模型<br>- 兼容VisualStudioCode、JetBrainsIDEs等主流IDE | - 对个人开发者不太友好，产品偏向企业客户<br>- 需要学习曲线，初次上手较为复杂<br>- 代码补全部分是基本可用的，有Copilot的70~80%的能力<br>- 对用户代码含义理解，尤其是对用户本地输入的代码，注释和用户提出的问题的理解，距离Copilot还有较大差距 |
| 代码小浣熊 (Raccoon)   | - 商汤科技推出的AI编程助手<br>- 特别针对初学者和中级开发者<br>- 提供智能代码生成和辅助工具<br>- 基于商汤自研大语言模型<br>- 支持Python、Java、JavaScript、C++、Go、SQL等30+主流编程语言<br>- 支持VSCode、IntelliJIDEA等主流IDE | - 目前功能较为基础，高级功能可能不如其他竞争产品<br>- 仅支持主流编程语言，未见对更复杂场景的深度支持<br>                             |
| 文心快码             | - 百度推出的AI编程助手<br>- 基于文心大模型（ERNIE）<br>- 支持超过100种编程语言<br>- 能够帮助开发者在多种语言环境下实现实时代码补全、生成和优化 | - 高级功能可能收费，对于个人开发者的成本较高<br>- 虽然支持多种语言，但在某些冷门语言中的表现尚待提升 |
| iflycode             | - 科大讯飞推出的智能编程助手<br>- 结合了其在自然语言处理和语音识别方面的技术优势<br>- 为开发者提供流畅、直观的代码生成和补全功能 | - 功能相对有限，复杂场景下可能表现不够稳定<br>- 语音编程在一些编程语言中可能不太适用 |
| CodeGeeX             | - 开源工具，免费提供全部功能<br>- 支持Python、C++、Java、JavaScript、Go等10多种主流编程语言<br>- 开发者可以自由下载使用，并根据需要进行修改和二次开发 | - 智能提示方面还有待提高                             |

## 写在最后

还是那句话，**不管你用哪款工具，我都建议你趁早把 AI 用起来。** 你可以选择停滞不前，但永远无法阻止时代前进的步伐。