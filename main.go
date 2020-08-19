package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello world!")
	})

	_ = engine.Run(":8080")
}
