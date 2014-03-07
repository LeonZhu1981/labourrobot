package resolver

import (
	"github.com/PuerkitoBio/goquery"
	"labourrobot/model"
	"strconv"
	"strings"
)

func ResolveHtml(htmlBody string, typeID string) (taxInsuranceMetaDataList []model.TaxInsuranceMetaData, err error) {
	var htmlTemplate = `<html>
		<head>
		</head>
		<body>
			<table id="datalist">
				{trdata}
			</table>
		</body>
	</html>`
	var trDatas = htmlBody
	htmlTemplate = strings.Replace(htmlTemplate, "{trdata}", trDatas, -1)
	ioreader := strings.NewReader(htmlTemplate)
	doc, err := goquery.NewDocumentFromReader(ioreader)
	if err != nil {
		return nil, err
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		cityName, _ := s.Find("td").Eq(1).Html()
		htmlText, _ := s.Find("td").Eq(3).Html()
		effectTime, _ := s.Find("td").Eq(4).Html()
		yearStr, _ := s.Find("td").Eq(2).Html()

		year, _ := strconv.Atoi(strings.Replace(yearStr, "年", "", -1))
		htmlText = strings.Replace(htmlText, "<br>", "<br/>", -1)
		htmlText = strings.Replace(htmlText, "线", "限", -1)

		data := model.TaxInsuranceMetaData{cityName, typeID, htmlText, 0.00, 0.00, effectTime, model.LookupTypeIdList[typeID], "", year}
		taxInsuranceMetaDataList = append(taxInsuranceMetaDataList, data)
	})

	return taxInsuranceMetaDataList, nil
}
