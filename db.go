package main

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(s *DBSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", s.UserName, s.Password, s.Host, s.DBName, s.Charset, s.ParseTime)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewElasticClient(s *ElasticSetting) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: s.Hosts,
	}
	if s.ServiceToken != "" {
		cfg.ServiceToken = s.ServiceToken
	}
	cli, err := elasticsearch.NewClient(cfg)
	return cli, err
}
