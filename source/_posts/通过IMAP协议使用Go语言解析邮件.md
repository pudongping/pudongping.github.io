---
title: 通过IMAP协议使用Go语言解析邮件
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
  - IMAP
abbrlink: '60016311'
date: 2024-08-02 14:18:46
img:
coverImg:
password:
summary:
---

电子邮件在现代通信中依然扮演着重要的角色。为了提升邮件处理的效率，使用编程语言进行自动化处理变得尤为重要。

本文将详细介绍如何使用 Go 语言从 IMAP 服务器读取邮件，解析邮件内容，并存储或处理所需信息。

说到邮件服务，我们就得先了解几个和邮件相关的协议。

## 什么是 POP3/IMAP/SMTP 服务
- POP3 （Post Office Protocol - Version 3）协议用于支持使用电子邮件客户端获取并删除在服务器上的电子邮件。
- IMAP （Internet Message Access Protocol）协议用于支持使用电子邮件客户端交互式存取服务器上的邮件。
- SMTP （Simple Mail Transfer Protocol）协议用于支持使用电子邮件客户端发送电子邮件。

## IMAP 和 POP 有什么区别

SMTP 协议就不用多说了，专门用于发送邮件，这个协议也是我们在编程开发中用的最多的协议之一。

POP 允许电子邮件客户端下载服务器上的邮件，但是你在电子邮件客户端上的操作（如：移动邮件、标记已读等）不会反馈到服务器上的，比如：你通过电子邮件客户端收取了 QQ 邮箱中的 3 封邮件并移动到了其他文件夹，这些移动动作是不会反馈到服务器上的，也就是说，QQ 邮箱服务器上的这些邮件是没有同时被移动的。

需要特别注意的是，第三方客户端通过 POP 收取邮件时，也是有可能同步删除服务端邮件。在第三方客户端设置 POP 时，请留意是否有 **保留邮件副本/备份** 相关选项。*如有该选项，且要保留服务器上的邮件，请勾选该选项。*

在 IMAP 协议上，电子邮件客户端的操作都会反馈到服务器上，你对邮件进行的操作（如：移动邮件、标记已读、删除邮件等）服务器上的邮件也会做相应的动作。也就是说，IMAP 是“双向”的。同时，IMAP 可以只下载邮件的主题，只有当你真正需要的时候，才会下载邮件的所有内容。**在 POP3 和 IMAP 协议上，QQ邮箱推荐你使用IMAP协议来存取服务器上的邮件。**

## 授权码

在我们开发之前，我们需要先准备好对应邮箱的**授权码**，这个授权码是邮箱用于登录第三方客户端/服务的专用密码，适用于登录以下服务：POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV 服务。

**不同的邮箱会有不同的获取方式，但是一般获取方式都非常简单**，可以自行通过搜索引擎检索一下即可。

比如：QQ 邮箱的授权码的获取方式是：

