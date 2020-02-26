package textractor

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// titleExtract 提取文章作者
// source 网页源码
// content 正文
func titleExtract(source *goquery.Selection, content *goquery.Selection) string {
	var title string
	titleNode := source.Find("title")
	if titleNode.Length() > 0 {
		re, _ := regexp.Compile("-|_|\\|")
		title = re.Split(titleNode.Text(), -1)[0]
		if strings.TrimSpace(title) != "" {
			return title
		}
	}

	title = titleExtractFromHTML(content)
	return title

}

// titleExtractFromHTML 从html标签中查找
func titleExtractFromHTML(content *goquery.Selection) string {
	title, _ := findHtag(content)
	return title
}

func findHtag(content *goquery.Selection) (string, bool) {
	parent := content.Parent()
	htag := parent.Find("h1,h2,h3,h4")
	if htag.Length() > 0 {
		return htag.Eq(0).Text(), true
	}
	return findHtag(parent)
}
