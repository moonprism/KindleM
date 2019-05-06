package lib

import (
	"fmt"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/package/util"
	"os"
)

func DownloadPicture(picture *model.Picture) {
	has, _ := XEngine().Get(picture)
	if !has {
		XEngine().Insert(picture)
	}

	if picture.Status {
		return
	}

	basePath := fmt.Sprintf("download/%s", picture.Path())

	// download image
	_ = os.MkdirAll(basePath, os.ModePerm)

	file := fmt.Sprintf("download/%s", picture.File())
	if err:= util.DownloadPicture(picture.Src, picture.Referer, file); err != nil {
		println(err.Error())
		return
	}

	picture.Status = true
	session := XEngine().NewSession()
	defer session.Close()
	_, err := session.ID(picture.Id).Cols("status").Update(picture)
	if err != nil {
		session.Rollback()
		return
	}
	_, err = session.ID(picture.ChapterId).Incr("count").Update(&model.Chapter{})
	if err != nil {
		session.Rollback()
		return
	}
	session.Commit()
}
