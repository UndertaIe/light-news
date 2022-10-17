package main

import (
	"log"
	"time"
)

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
	DataType   NewsType
	Parser     ParserType
	StoreType  []StoreType
	Cron       string
	DataSource string
	ListUrl    string
	RawListUrl string

	Item        string
	NewsUrl     string
	Title       string
	Rank        string
	Author      string
	Abstract    string
	PublishTime string
	IsHot       string
	ImgUrl      string
}

func (r Rule) Job() func() {
	return func() {
		log.Println("job init")
	}
}

func (r Rule) Key() string {
	return r.ListUrl
}

type NewsModel struct {
	NewsUrl     string    `json:"news_url" gorm:"column:news_url"` // required
	Title       string    `json:"title" gorm:"column:title"`       //  required
	Rank        int16     `json:"rank" gorm:"column:rank"`
	Author      string    `json:"author" gorm:"column:author"`
	Abstract    string    `json:"abstract" gorm:"column:abstract"`
	PublishTime time.Time `json:"publish_time" gorm:"column:publish_time"`
	IsHot       bool      `json:"is_hot" gorm:"column:is_hot"`
	ImgUrl      string    `json:"img_url" gorm:"column:img_url"`
	ListUrl     string    `json:"list_url" gorm:"column:list_url"`       // required
	RawListUrl  string    `json:"page_url" gorm:"column:raw_url"`        // required
	DataSource  string    `json:"data_source" gorm:"column:data_source"` // required
}

func (NewsModel) TableName() string {
	return "news_model"
}

func (m NewsModel) Create() error {
	return nil
}
