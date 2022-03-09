package main

import (
	"net/http"
	"web/web"
)

func main() {
	r := web.New()
	r.GET("/", func(c *web.Context) {
		c.String(http.StatusOK, "Hello web\n")
	})
	r.GET("/panic", func(c *web.Context) {
		names := []string{"web"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
