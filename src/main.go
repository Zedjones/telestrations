package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/zedjones/telestrations/telestrationsLib"
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

var gameManager telestrationsLib.GameState

func main() {
	/*
		gameManager = *telestrationsLib.CreateGameState()
		e := echo.New()
		renderer := &TemplateRenderer{
			templates: template.Must(template.ParseGlob("../templates/*.html")),
		}
		e.Renderer = renderer
		e.Static("/", "../templates/")
		e.GET("/game", drawPage)
		e.Start(":1234")
	*/
	telestrationsLib.AddUser("testing")
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func drawPage(c echo.Context) error {
	data := make(map[string]interface{})
	data["id"] = 11
	return c.Render(http.StatusOK, "game.html", data)
}

func index(c echo.Context) error {
	_, err := c.Cookie("id")
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	return nil
}

func login(c echo.Context) error {
	_, err := c.Cookie("id")
	if err == nil {
		c.Redirect(http.StatusFound, "/")
	}
	return c.Render(http.StatusOK, "login.html", nil)
}
