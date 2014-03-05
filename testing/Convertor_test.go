package testing

import (
	"labourrobot/convertor"
	"labourrobot/model"
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
)

func Test_BuildNonGroupedInsuranceParameter(t *testing.T) {
	typeID := "13"
	metaData := model.TaxInsuranceMetaData{CityName: "广州", TypeID: typeID, DataBody: "", Data1: 0.20, Data2: 0.80, EffectTime: "2013/7/1"}
	propertyName := model.LookupTypeIdList[typeID]
	ret := convertor.BuildNonGroupedInsuranceParameter(metaData, propertyName)
	if ret.PensionRateByCompany != 0.20 {
		t.Errorf("Test BuildNonGroupedInsuranceParameter failed. the actual value is: %f", ret.PensionRateByCompany)
	}
}

func Test_BuildGroupedInsuranceParameter(t *testing.T) {
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

	ret := convertor.BuildGroupedInsuranceParameter(metaDataList)
	if len(ret) != 3 {
		t.Errorf("The result slice len is not 3. the actual value is: %f", len(ret))
	}
	if ret[0].CityName != "广州" {
		t.Errorf("The result city name is not 广州. the actual value is: %s", ret[0].CityName)
	}
	if ret[0].PensionRateByCompany != 0.20 || ret[0].PensionRateByIndividual != 0.10 {
		t.Errorf("PensionRateByCompany or PensionRateByIndividual. the actual value is: %f or %f", ret[0].PensionRateByCompany, ret[0].PensionRateByIndividual)
	}
}
