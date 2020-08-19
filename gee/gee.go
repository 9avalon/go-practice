package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + request.URL.Path
	handle := engine.router[key]
	if handle == nil {
		_, _ = fmt.Fprint(writer, "404 not found")
		return
	}

	handle(writer, request)
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handleFunc HandleFunc) {
	key := method + pattern
	engine.router[key] = handleFunc
}

func (engine *Engine) GET(pattern string, handleFunc HandleFunc) {
	engine.addRoute("GET", pattern, handleFunc)
}

func (engine *Engine) POST(pattern string, handleFunc HandleFunc) {
	engine.addRoute("POST", pattern, handleFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
