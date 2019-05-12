package lib

import (
	"fmt"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/package/util"
	log "github.com/sirupsen/logrus"
	"os"
)

func DownloadPicture(picture *model.Picture) {
	has, _ := XEngine().Where("src=?", picture.Src).Get(picture)
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
	log.Debugf("start download picture %s fetch from : %s", file, picture.Src)
	if err:= util.DownloadPicture(picture.Src, picture.Referer, file); err != nil {
		log.Errorln(err.Error())
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
