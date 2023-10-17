// 打包
// go build -o mac-spider-blog-title spider_blog_title.go
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o spider-blog-title spider_blog_title.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	BlogBasePath = "https://pudongping.github.io"
)

const header = `
---
title: 文章目录
date: 2022-11-21 02:24:11
type: "menu"
layout: "menu"
---

# 文章目录


`

// SpiderMyHexoBlog 爬取文章标题和文章地址
func SpiderMyHexoBlog(webSite, baseUrl string) []string {
	fmt.Printf("start at %s\n", webSite)
	res, err := http.Get(webSite)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// 查看 html 文档
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var contents = make([]string, 0, 12)
	// 查找
	doc.Find("#articles .article-row .article").Each(func(i int, s *goquery.Selection) {
		item := s.Find(".card")
		url, _ := item.Find("a").Attr("href")
		url = baseUrl + url
		title := item.Find(".card-title").Text()

		category := item.Find(".post-category").Text()

		var tags []string
		item.Find(".article-tags a").Each(func(i int, s *goquery.Selection) {
			tag := s.Find(".chip").Text()
			tags = append(tags, tag)
		})
		tagName := strings.Join(tags, ",")

		content := fmt.Sprintf("%s== [%s](%s) 标签：%s \n", category, title, url, tagName)
		contents = append(contents, content)

	})

	return contents
}

// HandleContents 处理抓取后的内容，将内容处理成 markdown 格式
func HandleContents(contents []string) string {
	group := make(map[string][]string)
	for _, content := range contents {
		sp := strings.Split(content, "==")
		category := sp[0]
		title := fmt.Sprintf("- %s", sp[1])
		group[category] = append(group[category], title)
	}

	var tmp string
	for h1, titles := range group {
		tmp += "## " + h1 + "\n"
		for _, t := range titles {
			tmp += t
		}
		tmp += "\n"
	}

	return tmp
}

// BuildUrls 构建需要爬取的链接地址
func BuildUrls(maxPage int, baseUrl string) []string {
	urls := make([]string, 0, maxPage)
	// 分页中没有第一页（首页即为第一页）即不存在 `https://pudongping.github.io/page/1/`
	urls = append(urls, baseUrl)
	for i := 2; i <= maxPage; i++ {
		perPageUrl := fmt.Sprintf("%s/page/%d/", baseUrl, i)
		urls = append(urls, perPageUrl)
	}
	return urls
}

// FetchMaxPageNum 获取最大页码数
func FetchMaxPageNum(baseUrl string) int {
	fmt.Printf("fetch max page num %s\n", baseUrl)
	res, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("fetch max page num: status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 获取最大页码数
	pageText := doc.Find(".paging .center-align").Text()
	pageNums := strings.Split(pageText, "/")
	maxPageStr := strings.TrimSpace(pageNums[1])
	maxPage, _ := strconv.Atoi(maxPageStr)
	fmt.Printf("maxPage is %v\n", maxPage)
	return maxPage
}

func main() {
	now := time.Now()
	maxPage := FetchMaxPageNum(BlogBasePath)

	resChannel := make(chan []string, 2)

	urls := BuildUrls(maxPage, BlogBasePath)

	for _, url := range urls {
		url := url
		go func() {
			resChannel <- SpiderMyHexoBlog(url, BlogBasePath)
		}()
	}

	var results = make([]string, 0, len(urls)*12)
	for range urls {
		results = append(results, <-resChannel...)
	}

	output := HandleContents(results)
	ioutil.WriteFile("./blog-menu.md", []byte(header+output), 0777)

	fmt.Printf("程序执行耗时 %s \n", time.Since(now))
}
