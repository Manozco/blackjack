package main

import (
	"flag"

	"code.google.com/p/go-uuid/uuid"
	"github.com/Manozco/cards"
)

import "fmt"
import "strings"
import "math/rand"

func srand(uuidStr string) {
	alpha := "abcdef0123456789"
	var cnt int64
	if uuidStr == "" {
		uuidStr = uuid.New()
	}

	fmt.Println("UUID:", uuidStr)
	for i, v := range uuidStr {
		cnt += int64(i) * int64(strings.Index(alpha, string(v)))
	}
	rand.Seed(cnt)
}

func main() {
	var uuid string
	flag.StringVar(&uuid, "uuid", "", "an uuid which will seed the random generator used")
	flag.StringVar(&uuid, "u", "", "an uuid which will seed the random generator used")
	flag.Parse()
	srand(uuid)
	// c := cards.NewCardRandom()
	// v, _ := c.Value()
	// fmt.Println(c, v)
	d := cards.NewDeck()
	fmt.Println(d)
	d.Shuffle()
	fmt.Println(d)
	g := cards.NewGame()
	g.Shuffle()
	fmt.Println(g)
}
