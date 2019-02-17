package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zedjones/telestrations/telestrationsLib"
)

type TimeStatus struct {
	TimeLeft int `json:"time_left"`
}

type PlayerStatus struct {
	Players []string `json:"players"`
}

type User struct {
	Name string `json:"name"`
}

type TemplateRenderer struct {
	templates *template.Template
}

var gameManager telestrationsLib.GameState

func main() {
	gameManager = *telestrationsLib.CreateGameState()
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../templates/*.html")),
	}
	e.Renderer = renderer
	e.Use(middleware.Logger())
	e.Static("/", "../templates/")
	e.GET("/game", drawPage)
	e.GET("/", index)
	e.GET("/login", login)
	e.POST("/addUser", addUser)
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

func index(c echo.Context) error {
	_, err := c.Cookie("id")
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	return c.Render(http.StatusOK, "start.html", nil)
}

func login(c echo.Context) error {
	fmt.Println("login")
	_, err := c.Cookie("id")
	if err == nil {
		c.Redirect(http.StatusFound, "/")
	}
	return c.Render(http.StatusOK, "login.html", nil)
}

func addUser(c echo.Context) error {
	fmt.Println("adduser")
	name := c.FormValue("name")
	id := telestrationsLib.AddUser(name)
	gameManager.GMAddPlayer(id, name)
	cookie := new(http.Cookie)
	cookie.Name = "id"
	cookie.Value = strconv.Itoa(id)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}
