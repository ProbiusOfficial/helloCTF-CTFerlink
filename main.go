package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
)

type Friend struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Avatar      string       `json:"avatar"`
	URL         string       `json:"url"`
	RSS         string       `json:"rss"`
	Feed        *gofeed.Feed `json:"-"` // RSS Content
}

type FriendsData struct {
	Friends []*Friend `json:"friends"`
}

type Post struct {
	Title     string
	Author    string
	Date      string // YYYY-MM-DD
	Content   string // 保留前 200 字
	PostURL   string // 文章链接
	AuthorURL string // 作者链接
}

func main() {
	// 读取有哪些朋友
	friends := Read(Config.InputPath)

	// 获取朋友的RSS内容
	for _, friend := range friends.Friends {
		err := friend.FetchFeed()
		if err != nil {
			log.Printf("Error fetching feed for %s with %s : %v", friend.Name, friend.RSS, err)
		}
	}

	// 从 RSS 中解析文章
	posts := GetPosts(friends)
	for _, post := range posts {
		fmt.Printf("标题：%s\n", post.Title)
		fmt.Printf("作者：%s\n", post.Author)
		fmt.Printf("日期：%s\n", post.Date)
		fmt.Printf("内容：%s\n", post.Content)
		fmt.Printf("文章链接：%s\n", post.PostURL)
		fmt.Printf("作者链接：%s\n", post.AuthorURL)
		fmt.Println("---------------------")
	}

	// 渲染 markdown 和 json
	markdownBytes := RenderMarkdown(friends.Friends, posts)
	jsonBytes := RenderJson(friends.Friends, posts)
	//fmt.Println(string(markdownBytes))

	// 写入文件
	err := Write(markdownBytes, Config.OutputMarkdownPath)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
	fmt.Println("Markdown content written to file successfully!")

	err = Write(jsonBytes, Config.OutputJsonPath)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Println("Json content written to file successfully!")
}
