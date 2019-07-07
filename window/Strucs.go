package window

type TumblingWindow struct {
	Duration int64
	Uom      int64
}

type SlidingWindow struct {
	Duration int64
	Interval int64
}

type TickData struct {
	Symbol    int
	TimeStamp int64
	Price     float64
}
