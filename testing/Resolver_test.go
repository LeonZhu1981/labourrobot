package testing

import (
	_ "github.com/PuerkitoBio/goquery"
	"labourrobot/resolver"
	"testing"
)

func Test_NewResolver_PensionRate_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.PensionRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("单位缴费:12%<br>个人缴费:8%")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 12.0 || personalRate != 8.0 {
			t.Error("test failed")
		}
	}
}

func Test_NewResolver_PensionBase_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.PensionBaseTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("缴费基数:5223元<br>缴费基数(上下限):上限15669元、下限2089元。")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 15669.0 || personalRate != 2089.0 {
			t.Error("test failed")
		}
	}
}
