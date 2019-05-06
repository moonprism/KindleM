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
	"log"
)

func WorkPicture() {
	for pic := range lib.PictureDownloadChan {
		lib.DownloadPicture(pic)
	}
}

func WorkChapter() {
	for cha := range lib.ChapterFetchChan {
		err := manhuagui.ChapterProcess(cha)
		if err != nil {
			log.Print(err.Error())
		}
		lib.XEngine().Id(cha.Id).Update(cha)
	}
}

func Work(num int) {
	for i := 0; i < num; i++ {
		go WorkPicture()
	}
	go WorkChapter()
}

// @title kindleM API
// @version 0.0.1
// @description
func main() {
	lib.InitLogrus()

	Work(5)

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
	r.POST("/download", api.DownLoad)

	if err := r.Run(":8001"); err != nil {
		fmt.Printf("run gin : %v\n", err)
	}
}