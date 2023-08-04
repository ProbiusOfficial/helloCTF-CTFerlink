package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"
)

type RenderStruct struct {
	Friends []*Friend `json:"friends"`
	Posts   []Post    `json:"posts"`
}

// RenderMarkdown 将朋友信息和最近的文章渲染成Markdown格式
func RenderMarkdown(friends []*Friend, posts []Post) []byte {
	markdownTemplate, err := os.ReadFile(Config.TemplatePath)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	tmpl, err := template.New("markdown").Parse(string(markdownTemplate))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	var buffer bytes.Buffer
	data := RenderStruct{
		Friends: friends,
		Posts:   posts,
	}

	if err := tmpl.Execute(&buffer, data); err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	return buffer.Bytes()
}

func RenderJson(friends []*Friend, posts []Post) []byte {
	data := RenderStruct{
		Friends: friends,
		Posts:   posts,
	}

	marshal, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	return marshal
}
