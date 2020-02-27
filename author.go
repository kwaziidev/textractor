package textractor

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

const patternSuffix = `[：|:| |丨|/]\s*([\p{Han}]{2,5}|[\w]{2,20})`

var authorPattern []string = []string{
	`责编` + patternSuffix,
	`责任编辑` + patternSuffix,
	`作者` + patternSuffix,
	`编辑` + patternSuffix,
	`原创` + patternSuffix,
	`文` + patternSuffix,
	`撰文` + patternSuffix,
	`来源` + patternSuffix,
	`稿件` + patternSuffix,
	`发稿人` + patternSuffix,
	`投稿人` + patternSuffix,
	`投稿` + patternSuffix,
	`来自` + patternSuffix,
}

// authorExtract 提取文章作者
func authorExtract(body *goquery.Selection) string {
	var text []string
	for _, v := range iterator(body) {
		if goquery.NodeName(v) == "#text" {
			t := strings.TrimSpace(v.Text())
			length := utf8.RuneCountInString(t)
			if t != "" && length >= 4 && length <= 25 {
				text = append(text, t)
			}
		}
	}
	for _, t := range text {
		if author, ok := matchAuthor(t); ok {
			return author
		}
	}
	return ""
}

func matchAuthor(text string) (string, bool) {
	for _, v := range authorPattern {
		ok, err := regexp.MatchString(v, text)
		if err == nil && ok {
			re, _ := regexp.Compile(v)
			return re.ReplaceAllString(re.FindString(text), "$1"), true
		}
	}
	return "", false
}
