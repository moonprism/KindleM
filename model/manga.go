package model

const (
	SOURCE_MANHUAGUI int = iota
)

type Manga struct {
	Id	int64	`xorm:"pk autoincr"`
	Name	string	`xorm:"varchar(250) 'name'"`
	Author	string	`xorm:"varchar(250) 'author'"`
	Link	string	`xorm:"varchar(250)	'link'"`
	Alias	string	`xorm:"varchar(250) 'alias'"`
	//Tags	string	`xorm:"varchar(250) 'tags'"`
	Intro	string	`xorm:"text 'intro'"`
	Cover	string	`xorm:"varchar(250) 'cover'"`
	Source	int	`xorm:"TINYINT 'source'"`
}