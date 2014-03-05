package convertor

import (
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"labourrobot/model"
	"reflect"
	"strings"
)

const (
	CompanySuffix    = "ByCompany"
	IndividualSuffix = "ByIndividual"
	MaxSuffix        = "Max"
	MinSuffix        = "Min"
)

func BuildGroupedInsuranceParameter(metaDataList []*model.TaxInsuranceMetaData) []model.InsuranceParameter {
	var completedInsuranceParameters []*model.InsuranceParameter
	var groupedInsuranceParameters []model.InsuranceParameter

	for k, v := range model.LookupTypeIdList {
		results, err := From(metaDataList).Where(func(t T) (bool, error) {
			return t.(*model.TaxInsuranceMetaData).TypeID == k, nil
		}).Select(func(t T) (T, error) {
			return t.(*model.TaxInsuranceMetaData), nil
		}).Results()

		if err != nil {
			panic(err.Error())
		}

		for _, result := range results {
			metatData, _ := result.(model.TaxInsuranceMetaData)
			item := BuildNonGroupedInsuranceParameter(metatData, v)
			completedInsuranceParameters = append(completedInsuranceParameters, item)
		}
	}

	allCityNames, err := From(completedInsuranceParameters).DistinctBy(func(t1 T, t2 T) (bool, error) {
		return t1.(*model.InsuranceParameter).CityName == t2.(*model.InsuranceParameter).CityName, nil
	}).Select(func(t T) (T, error) {
		return t.(*model.InsuranceParameter).CityName, nil
	}).Results()

	if err != nil {
		panic(err.Error())
	}

	for _, city := range allCityNames {
		results, err := From(completedInsuranceParameters).Where(func(t T) (bool, error) {
			return t.(*model.InsuranceParameter).CityName == city, nil
		}).Select(func(t T) (T, error) {
			return t.(*model.InsuranceParameter), nil
		}).Results()

		if err != nil {
			panic(err.Error())
		}

		var groupedInsuranceParameter model.InsuranceParameter = model.InsuranceParameter{}
		groupedInsuranceParameter.CityName = city.(string)

		for _, result := range results {
			item := result.(*model.InsuranceParameter)
			groupedInsuranceParameter.EffectTime = item.EffectTime

			val := reflect.ValueOf(item).Elem()
			for i := 0; i < val.NumField(); i++ {
				if val.Type().Name() == "float32" {
					fieldValue := val.Field(i).Interface()
					if fieldValue != 0.00 {
						fieldName := val.Type().Field(i).Name
						reflect.ValueOf(&groupedInsuranceParameter).FieldByName(fieldName).SetFloat(float64(fieldValue.(float32)))
					}
				}
			}
		}

		groupedInsuranceParameters = append(groupedInsuranceParameters, groupedInsuranceParameter)
	}

	return groupedInsuranceParameters
}

func BuildNonGroupedInsuranceParameter(metaData model.TaxInsuranceMetaData, propertyName string) *model.InsuranceParameter {
	var result *model.InsuranceParameter = &model.InsuranceParameter{}
	var fullPropertyName1, fullPropertyName2 string
	if strings.Index(propertyName, "Rate") > -1 {
		fullPropertyName1 = fmt.Sprintf(propertyName, CompanySuffix)
		fullPropertyName2 = fmt.Sprintf(propertyName, IndividualSuffix)
	} else {
		fullPropertyName1 = fmt.Sprintf(propertyName, MaxSuffix)
		fullPropertyName2 = fmt.Sprintf(propertyName, MinSuffix)
	}

	result.CityName = metaData.CityName
	result.EffectTime = metaData.EffectTime
	reflect.ValueOf(result).Elem().FieldByName(fullPropertyName1).SetFloat(float64(metaData.Data1))
	reflect.ValueOf(result).Elem().FieldByName(fullPropertyName2).SetFloat(float64(metaData.Data2))

	return result
}
