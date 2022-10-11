package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/UndertaIe/go-eden/utils"
)

func RunServer() {
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
