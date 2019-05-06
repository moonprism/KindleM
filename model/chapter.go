package model

import "time"

type Chapter struct {
	Id	int64	`xorm:"pk autoincr 'id'" json:"id"`
	ChapterRow `xorm:"extends -"`
	Total	int	`json:"total"`
	Count	int `json:"count"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

func (ch *Chapter) Status()  bool {
	return ch.Total != 0 && ch.Count == ch.Total
}

// ChapterRow is chapter base info
type ChapterRow struct {
	MangaId	int64	`xorm:"index 'manga_id'" json:"manga_id"`
	Title	string	`json:"title"`
	Link	string	`xorm:"unique" json:"link"`
}

type ChapterRowList []ChapterRow

type ChapterList []Chapter 