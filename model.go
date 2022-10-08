package main

import "fmt"

type NewsType string

const (
	HTML NewsType = "html"
	JSON NewsType = "json"
)

func (t NewsType) IsHTML() bool {
	return t == HTML
}

func (t NewsType) IsJSON() bool {
	return t == JSON
}

type Rule struct {
	Type        NewsType
	Name        string
	ListUrl     string
	Title       string
	Rank        string
	Author      string
	PublishTime string
	Abstract    string
	Cron        string
}

func (r Rule) Job() func() {
	return func() {
		fmt.Println("job init")
	}
}

func (r Rule) Key() string {
	return r.ListUrl
}

type NewsModel struct {
	NewsUrl     string `json:"news_url" gorm:"column:news_url"`
	Title       string `json:"title" gorm:"column:title"`
	Rank        string `json:"rank" gorm:"column:rank"`
	Author      string `json:"author" gorm:"column:author"`
	Abstract    string `json:"abstract" gorm:"column:abstract"`
	PublishTime string `json:"publish_time" gorm:"column:publish_time"`
	ListUrl     string `json:"list_url" gorm:"column:list_url"`
	PageUrl     string `json:"page_url" gorm:"column:page_url"`
}

func (NewsModel) TableName() string {
	return "news_model"
}

func (m NewsModel) Create() error {
	return nil
}
