package labourrobot

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func Test_GoQuery_1(t *testing.T) {
	var htmlTemplate = `<html>
<head>
</head>
<body>
	<table id="datalist">
		{trdata}
	</table>
</body>
</html>`
	var trDatas = `<tr ><td><input class="currentPage" type="hidden" value="2">养老保险缴费比例</td><td>济南</td><td>2014年</td><td>单位缴费:18%<br />个人缴费:8%</td><td>2014/1/1</td><td class="popfile"></td><td class="popfile"><a href='lawshow-95525.html'>关于降低企业职工基本养老保险单位缴费比例的通知(2014)</a></td></tr><tr class='even'><td>养老保险缴费比例</td><td>江西</td><td>2013年</td><td>单位缴费:20%<br />个人缴费:8%</td><td>2013/1/1</td><td class="popfile"></td><td class="popfile"></td></tr><tr ><td>养老保险缴费比例</td><td>安徽</td><td>2013年</td><td>单位缴费:20%<br />个人缴费:8%</td><td>2013/7/1</td><td class="popfile"></td><td class="popfile"></td></tr><tr class='even'><td>养老保险缴费比例</td><td>湖州</td><td>2013年</td><td>单位缴费:14%<br />个人缴费:8%</td><td>2013/7/1</td><td class="popfile"></td><td class="popfile"></td></tr><tr ><td>养老保险缴费比例</td><td>温州</td><td>2013年</td><td>单位缴费:14%<br />个人缴费:8%</td><td>2013/4/1</td><td class="popfile"></td><td class="popfile"><a href='lawshow-95476.html'>关于公布市区2013年度社会保险缴费标准的通知</a></td></tr><tr class='even'><td>养老保险缴费比例</td><td>上海</td><td>2013年</td><td>单位缴费:22%<br />个人缴费:8%</td><td>2013/1/1</td><td class="popfile"></td><td class="popfile"></td></tr><tr ><td>养老保险缴费比例</td><td>汕头</td><td>2013年</td><td>单位缴费:15%<br />个人缴费:8%</td><td>2013/7/1</td><td class="popfile"></td><td class="popfile"><a href='lawshow-95136.html'>关于调整我市2013至2014年度社会保险费征收和社会保险待遇标准的通知</a></td></tr><tr class='even'><td>养老保险缴费比例</td><td>重庆</td><td>2013年</td><td>单位缴费:17%<br />个人缴费:8%</td><td>2013/1/1</td><td class="popfile"><a href='newsshow-187908.html'>重庆社保缴费基数调为3783元 个人增19.58元</a></td><td class="popfile"></td></tr><tr ><td>养老保险缴费比例</td><td>北京</td><td>2013年</td><td>单位缴费:20%<br />个人缴费:8%</td><td>2013/7/1</td><td class="popfile"></td><td class="popfile"><a href='lawshow-95182.html'>关于统一2013年度各项社会保险缴费工资基数和缴费金额的通知</a></td></tr><tr class='even'><td>养老保险缴费比例</td><td>广州</td><td>2013年</td><td>单位缴费:12%<br />个人缴费:8%</td><td>2013/7/1</td><td class="popfile"></td><td class="popfile"><a href='lawshow-95175.html'>关于2013社保年度缴费基数调整的通知</a></td></tr>`
	htmlTemplate = strings.Replace(htmlTemplate, "{trdata}", trDatas, -1)
	ioreader := strings.NewReader(htmlTemplate)
	doc, err := goquery.NewDocumentFromReader(ioreader)
	if err != nil {
		t.Error(err.Error())
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("Content of 1 cell %d: %s\n", i, s.Find("td").Eq(0).Text())
		fmt.Printf("Content of 2 cell %d: %s\n", i, s.Find("td").Eq(1).Text())
		fmt.Printf("Content of 3 cell %d: %s\n", i, s.Find("td").Eq(2).Text())
		fmt.Printf("Content of 4 cell %d: %s\n", i, s.Find("td").Eq(3).Text())
		// doc.Find("td").Each(func(i int, s *goquery.Selection) {
		// 	fmt.Printf("Content of cell %d: %s\n", i, s.Text())
		// })
	})
}
