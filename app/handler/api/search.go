package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/typomedia/echodemo/app/structs"
	"io"
	"net/http"
)

func Search(c echo.Context) error {
	search := c.QueryParam("search")
	url := "https://dummyjson.com/users/search"

	response, err := http.Get(url + "?q=" + search)
	if err != nil {
		log.Error(err)
	}
	defer response.Body.Close()

	// body to bytes
	body, err := io.ReadAll(response.Body)

	var users structs.Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Error(err)
	}

	return c.Render(http.StatusOK, "table.html", echo.Map{
		"Users": users.Users,
	})
}
