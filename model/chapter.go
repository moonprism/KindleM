package model

import "time"

type Chapter struct {
	Id	int64	`xorm:"pk autoincr 'id'" json:"id"`
	MangaId	int64	`xorm:"index 'manga_id'" json:"manga_id"`
	Title	string	`xorm:"varchar(250) 'title'" json:"title"`
	Link	string	`xorm:"varchar(250)	'link'" json:"link"`
	Total	int	`xorm:"'total'" json:"total"`
	Status	bool	`xorm:"'status'" json:"status"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}