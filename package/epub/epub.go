package epub

import (
	"fmt"
	"github.com/bmaupin/go-epub"
	"github.com/moonprism/kindleM/model"
	"log"
)

type MakeEpubInfo struct {
	Resource	[] model.Chapter
	Title	string
	Author	string
	Cover	string
}

func (info *MakeEpubInfo) make() {
	e := epub.NewEpub(info.Title)

	if info.Author != "" {
		e.SetAuthor(info.Author)
	}

	if info.Cover != "" {
		e.SetCover(info.Cover, "")
	}

	imgTag := `<img src="%s" alt="Cover Image" />`

	for _, chapter := range info.Resource {

		// todo search all picture in the chapter
		imgpath, err := e.AddImage(folder+file.Name(), "")
		if err != nil {
			log.Fatal(err)
		}
		_, err = e.AddSection(fmt.Sprintf(imgTag, imgpath), "", "", "")
		if err != nil {
			log.Fatal(err)
		}
	}
	err := e.Write("xy.epub")
	if err != nil {
		log.Fatal(err)
	}
}
