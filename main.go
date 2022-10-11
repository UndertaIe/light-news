package main

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var Path = flag.String("config", "config", "配置文件目录")

var svp *viper.Viper // server viper
var rvp *viper.Viper // rule viper

var ss ServerSetting

func main() {
	InitSettings()

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

func InitSettings() {
	var err error
	err = svp.UnmarshalKey("MySQL", &dbSettings)
	if err != nil {
		log.Fatalln("svp.UnmarshalKey err: ", err)
	}
	err = svp.UnmarshalKey("Server", &ss)
	if err != nil {
		log.Fatalln("svp.UnmarshalKey err: ", err)
	}
}
