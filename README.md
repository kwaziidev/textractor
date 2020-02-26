# textractor
从html文本中提取标题,正文,图片,作者,时间等信息,适用于新闻类网页  


## 安装
```
    go get github.com/gloomyzerg/textractor
```
## 使用

```golang
package main

import (
    "io/ioutil"
	"log"
	"net/http"

	"github.com/gloomyzerg/textractor"
)

func main(){
    url := "http://www.xxx.com/xxx"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	source, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
    }
    // 这只是一个例子
    // textractor.Extract 接收一个html的字符串
    // 可根据需求自行选择如何获取一个html字符串
    // 例如带分页的页面,可自行获取所有分页内容,拼接后传入
    result, _ := textractor.Extract(string(source))
    fmt.Printf("%+v", result)
}
```

## 命令行使用
```
    go get -u github.com/gloomyzerg/textractor/cmd/...
```
```
    textractor [url]
```

## 说明
textractor使用的[《基于文本及符号密度的网页正文提取方法》](https://kns.cnki.net/KCMS/detail/detail.aspx?dbcode=CJFQ&dbname=CJFDLAST2019&filename=GWDZ201908029&v=MDY4MTRxVHJXTTFGckNVUkxPZmJ1Wm5GQ2poVXJyQklqclBkTEc0SDlqTXA0OUhiWVI4ZVgxTHV4WVM3RGgxVDM=)对于一般的中文新闻类网页有较高的准确率,根据论文结论可知准确率高达99%以上.但由于样本条件限制作者并未测试足够多的样本来验证准确率.  
由于网页代码的多样性,任何提取算法都不可能覆盖所有网页.如遇到不能正确提取的网页,欢迎在issue中留下网页地址,具体问题具体分析.作者尽可能的去完善,以覆盖更多的页面.

textractor 命令行是为了方便测试和调试使用, 只是简单的 `wget + extract` , 并不能解析由js生成的动态页面, 动态页面可自行选择使用合适的解析办法.

## 感谢
本项目受到 [github.com/kingname/GeneralNewsExtractor](https://github.com/kingname/GeneralNewsExtractor) 的启发,并参考使用了它的测试用例用进行开发和测试
