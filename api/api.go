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
// @Success 200 {object} model.MangaInfo
// @Router /chapters [get]
func Chapters(context *gin.Context) {
	mangaUrl := context.Query("manga_url")

	manga := &model.Manga{Link:mangaUrl}

	has, _ := lib.XEngine().Get(manga)
	if !has {
		lib.XEngine().Insert(manga)
	}
	result := manhuagui.ChapterList(manga)

	if manga.Name != "" {
		lib.XEngine().Insert(manga)
	}

	var response model.MangaInfo
	response.Manga = manga
	for _, r := range result {
		response.ChapterRowList = append(response.ChapterRowList, r)
	}

	context.JSON(200, response)
}

// @Summary download chapter list
// @Produce  json
// @Param download_list body model.ChapterRowList true " "
// @Success 200 {object} model.ChapterList
// @Router /download [POST]
func DownLoad(context *gin.Context) {
//	mangaUrl := context.Query("manga_url")

//	manga := &model.Manga{Link:mangaUrl}

	// todo
}