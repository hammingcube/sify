package blog

import (
	"github.com/labstack/echo"
	"net/http"
)

func Router() *echo.Echo {
	//------
	// Blog
	//------
	blog := echo.New()
	blog.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blog")
	})
	return blog
}
