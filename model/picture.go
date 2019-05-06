package model

import (
	"fmt"
	"net/url"
	"path"
	"time"
)

type Picture struct {
	Id      int64 `xorm:"pk autoincr 'id'"`
	MangaId int64 `xorm:"index 'manga_id' notnull"`
	ChapterId int64 `xorm:"index 'chapter_id' notnull"`
	Src	string
	Status	bool
	Referer	string
	Index	int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type PictureList []Picture

func (pic *Picture) Path() string {
	return fmt.Sprintf("%d/%d", pic.MangaId, pic.ChapterId)
}

func (pic *Picture) File() string {
	u, _ := url.Parse(pic.Src)
	return fmt.Sprintf("%s/%d%s", pic.Path(), pic.Index, path.Ext(u.Path))
}