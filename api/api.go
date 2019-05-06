package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moonprism/kindleM/lib"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/site/manhuagui"
	"net/http"
)

// @Summary search manga
// @ID search-manga
// @Produce  json
// @Param query path string true " "
// @Success 200 {object} model.Manga
// @Router /search/{query} [get]
func Search(context *gin.Context) {
	result := manhuagui.Search(context.Param("query"))
	context.JSON(200, result)
}

// @Summary get manga chapter list
// @Produce  json
// @Param manga_url query string true " "
// @Success 200 {object} model.MangaInfo
// @Router /chapters [get]
func Chapters(context *gin.Context) {
	mangaUrl := context.Query("manga_url")
	manga := &model.Manga{Link:mangaUrl}
	has, _ := lib.XEngine().Get(manga)
	IManga, err := manhuagui.NewManga(manga)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err = IManga.SyncInfo(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if has {
		lib.XEngine().Id(manga.Id).Update(manga)
	} else {
		lib.XEngine().Insert(manga)
	}

	var response model.MangaInfo
	response.Manga = manga
	response.ChapterRowList, _ = IManga.FetchChapterRowList()

	context.JSON(200, response)
}

// @Summary download chapter list
// @Produce  json
// @Param download_list body model.ChapterRowList true " "
// @Success 200 {object} model.ChapterList
// @Router /download [POST]
func Download(context *gin.Context) {
	var chapterRowList model.ChapterRowList
	if err := context.BindJSON(&chapterRowList); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}

	var chapterList model.ChapterList

	for _, chapterRow := range chapterRowList {
		var chapter model.Chapter
		chapter.ChapterRow = chapterRow
		has, _ := lib.XEngine().Get(&chapter)
		if !has {
			lib.XEngine().Insert(&chapter)
		}
		if !chapter.Status() {
			lib.ChapterFetchChan <- &chapter
		}
		chapterList = append(chapterList, chapter)
	}

	context.JSON(http.StatusOK, chapterList)
}

// @Summary check chapter all picture download again
// @Produce  json
// @Param chapter_id path int true " "
// @Success 200 {object} model.PictureList
// @Router /check/chapter/{chapter_id} [POST]
func ReDwonloadChapter(ctx *gin.Context) {
	chapterId := ctx.Param("chapter_id")

	var chapter model.Chapter
	has, _ := lib.XEngine().Id(chapterId).Get(&chapter)
	if !has {
		ctx.JSON(http.StatusNotFound, gin.H{"error":"chapter_id "+chapterId})
		return
	}

	picture := new(model.Picture)
	total, _ := lib.XEngine().Where("chapter_id = ?", chapter.Id).Count(picture)

	if total != int64(chapter.Total) {
		ctx.JSON(http.StatusForbidden, gin.H{"error":"chapter_id "+chapterId})
		return
	}

	if total == int64(chapter.Count) {
		ctx.JSON(http.StatusAlreadyReported, gin.H{"error":"already download"})
		return
	}

	reDownloadPicList := make([]model.Picture, 0)
	lib.XEngine().Where("chapter_id = ? AND status = ?", chapter.Id, false).Find(&reDownloadPicList)

	for _, pic := range reDownloadPicList {
		func(pic model.Picture) {
			lib.PictureDownloadChan <- &pic
		}(pic)
	}
	ctx.JSON(http.StatusOK, reDownloadPicList)
}