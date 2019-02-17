package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type TimeStatus struct {
	TimeLeft int `json:"time_left"`
}

type PlayerStatus struct {
	Players []string `json:"players"`
}

type TemplateRenderer struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../templates/*.html")),
	}
	e.Renderer = renderer
	e.Static("/", "..")
	e.GET("/game", drawPage)
	e.Start(":1234")
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func drawPage(c echo.Context) error {
	data := make(map[string]interface{})
	data["id"] = 11
	return c.Render(http.StatusOK, "game.html", data)
}
