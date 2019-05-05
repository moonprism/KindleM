package lib

import (
	"fmt"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/package/util"
	"os"
	"path"
	"strings"
)

func DownloadPicture(picture *model.Picture) {
	has, _ := XEngine().Get(picture)
	if !has {
		XEngine().Insert(picture)
	}

	if picture.Status {
		return
	}

	basePath := fmt.Sprintf("download/%d/%d", picture.MangaId, picture.ChapterId)

	// download image
	_ = os.MkdirAll(basePath, os.ModePerm)

	srcArr := strings.Split(picture.Src, "?")

	file := fmt.Sprintf("%s/%d%s", basePath, picture.Id, path.Ext(srcArr[0]))
	if err:= util.DownloadPicture(picture.Src, picture.Referer, file); err != nil {
		println(err.Error())
		return
	}

	picture.File = file
	picture.Status = true
	XEngine().Id(picture.Id).Update(picture)
}
