package main

import (
	"context"
	"log"

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
		return NewMySQLStorer(&dbSettings), nil
	case MongoStorerT:
		return nil, ErrStorerNotFound
	case KafkaStorerT:
		return nil, ErrStorerNotFound
	case ESStoreT:
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
	return db.Model(&t).Where("news_url = ? and rank < ?", m.NewsUrl, m.Rank).Update("rank", m.Rank).Error
}

type DummyStorer struct {
}

func NewDummyStorer() *DummyStorer {
	return &DummyStorer{}
}

func (ds *DummyStorer) Store(m *NewsModel) error {
	return nil
}
