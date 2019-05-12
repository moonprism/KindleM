package model

import "time"

type Mobi struct {
	Id	int64	`xorm:"pk autoincr 'id'"`
	MobiMeta	`xorm:"extends -"`
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
	// swag wtf?
	ChapterIdList ChapterIdList	`json:"chapter_id_list"`
}