package epub

import (
	"fmt"
	"github.com/bmaupin/go-epub"
	"log"
)

type MakeEpubInfo struct {
	epub 	*epub.Epub
	Title	string
	Author	string
	Cover	string
}

func NewEpub(title string) *MakeEpubInfo {
	return &MakeEpubInfo{
		epub: epub.NewEpub(title),
		Title: title,
	}
}

func (info *MakeEpubInfo) MakeChapter(name string) {

}

func (info *MakeEpubInfo) AddImage(fileName string) {
	imgTag := `<img src="%s" alt="Cover Image" />`
	imgpath, _ := info.epub.AddImage(fileName, "")
	_, _ = info.epub.AddSection(fmt.Sprintf(imgTag, imgpath), "", "", "")
}

func (info *MakeEpubInfo) Generate(fileName string) {
	if info.Author != "" {
		info.epub.SetAuthor(info.Author)
	}

	if info.Cover != "" {
		info.epub.SetAuthor(info.Cover)
	}

	err := info.epub.Write(fileName)
	if err != nil {
		log.Fatal(err)
	}
}
