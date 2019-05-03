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

type ManHuaGui struct {

}

func Search(query string) (result []model.Manga) {
	doc, err := util.GetFetchDocument(fmt.Sprintf("https://www.manhuagui.com/s/%s.html", query))
	if err != nil {
		log.Printf("http request error: %v", err)
	}

	doc.Find(".book-detail").Each(func(i int, selection *goquery.Selection) {
		node := model.Manga{
			Name: selection.Find("a").First().Text(),
			Link: MANHUAGUI_URL + selection.Find("a").First().AttrOr("href", ""),
			Author: selection.Find(".tags").Eq(2).Find("a").Text(),
			Alias: selection.Find(".tags").Eq(3).Find("a").First().Text(),
			Cover: selection.Find(".bcover").Find("img").AttrOr("src", ""),
			Source: model.SOURCE_MANHUAGUI,
		}
		result = append(result, node)
	})

	return result
}

func ChapterList(manga *model.Manga) (result []model.Chapter)  {
	doc, err := util.GetFetchDocument(manga.Link)
	if err != nil {
		log.Printf("http request error: %v", err)
	}

	var resBack []model.Chapter

	for eqIndex := 1; eqIndex >= 0; eqIndex-- {
		doc.Find(".chapter-list").Eq(eqIndex).Find("ul").Each(func(i int, UlSelection *goquery.Selection) {
			UlSelection.Find("li").Each(func(i2 int, LiSelection *goquery.Selection) {
				node := model.Chapter{
					MangaId: manga.Id,
					Link:    MANHUAGUI_URL + LiSelection.Find("a").First().AttrOr("href", ""),
					Title:   LiSelection.Find("a").First().AttrOr("title", ""),
				}
				resBack = append(resBack, node)
			})
			for i3 := len(resBack)-1; i3 >= 0; i3-- {
				result = append([]model.Chapter{resBack[i3]}, result...)
			}
			resBack = []model.Chapter{}
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
			chromedp.WaitVisible("#mangaFile", chromedp.ByQuery),
			chromedp.AttributeValue(`#mangaFile`, "src", &node.Src, nil, chromedp.ByQuery),
		)
		result = append(result, node)
	}

	return
}