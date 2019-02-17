package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
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

var colorMap = map[int]string{
	1:  "#ef9a9a",
	2:  "#f48fb1",
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
	gob.Register(&telestrationsLib.Player{})
	e.Renderer = renderer
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Static("/", "../templates/")
	e.GET("/game", drawPage)
	e.GET("/", index)
	e.GET("/login", login)
	e.POST("/addUser", addUser)
	e.GET("/getPlayers", getPlayers)
	e.GET("/getTime", getTime)
	e.GET("/drawPage", drawPage)
	e.POST("/startGame", startGame)
	e.Start(":1234")
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func drawPage(c echo.Context) error {
	sess, _ := session.Get("session", c)
	id, _ := sess.Values["id"]
	var aid *telestrationsLib.Player
	aid = id.(*telestrationsLib.Player)
	oddeven := len(gameManager.AllPlayers) % 2
	if oddeven == 1 {
		// odd page
	} else {
		// even page
	}
	data := make(map[string]interface{})
	data["id"] = strconv.Itoa(aid.ID)
	return c.Render(http.StatusOK, "game.html", data)
}

func index(c echo.Context) error {
	fmt.Println("index")
	sess, _ := session.Get("session", c)
	id, _ := sess.Values["id"]
	fmt.Println(gameManager.AllPlayers)
	if id == nil {
		// fmt.Println(id)
		return c.Redirect(http.StatusFound, "/login")
	}
	var aid *telestrationsLib.Player
	//json.Unmarshal(id.([]byte), aid)
	aid = id.(*telestrationsLib.Player)
	fmt.Println(gameManager.GMPlayerExists(aid.ID))
	if !gameManager.GMPlayerExists(aid.ID) {
		id := telestrationsLib.AddUser(aid.Name, "test")
		gameManager.GMAddPlayer(id, aid.Name)
	}
	fmt.Println(gameManager.AllPlayers)
	data := make(map[string]interface{})
	data["colorMap"] = colorMap
	data["players"] = gameManager.GMGetPlayersAsArray()
	data["id"] = strconv.Itoa(aid.ID)
	//fmt.Println(data["players"])
	return c.Render(http.StatusOK, "start.html", data)
}

func login(c echo.Context) error {
	fmt.Println("login")
	sess, _ := session.Get("session", c)
	id, _ := sess.Values["id"]
	if id != nil {
		c.Redirect(http.StatusFound, "/")
	}
	return c.Render(http.StatusOK, "login.html", nil)
}

func addUser(c echo.Context) error {
	fmt.Println("adduser")
	name := c.FormValue("name")
	id := telestrationsLib.AddUser(name, "test")
	gameManager.GMAddPlayer(id, name)
	sess, _ := session.Get("session", c)
	sess.Values["id"] = *gameManager.GMGetPlayer(id) //strconv.Itoa(id)
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

func getTime(c echo.Context) error {
	return c.JSON(http.StatusOK, gameManager.TimeLeft)
}

func startGame(c echo.Context) error {
	gameManager.GMStartGame()
	return c.NoContent(http.StatusOK)
}

func checkForStart(c echo.Context) error {
	if gameManager.State == telestrationsLib.StateProgress {
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusForbidden)
	}
}
