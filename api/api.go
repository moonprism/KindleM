package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moonprism/kindleM/lib"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/site/manhuagui"
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
// @Success 200 {object} model.MangaDetail
// @Router /chapters [get]
func Chapters(context *gin.Context) {
	mangaUrl := context.Query("manga_url")

	manga := &model.Manga{Link:mangaUrl}
	has, _ := lib.XEngine().Get(manga)

	result := manhuagui.ChapterList(manga)

	if !has {
		lib.XEngine().Insert(manga)
	} else {
		lib.XEngine().Id(manga.Id).Update(manga)
	}

	var response model.MangaDetail
	response.Manga = manga
	for _, r := range result {
		response.Chapters = append(response.Chapters, r)
	}

	context.JSON(200, response)
}

// @Summary download chapter list
// @Produce  json
// @Param download_list body model.ChapterRowList true " "
// @Success 200 {object} model.ChapterList
// @Router /download [get]
func DownLoad(context *gin.Context) {
//	mangaUrl := context.Query("manga_url")

//	manga := &model.Manga{Link:mangaUrl}


}