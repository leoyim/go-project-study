package goo

import (
	"net/http"
)

// 定义请求handler供goo使用
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// 实现ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc
}

//
