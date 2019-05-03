package main

import (
	"fmt"
	"github.com/moonprism/kindleM/site/manhuagui"
)

func main() {

	result :=  manhuagui.Search("剑风")

	chapters := manhuagui.ChapterList(&result[0])

	fmt.Printf("%v", chapters)
}
