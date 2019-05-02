// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// run task list
	var downloadSrc string
	var pagesHtml string

	var chapterHtml string

	var chapter int

	_ = chromedp.Run(ctx,
		chromedp.Navigate(`https://manhua.dmzj.com/xingyebishangyan`),

		chromedp.InnerHTML(`body > div.wrap > div.middleright > div:nth-child(1) > div.cartoon_online_border`, &chapterHtml, chromedp.ByQuery),
	)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(chapterHtml))
	doc.Find("li").Each(func(i int, selection *goquery.Selection) {
		chapter++

		ref, _ := selection.Find("a").First().Attr("href")
		ref = "https://manhua.dmzj.com" + ref
		println(ref)
		_ = chromedp.Run(ctx,
			chromedp.Navigate(ref),
		)

		currentPage := 0

		for {
			_ = chromedp.Run(ctx,
				chromedp.WaitVisible(`#center_box > img`, chromedp.ByQuery),
				chromedp.AttributeValue(`#center_box > img`, "src", &downloadSrc, nil, chromedp.ByQuery),
				chromedp.Click(`#center_box > a.img_land_next`, chromedp.ByQuery),
				chromedp.InnerHTML(`#page_select`, &pagesHtml, chromedp.ByQuery),
			)

			currentPage++

			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(pagesHtml))
			countPage := 0
			doc.Find("option").Each(func(i int, selection *goquery.Selection) {
				countPage++
			})

			println(currentPage, "/", countPage)

			if currentPage == countPage {
				break
			}

			downloadSrc = "https:" + downloadSrc
			downloadDMZJ(downloadSrc, ref, fmt.Sprintf("%03d_%03d_", chapter, currentPage))
		}
	})

}

func downloadDMZJ(url string, ref string, prefix string) {
	println(url, ref, prefix)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Referer", ref)
	res, _ := client.Do(req)

	file, _ := os.Create(prefix + path.Base(url))
	io.Copy(file, res.Body)
}
