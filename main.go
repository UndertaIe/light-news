package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/UndertaIe/go-eden/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var Path = flag.String("config", "config", "配置文件目录")

var svp *viper.Viper // server viper
var rvp *viper.Viper // rule viper

var DB *gorm.DB

func main() {
	InitDatabase()

	RunScheduler()
	
	RunServer()
}

func init() {
	flag.Parse()
	var err error
	svp = viper.New()
	svp.AddConfigPath(*Path)
	svp.SetConfigName("config")
	err = svp.ReadInConfig()
	if err != nil {
		log.Fatalln("svp.ReadInConfig() err: ", err)
	}

	rvp = viper.New()
	rvp.AddConfigPath(*Path)
	rvp.SetConfigName("rule")
	err = rvp.ReadInConfig()
	if err != nil {
		log.Fatalln("rvp.ReadInConfig() err: ", err)
	}
}

func RunScheduler() {
	var rules []Rule
	err := rvp.UnmarshalKey("News", &rules)
	if err != nil {
		log.Fatalln("vp.UnmarshalKey(\"News\", &rules) err: ", err)
	}
	for _, r := range rules {
		DefaultScheduler.AddJob(r)
	}
	DefaultScheduler.Start()
}

func RunServer() {
	var ss ServerSetting
	err := svp.UnmarshalKey("Server", &ss)
	if err != nil {
		log.Fatalln("vp.UnmarshalKey(\"Server\", &s) err: ", err)
	}
	handlers := Handlers()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(ss.HttpPort),
		Handler:        handlers,
		ReadTimeout:    ss.ReadTimeout * time.Second,
		WriteTimeout:   ss.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	utils.ListenAndServe(s)
}

func InitDatabase() {
	var dbSettings DBSetting
	err := svp.UnmarshalKey("MySQL", &dbSettings)
	if err != nil {
		log.Fatalln("UnmarshalKey(\"MySQL\", &rules) err: ", err)
	}
	DB, err = NewDBEngine(&dbSettings)
	if err != nil {
		log.Fatalln("NewDBEngine(*dbSettings) err: ", err)
	}
}
