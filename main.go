package main

import (
	"github.com/labstack/echo/v4"
	"github.com/typomedia/echodemo/app/handler"
	"github.com/typomedia/echodemo/app/handler/api"
	"html/template"
	"io"
	"log"
)

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(writer io.Writer, name string, data interface{}, context echo.Context) error {
	return r.templates.ExecuteTemplate(writer, name, data)
}

func main() {
	app := echo.New()

	views, err := template.ParseGlob("app/views/*.html")
	if err != nil {
		log.Fatal(err)
	}

	app.Renderer = &Renderer{
		templates: template.Must(views, err),
	}

	app.GET("/", handler.Index)
	app.GET("/api/html/search", api.Search)
	app.Static("/", "public")

	app.Logger.Fatal(app.Start(":4000"))
}
