package telestrationsLib

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // need pq for postgresql
)

type User struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	StartWord string `db:"starting_word"`
}

type Guess struct {
	UserID int    `db:"user_id"`
	Word   string `db:"word"`
	Round  int    `db:"round"`
}

type Picture struct {
	UserID int    `db:"user_id"`
	SVG    string `db:"svg"`
	Round  int    `db:"round"`
}

type Settings struct {
	ID         int    `db:"id"`
	TimeLimit  int    `db:"time_limit"`
	Difficulty string `db:"difficulty"`
}

const connStr = "postgres://zedjones:password@localhost/telestrations?sslmode=disable"

const createUserStr = "INSERT INTO \"user\" (name, starting_word) VALUES ($1, $2)"
const getUserStr = "SELECT id, name, starting_word FROM \"user\" WHERE id = $1"
const resetIncStr = "ALTER SEQUENCE user_id_seq RESTART with 1"
const getLastUser = "SELECT id, name, starting_word FROM \"user\" ORDER BY id DESC"
const deleteUsersStr = "DELETE FROM \"user\""

const insertPictureStr = "INSERT INTO picture (user_id, svg, round) VALUES ($1, $2, $3)"
const getPictureStr = "SELECT user_id, svg, round FROM picture " +
	"	WHERE user = $1 and round = $2"

const insertGuessStr = "INSERT INTO guess (user_id, word, round) VALUES ($1, $2, $3)"
const getGuessStr = "SELECT user_id, guess, round FROM picture " +
	"	WHERE user = $1 and round = $2"

const changeSettingsStr = "UPDATE settings SET time_limit = $1, word_difficulty = $2" +
	"	WHERE id = $3"
const initSettingsStr = "INSERT INTO \"settings\" (id) VALUES ($1)"
const getSettingsStr = "SELECT id, name, starting_word FROM settings WHERE id = $1"

func Connect() *sqlx.DB {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database.")
	}
	return db
}

func ResetInc() {
	db := Connect()
	if _, err := db.Exec(resetIncStr); err != nil {
		fmt.Println(err)
	}
}

func AddUser(name string, word string) int {
	db := Connect()
	if _, err := db.Exec(createUserStr, name, word); err != nil {
		fmt.Println(err)
	}
	users := []User{}
	if err := db.Select(&users, getLastUser); err != nil {
		fmt.Println(err)
	}
	return users[0].ID
}

func DeleteAllUsers() {
	db := Connect()
	if _, err := db.Exec(deleteUsersStr); err != nil {
		fmt.Println(err)
	}
}

func GetName(user int) string {
	db := Connect()
	users := []User{}
	if err := db.Select(&users, getUserStr, user); err != nil {
		fmt.Println(err)
	}
	return users[0].Name
}

func GetStartingWord(user int) string {
	db := Connect()
	users := []User{}
	if err := db.Select(&users, getUserStr, user); err != nil {
		fmt.Println(err)
	}
	return users[0].StartWord
}

func AddPicture(user int, svg string, round int) {
	db := Connect()
	if _, err := db.Exec(insertPictureStr, user, svg, round); err != nil {
		fmt.Println(err)
	}
}

func GetPicture(user int, round int) Picture {
	db := Connect()
	pictures := []Picture{}
	if err := db.Select(&pictures, getPictureStr, user, round); err != nil {
		fmt.Println(err)
	}
	return pictures[0]
}

func AddGuess(user int, guess string, round int) {
	db := Connect()
	if _, err := db.Exec(insertGuessStr, user, guess, round); err != nil {
		fmt.Println(err)
	}
}

func GetGuess(user int, round int) Guess {
	db := Connect()
	guesses := []Guess{}
	if err := db.Select(&guesses, getGuessStr, user, round); err != nil {
		fmt.Println(err)
	}
	return guesses[0]
}

func InitSettings(game int) {
	db := Connect()
	if _, err := db.Exec(initSettingsStr, game); err != nil {
		fmt.Println(err)
	}
}

func ChangeSettings(game int, timeLimit int, difficulty string) {
	db := Connect()
	if _, err := db.Exec(changeSettingsStr, timeLimit, difficulty, game); err != nil {
		fmt.Println(err)
	}
}

func GetSettings(game int) Settings {
	db := Connect()
	settings := []Settings{}
	if err := db.Select(&settings, getSettingsStr); err != nil {
		fmt.Println(err)
	}
	return settings[0]
}
