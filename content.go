package textractor

import (
	"math"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type textDensity struct {
	density float64
	tiText  string
	ti      int
	lti     int
	tgi     int
	ltgi    int
}

type nodeInfo struct {
	density      textDensity
	node         *goquery.Selection
	sbdi         float64
	textTagCount int
	score        float64
}

type infos []*nodeInfo

func (in infos) Len() int {
	return len(in)
}

func (in infos) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in infos) Less(i, j int) bool {
	return in[i].score > in[j].score
}

// contentExtract 提取正文
func contentExtract(body *goquery.Selection) *nodeInfo {
	var info infos
	for _, v := range iterator(body) {
		density := calcTextDensity(v)
		textTagCount := countTextTag(v)
		sbdi := calcSbdi(density)
		info = append(info, &nodeInfo{
			density:      density,
			node:         v,
			sbdi:         sbdi,
			textTagCount: textTagCount,
		})
	}
	std := calcDensityStd(info)
	calcNewScore(info, std)
	sort.Sort(info)
	return info[0]
}

// calcTextDensity 计算文本密度
//        Ti - LTi
// TDi = -----------
//       TGi - LTGi
// Ti:节点 i 的字符串字数
// LTi：节点 i 的带链接的字符串字数
// TGi：节点 i 的标签数
// LTGi：节点 i 的带连接的标签数
func calcTextDensity(s *goquery.Selection) textDensity {
	tiText := strings.Join(getAllTiText(s), "\n")
	ti := utf8.RuneCountInString(tiText)
	lti := utf8.RuneCountInString((strings.Join(getAllLtiText(s), "\n")))
	tgi := s.Find("*").Length()
	ltgi := s.Find("a").Length()
	var density float64 = 0
	if tgi-ltgi != 0 {
		density = float64(ti-lti) / float64(tgi-ltgi)
	}
	result := textDensity{
		density: density,
		tiText:  tiText,
		ti:      ti,
		lti:     lti,
		tgi:     tgi,
		ltgi:    ltgi,
	}
	return result
}

func getAllTiText(s *goquery.Selection) []string {
	var result []string
	for _, v := range iterator(s) {
		if v.Get(0).Type == html.TextNode {
			text := v.Text()
			text = regexp.MustCompile(`/[\n\r]/g`).ReplaceAllString(text, " ")
			text = regexp.MustCompile(`/\s{1,}/g`).ReplaceAllString(text, " ")
			text = strings.TrimSpace(text)
			if len(text) > 0 {
				result = append(result, text)
			}
		}
	}
	return result
}

func getAllLtiText(s *goquery.Selection) []string {
	var result []string
	for _, v := range iterator(s) {
		if goquery.NodeName(s) == "a" {
			text := v.Text()
			text = regexp.MustCompile(`/[\n\r]/g`).ReplaceAllString(text, " ")
			text = regexp.MustCompile(`/\s{1,}/g`).ReplaceAllString(text, " ")
			text = strings.TrimSpace(text)
			if len(text) > 0 {
				result = append(result, text)
			}
		}
	}
	return result
}

func countTextTag(s *goquery.Selection) int {
	return s.Find("p").Length()
}

// calcSbdi 计算符号密度
//          Ti - LTi
// SbDi = --------------
//          Sbi + 1
// SbDi: 符号密度
// Sbi：符号数量
func calcSbdi(density textDensity) float64 {
	sbi := countPunctuation(density.tiText)
	sbdi := float64(density.ti-density.lti) / float64(sbi+1)
	if sbdi <= 0 {
		sbdi = 1
	}
	return sbdi
}

func countPunctuation(text string) int {
	result := 0
	punctuation := `！，。？、；：“”‘’《》%（）,.?:;'"!%()`
	for _, v := range strings.Split(text, "") {
		if strings.ContainsAny(v, punctuation) {
			result++
		}
	}
	return result
}

// calcDensityStd 密度标准差
func calcDensityStd(info infos) float64 {
	var score []float64
	for _, v := range info {
		score = append(score, v.density.density)
	}
	std := std(score)
	return std
}

// calcNewScore 计算得分
//score = log(std) * ndi * log10(textTagCount + 2) * log(sbdi)
//std：每个节点文本密度的标准差
//ndi：节点 i 的文本密度
//textTagCount: 正文所在标签数。例如正文在<p></p>标签里面，这里就是 p 标签数。（目前把所有 p 内的 span、div 转成了 p 标签）
//sbdi：节点 i 的符号密度
func calcNewScore(info infos, std float64) {
	for _, v := range info {
		v.score = math.Log(std) * v.density.density * math.Log10(float64(v.textTagCount+2)) * math.Log(v.sbdi)
	}
}
