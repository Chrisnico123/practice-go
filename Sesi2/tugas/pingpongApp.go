package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

const BreakPoint = 11

func finish(done chan *Player) {
	for {
		select {
		case d := <-done: 
			fmt.Println("Player " ,d.Name, " kalah pada pukulan ke ", d.Hit)
			return
		}
	}
}

func play(name string, player, done chan *Player) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	for {
		select {
		case k := <-player:
			v := rand.Intn(max-min) + min
			time.Sleep(500 * time.Millisecond)
			k.Hit++
			k.Name = name
			fmt.Println("Player ", k.Name, " dengan pukulan ke ", k.Hit, "dengan counter ", v)
			if v%BreakPoint == 0 {
				done <- k
				return
			}

			player <- k
		}
	}
}

func main() {
	player := make(chan *Player)
	done := make(chan *Player)
	players := []string{"A", "B"}

	for _, p := range players {
		go play(p, player , done)
	}

	player <- new(Player)

	finish(done)
}