package gee

type Router struct {
	handlers map[string]HandleFunc
}

func newRouter() * Router {
	return &Router{handlers: make(map[string]HandleFunc)}
}

func (router *Router) addRouter(method string, pattern string, handle HandleFunc) {
	router.handlers[method + pattern] = handle
}

func (router *Router) handle(context *Context) {
	key := context.Method + context.Path
	handle := router.handlers[key]
	if handle != nil {
		handle(context)
	} else {
		context.String(404, "[%s]method not found", context.Path)
	}
}
