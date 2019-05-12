package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moonprism/kindleM/lib"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/package/epub"
	"github.com/moonprism/kindleM/package/util"
	"github.com/moonprism/kindleM/site/manhuagui"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/exec"
	"sort"
)

// @Summary search manga
// @ID search-manga
// @Produce  json
// @Param query path string true " "
// @Success 200 {object} model.Manga
// @Router /search/{query} [get]
func Search(ctx *gin.Context) {
	result := manhuagui.Search(ctx.Param("query"))
	ctx.JSON(200, result)
}

// @Summary get manga chapter list
// @Produce  json
// @Param manga_url query string true " "
// @Success 200 {object} model.MangaInfo
// @Router /chapters [get]
func Chapters(ctx *gin.Context) {
	mangaUrl := ctx.Query("manga_url")
	manga := &model.Manga{Link:mangaUrl}
	has, _ := lib.XEngine().Get(manga)
	IManga, err := manhuagui.NewManga(manga)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err = IManga.SyncInfo(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
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

	ctx.JSON(200, response)
}

// @Summary download chapter list
// @Produce  json
// @Param download_list body model.ChapterRowList true " "
// @Success 200 {object} model.ChapterList
// @Router /download [POST]
func Download(ctx *gin.Context) {
	var chapterRowList model.ChapterRowList
	if err := ctx.BindJSON(&chapterRowList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
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

	ctx.JSON(http.StatusOK, chapterList)
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
			manhuagui.PictureDownloadChan <- &pic
		}(pic)
	}
	ctx.JSON(http.StatusOK, reDownloadPicList)
}

// @Summary CountProcess
// @Produce  json
// @Success 200 {object} model.ProcessCount
// @Router /count/process [GET]
func CountProcess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.ProcessCount{
		Chromedp: len(lib.ChapterFetchChan),
		Picture: len(manhuagui.PictureDownloadChan),
	})
}

// @Summary download manga list
// @Produce  json
// @Success 200 {object} model.MangaDetailList
// @Router /manga [GET]
func DownloadMangaList(ctx *gin.Context) {
	mangaMap := make(map[int64]model.Manga)
	lib.XEngine().Limit(5, 0).OrderBy("id desc").Find(&mangaMap)
	keys := make([]int64, 0, len(mangaMap))
	for k := range mangaMap {
		keys = append(keys, k)
	}
	chapterList := make(model.ChapterList, 0)
	lib.XEngine().In("manga_id", keys).OrderBy("id desc").Find(&chapterList)

	chaptersMap := map[int64]model.ChapterList{}
	var response model.MangaDetailList
	for _, chapter := range chapterList {
		chaptersMap[chapter.MangaId] = append(chaptersMap[chapter.MangaId], chapter)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i:=len(keys)-1; i>=0 ; i-- {
		response = append(response, model.MangaDetail{
			Manga: mangaMap[keys[i]],
			ChapterList: chaptersMap[keys[i]],
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary generate manga
// @Accept json
// @Produce  json
// @Param mobi_info body model.MobiInfo true " "
// @Success 200 {object} model.Mobi
// @Router /manga/generate [POST]
func GenerateManga(ctx *gin.Context) {
	var mobiInfo model.MobiInfo
	var mobi model.Mobi
	if err := ctx.BindJSON(&mobiInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}

	epubInfo := epub.NewEpub(mobiInfo.Title)
	cover := "./runtime/cover.jpg"
	util.DownloadPicture(mobiInfo.Cover, "", cover)
	epubInfo.Cover = cover
	epubInfo.Author = mobiInfo.Author

	mobi.MobiMeta = mobiInfo.MobiMeta
	_, err := lib.XEngine().Insert(&mobi)
	if err != nil {
		logrus.Errorln(err.Error())
	}

	for _, chapterId := range mobiInfo.ChapterIdList {
		var pictureList model.PictureList
		var mobiPic model.MobiXChapter

		chapter := model.Chapter{Id:chapterId}

		_, err := lib.XEngine().Get(&chapter)
		if err != nil {
			logrus.Errorln(err.Error())
		}

		err = lib.XEngine().Where("chapter_id=?", chapterId).OrderBy("`index` asc").Find(&pictureList)
		if err != nil {
			logrus.Errorln(err.Error())
		}

		mobiPic.MobiId = mobi.Id
		mobiPic.ChapterId = chapterId
		lib.XEngine().Insert(&mobiPic)

		epubInfo.MakeChapter("")
		for _, pic := range pictureList {
			pngFile := fmt.Sprintf("./runtime/%d.png", pic.Id)
			cmd := exec.Command("dwebp", "./download/"+pic.File(), "-o", pngFile)
			cmd.Run()
			defer os.Remove(pngFile)
			epubInfo.AddImage(pngFile)
		}
	}

	epubInfo.Generate(fmt.Sprintf("./runtime/%d.epub", mobi.Id))
	ctx.JSON(http.StatusOK, mobi)
}