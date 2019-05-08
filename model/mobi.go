package model

import "time"

type Mobi struct {
	Id	int64	`xorm:"pk autoincr 'id'"`
	MobiMeta
	MobiFile	string	`xorm:"'mobi_file'" json:"mobi_file"`
	EpubFile	string	`xorm:"'epub_file'" json:"epub_file"`
	ProcessInfo	string	`xorm:"text 'process_info'"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// MobiMeta is base book file info
type MobiMeta struct {
	Title	string
	Author	string
	Cover	string
}

type MobiXChapter struct {
	Id	int64	`xorm:"pk autoincr 'id'"`
	MobiId int64 `xorm:"index 'mobi_id'"`
	ChapterId int64 `xorm:"index 'chapter_id'"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type MobiInfo struct {
	MobiMeta
	ChapterIdList	`json:"chapter_id_list"`
}