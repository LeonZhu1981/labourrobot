package convertor

import (
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/axgle/pinyin"
	"labourrobot/model"
	"reflect"
	//"sort"
	"strings"
)

const (
	CompanySuffix    = "ByCompany"
	IndividualSuffix = "ByIndividual"
	MaxSuffix        = "Max"
	MinSuffix        = "Min"
)

func BuildGroupedAndApplyProvinceInsuranceParameter(metaDataList []*model.TaxInsuranceMetaData, provinceCityList []model.ProvinceCityList) []model.InsuranceParameter {
	var completedInsuranceParameters []*model.InsuranceParameter
	var groupedInsuranceParameters []model.InsuranceParameter
	municipalityList := []string{"北京", "上海", "天津", "重庆"}

	for k, v := range model.LookupTypeIdList {
		results, err := From(metaDataList).Where(func(t T) (bool, error) {
			return t.(*model.TaxInsuranceMetaData).TypeID == k, nil
		}).Results()

		if err != nil {
			panic(err.Error())
		}

		for _, result := range results {
			metatData, _ := result.(*model.TaxInsuranceMetaData)
			item := BuildNonGroupedInsuranceParameter(metatData, v)
			if item.CityName == "达川" {
				item.CityName = "达州"
			}
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
		groupedInsuranceParameter.CityShortName = pinyin.Convert(groupedInsuranceParameter.CityName)

		if len(provinceCityList) > 0 {
			for _, provinceCity := range provinceCityList {
				if provinceCity.ProvinceName == groupedInsuranceParameter.CityName {
					groupedInsuranceParameter.ProvinceName = provinceCity.ProvinceName
					groupedInsuranceParameter.ProvinceShortName = pinyin.Convert(groupedInsuranceParameter.ProvinceName)
					groupedInsuranceParameter.IsProvince = true
					continue
				}
				for _, city := range provinceCity.CityList {
					if strings.Contains(city.CityName, groupedInsuranceParameter.CityName) {
						groupedInsuranceParameter.ProvinceName = provinceCity.ProvinceName
						groupedInsuranceParameter.ProvinceShortName = pinyin.Convert(groupedInsuranceParameter.ProvinceName)
						break
					}
				}
			}
			for _, municipality := range municipalityList {
				if municipality == groupedInsuranceParameter.ProvinceName {
					groupedInsuranceParameter.IsProvince = false
				}
			}
		}

		for _, result := range results {
			item := result.(*model.InsuranceParameter)
			groupedInsuranceParameter.EffectTime = item.EffectTime
			val := reflect.ValueOf(item).Elem()
			for i := 0; i < val.NumField(); i++ {
				if val.Field(i).Kind().String() == "float32" {
					fieldValue := val.Field(i).Interface()
					rateValue, _ := fieldValue.(float32)
					if rateValue == 0.00 {
						continue
					}
					fieldName := val.Type().Field(i).Name
					reflect.ValueOf(&groupedInsuranceParameter).Elem().FieldByName(fieldName).SetFloat(float64(rateValue))
				}
			}
		}
		if len(groupedInsuranceParameter.CityName) > 0 {
			groupedInsuranceParameters = append(groupedInsuranceParameters, groupedInsuranceParameter)
		}
	}
	groupedInsuranceParameters = OrderProvinceAndCity(Slice2PointSlice(groupedInsuranceParameters))
	return groupedInsuranceParameters
}

func BuildNonGroupedInsuranceParameter(metaData *model.TaxInsuranceMetaData, propertyName string) *model.InsuranceParameter {
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

func BuildTaxParameters() []model.TaxParameter {
	return []model.TaxParameter{
		model.TaxParameter{DeductAmount: 0, TaxRate: 0.03, MaxWholeMonthTaxAmount: 1500, MinWholeMonthTaxAmount: 0},
		model.TaxParameter{DeductAmount: 105, TaxRate: 0.10, MaxWholeMonthTaxAmount: 4500, MinWholeMonthTaxAmount: 1500},
		model.TaxParameter{DeductAmount: 555, TaxRate: 0.20, MaxWholeMonthTaxAmount: 9000, MinWholeMonthTaxAmount: 4500},
		model.TaxParameter{DeductAmount: 1005, TaxRate: 0.25, MaxWholeMonthTaxAmount: 35000, MinWholeMonthTaxAmount: 9000},
		model.TaxParameter{DeductAmount: 2755, TaxRate: 0.30, MaxWholeMonthTaxAmount: 55000, MinWholeMonthTaxAmount: 35000},
		model.TaxParameter{DeductAmount: 5505, TaxRate: 0.35, MaxWholeMonthTaxAmount: 80000, MinWholeMonthTaxAmount: 55000},
		model.TaxParameter{DeductAmount: 13505, TaxRate: 0.45, MaxWholeMonthTaxAmount: 9999999999, MinWholeMonthTaxAmount: 80000},
	}
}

func BuildLastEffectiveTaxInsuranceMetaDataList(metaDataList []*model.TaxInsuranceMetaData) []*model.TaxInsuranceMetaData {
	var result []*model.TaxInsuranceMetaData

	allCityNames, err := From(metaDataList).DistinctBy(func(t1 T, t2 T) (bool, error) {
		return t1.(*model.TaxInsuranceMetaData).CityName == t2.(*model.TaxInsuranceMetaData).CityName && len(strings.TrimSpace(t1.(*model.TaxInsuranceMetaData).CityName)) > 0, nil
	}).Select(func(t T) (T, error) {
		return t.(*model.TaxInsuranceMetaData).CityName, nil
	}).Results()

	if err != nil {
		panic(err.Error())
	}

	for _, city := range allCityNames {
		v, _, err := From(metaDataList).Where(func(t T) (bool, error) {
			return t.(*model.TaxInsuranceMetaData).CityName == city && t.(*model.TaxInsuranceMetaData).Year > 2009, nil
		}).OrderBy(func(this, that T) bool {
			return this.(*model.TaxInsuranceMetaData).Year > that.(*model.TaxInsuranceMetaData).Year
		}).First()

		if err != nil {
			panic(err.Error())
		}
		if v != nil {
			result = append(result, v.(*model.TaxInsuranceMetaData))
		}
	}
	return result
}

func OrderProvinceAndCity(metaDataList []*model.InsuranceParameter) []model.InsuranceParameter {
	var results []model.InsuranceParameter
	orderedList, _ := From(metaDataList).OrderBy(func(this, that T) bool {
		return this.(*model.InsuranceParameter).CityShortName < that.(*model.InsuranceParameter).CityShortName
	}).Results()
	// .OrderBy(func(this, that T) bool {
	// 	return this.(*model.InsuranceParameter).CityShortName < that.(*model.InsuranceParameter).CityShortName
	// })

	for _, p := range orderedList {
		item := p.(*model.InsuranceParameter)
		result := model.InsuranceParameter{
			CityName:                    item.CityName,
			CityShortName:               item.CityShortName,
			ProvinceName:                item.ProvinceName,
			ProvinceShortName:           item.ProvinceShortName,
			EffectTime:                  item.EffectTime,
			PensionRateByCompany:        item.PensionRateByCompany,
			PensionRateByIndividual:     item.PensionRateByIndividual,
			PensionMaxBase:              item.PensionMaxBase,
			PensionMinBase:              item.PensionMinBase,
			JoblessRateByCompany:        item.JoblessRateByCompany,
			JoblessRateByIndividual:     item.JoblessRateByIndividual,
			JoblessMaxBase:              item.JoblessMaxBase,
			JoblessMinBase:              item.JoblessMinBase,
			MedicalRateByCompany:        item.MedicalRateByCompany,
			MedicalRateByIndividual:     item.MedicalRateByIndividual,
			MedicalMaxBase:              item.MedicalMaxBase,
			MedicalMinBase:              item.MedicalMinBase,
			WorkInjuryRateByCompany:     item.WorkInjuryRateByCompany,
			WorkInjuryRateByIndividual:  item.WorkInjuryRateByIndividual,
			WorkInjuryMaxBase:           item.WorkInjuryMaxBase,
			WorkInjuryMinBase:           item.WorkInjuryMinBase,
			MaternityRateByCompany:      item.MaternityRateByCompany,
			MaternityRateByIndividual:   item.MaternityRateByIndividual,
			MaternityMaxBase:            item.MaternityMaxBase,
			MaternityMinBase:            item.MaternityMinBase,
			HousingFundRateByCompany:    item.HousingFundRateByCompany,
			HousingFundRateByIndividual: item.HousingFundRateByIndividual,
			HousingFundMaxBase:          item.HousingFundMaxBase,
			HousingFundMinBase:          item.HousingFundMinBase,
			IsProvince:                  item.IsProvince,
		}
		results = append(results, result)
	}
	return results
}

func PointSlice2Slice(metaDataList []*model.InsuranceParameter) []model.InsuranceParameter {
	var results []model.InsuranceParameter
	for _, item := range metaDataList {
		result := model.InsuranceParameter{
			CityName:                    item.CityName,
			CityShortName:               item.CityShortName,
			ProvinceName:                item.ProvinceName,
			ProvinceShortName:           item.ProvinceShortName,
			EffectTime:                  item.EffectTime,
			PensionRateByCompany:        item.PensionRateByCompany,
			PensionRateByIndividual:     item.PensionRateByIndividual,
			PensionMaxBase:              item.PensionMaxBase,
			PensionMinBase:              item.PensionMinBase,
			JoblessRateByCompany:        item.JoblessRateByCompany,
			JoblessRateByIndividual:     item.JoblessRateByIndividual,
			JoblessMaxBase:              item.JoblessMaxBase,
			JoblessMinBase:              item.JoblessMinBase,
			MedicalRateByCompany:        item.MedicalRateByCompany,
			MedicalRateByIndividual:     item.MedicalRateByIndividual,
			MedicalMaxBase:              item.MedicalMaxBase,
			MedicalMinBase:              item.MedicalMinBase,
			WorkInjuryRateByCompany:     item.WorkInjuryRateByCompany,
			WorkInjuryRateByIndividual:  item.WorkInjuryRateByIndividual,
			WorkInjuryMaxBase:           item.WorkInjuryMaxBase,
			WorkInjuryMinBase:           item.WorkInjuryMinBase,
			MaternityRateByCompany:      item.MaternityRateByCompany,
			MaternityRateByIndividual:   item.MaternityRateByIndividual,
			MaternityMaxBase:            item.MaternityMaxBase,
			MaternityMinBase:            item.MaternityMinBase,
			HousingFundRateByCompany:    item.HousingFundRateByCompany,
			HousingFundRateByIndividual: item.HousingFundRateByIndividual,
			HousingFundMaxBase:          item.HousingFundMaxBase,
			HousingFundMinBase:          item.HousingFundMinBase,
			IsProvince:                  item.IsProvince,
		}
		results = append(results, result)
	}
	return results
}

func Slice2PointSlice(metaDataList []model.InsuranceParameter) []*model.InsuranceParameter {
	var results []*model.InsuranceParameter
	for _, item := range metaDataList {
		result := &model.InsuranceParameter{
			CityName:                    item.CityName,
			CityShortName:               item.CityShortName,
			ProvinceName:                item.ProvinceName,
			ProvinceShortName:           item.ProvinceShortName,
			EffectTime:                  item.EffectTime,
			PensionRateByCompany:        item.PensionRateByCompany,
			PensionRateByIndividual:     item.PensionRateByIndividual,
			PensionMaxBase:              item.PensionMaxBase,
			PensionMinBase:              item.PensionMinBase,
			JoblessRateByCompany:        item.JoblessRateByCompany,
			JoblessRateByIndividual:     item.JoblessRateByIndividual,
			JoblessMaxBase:              item.JoblessMaxBase,
			JoblessMinBase:              item.JoblessMinBase,
			MedicalRateByCompany:        item.MedicalRateByCompany,
			MedicalRateByIndividual:     item.MedicalRateByIndividual,
			MedicalMaxBase:              item.MedicalMaxBase,
			MedicalMinBase:              item.MedicalMinBase,
			WorkInjuryRateByCompany:     item.WorkInjuryRateByCompany,
			WorkInjuryRateByIndividual:  item.WorkInjuryRateByIndividual,
			WorkInjuryMaxBase:           item.WorkInjuryMaxBase,
			WorkInjuryMinBase:           item.WorkInjuryMinBase,
			MaternityRateByCompany:      item.MaternityRateByCompany,
			MaternityRateByIndividual:   item.MaternityRateByIndividual,
			MaternityMaxBase:            item.MaternityMaxBase,
			MaternityMinBase:            item.MaternityMinBase,
			HousingFundRateByCompany:    item.HousingFundRateByCompany,
			HousingFundRateByIndividual: item.HousingFundRateByIndividual,
			HousingFundMaxBase:          item.HousingFundMaxBase,
			HousingFundMinBase:          item.HousingFundMinBase,
			IsProvince:                  item.IsProvince,
		}
		results = append(results, result)
	}
	return results
}
