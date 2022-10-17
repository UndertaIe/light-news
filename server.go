package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/UndertaIe/go-eden/utils"
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

var serverSettings ServerSetting

var (
	elasticClient *elasticsearch.Client
	DB            *gorm.DB
)

func RunServer() {
	Init()
	handlers := Handlers()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(serverSettings.HttpPort),
		Handler:        handlers,
		ReadTimeout:    serverSettings.ReadTimeout * time.Second,
		WriteTimeout:   serverSettings.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	utils.ListenAndServe(s)
}

func Init() {
	InitElasticClient()
	InitDBEngine()
}

func InitElasticClient() {
	var err error
	elasticClient, err = NewElasticClient(elasticSettings)
	if err != nil {
		log.Fatalf("InitElasticClient error: %s\n", err)
	}
}

func InitDBEngine() {
	var err error
	DB, err = NewDBEngine(dbSettings)
	if err != nil {
		log.Fatalf("InitDBEngine error: %s\n", err)
	}
}
