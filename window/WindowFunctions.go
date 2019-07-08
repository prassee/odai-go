package window

import (
	"sync"
)

/*
OnTumblingWindow -
*/
func OnTumblingWindow(wg *sync.WaitGroup,
	intStrm chan TickData,
	window TumblingWindow,
	fn func(batch []TickData)) {
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
}

/*
OnSlidingWindow -
*/
func OnSlidingWindow(wg *sync.WaitGroup,
	intStrm chan TickData,
	window SlidingWindow,
	fn func([]TickData)) {
	defer wg.Done()
	data := map[int64][]TickData{}
	for i := range intStrm {
		y := (i.TimeStamp / window.Interval)
		vl, isExist := data[y]
		if !isExist {
			ky := []TickData{}
			ky = append(ky, i)
			data[y] = ky
		} else {
			vl = append(vl, i)
			data[y] = vl
		}
		if len(data) > int(window.Duration/1000) {
			evictKey := (y - (window.Duration / 1000))
			delete(data, evictKey)
			allData := []TickData{}
			for _, v := range data {
				allData = append(allData, v...)
			}
			fn(allData)
		}
	}
}
