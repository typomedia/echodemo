package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	search := c.QueryParam("search")

	return c.Render(http.StatusOK, "index.html", echo.Map{
		"Search": search,
	})
}
