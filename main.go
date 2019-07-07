package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TickData struct {
	Symbol    int
	TimeStamp int64
	Price     float64
}

type TumblingWindow struct {
	Duration int64
	Uom      int64
}

type SlidingWindow struct {
	Duration int64
	Interval int64
}

func main() {
	intStrm := make(chan TickData, 2)
	symbols := []int{234, 789, 345}
	wg := sync.WaitGroup{}

	wg.Add(2)

	aggFn := func(batch []TickData) {
		fmt.Printf("received tick data batch %v \n", batch)
	}

	// start the stream by sourcing
	go func() {
		defer wg.Done()
		for true {
			td := TickData{
				Symbol:    symbols[0],
				TimeStamp: time.Now().UnixNano() / 1000000,
				Price:     (float64(rand.Intn(100)) / 100.0) * 100.0}
			intStrm <- td
			time.Sleep(time.Duration(200+rand.Intn(300)) * time.Millisecond)
		}
		close(intStrm)
	}()

	go OnTumblingWindow(&wg, intStrm, TumblingWindow{Duration: 1, Uom: (1000 * 60)}, aggFn)
	// 	go OnSlidingWindow(&wg, intStrm, SlidingWindow{Duration: 3 * 1000, Interval: 1 * 1000}, aggFn)

	wg.Wait()
}
