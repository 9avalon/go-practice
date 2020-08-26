package main

import (
	"gee"
)

func main() {
	r := gee.New()

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func (c *gee.Context) {
			c.String(200, "hello! %s", "miguel")
		})
	}

	_ = r.Run(":8080")
}
