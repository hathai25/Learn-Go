package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(philoNum int) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", philoNum+1)
		time.Sleep(time.Millisecond * 500)
		fmt.Println("finishing eating ", philoNum+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			philos[i].eat(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
