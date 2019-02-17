package telestrationsLib

import (
	"math/rand"
	"time"
)

func GetWord() string {
	words := []string{"Angel", "Pancake", "Bike", "Tree", "Water", "Ocean", "Grass", "Computer", "Food", "Phone", "Pencil", "Drink", "Trash", "Chair", "Table", "Brick"}
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(len(words))
	return words[r]
}
