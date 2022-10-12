package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/UndertaIe/go-eden/str"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"golang.org/x/net/html/charset"
)

type ParserType string

const (
	CSSParserT  ParserType = "css"
	XMLParserT  ParserType = "xml"
	JSONParserT ParserType = "json"
)

type Parser interface {
	Parse(*Rule) ([]*NewsModel, error)
}

func SelectParser(r ParserType) (Parser, error) {
	switch r {
	case CSSParserT:
		return &CSSParser{}, nil
	case XMLParserT:
		return &CSSParser{}, nil
	case JSONParserT:
		return &JSONParser{}, nil
	}
	return nil, ErrParserNotFound
}

type CSSParser struct{}

func (cp *CSSParser) Parse(r *Rule) ([]*NewsModel, error) {
	var models []*NewsModel
	resp, err := http.Get(r.ListUrl)
	if err != nil {
		return nil, err
	}
	rd, err := charset.NewReader(resp.Body, "utf8")
	if err != nil {
		return nil, err
	}
	docu, err := goquery.NewDocumentFromReader(rd)

	if err != nil {
		return nil, err
	}
	docu.Find(r.Item).Each(func(i int, s *goquery.Selection) {
		model := new(NewsModel)
		model.NewsUrl = trim(s.Find(r.NewsUrl).AttrOr("href", ""))
		model.Title = trim(s.Find(r.Title).Text())
		model.Rank = cast.ToInt16(trim(s.Find(r.Rank).Text()))
		if model.Rank == 0 {
			model.Rank = cast.ToInt16(i) + 1
		}
		model.Author = trim(s.Find(r.Author).Text())
		model.Abstract = trim(s.Find(r.Abstract).Text())
		model.PublishTime = time.Now() // s.Find(r.PublishTime).Text()
		model.IsHot = isHot(s.Find(r.IsHot).Text())
		model.ImgUrl = trim(s.Find(r.ImgUrl).AttrOr("src", ""))
		model.ListUrl = r.ListUrl
		model.RawListUrl = r.RawListUrl
		model.DataSource = r.DataSource
		models = append(models, model)
	})
	return models, nil
}

type XMLParser struct{}

func (cp *XMLParser) Parse(r *Rule) ([]*NewsModel, error) {
	return nil, nil
}

type JSONParser struct{}

func (cp *JSONParser) Parse(r *Rule) ([]*NewsModel, error) {
	var models []*NewsModel
	resp, err := http.Get(r.ListUrl)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tree := gjson.ParseBytes(jsonp2json(bytes))
	items := tree.Get(r.Item)
	items.ForEach(func(key, item gjson.Result) bool {
		model := new(NewsModel)
		model.NewsUrl = trim(item.Get(r.NewsUrl).String())
		model.Title = trim(item.Get(r.Title).String())
		model.Rank = cast.ToInt16(trim(item.Get(r.Rank).String()))
		if model.Rank == 0 {
			model.Rank = cast.ToInt16(key.String()) + 1
		}
		model.Author = trim(item.Get(r.Author).String())
		model.Abstract = trim(item.Get(r.Abstract).String())
		if str.Slen(model.Abstract) > 512 {
			model.Abstract, _ = str.SubString(model.Abstract, 0, 512)
		}
		model.PublishTime = time.Now()
		model.IsHot = isHot(item.Get(r.IsHot).String())
		model.ImgUrl = trim(item.Get(r.ImgUrl).String())
		model.ListUrl = r.ListUrl
		model.RawListUrl = r.RawListUrl
		if model.NewsUrl != "" || model.Title != "" {
			models = append(models, model)
		}
		return true
	})
	if len(models) == 0 {
		return nil, ErrNoNewsModelParsed
	}
	return models, nil
}
