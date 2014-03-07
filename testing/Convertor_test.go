package testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"labourrobot/convertor"
	"labourrobot/model"
	"reflect"
	"testing"
)

var (
	LookupTypeIdList = map[string]string{
		"17": "Pension%sBase",
		"13": "PensionRate%s",
		"18": "Jobless%sBase",
		"8":  "JoblessRate%s",
		"19": "Medical%sBase",
		"11": "MedicalRate%s",
		"20": "WorkInjury%sBase",
		"6":  "WorkInjuryRate%s",
		"21": "Maternity%sBase",
		"3":  "MaternityRate%s",
		"22": "HousingFund%sBase",
		"14": "HousingFundRate%s",
	}

	defaultData = model.InsuranceParameter{
		ProvinceName:                "全国平均",
		ProvinceShortName:           "QuanGuo",
		CityName:                    "全国平均",
		CityShortName:               "QuanGuo",
		EffectTime:                  "2013/7/1",
		PensionRateByCompany:        0.2,
		PensionRateByIndividual:     0.08,
		PensionMaxBase:              15669,
		PensionMinBase:              2089,
		JoblessRateByCompany:        0.01,
		JoblessRateByIndividual:     0.002,
		JoblessMaxBase:              15669,
		JoblessMinBase:              2089,
		MedicalRateByCompany:        0.1,
		MedicalRateByIndividual:     0.02,
		MedicalMaxBase:              15669,
		MedicalMinBase:              3134,
		WorkInjuryRateByCompany:     0.005,
		WorkInjuryRateByIndividual:  0,
		WorkInjuryMaxBase:           15669,
		WorkInjuryMinBase:           3134,
		MaternityRateByCompany:      0.008,
		MaternityRateByIndividual:   0,
		MaternityMaxBase:            15669,
		MaternityMinBase:            3134,
		HousingFundRateByCompany:    0.08,
		HousingFundRateByIndividual: 0.12,
		HousingFundMaxBase:          15669,
		HousingFundMinBase:          1400,
	}
)

func Test_BuildNonGroupedInsuranceParameter(t *testing.T) {
	typeID := "13"
	metaData := &model.TaxInsuranceMetaData{CityName: "广州", TypeID: typeID, DataBody: "", Data1: 0.20, Data2: 0.80, EffectTime: "2013/7/1"}
	propertyName := model.LookupTypeIdList[typeID]
	ret := convertor.BuildNonGroupedInsuranceParameter(metaData, propertyName)
	if ret.PensionRateByCompany != 0.20 {
		t.Errorf("Test BuildNonGroupedInsuranceParameter failed. the actual value is: %f", ret.PensionRateByCompany)
	}
}

