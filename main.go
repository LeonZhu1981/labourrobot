package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"labourrobot/convertor"
	"labourrobot/http"
	"labourrobot/model"
	"labourrobot/resolver"
	"os"
	"reflect"
	"strings"
)

func buildDefaultInsurceParamter() model.InsuranceParameter {
	var defaultData model.InsuranceParameter
	bytes, err := ioutil.ReadFile("defaultInsuranceData.json")

	if err != nil {
		fmt.Errorf("Read %s file error: %s", "defaultInsuranceData.json", err.Error())
	}

	if err := json.Unmarshal(bytes, &defaultData); err != nil {
		fmt.Errorf("Unmarshal json file error")
	}
	return defaultData
}

func persistenceFileByJsonFormat(t interface{}, fileName string) error {
	retByte, err := json.Marshal(t)
	if err != nil {
		return err
	}

	fout, err := os.Create(fileName)
	defer fout.Close()

	if err != nil {
		return err
	}

	fout.WriteString(string(retByte))
	return nil
}

func loadProvinceCityListFromFile() []model.ProvinceCityList {
	var provinceCityList []model.ProvinceCityList
	bytes, err := ioutil.ReadFile("provincecityList.json")

	if err != nil {
		fmt.Errorf("Read %s file error: %s", "provincecityList.json")
	}

	if err := json.Unmarshal(bytes, &provinceCityList); err != nil {
		fmt.Errorf("Unmarshal json file error")
	}
	return provinceCityList
}

func loadDataFromFile() []model.TaxInsurance {
	var data []model.TaxInsurance
	bytes, err := ioutil.ReadFile("TaxInsuranceMetaData.json")

	if err != nil {
		fmt.Errorf("Read %s file error: %s", "TaxInsuranceMetaData.json")
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Errorf("Unmarshal json file error")
	}
	return data
}

