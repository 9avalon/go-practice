package gee

import (
	"net/http"
)

type HandleFunc func(context *Context)

type Engine struct {
	router *Router
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	context := newContext(writer, request)
	engine.router.handle(context)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) GET(pattern string, handleFunc HandleFunc) {
	engine.router.addRouter("GET", pattern, handleFunc)
}

func (engine *Engine) POST(pattern string, handleFunc HandleFunc) {
	engine.router.addRouter("POST", pattern, handleFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
