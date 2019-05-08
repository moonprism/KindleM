package manhuagui

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/moonprism/kindleM/lib"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/package/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const MANHUAGUI_URL  = "https://www.manhuagui.com"

var PictureDownloadChan = make(chan *model.Picture, 500)

// init start works(chan) process image download
func init() {
	for i := 0; i < 1; i++ {
		go func() {
			for pic := range PictureDownloadChan {
				lib.DownloadPicture(pic)
			}
		}()
	}
}

type MangaManage struct {
	*model.Manga
	Doc *goquery.Document
}

func NewManga(model *model.Manga) (IManga *MangaManage, err error) {
	IManga = new(MangaManage)
	IManga.Manga = model
	IManga.Doc, err = util.GetFetchDocument(model.Link)
	return
}

func Search(query string) (result []model.Manga) {
	doc, err := util.GetFetchDocument(fmt.Sprintf("https://www.manhuagui.com/s/%s.html", query))
	if err != nil {
		log.Errorf("http request error: %v", err)
	}

	doc.Find(".book-result").First().Find(".cf").Each(func(i int, selection *goquery.Selection) {
		node := model.Manga{
			Name: selection.Find(".book-detail").First().Find("a").First().Text(),
			Link: MANHUAGUI_URL + selection.Find("a").First().AttrOr("href", ""),
			Author: selection.Find(".tags").Eq(2).Find("a").Text(),
			Alias: selection.Find(".tags").Eq(3).Find("a").First().Text(),
			Intro: selection.Find(".intro").First().Text(),
			Cover: selection.Find(".bcover").First().Find("img").First().AttrOr("src", ""),
		}
		result = append(result, node)
	})

	return result
}

func (IManga *MangaManage) SyncInfo() (err error) {
	mangaDoc := IManga.Doc.Find(".book-cont").First()
	manga := IManga.Manga
	if mangaDoc.Nodes != nil {
		manga.Name = mangaDoc.Find("h1").First().Text()
		manga.Author = mangaDoc.Find(".detail-list").First().Find("li").Eq(1).Find("span").Eq(1).Find("a").First().Text()
		manga.Cover = mangaDoc.Find(".hcover").First().Find("img").First().AttrOr("src", "")
		mangaDoc.Find(".detail-list").First().Find("li").Eq(2).Find("strong").First().Remove()
		manga.Alias = mangaDoc.Find(".detail-list").First().Find("li").Eq(2).Find("a").First().Text()
		manga.Intro = mangaDoc.Find("#intro-cut").Text()
		manga.Source = model.SOURCE_MANHUAGUI
	} else {
		err = errors.New("go-query .book-cont not found in "+manga.Link)
	}
	log.Debugf("manga %s is sync", manga.Name)
	return
}

func (IManga *MangaManage) FetchChapterRowList() (chapterRowList model.ChapterRowList, err error) {
	var resBack []model.ChapterRow
	for eqIndex := 2; eqIndex >= 0; eqIndex-- {
		IManga.Doc.Find(".chapter-list").Eq(eqIndex).Find("ul").Each(func(i int, UlSelection *goquery.Selection) {
			UlSelection.Find("li").Each(func(i2 int, LiSelection *goquery.Selection) {
				var node  model.ChapterRow
				node.MangaId = IManga.Id
				node.Link = MANHUAGUI_URL + LiSelection.Find("a").First().AttrOr("href", "")
				node.Title = LiSelection.Find("a").First().AttrOr("title", "")
				resBack = append(resBack, node)
			})
			for i3 := len(resBack)-1; i3 >= 0; i3-- {
				chapterRowList = append([]model.ChapterRow{resBack[i3]}, chapterRowList...)
			}
			resBack = []model.ChapterRow{}
		})
	}
	for index := range chapterRowList {
		chapterRowList[index].Index = index
	}
	return
}

func ChapterProcess(chapter *model.Chapter) (err error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 500 * time.Second)
	defer cancel()

	log.Debugf("start chapter process in chromedp : %s", chapter.Link)
	var pageSelectHtml string
	err = chromedp.Run(ctx,
		chromedp.Navigate(chapter.Link),
		// Action wait for cookie setup completed
		chromedp.ActionFunc(func(ctx context.Context, h cdp.Executor) (err error) {
			var cookies []*network.Cookie
			for i:=0; i<10; i++ {
				cookies, err = network.GetAllCookies().Do(ctx, h)
				if err != nil {
					return err
				}
				if len(cookies) < 3 {
					time.Sleep(1*time.Second)
				} else {
					break
				}
			}
			if len(cookies) < 3 {
				log.Errorf("set cookie error in chromedp : %s", chapter.Link)
			}
			return
		}),
		chromedp.Reload(),
		chromedp.WaitVisible("#pageSelect", chromedp.ByQuery),
		chromedp.InnerHTML(`#pageSelect`, &pageSelectHtml, chromedp.ByQuery),
	)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageSelectHtml))
	if err != nil {
		return
	}

	log.Debugln(doc.Html())

	chapter.Total, err = strconv.Atoi(doc.Find("option").Last().AttrOr("value", "0"))
	if err != nil {
		return err
	}

	log.Debugf("fetch total page %d from : %s", chapter.Total, chapter.Link)

	lib.XEngine().Id(chapter.Id).Update(chapter)
	for page := 1; page <= chapter.Total; page++ {
		picture := model.Picture{
			MangaId: chapter.MangaId,
			ChapterId: chapter.Id,
			Index: page,
			Referer: chapter.Link,
		}
		err = chromedp.Run(ctx,
			chromedp.WaitVisible(`#mangaFile`, chromedp.ByQuery),
			chromedp.AttributeValue(`#mangaFile`, "src", &picture.Src, nil, chromedp.ByQuery),
			chromedp.Click(`#next`, chromedp.ByQuery),
		)
		if err != nil {
			return
		}
		log.Debugf("process picture dwonload : %s", picture.Src)
		PictureDownloadChan <- &picture
	}

	return
}