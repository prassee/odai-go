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

	// start the stream by sourcing
	go func() {
		defer wg.Done()
		for true {
			td := TickData{
				Symbol:    symbols[0],
				TimeStamp: time.Now().UnixNano() / 1000000,
				Price:     (float64(rand.Intn(100)) / 100.0) * 100.0}
			intStrm <- td
			time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
		}
		close(intStrm)
	}()

	go func(window TumblingWindow, fn func(batch []TickData)) {
		defer wg.Done()
		data := []TickData{}
		for i := range intStrm {
			den := (window.Duration * window.Uom)
			data = append(data, i)
			if i.TimeStamp%den < 1000 {
				fn(data)
				data = []TickData{}
			}
		}
	}(TumblingWindow{Duration: 1, Uom: (1000 * 60)}, func(batch []TickData) {
		fmt.Printf("received tick data batch %v \n", batch)
	})

	go func(window SlidingWindow) {
		for i := range intStrm {
			println(i)
		}
	}(SlidingWindow{Duration: 3 * 1000, Interval: 1 * 1000})

	wg.Wait()
}
