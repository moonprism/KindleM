package model

import "time"

type Mobi struct {
	Id	int64	`xorm:"pk autoincr 'id'"`
	MobiFile	string	`xorm:"varchar(250) 'mobi_file'"'`
	EpubFile	string	`xorm:"varchar(250) 'epub_file'"'`
	ProcessInfo	string	`xorm:"text(250) 'process_info'"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type MobiXChapter struct {
	Id	int64	`xorm:"pk autoincr 'id'"`
	MobiId int64 `xorm:"index 'mobi_id'"`
	ChapterId int64 `xorm:"index 'chapter_id'"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}