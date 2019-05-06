package util

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

var HttpClient *http.Client
func init () {
	HttpClient = &http.Client{
		Timeout:   33 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func SetProxy(host string, port string) (err error) {
	dialer, err := proxy.SOCKS5("tcp", host+":"+port,
		nil,
		&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	)
	if err != nil {
		return
	}

	HttpClient.Transport = &http.Transport{
		Proxy:               nil,
		Dial:                dialer.Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return
}

func GetFetchDocument(url string) (doc *goquery.Document, err error) {
	req, err := http.NewRequest("GET", url, nil)

	res, err := HttpClient.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err = res.Body.Close()
	}()

	doc, err = goquery.NewDocumentFromReader(res.Body)
	return
}

func DownloadPicture(url string, referer string, fileName string) error {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Referer", referer)
	res, err := HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err = res.Body.Close()
	}()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, res.Body)
	return err
}
