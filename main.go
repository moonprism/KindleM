package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/moonprism/kindleM/api"
	"github.com/moonprism/kindleM/lib"

	_ "github.com/moonprism/kindleM/docs"
	"github.com/swaggo/gin-swagger"              // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

// @title kindleM API
// @version 0.0.1
// @description
func main() {
	lib.InitLogrus()

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