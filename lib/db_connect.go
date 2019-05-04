package lib

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/moonprism/kindleM/model"
	"log"
)

var xormEngine *xorm.Engine

func XEngine () *xorm.Engine {
	if xormEngine == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			Config.Storage.Username,
			Config.Storage.Password,
			Config.Storage.Host,
			Config.Storage.Port,
			Config.Storage.Database,
		)

		var err error
		xormEngine, err = xorm.NewEngine("mysql", dsn)
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		if err = xormEngine.Sync2(new(model.Picture)); err != nil {
			log.Fatalf("picture model : %v\n", err)
		}
		if err = xormEngine.Sync2(new(model.Mobi)); err != nil {
			log.Fatalf("mobi model : %v\n", err)
		}
		if err = xormEngine.Sync2(new(model.MobiXChapter)); err != nil {
			log.Fatalf("mobixchapter model : %v\n", err)
		}
		if err = xormEngine.Sync2(new(model.Chapter)); err != nil {
			log.Fatalf("chapter model : %v\n", err)
		}
		if err = xormEngine.Sync2(new(model.Manga)); err != nil {
			log.Fatalf("manga model : %v\n", err)
		}
	}

	return xormEngine
}