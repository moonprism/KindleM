package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/site/manhuagui"
)

// Search common
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

// Chapters common
// @Summary get manga chapter list
// @Produce  json
// @Param manga_url query string true " "
// @Success 200 {object} model.MangaDetail
// @Router /chapter [get]
func Chapters(context *gin.Context) {
	mangaUrl := context.Query("manga_url")

	manga := &model.Manga{Link:mangaUrl}
	result := manhuagui.ChapterList(manga)

	var response model.MangaDetail
	response.Manga = manga
	for _, r := range result {
		response.Chapters = append(response.Chapters, r)
	}

	context.JSON(200, response)
}
