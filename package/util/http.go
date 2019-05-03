package util

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
)

func GetFetchDocument(url string) (doc *goquery.Document, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer func() {
		err = res.Body.Close()
	}()

	doc, err = goquery.NewDocumentFromReader(res.Body)
	return
}

func DownloadPicture(url string, referer string, fileNmae string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Referer", referer)
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer func() {
		err = res.Body.Close()
	}()

	file, err := os.Create(fileNmae)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, res.Body)
	return err
}
