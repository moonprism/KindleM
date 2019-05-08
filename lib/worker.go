package lib

import (
	"github.com/moonprism/kindleM/model"
)

// simple async works queue

var ChapterFetchChan = make(chan *model.Chapter, 10)
var MobiGenerateChan = make(chan *model.Mobi, 5)