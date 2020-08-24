package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()

	engine.GET("/json", func(c *gee.Context) {
		reqMap := make(gee.H)
		for value := range c.Req.URL.Query() {
			reqMap[value] = c.Req.URL.Query()[value][0]
		}

		c.JSON(200, reqMap)
	})

	engine.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "%s good", "very")
	})

	engine.GET("/like/:name", func(c *gee.Context) {
		c.String(http.StatusOK, ":name is %s", c.Param("name"))
	})

	_ = engine.Run(":8080")
}