在邮箱**[帐号与安全](https://wx.mail.qq.com/account)**点击 **设备管理 > 授权码管理**，对授权码进行管理，即可获得。

## 实战

今天我们就通过 Go 语言来演示一下如何解析邮件。

首先我们先下载第三方 imap 协议的插件包：

```bash
go get -v github.com/emersion/go-imap@v1.2.1
```

接下来的就是示例代码，很多重要信息，在代码里都有注释信息，因此请多留意注释：

下面的代码逻辑大致是：读取指定邮箱中的收件箱邮件，每次读取 2 封邮件，并解析出邮件的主题、收件人、发件人、收件时间、邮件正文，读取完毕之后，给每封邮件标记已读。

```go
package mail_parse

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const (
	Addr          string = "imap.qq.com:993"
	UserName      string = "123456789@qq.com" // 邮箱地址
	Password      string = ""                 // 这里的密码是使用开启 imap 协议后对应的服务商给到的密码，不是邮箱账号密码
	Folder        string = "INBOX"            // 邮箱文件夹，比如： INBOX 收件箱、Sent Messages 发件箱、Drafts 草稿箱、Trash、Junk 垃圾箱
	ReadBatchSize int    = 2                  // 每次读取的邮件数量
)

// IMAP（Internet Message Access Protocol）是一种用于在互联网上访问电子邮件的协议。
// 它允许用户通过 Internet 访问他们在邮件服务器上存储的电子邮件。
// Go 语言的 go-imap 库是一个用于从 IMAP 服务器获取电子邮件的库，它可以帮助你在 Go 代码中访问 IMAP 协议

func ReadEmail() {
	log.Println("开始连接服务器")

	// 建立与 IMAP 服务器的连接
	c, err := client.DialTLS(Addr, nil)
	if err != nil {
		log.Fatalf("连接 IMAP 服务器失败: %+v \n", err)
	}
	log.Println("连接成功！")
	// 最后一定不要忘记退出登录
	defer c.Logout()

	// 登录
	if err := c.Login(UserName, Password); err != nil {
		log.Fatalf("邮箱[%s] 登录失败: %v \n", Addr, err)
	}
	log.Printf("邮箱[%s] 登录成功！\n", UserName)

	// 列出当前邮箱中的文件夹
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1) // 记录错误的 chan
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	log.Println("-->当前邮箱的文件夹 Mailboxes:")
	var folderExists bool
	for m := range mailboxes {
		log.Println("* ", m.Name)
		if m.Name == Folder {
			folderExists = true
		}
	}
	if err := <-done; err != nil {
		log.Fatalf("列出邮箱列表时，出现错误：%v \n", err)
	}
	log.Println("-->列出邮箱列表完毕！")
	if !folderExists {
		log.Fatalf("文件夹[%s] 不存在 \n", Folder)
	}

	// 选择指定的文件夹
	mbox, err := c.Select(Folder, false)
	if err != nil {
		log.Fatalf("选择邮件箱失败: %v \n", err)
	}
	log.Printf("mbox %+v \n", mbox)
	log.Printf("当前文件夹[%s]中，总共有 %d 封邮件 \n", Folder, mbox.Messages)
	if mbox.Messages == 0 {
		log.Fatalf("当前文件夹[%s]中没有邮件", Folder)
	}

	// 创建一个序列集，用于批量读取邮件
	seqset := new(imap.SeqSet)

	// 假设需要获取最后4封邮件时
	// from := uint32(1)
	// to := mbox.Messages // 此文件下的邮件总数
	// if mbox.Messages > 3 {
	// 	from = mbox.Messages - 3
	// }
	// seqset.AddRange(from, to) // 添加指定范围内的邮件编号

	// 搜索指定状态的邮件
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag} // 未读邮件标记
	// criteria.WithFlags = []string{imap.SeenFlag} // 已读邮件标记
	uids, err := c.Search(criteria)
	// 在这里也可以使用 UidSearch 方法，但是用了 UidSearch 方法后，下面的很多方法都需要使用 Uid 开头的方法
	// 也就是说 Fetch -> UidFetch，Store -> UidStore，Copy -> UidCopy，Move -> UidMove，Search -> UidSearch
	// uids, err := c.UidSearch(criteria)
	// 关于 Store 方法和 UidStore 方法
	// Store 和 UidStore 方法都是用于在 IMAP 中更新邮件标志的，但它们有一些区别：
	//
	// Store：使用的是消息序列号（message sequence number）来标识邮件。序列号是动态的，每次邮件删除或添加时，序列号可能会改变。序列号从1开始，按邮件在邮箱中的位置进行排序。
	// UidStore：使用的是消息的唯一标识符（UID）来标识邮件。UID 是固定的，不会因为邮件的添加或删除而改变，适合于需要确保唯一标识邮件的操作。
	// 在标记为已读时，使用 UidStore 方法更为安全和可靠，因为它使用邮件的唯一标识符，可以避免由于序列号变化导致的潜在问题。
	if err != nil {
		log.Fatalf("搜索邮件时出现错误：%v \n", err)
	}
	log.Printf("搜索到的邮件 uids: %+v \n", uids)
	if len(uids) == 0 {
		log.Println("没有搜索到邮件")
		return
	}
	log.Printf("搜索到的邮件总共有 %v 封 %+v \n", len(uids), uids)

	// 获取整个消息正文
	// imap.FetchEnvelope：请求获取邮件的信封数据（例如发件人、收件人、主题等元数据）。
	// imap.FetchRFC822：请求获取完整的邮件内容，包括所有头部和正文。
	items := []imap.FetchItem{imap.FetchFlags, imap.FetchEnvelope, imap.FetchRFC822}

	for i, uidsCount := 0, len(uids); i < uidsCount; i += ReadBatchSize {
		// 清空序列集中的所有邮件编号，以便添加新的邮件编号。每次循环开始时调用此方法，确保序列集中只有当前批次的邮件编号
		seqset.Clear()

		// 添加一批邮件到序列集中
		if i+ReadBatchSize < uidsCount {
			seqset.AddNum(uids[i : i+ReadBatchSize]...) // 添加指定范围内的邮件编号
		} else {
			seqset.AddNum(uids[i:]...) // 添加剩余的邮件编号
		}

		// 获取邮件内容 Start
		messages := make(chan *imap.Message, ReadBatchSize) // 创建一个通道，用于接收邮件消息
		fetchDone := make(chan error, 1)                    // 创建一个通道，用于接收错误消息
		go func() {
			// Fetch方法用于从服务器获取邮件数据，这里请求了邮件的信封和完整内容
			fetchDone <- c.Fetch(seqset, items, messages)
		}()
		log.Println("开始读取邮件内容")
		for msg := range messages {
			readEveryMsg(msg)
		}
		if err := <-fetchDone; err != nil {
			log.Fatalf("获取邮件信息出现错误：%v \n", err)
		}
		// 获取邮件内容 End

		// 给邮件打标记 Start
		item := imap.FormatFlagsOp(imap.AddFlags, true) // 标记为已读
		// item := imap.FormatFlagsOp(imap.RemoveFlags, true) // 标记为未读
		flags := []interface{}{imap.SeenFlag}
		log.Printf("即将给这些邮件 [%s] 打标记 \n", seqset)
		if err := c.Store(seqset, item, flags, nil); err != nil {
			log.Fatalf("给邮件打标记失败：%v \n", err)
		}
		// 给邮件打标记 End

		time.Sleep(time.Second * 10) // 休眠10秒
	}

	log.Println("读取了所有邮件，完毕！")

}

// document link: https://github.com/emersion/go-imap/wiki/Fetching-messages
func readEveryMsg(msg *imap.Message) {
	log.Printf("每一封邮件的消息序列号 %+v \n", msg.SeqNum)
	log.Println("-------------------------")
	// 获取邮件正文
	r := msg.GetBody(&imap.BodySectionName{})
	if r == nil {
		log.Fatal("服务器没有返回消息内容")
	}

	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatalf("邮件读取时出现错误： %v \n", err)
	}
	if date, err := mr.Header.Date(); err == nil {
		log.Println("收件时间 Date:", date)
	}
	if from, err := mr.Header.AddressList("From"); err == nil {
		log.Println("发件人 From:", from)
	}
	if to, err := mr.Header.AddressList("To"); err == nil {
		log.Println("收件人 To:", to)
	}
	if subject, err := mr.Header.Subject(); err == nil {
		log.Println("邮件主题 Subject:", subject)
	}
	log.Printf("抄送 Cc: %+v \n", msg.Envelope.Cc)

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("读取邮件内容时出现错误：%v \n", err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// 这是消息的文本（可以是纯文本或 HTML）
			contentType := h.Get("Content-Type")
			b, _ := ioutil.ReadAll(p.Body)
			if strings.HasPrefix(contentType, "text/plain") {
				log.Printf("得到正文 -> TEXT: %v \n", string(b))
			} else if strings.HasPrefix(contentType, "text/html") {
				log.Printf("得到正文 -> HTML: %v \n", len(b))
			}
			break
		case *mail.AttachmentHeader:
			// 这是一个附件
			filename, _ := h.Filename()
			log.Printf("得到附件: %v \n", filename)
			break
		}
	}

	log.Println("一封邮件读取完毕")
	log.Printf("------------------------- \n\n")
}
```

## 值得一提

需要说明一下的是，上面代码中，我们给邮件标记已读时，采用的是 `Store` 方法，其实也可以使用 `UidStore` 方法，那么这两个方法有什么区别呢？

### 关于 Store 方法和 UidStore 方法

`Store` 和 `UidStore` 方法都是用于在 IMAP 中更新邮件标志的，但它们有一些区别：

**Store**：使用的是消息序列号（message sequence number）来标识邮件。序列号是动态的，每次邮件删除或添加时，序列号可能会改变。序列号从1开始，按邮件在邮箱中的位置进行排序。

**UidStore**：使用的是消息的唯一标识符（UID）来标识邮件。UID 是固定的，不会因为邮件的添加或删除而改变，适合于需要确保唯一标识邮件的操作。

在标记为已读时，使用 `UidStore` 方法更为安全和可靠，因为它使用邮件的唯一标识符，可以避免由于序列号变化导致的潜在问题。但是经过我的测试，发现使用 `Store` 方法也没啥太大的问题，但是**使用的时候一定要配套使用**，也就是说，**要是你使用了 `Uid` 开头的方法时，很多方法你都需要改成 `Uid` 开头的方法才能有效使用**，比如： Fetch -> UidFetch，Store -> UidStore，Copy -> UidCopy，Move -> UidMove，Search -> UidSearch。否则，可能会有一些意料之外的事情发生。这是我看文档以及自己摸索得出来的结论，如果你觉得我理解的不对，也可以予以纠正。

好了，聊到这里基本上就结束了。本文主要还是以代码为主，毕竟没有什么比几行代码来得干脆了。不过，可不要将上面的代码直接放到项目中跑呀，放到自己的项目中还是需要稍作调整的，上面代码只是为了方便我在本地调试，因此才有大批量的 log 输出。

如果刚好你也有类似的需求，希望这篇文章可以帮得到你。
