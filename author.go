package textractor

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

const patternSuffix = `[：|:| |丨|/]\s*([\p{Han}]{2,5}|[\w]{2,20})`

var authorPattern = []string{
	`责编`,
	`责任编辑`,
	`作者`,
	`编辑`,
	`原创`,
	`文`,
	`撰文`,
	`来源`,
	`稿件`,
	`发稿人`,
	`投稿人`,
	`投稿`,
	`来自`,
}

var authorPatternRx []*regexp.Regexp

func init() {

	if _, err := regexp.Compile(patternSuffix); err != nil {
		panic(err)
	}

	for _, v := range authorPattern {
		if rx, err := regexp.Compile(v + patternSuffix); err == nil {
			authorPatternRx = append(authorPatternRx, rx)
		}
	}
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
	for _, rx := range authorPatternRx {
		if rx.MatchString(text) {
			return rx.ReplaceAllString(rx.FindString(text), "$1"), true
		}
	}
	return "", false
}
