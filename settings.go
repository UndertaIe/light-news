package main

import "time"

// run init in main
var (
	dbSettings      *DBSetting
	elasticSettings *ElasticSetting
)

type ServerSetting struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DBSetting struct {
	DBType    string
	UserName  string
	Password  string
	Host      string
	DBName    string
	Charset   string
	ParseTime bool
}

type RedisSetting struct {
	Host              string
	Db                int
	Password          string
	DefaultExpireTime int
}

type ElasticSetting struct {
	Hosts        []string
	ServiceToken string
	Index        string
}
