package model

type Chapter struct {
	Id	int64	`xorm:"pk autoincr 'id'"`
	MangaId	int64	`xorm:"index 'manga_id'"`
	Title	string	`xorm:"varchar(250) 'title'"`
	Link	string	`xorm:"varchar(250)	'link'"`
	Total	int	`xorm:"'total'"`
	Status	bool	`xorm:"'status'"`
}