func main() {
	fmt.Println("Process starting......")
	var successfulInsuranceList []model.TaxInsuranceMetaData = []model.TaxInsuranceMetaData{}
	var copyedInsuranceList []*model.TaxInsuranceMetaData = []*model.TaxInsuranceMetaData{}
	var failedInsuranceList []model.TaxInsuranceMetaData = []model.TaxInsuranceMetaData{}
	var result model.TaxInsurance = model.TaxInsurance{}
	var convertedInsuranceList []model.InsuranceParameter = []model.InsuranceParameter{}

	for lookupTypeId, _ := range model.LookupTypeIdList {
		var copyedTaxInsuranceMetaDataList []*model.TaxInsuranceMetaData = []*model.TaxInsuranceMetaData{}
		fmt.Printf("Http request start...[type id]: %s\r\n", lookupTypeId)
		responseBody, err := http.RequestInsuranceDataByHttp(lookupTypeId)
		if err != nil {
			fmt.Errorf("Http Request error: %v, type id: %s", err, lookupTypeId)
		}
		fmt.Printf("Http request ending...[type id]: %s\r\n", lookupTypeId)

		fmt.Printf("Resolve html start...[type id]: %s\r\n", lookupTypeId)
		taxInsuranceMetaDataList, err := resolver.ResolveHtml(responseBody, lookupTypeId)
		if err != nil {
			fmt.Errorf("Resolve html error: %v, type id: %s", err, lookupTypeId)
		}
		fmt.Printf("Resolve html end...[type id]: %s\r\n", lookupTypeId)

		for _, val := range taxInsuranceMetaDataList {
			item := &model.TaxInsuranceMetaData{CityName: val.CityName, TypeID: val.TypeID, DataBody: val.DataBody, Data1: val.Data1, Data2: val.Data2, EffectTime: val.EffectTime, TypeName: val.TypeName, FailedReason: "", Year: val.Year}
			copyedTaxInsuranceMetaDataList = append(copyedTaxInsuranceMetaDataList, item)
		}

		taxInsuranceMetaDataList = taxInsuranceMetaDataList[:0]

		lastEffectiveTaxInsuranceMetaData := convertor.BuildLastEffectiveTaxInsuranceMetaDataList(copyedTaxInsuranceMetaDataList)
		for _, lastEffectItemPoint := range lastEffectiveTaxInsuranceMetaData {
			lastEffectItem := model.TaxInsuranceMetaData{CityName: lastEffectItemPoint.CityName, TypeID: lastEffectItemPoint.TypeID, DataBody: lastEffectItemPoint.DataBody, Data1: lastEffectItemPoint.Data1, Data2: lastEffectItemPoint.Data2, EffectTime: lastEffectItemPoint.EffectTime, TypeName: lastEffectItemPoint.TypeName, FailedReason: "", Year: lastEffectItemPoint.Year}
			taxInsuranceMetaDataList = append(taxInsuranceMetaDataList, lastEffectItem)
		}

		fmt.Printf("Resolve text start...[type id]: %s\r\n", lookupTypeId)
		for _, metaData := range taxInsuranceMetaDataList {
			r, err := resolver.NewResolver(lookupTypeId)
			if err != nil {
				metaData.FailedReason = fmt.Sprintf("NewResolver failed, type id:%s", lookupTypeId)
				failedInsuranceList = append(failedInsuranceList, metaData)
				continue
			}

			data1, data2, err := r.Resolve(metaData.DataBody)
			if err != nil {
				metaData.FailedReason = err.Error()
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
		fmt.Printf("Resolve text end...[type id]: %s\r\n", lookupTypeId)
	}

	for _, val := range successfulInsuranceList {
		item := &model.TaxInsuranceMetaData{CityName: val.CityName, TypeID: val.TypeID, DataBody: "", Data1: val.Data1, Data2: val.Data2, EffectTime: val.EffectTime}
		copyedInsuranceList = append(copyedInsuranceList, item)
	}

	fmt.Println("Convert text start...")
	provinceCityList := loadProvinceCityListFromFile()
	groupedInsuranceList := convertor.BuildGroupedAndApplyProvinceInsuranceParameter(copyedInsuranceList, provinceCityList, buildDefaultInsurceParamter())
	for _, item := range groupedInsuranceList {
		convertedInsuranceList = append(convertedInsuranceList, item)
	}
	fmt.Println("Convert text end...")

	result.ChinaTaxThreshold = 3500
	result.ForeignTaxThreshold = 4800
	result.InsuranceParameterList = convertedInsuranceList
	result.TaxParameterList = convertor.BuildTaxParameters()

	fmt.Println("Persistence file start")
	persistenceFileByJsonFormat(result, "TaxInsuranceMetaData.json")
	persistenceFileByJsonFormat(failedInsuranceList, "FailedData.json")
	fmt.Println("Persistence file end")
	buildCheckData(result)
	fmt.Println("Process completed. Please reference [TaxInsuranceMetaData.json] and [FailedData.json].")
}

func buildCheckData(data model.TaxInsurance) {
	insurceDataList := data.InsuranceParameterList
	var result []model.InsuranceParameter
	MaxSuffix := "Max"
	MinSuffix := "Min"

	for _, item := range insurceDataList {
		for _, propertyName := range model.LookupTypeIdList {
			if strings.Index(propertyName, "Base") > -1 {
				fullPropertyName1 := fmt.Sprintf(propertyName, MaxSuffix)
				fullPropertyName2 := fmt.Sprintf(propertyName, MinSuffix)
				maxValue := reflect.ValueOf(&item).Elem().FieldByName(fullPropertyName1).Float()
				minValue := reflect.ValueOf(&item).Elem().FieldByName(fullPropertyName2).Float()
				if minValue > maxValue || minValue > 100000 || maxValue > 100000 {
					result = append(result, item)
					break
				}
			}
		}
	}
	persistenceFileByJsonFormat(result, "needCheckThisResult.json")
}
