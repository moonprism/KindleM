package lib

import "sync"

// Mutex is global mutex lock for control chromedp operating
var Mutex *sync.Mutex

func init() {
	Mutex = new(sync.Mutex)
}
