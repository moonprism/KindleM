package model


type ProcessInfo struct {
	ProcessCount
}

type ProcessCount struct {
	Chromedp 	int `json:"chromedp"`
	Picture 	int `json:"picture"`
}