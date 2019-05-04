package manhuagui

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/moonprism/kindleM/model"
	"github.com/moonprism/kindleM/package/util"
	"log"
	"strconv"
	"strings"
)

const MANHUAGUI_URL  = "https://www.manhuagui.com"

func Search(query string) (result []model.Manga) {
	doc, err := util.GetFetchDocument(fmt.Sprintf("https://www.manhuagui.com/s/%s.html", query))
	if err != nil {
		log.Printf("http request error: %v", err)
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

func setMangaInfo(manga *model.Manga, doc *goquery.Document){
	mangaDoc := doc.Find(".book-cont").First()
	manga.Name = mangaDoc.Find("h1").First().Text()
	manga.Author = mangaDoc.Find(".detail-list").First().Find("li").Eq(1).Find("span").Eq(1).Find("a").First().Text()
	manga.Cover = mangaDoc.Find(".hcover").First().Find("img").First().AttrOr("src", "")
	mangaDoc.Find(".detail-list").First().Find("li").Eq(2).Find("strong").First().Remove()
	manga.Alias = mangaDoc.Find(".detail-list").First().Find("li").Eq(2).Find("a").First().Text()
	manga.Intro = mangaDoc.Find("#intro-cut").Text()
	manga.Source = model.SOURCE_MANHUAGUI
}

func ChapterList(manga *model.Manga) (result []model.ChapterRow)  {
	doc, err := util.GetFetchDocument(manga.Link)
	if err != nil {
		log.Printf("http request error: %v", err)
	}

	setMangaInfo(manga, doc)

	var resBack []model.ChapterRow
	for eqIndex := 1; eqIndex >= 0; eqIndex-- {
		doc.Find(".chapter-list").Eq(eqIndex).Find("ul").Each(func(i int, UlSelection *goquery.Selection) {
			UlSelection.Find("li").Each(func(i2 int, LiSelection *goquery.Selection) {
				var node  model.ChapterRow
				node.MangaId = manga.Id
				node.Link = MANHUAGUI_URL + LiSelection.Find("a").First().AttrOr("href", "")
				node.Title = LiSelection.Find("a").First().AttrOr("title", "")
				resBack = append(resBack, node)
			})
			for i3 := len(resBack)-1; i3 >= 0; i3-- {
				result = append([]model.ChapterRow{resBack[i3]}, result...)
			}
			resBack = []model.ChapterRow{}
		})
	}

	return
}

func PictureList(chapter *model.Chapter) (result []model.Picture) {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	var pageSelectHtml string

	err := chromedp.Run(ctx,
		chromedp.Navigate(chapter.Link),
		chromedp.WaitVisible("#pageSelect", chromedp.ByQuery),
		chromedp.InnerHTML(`#pageSelect`, &pageSelectHtml, chromedp.ByQuery),
	)

	if err != nil {
		log.Printf("http request error: %v", err)
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(pageSelectHtml))

	chapter.Total, _ = strconv.Atoi(doc.Find("option").Last().AttrOr("value", "0"))

	for page := 1; page <= chapter.Total; page++ {
		node := model.Picture{
			MangaId: chapter.MangaId,
			ChapterId: chapter.Id,
			Index: page,
		}
		err = chromedp.Run(ctx,
			chromedp.WaitVisible(`#mangaFile`, chromedp.ByQuery),
			chromedp.AttributeValue(`#mangaFile`, "src", &node.Src, nil, chromedp.ByQuery),
			chromedp.Click(`#next`, chromedp.ByQuery),
		)
		result = append(result, node)
	}

	return
}