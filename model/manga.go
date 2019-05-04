package model

import "time"

const (
	SOURCE_MANHUAGUI int = iota
)

type Manga struct {
	Id	int64	`xorm:"pk autoincr" json:"id"`
	Name	string	`json:"name"`
	Author	string	`json:"author"`
	Link	string	`xorm:"unique" json:"link"`
	Alias	string	`json:"alias"`
	Intro	string	`xorm:"text" json:"intro"`
	Cover	string	`json:"cover"`
	Source	int	`xorm:"TINYINT 'source'" json:"source"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

type MangaDetail struct {
	*Manga
	Chapters []Chapter `json:"chapters"`
}