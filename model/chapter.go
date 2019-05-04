package model

import "time"

type Chapter struct {
	Id	int64	`xorm:"pk autoincr 'id'" json:"id"`
	ChapterRow `xorm:"extends -"`
	Total	int	`json:"total"`
	Status	bool	`json:"status"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

// ChapterRow is chapter base info
type ChapterRow struct {
	MangaId	int64	`xorm:"index 'manga_id'" json:"manga_id"`
	Title	string	`json:"title"`
	Link	string	`xorm:"unique" json:"link"`
}

type ChapterRowList []ChapterRow

type ChapterList []Chapter 