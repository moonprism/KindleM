package lib

import (
	"github.com/moonprism/kindleM/model"
)

// simple async works queue

var PictureDownloadChan = make(chan *model.Picture, 100)
var ChapterFetchChan = make(chan *model.Chapter, 10)