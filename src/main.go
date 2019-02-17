package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
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
var session *sessions.CookieStore

var colorMap = map[int]string{
	1:  "#ef9a9a",
	2:  "#f48fb",
	3:  "#ce93d8",
	4:  "#b39ddb",
	5:  "#9fa8da",
	6:  "#90caf9",
	7:  "#80deea",
	8:  "#80cbc4",
	9:  "#a5d6a7",
	10: "#fff59d",
	11: "#ffcc80",
	12: "#bcaaa4"}

func main() {
	telestrationsLib.DeleteAllUsers()
	go telestrationsLib.ResetInc()
	gameManager = *telestrationsLib.CreateGameState()
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../templates/*.html")),
	}
	session = sessions.NewCookieStore([]byte("secret"))
	e.Renderer = renderer
	e.Use(middleware.Logger())
	e.Static("/", "../templates/")
	e.GET("/game", drawPage)
	e.GET("/", index)
	e.GET("/login", login)
	e.POST("/addUser", addUser)
	e.GET("/getPlayers", getPlayers)
	e.Start(":1234")
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func drawPage(c echo.Context) error {
	val, err := c.Cookie("id")
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["id"] = val.Value
	return c.Render(http.StatusOK, "game.html", data)
}

func index(c echo.Context) error {
	fmt.Println("index")
	sess, _ := session.Get(c.Request(), "session")
	// sess.Options.MaxAge = -1
	// return nil
	id, _ := sess.Values["id"]
	if id == nil {
		fmt.Println(id)
		return c.Redirect(http.StatusFound, "/login")
	}
	var aid telestrationsLib.Player
	json.Unmarshal(id.([]byte), aid)
	if !gameManager.GMPlayerExists(aid.ID) {
		gameManager.GMAddPlayer(aid.ID, aid.Name)
	}
	data := make(map[string]interface{})
	data["colorMap"] = colorMap
	data["players"] = gameManager.GMGetPlayersAsArray()
	data["id"] = strconv.Itoa(aid.ID)
	//fmt.Println(data["players"])
	return c.Render(http.StatusOK, "start.html", data)
}

func login(c echo.Context) error {
	fmt.Println("login")
	sess, _ := session.Get(c.Request(), "session")
	id, _ := sess.Values["id"]
	if id != nil {
		c.Redirect(http.StatusFound, "/")
	}
	return c.Render(http.StatusOK, "login.html", nil)
}

func addUser(c echo.Context) error {
	fmt.Println("adduser")
	name := c.FormValue("name")
	id := telestrationsLib.AddUser(name)
	gameManager.GMAddPlayer(id, name)
	sess, _ := session.Get(c.Request(), "session")
	sess.Values["id"], _ = json.Marshal(*gameManager.GMGetPlayer(id)) //strconv.Itoa(id)
	sess.Options.MaxAge = 99999
	fmt.Println(sess.Values["id"])
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/")
}

func getPlayers(c echo.Context) error {
	players := gameManager.GMGetPlayersAsArray()
	//fmt.Println(players)
	return c.JSON(http.StatusOK, players)
}