func Test_BuildGroupedAndApplyProvinceInsuranceParameter(t *testing.T) {
	var provinceCityList []model.ProvinceCityList
	bytes, err := ioutil.ReadFile("../provincecityList.json")
	if err != nil {
		t.Errorf("read file failed: %s", err.Error())
	}
	json.Unmarshal(bytes, &provinceCityList)

	pensionBaseTypeID := "17"
	pensionRateTypeID := "13"
	joblessBaseTypeID := "18"
	joblessRateTypeID := "8"
	medicalBaseTypeID := "19"
	medicalRateTypeID := "11"
	workInjuryBaseTypeID := "20"
	workInjuryRateTypeID := "6"
	maternityBaseTypeID := "21"
	maternityRateTypeID := "3"
	housingFundBaseTypeID := "22"
	housingFundRateTypeID := "14"

	metaDataList := []*model.TaxInsuranceMetaData{
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: pensionRateTypeID, DataBody: "", Data1: 0.20, Data2: 0.10, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: pensionBaseTypeID, DataBody: "", Data1: 18000, Data2: 5000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: joblessRateTypeID, DataBody: "", Data1: 0.002, Data2: 0.002, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: joblessBaseTypeID, DataBody: "", Data1: 12000, Data2: 3000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: medicalRateTypeID, DataBody: "", Data1: 0.12, Data2: 0.12, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: medicalBaseTypeID, DataBody: "", Data1: 18000, Data2: 5000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: workInjuryRateTypeID, DataBody: "", Data1: 0.004, Data2: 0.00, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: workInjuryBaseTypeID, DataBody: "", Data1: 11000, Data2: 2000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: maternityRateTypeID, DataBody: "", Data1: 0.002, Data2: 0.00, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: maternityBaseTypeID, DataBody: "", Data1: 8000, Data2: 3500, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: housingFundRateTypeID, DataBody: "", Data1: 0.07, Data2: 0.07, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "广州", TypeID: housingFundBaseTypeID, DataBody: "", Data1: 6000, Data2: 2000, EffectTime: "2013/7/1"},

		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: pensionRateTypeID, DataBody: "", Data1: 0.20, Data2: 0.10, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: pensionBaseTypeID, DataBody: "", Data1: 18000, Data2: 5000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: joblessRateTypeID, DataBody: "", Data1: 0.002, Data2: 0.002, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: joblessBaseTypeID, DataBody: "", Data1: 12000, Data2: 3000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: medicalRateTypeID, DataBody: "", Data1: 0.12, Data2: 0.12, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: medicalBaseTypeID, DataBody: "", Data1: 18000, Data2: 5000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: workInjuryRateTypeID, DataBody: "", Data1: 0.004, Data2: 0.00, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: workInjuryBaseTypeID, DataBody: "", Data1: 11000, Data2: 2000, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: maternityRateTypeID, DataBody: "", Data1: 0.002, Data2: 0.00, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: maternityBaseTypeID, DataBody: "", Data1: 8000, Data2: 3500, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: housingFundRateTypeID, DataBody: "", Data1: 0.12, Data2: 0.12, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "成都", TypeID: housingFundBaseTypeID, DataBody: "", Data1: 6000, Data2: 2000, EffectTime: "2013/7/1"},

		&model.TaxInsuranceMetaData{CityName: "上海", TypeID: pensionRateTypeID, DataBody: "", Data1: 0.22, Data2: 0.11, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "上海", TypeID: pensionBaseTypeID, DataBody: "", Data1: 13000, Data2: 7500, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "上海", TypeID: joblessRateTypeID, DataBody: "", Data1: 0.04, Data2: 0.04, EffectTime: "2013/7/1"},
		&model.TaxInsuranceMetaData{CityName: "上海", TypeID: joblessBaseTypeID, DataBody: "", Data1: 8000, Data2: 2000, EffectTime: "2013/7/1"},
	}

	ret := convertor.BuildGroupedAndApplyProvinceInsuranceParameter(metaDataList, provinceCityList, defaultData)
	if len(ret) != 3 {
		t.Errorf("The result slice len is not 3. the actual value is: %f", len(ret))
	}
	if ret[0].CityName != "成都" {
		t.Errorf("The result city name is not 广州. the actual value is: %s", ret[0].CityName)
	}
	if ret[0].PensionRateByCompany != 0.20 || ret[0].PensionRateByIndividual != 0.10 {
		t.Errorf("PensionRateByCompany or PensionRateByIndividual. the actual value is: %f or %f", ret[0].PensionRateByCompany, ret[0].PensionRateByIndividual)
	}
	if ret[0].CityShortName != "ChengDu" {
		t.Errorf("CityShortName. the actual value is: %s", ret[0].CityShortName)
	}

	if ret[0].JoblessRateByCompany != 0.002 || ret[0].JoblessRateByIndividual != 0.002 {
		t.Errorf("JoblessRateByCompany or JoblessRateByIndividual. the actual value is: %f or %f", ret[0].JoblessRateByCompany, ret[0].JoblessRateByIndividual)
	}

	if ret[0].HousingFundRateByCompany != 0.12 || ret[0].HousingFundRateByIndividual != 0.12 {
		t.Errorf("HousingFundRateByCompany or HousingFundRateByIndividual. the actual value is: %f or %f", ret[0].HousingFundRateByCompany, ret[0].HousingFundRateByIndividual)
	}

	if ret[2].CityName != "上海" {
		t.Errorf("The result city name is not 上海. the actual value is: %s", ret[2].CityName)
	}
	if ret[2].PensionRateByCompany != 0.22 || ret[2].PensionRateByIndividual != 0.11 {
		t.Errorf("PensionRateByCompany or PensionRateByIndividual. the actual value is: %f or %f", ret[2].PensionRateByCompany, ret[2].PensionRateByIndividual)
	}

	if ret[2].JoblessMaxBase != 8000 || ret[2].JoblessMinBase != 2000 {
		t.Errorf("JoblessMaxBase or JoblessMinBase. the actual value is: %f or %f", ret[2].JoblessMaxBase, ret[2].JoblessMinBase)
	}
	assertHasZeroRate(ret, t)
}

func Test_ApplyMissingInsuranceParameter(t *testing.T) {
	testData := []model.InsuranceParameter{
		model.InsuranceParameter{ProvinceName: "四川", CityName: "成都", HousingFundRateByIndividual: 0.07, PensionRateByIndividual: 0.00, JoblessMaxBase: 0.00},
		model.InsuranceParameter{ProvinceName: "四川", CityName: "四川", HousingFundRateByIndividual: 0.06, PensionRateByIndividual: 1000, IsProvince: true},
	}

	result := convertor.ApplyMissingInsuranceParameter(testData, defaultData)

	if len(result) != 1 {
		t.Errorf("test failed: the length is: %d", len(result))
	}

	item := result[0]
	if item.HousingFundRateByIndividual != 0.07 {
		t.Errorf("test failed: the length is: %f", item.HousingFundRateByIndividual)
	}
	if item.PensionRateByIndividual != 1000 {
		t.Errorf("test failed: the length is: %f", item.PensionRateByIndividual)
	}
	if item.JoblessMaxBase != 15669 {
		t.Errorf("test failed: the length is: %f", item.JoblessMaxBase)
	}
	if item.JoblessMinBase != 2089 {
		t.Errorf("test failed: the length is: %f", item.JoblessMinBase)
	}
	assertHasZeroRate(result, t)
}

func assertHasZeroRate(testDataList []model.InsuranceParameter, t *testing.T) {
	for _, item := range testDataList {
		val := reflect.ValueOf(&item).Elem()
		for i := 0; i < val.NumField(); i++ {
			if val.Field(i).Kind().String() == "float32" {
				fieldValue := val.Field(i).Interface()
				rateValue, _ := fieldValue.(float32)
				fieldName := val.Type().Field(i).Name
				if rateValue == 0.00 && fieldName != "WorkInjuryRateByIndividual" && fieldName != "MaternityRateByIndividual" {
					t.Error(fmt.Sprintf("City Name:%s, fieldValue: %f, fieldName: %s", item.CityName, rateValue, fieldName))
				}
			}
		}
	}
}
