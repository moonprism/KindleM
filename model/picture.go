package model

type Picture struct {
	Id      int64 `xorm:"pk autoincr 'id'"`
	MangaId int64 `xorm:"index 'manga_id'"`
	ChapterId int64 `xorm:"index 'chapter_id'"`
	Src	string	`xorm:"varchar(250) 'src'"`
	File	string	`xorm:"varchar(250) 'file'"`
	Index	int	`xorm:"'index'"`
}