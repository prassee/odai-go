package main

import (
	"fmt"
	"github.com/prassee/odai/source"
	"github.com/prassee/odai/window"
	"math/rand"
	"sync"
	"time"
)

func main() {

	intStrm := make(chan window.TickData, 2)
	symbols := []int{234, 789, 345}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Add(2)

	aggFn := func(batch []window.TickData) {
		fmt.Printf("received tick data batch %v \n", batch)
	}

	// start the stream by sourcing
	go source.FromStream(&wg, intStrm, func() {
		td := window.TickData{
			Symbol:    symbols[0],
			TimeStamp: time.Now().UnixNano() / 1000000,
			Price:     (float64(rand.Intn(100)) / 100.0) * 100.0}
		intStrm <- td
		time.Sleep(time.Duration(200+rand.Intn(300)) * time.Millisecond)
	})

	go window.OnTumblingWindow(&wg, intStrm, window.TumblingWindow{Duration: 1, Uom: (1000 * 60)}, aggFn)
	// go window.OnSlidingWindow(&wg, intStrm, window.SlidingWindow{Duration: 3 * 1000, Interval: 1 * 1000}, aggFn)

	wg.Wait()
}
