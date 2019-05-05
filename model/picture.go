package model

import "time"

type Picture struct {
	Id      int64 `xorm:"pk autoincr 'id'"`
	MangaId int64 `xorm:"index 'manga_id' notnull"`
	ChapterId int64 `xorm:"index 'chapter_id' notnull"`
	Src	string
	Status	bool	`xorm:"default false notnull"`
	Referer	string
	File	string
	Index	int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type PictureList []Picture