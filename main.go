package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/moonprism/kindleM/api"
	_ "github.com/moonprism/kindleM/docs"
	"github.com/moonprism/kindleM/lib"
	"github.com/moonprism/kindleM/site/manhuagui"
	"github.com/swaggo/gin-swagger"              // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func workChromedp() {
	for cha := range lib.ChapterFetchChan {
		manhuagui.ChapterProcess(cha)
		lib.XEngine().Id(cha.Id).Update(cha)
	}
}

func workGenerateMobi() {
	for _ = range lib.MobiGenerateChan {

	}
}

// @title kindleM API
// @version 0.0.1
// @description
func main() {
	lib.InitLogrus()

	go workChromedp()

	// util.SetProxy("127.0.0.1", "1080")
	//file, _ := os.OpenFile(lib.Config.Log.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//
	//gin.DisableConsoleColor()
	//gin.DefaultWriter = io.MultiWriter(file)
	//gin.DefaultErrorWriter = io.MultiWriter(file)
	// run
	r := gin.Default()

	// swagger
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.GET("/search/:query", api.Search)
	r.GET("/chapters", api.Chapters)
	r.POST("/download", api.Download)
	r.POST("/check/chapter/:chapter_id", api.ReDwonloadChapter)

	r.GET("/count/process", api.CountProcess)

	// generate file
	r.GET("/manga", api.DownloadMangaList)
	r.POST("/manga/generate", )

	if err := r.Run(":8001"); err != nil {
		fmt.Printf("run gin : %v\n", err)
	}
}