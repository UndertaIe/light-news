package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/UndertaIe/go-eden/app"
	"github.com/UndertaIe/go-eden/errcode"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		v1.GET("/search/:key", SearchFunc)
		v1.GET("/hot", HotNewsFunc)
		v1.GET("/news", NewsFunc)
	}

	return r
}

// 关键词搜索，按照时间排序
//! 数据冗余
func SearchFunc(c *gin.Context) {
	resp := app.NewResponse(c)
	keywd := c.Param("key")
	if keywd == "" {
		resp.ToError(errcode.InvalidParams)
		return
	}
	pager := app.NewPager(c)

	reader := QueryMatchReader(keywd)
	esResp, err := elasticClient.Search(
		elasticClient.Search.WithIndex(elasticSettings.Index),
		elasticClient.Search.WithBody(reader),
		elasticClient.Search.WithSort("publish_time:desc", "_score:desc"),
		elasticClient.Search.WithFrom(pager.Offset()),
		elasticClient.Search.WithSize(pager.Limit()),
		elasticClient.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("[ERROR] elasticClient.Search err: %v \n", err)
		resp.ToError(errcode.ErrorService.WithDetails(err.Error()))
	}
	arrs, err := unpack(esResp)
	if err != nil {
		log.Printf("[ERROR] unpack err: %v \n", err)
		resp.ToError(errcode.ErrorService.WithDetails(err.Error()))
	}
	resp.ToList(arrs, pager)
}

func HotNewsFunc(c *gin.Context) {
	// TODO:
}

func NewsFunc(c *gin.Context) {
	// TODO:
}

// 多字段查询
// {"query": { "multi_match" : { "query": "<KEY>","fields": [ "title", "abstract" ]}}}
func QueryMatchReader(keywd string) io.Reader {
	var bb = new(bytes.Buffer)
	bb.WriteString(`{"query": { "multi_match" : { "query": "`)
	bb.WriteString(keywd)
	bb.WriteString(`", "fields": [ "title", "abstract" ]}}}`)
	return bb
}

func unpack(esResp *esapi.Response) (arrs []any, err error) {
	var raw map[string]any
	if err = json.NewDecoder(esResp.Body).Decode(&raw); err != nil {
		log.Printf("[ERROR] parsing the response body: %s\n", err)
		return nil, err
	}
	hits := raw["hits"].(map[string]any)["hits"].([]any)
	for _, hit := range hits {
		arrs = append(arrs, hit.(map[string]any)["_source"])
	}
	return
}
