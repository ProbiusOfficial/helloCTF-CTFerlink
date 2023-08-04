package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"sort"
	"strings"
)

func GetPosts(friends FriendsData) []Post {
	var posts []Post

	for _, friend := range friends.Friends {
		//fmt.Println(friend.Feed)
		if friend.Feed == nil {
			// 跳过没有获取到RSS内容的朋友
			log.Printf("跳过没有获取到RSS内容的朋友: %s", friend.Name)
			continue
		}

		for _, item := range friend.Feed.Items {
			// 简化示例，假设item.Published为格式为"2006-01-02T15:04:05Z"的时间字符串
			publishedTime := item.PublishedParsed.Format("2006-01-02")
			posts = append(posts, Post{
				Title:  item.Title,
				Author: friend.Name,
				Date:   publishedTime,
				Content: filterPosts(getShortContent( /*item.Content*/
					func() string {
						if item.Content != "" {
							return item.Content
						}
						return item.Description
					}(), Config.ContentLength)),
				PostURL:   item.Link,
				AuthorURL: friend.URL,
			})
		}
	}

	// 按照日期降序排序文章
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	// 获取最近的 Config.NumOutputPost 篇文章
	var recentPosts []Post
	if len(posts) <= Config.NumOutputPost {
		recentPosts = posts
	} else {
		recentPosts = posts[:Config.NumOutputPost]
	}

	return recentPosts
}

func filterPosts(content string) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n", "")
	content = strings.ReplaceAll(content, "\r", "")
	content = strings.ReplaceAll(content, "\t", "")
	return content
}

// getShortContent 根据中文字符数截取文本，并去除收尾空白字符
func getShortContent(content string, length int) string {

	// 使用golang.org/x/net/html解析HTML标签
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		// 解析失败，则直接截取前length个字符
		fmt.Println("解析失败，则直接截取前", length, "个字符")
		return truncateText(content, length)
	}

	// 遍历HTML节点，将文字内容提取出来
	var buf strings.Builder
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractText(c)
		}
	}
	extractText(doc)

	// 提取的文字内容进行截取
	truncatedText := truncateText(buf.String(), length)

	return truncatedText
}

// 根据中文字符数截取文本
func truncateText(text string, length int) string {
	var buf strings.Builder
	var count int
	for _, r := range text {
		buf.WriteRune(r)
		count++
		if count >= length {
			break
		}
	}
	return buf.String()
}
