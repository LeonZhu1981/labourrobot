package main

import (
	"labourrobot/crawler"
	"labourrobot/model"
	"labourrobot/resolver"
)

func main() {
	var successfulInsuranceList []model.TaxInsuranceMetaData = []model.TaxInsuranceMetaData{}
	var failedInsuranceList []model.TaxInsuranceMetaData = []model.TaxInsuranceMetaData{}
	for _, lookupTypeId := range model.LookupTypeIdList {
		responseBody, _ := crawler.RequestInsuranceDataByHttp(lookupTypeId)
		taxInsuranceMetaDataList, _ := resolver.ResolveHtml(responseBody, lookupTypeId)
		for _, metaData := range taxInsuranceMetaDataList {
			r, err := resolver.NewResolver(lookupTypeId)
			if err != nil {
				failedInsuranceList = append(failedInsuranceList, metaData)
				continue
			}

			data1, data2, err := r.Resolve(metaData.DataBody)
			if err != nil {
				failedInsuranceList = append(failedInsuranceList, metaData)
				continue
			}

			metaData.Data1 = data1
			metaData.Data2 = data2
			if metaData.Data1 == 0.00 && metaData.Data2 == 0.00 {
				failedInsuranceList = append(failedInsuranceList, metaData)
				continue
			}

			metaData.DataBody = ""
			successfulInsuranceList = append(successfulInsuranceList, metaData)
		}
	}
}
