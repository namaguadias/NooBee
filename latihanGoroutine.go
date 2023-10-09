package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

func appsPingPong(playerName string, player *Player, fightPlayer chan int, fightPlayerHit chan int, hitValue chan int, numchan chan int, wg *sync.WaitGroup) {

	counter := 0
	var totalHit int

	for {
		select {
		case num := <-numchan:
			log.Println(num)
		case numchan := <-hitValue:
			value := rand.Intn(100-1) + 1
			counter += value
			totalHit += numchan
			fmt.Printf("%s = Hit %d // counter %d\n", playerName, numchan, counter)

			if counter%11 == 0 {
				fmt.Printf("%s lose, total Hit ; %d, lose in counter %d\n", playerName, totalHit, counter)
				return
			}

			fightPlayer <- numchan
		case fightHit := <-fightPlayerHit:
			counter += fightHit
			totalHit += fightHit
			fmt.Printf("%s = Hit %d // counter %d\n", playerName, fightHit, counter)

			if counter%11 == 0 {
				fmt.Printf("%s lose, total Hit : %d, lose in counter %d\n", playerName, totalHit, counter)
				return
			}

			fightPlayer <- fightHit

			time.Sleep(1 * time.Second)
			defer wg.Done()
		}
	}
}

func main() {

	rincember := Player{
		Name: "rincember",
		Hit:  0,
	}

	ochaichihii := Player{
		Name: "ochaichihii",
		Hit:  0,
	}

	var wg sync.WaitGroup

	rincemberHit := make(chan int)
	ochaichihiiHit := make(chan int)
	KnockOut := make(chan int)
	numberHit := make(chan int)

	wg.Add(1)
	go appsPingPong("rincember", &rincember, rincemberHit, ochaichihiiHit, KnockOut, numberHit, &wg)

	wg.Add(1)
	go appsPingPong("ochaichihii", &ochaichihii, ochaichihiiHit, rincemberHit, KnockOut, numberHit, &wg)

	KnockOut <- 1

	wg.Wait()
}
