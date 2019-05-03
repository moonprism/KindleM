package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/moonprism/kindleM/api"
	"github.com/moonprism/kindleM/model"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
	"log"

	_ "github.com/moonprism/kindleM/docs"
)

// @title kindleM API
// @version 0.0.1
// @description
func main() {

	engine, err := xorm.NewEngine("mysql", "root:123456@/app?charset=utf8")

	if err != nil {
		log.Printf("%v", err)
	}

	err = engine.Sync2(new(model.Picture))
	err = engine.Sync2(new(model.Manga))
	err = engine.Sync2(new(model.Chapter))
	err = engine.Sync2(new(model.Mobi))

	if err != nil {
		log.Printf("%v\n", err)
	}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))


	r.GET("/search/:query", api.Search)
	r.GET("/chapter", api.Chapters)

	r.Run(":8001")

}