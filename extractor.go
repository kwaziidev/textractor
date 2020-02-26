package textractor

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ignoreTag 需要忽略的标签
var ignoreTag []string = []string{"style", "script", "link", "video", "iframe", "source", "picture", "header", "noscript"}

// ignoreClass 需要忽略的class, 当标签的class属性中包含下列值时, 忽略该标签
var ignoreClass []string = []string{"share", "contribution", "copyright", "copy-right", "disclaimer", "recommend", "related", "footer", "comment", "social", "submeta", "report-infor"}

// canBeRemoveIfEmpty 可以移除的空标签
var canBeRemoveIfEmpty []string = []string{"section", "h1", "h2", "h3", "h4", "h5", "h6", "span"}

// Text 提取的文本信息
type Text struct {
	// Title 标题
	Title string `json:"title,omitempty"`
	// Image 图片
	Image []string `json:"image,omitempty"`
	// Author 作者
	Author string `json:"author,omitempty"`
	// PublishTime 发布时间
	PublishTime string `json:"publish_time,omitempty"`
	// Content 正文
	Content string `json:"content,omitempty"`
	// ContentHTML 正文源码
	ContentHTML string `json:"content_html,omitempty"`
}

// Extract 提取信息
func Extract(source string) (*Text, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(source))
	if err != nil {
		return nil, err
	}
	body := dom.Find("body")
	normalize(body)
	content := contentExtract(body)
	publishTime := timeExtract(body)
	author := authorExtract(body)
	title := titleExtract(dom.Selection, content.node)
	result := &Text{}
	result.Content = content.density.tiText
	result.ContentHTML, _ = content.node.Html()
	var imgs []string
	content.node.Find("img").Each(func(i int, s *goquery.Selection) {
		if src, ok := s.Attr("src"); ok {
			imgs = append(imgs, src)
		}
	})
	result.Image = imgs
	result.PublishTime = publishTime
	result.Author = author
	result.Title = title
	return result, nil
}

// normalize 初始化节点
func normalize(element *goquery.Selection) {
	for _, v := range ignoreTag {
		element.Find(v).Remove()
	}
	for _, v := range iterator(element) {
		tagName := goquery.NodeName(v)
		// 删除注释
		if tagName == "#comment" {
			v.Remove()
			continue
		}
		// 删除一些可以删除的空标签
		if canBeRemove(v) {
			v.Remove()
			continue
		}
		// 删除标签class中含有ignoreClass的标签
		if val, ok := v.Attr("class"); ok {
			for _, class := range ignoreClass {
				if strings.Contains(val, class) {
					v.Remove()
					continue
				}
			}
		}
		// 去除p标签中的span, strong, em, b
		if tagName == "p" {
			v.Find("span,strong,em,b").Each(func(i int, child *goquery.Selection) {
				text := child.Text()
				child.ReplaceWithHtml(text)
			})
		}
		// 将没有子节点的div转换为p
		if tagName == "div" && v.Children().Length() <= 0 {
			v.Get(0).Data = "p"
		}
		// 去除空的p, 因上一步此处必须重新获取tagName
		if goquery.NodeName(v) == "p" {
			if v.Children().Length() <= 0 && len(strings.TrimSpace(v.Text())) == 0 {
				v.Remove()
			}
		}
	}
}

// iterator 遍历所有节点
func iterator(s *goquery.Selection) []*goquery.Selection {
	var result []*goquery.Selection
	iteratorNode(s, func(child *goquery.Selection) {
		result = append(result, child)
	})
	return result
}

func iteratorNode(s *goquery.Selection, fn func(*goquery.Selection)) {
	if s == nil {
		return
	}
	fn(s)
	s.Contents().Each(func(i int, c *goquery.Selection) {
		iteratorNode(c, fn)
	})
}

// canBeRemove 判断节点是否可以移除
func canBeRemove(s *goquery.Selection) bool {
	for _, v := range canBeRemoveIfEmpty {
		if strings.ToLower(goquery.NodeName(s)) == v {
			if s.Children().Length() <= 0 && strings.TrimSpace(s.Text()) == "" {
				return true
			}
		}
	}
	return false
}
