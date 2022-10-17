package main

import (
	"bytes"
	"context"
	"log"
	"net/url"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/goccy/go-json"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type StoreType string

const (
	DummyStoreT  StoreType = "dummy"
	MySQLStorerT StoreType = "mysql"
	MongoStorerT StoreType = "mongo"
	KafkaStorerT StoreType = "kafka"
	ESStoreT     StoreType = "es"
)

type Storer interface {
	Store(*NewsModel) error
}

func SelectStorer(t StoreType) (Storer, error) {
	switch t {
	case DummyStoreT:
		return NewDummyStorer(), nil
	case MySQLStorerT:
		return NewMySQLStorer(dbSettings), nil
	case ESStoreT:
		return NewElasticStorerr(elasticSettings), nil
	case MongoStorerT:
		return nil, ErrStorerNotFound
	case KafkaStorerT:
		return nil, ErrStorerNotFound
	}
	return nil, ErrStorerNotFound
}

type MySQLStorer struct {
	db *gorm.DB
}

func NewMySQLStorer(s *DBSetting) *MySQLStorer {
	db, err := NewDBEngine(s)
	if err != nil {
		log.Fatalf("NewDBEngine(s)[s: %v] err: %v\n", s, err)
	}
	return &MySQLStorer{db: db}
}

func (ms *MySQLStorer) Store(m *NewsModel) error {
	var t NewsModel
	var err error
	db := ms.db.WithContext(context.Background())
	var n int64
	err = db.Model(&t).Where("news_url = ?", m.NewsUrl).Count(&n).Error
	if err != nil {
		return err
	}
	if n == 0 {
		return db.Create(m).Error
	}
	return db.Model(&t).Where("news_url = ? and rank > ?", m.NewsUrl, m.Rank).Update("rank", m.Rank).Error
}

type DummyStorer struct {
}

func NewDummyStorer() *DummyStorer {
	return &DummyStorer{}
}

func (ds *DummyStorer) Store(m *NewsModel) error {
	return nil
}

type ElasticStorer struct {
	cli       *elasticsearch.Client
	eSettings *ElasticSetting
}

func NewElasticStorerr(s *ElasticSetting) *ElasticStorer {
	cli, err := NewElasticClient(s)
	if err != nil {
		panic(err)
	}
	return &ElasticStorer{cli: cli, eSettings: s}
}

var updateTmpl = `{
	"script": {
		"source": "if (ctx._source.rank > params.rank ) {ctx._source.rank = params.rank}",
		"params": {"rank": <RANK>}
	}
}`

func (es *ElasticStorer) Store(m *NewsModel) error {
	payload, _ := json.Marshal(m)
	r := bytes.NewReader(payload)
	Index := es.eSettings.Index
	escapedUrl := url.PathEscape(m.NewsUrl)
	resp, err := es.cli.Get(Index, escapedUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == 404 {
		_, err := es.cli.Create(Index, escapedUrl, r)
		if err != nil {
			return err
		}
	}
	_, err = es.cli.Update(
		Index,
		escapedUrl,
		strings.NewReader(strings.ReplaceAll(updateTmpl, "<RANK>", cast.ToString(m.Rank))),
		es.cli.Update.WithPretty(),
	)
	return err
}
