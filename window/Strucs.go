package window

/*
TumblingWindow -
*/
type TumblingWindow struct {
	Duration int64
	Uom      int64
}

/*
SlidingWindow -
*/
type SlidingWindow struct {
	Duration int64
	Interval int64
}

/*
TickData -
*/
type TickData struct {
	Symbol    int
	TimeStamp int64
	Price     float64
}
