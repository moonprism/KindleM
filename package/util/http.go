package util

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"time"
)

func newClient() *http.Client {
	return &http.Client{
		Timeout:   15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func GetFetchDocument(url string) (doc *goquery.Document, err error) {
	req, err := http.NewRequest("GET", url, nil)

	res, err := newClient().Do(req)
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
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Referer", referer)
	res, err := newClient().Do(req)

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
