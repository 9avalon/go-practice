package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Write http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	Params map[string]string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Write: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) setHeader(key string, value string) {
	c.Write.Header().Set(key, value)
}

func (c *Context) Status(status int) {
	c.StatusCode = status
	c.Write.WriteHeader(status)
}

func (c *Context) String(status int, format string, values ...interface{}) {
	c.setHeader("Context", "text/plain")
	c.Status(status)
	_, _ = c.Write.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) JSON(status int, obj interface{}) {
	c.setHeader("Content-Type", "text/plain")
	c.Status(status)

	encoder := json.NewEncoder(c.Write)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Write, err.Error(), 500)
	}
}

