package resolver

import (
	"github.com/PuerkitoBio/goquery"
	"labourrobot/model"
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
		column2, _ := s.Find("td").Eq(1).Html()
		column4, _ := s.Find("td").Eq(3).Html()
		column5, _ := s.Find("td").Eq(4).Html()

		data := model.TaxInsuranceMetaData{column2, typeID, column4, 0.00, 0.00, column5}
		taxInsuranceMetaDataList = append(taxInsuranceMetaDataList, data)
	})

	return taxInsuranceMetaDataList, nil
}
