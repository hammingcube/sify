package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/maddyonline/sify/blog"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Hosts
	hosts := make(map[string]*Host)

	//------
	// Blog
	//------

	blog := blog.Router()
	blog.Use(middleware.Logger())
	blog.Use(middleware.Recover())

	hosts["localhost:1323"] = &Host{blog}

	// Server
	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host()]
		fmt.Printf("HOST: %s\n", req.Host())

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(req, res)
		}

		return
	})
	e.Run(standard.New(":1323"))
}
