package main

import (
	"errors"
)

var (
	ErrParserNotFound         = errors.New("Parser Not Found. Supported Parser: css, xml, json")
	ErrStorerNotFound         = errors.New("Storer Not Found. Supported Storer: mysql")
	ErrGenerateDocument       = errors.New("generate document error")
	ErrModelFieldEmpty        = errors.New("required field is empty")
	ErrModelFieldTitleEmpty   = errors.New("required field Title is empty")
	ErrModelFieldNewsURLEmpty = errors.New("required field NewsUrl is empty")
	ErrNoNewsModelParsed      = errors.New("no NewsModel parsed")
)
