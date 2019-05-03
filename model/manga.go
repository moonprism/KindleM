package model

import "time"

const (
	SOURCE_MANHUAGUI int = iota
)

type Manga struct {
	Id	int64	`xorm:"pk autoincr" json:"id"`
	Name	string	`xorm:"varchar(250) 'name'" json:"name"`
	Author	string	`xorm:"varchar(250) 'author'" json:"author"`
	Link	string	`xorm:"varchar(250)	'link'" json:"link"`
	Alias	string	`xorm:"varchar(250) 'alias'" json:"alias"`
	Intro	string	`xorm:"text 'intro'" json:"intro"`
	Cover	string	`xorm:"varchar(250) 'cover'" json:"cover"`
	Source	int	`xorm:"TINYINT 'source'" json:"source"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

type MangaDetail struct {
	*Manga
	Chapters []Chapter `json:"chapters"`
}