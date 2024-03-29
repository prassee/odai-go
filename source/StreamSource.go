package source

import (
	"github.com/prassee/odai/window"
	"sync"
)

/*
FromStream -
*/
func FromStream(wg *sync.WaitGroup, intStrm chan window.TickData, fn func()) {
	defer wg.Done()
	for true {
		fn()
	}
	close(intStrm)
}
