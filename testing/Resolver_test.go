package testing

import (
	_ "github.com/PuerkitoBio/goquery"
	"labourrobot/resolver"
	"testing"
)

func Test_NewResolver_CommonRate_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.PensionRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("单位缴费:12%12313<br>个人缴费:8.1%4.512")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.12 || personalRate != 0.081 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费不缴费哦!")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:12%<br>个人缴费:")

		if companyRate != 0.12 || personalRate != 0.00 {
			t.Error("test failed")
		}
	}
}

func Test_NewResolver_CommonBase_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.PensionBaseTypeId)
	if err == nil {
		maxRate, minRate, err := r1.Resolve("缴费基数:5223元<br>缴费基数(上下限):上限15669.32元、下限2089.33元。")
		if err != nil {
			t.Error(err.Error())
		}

		if maxRate != 15669.32 || minRate != 2089.33 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("单位缴费不缴费哦!")

		if maxRate != 0.00 || maxRate != 0.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("")

		if maxRate != 0.00 || minRate != 0.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):下限1140上限5760")

		if maxRate != 5760.00 || minRate != 1140.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):868.2~4341")
		if maxRate != 4341.00 || minRate != 868.20 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):最高缴费基数为7611元，最低缴费基数为1522元")
		if maxRate != 7611.00 || minRate != 1522.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):最高缴费基数为7611元,最低缴费基数为1522元")
		if maxRate != 7611.00 || minRate != 1522.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):固定值2138元")
		if maxRate != 0.00 || minRate != 0.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):1077")
		if maxRate != 0.00 || minRate != 0.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数(上下限):基数下限按不低于全省在岗职工平均工资的60%（1369元）确定，其中就业困难人员以个人形式参保缴费的暂按不低于1200元确定；月缴费工资基数上限不超过全省在岗职工平均工资的300%（6843元）确定")
		t.Log(maxRate)
		t.Log(minRate)
		if maxRate != 0.00 || minRate != 0.00 {
			t.Error("test failed")
		}
	}
}

func Test_NewResolver_JobLessRate_Special_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.JoblessRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("单位缴费:1%<br>个人缴费:农村：不缴费 城镇：0.2%")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 0.01 || personalRate != 0.002 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费不缴费哦!")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}
	}
}

func Test_NewResolver_MedicalRate_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.MedicalRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("单位缴费:8.0%<br>单位大额医疗费用互助资金:<br>个人缴费:2%<br>个人大额医疗费用互助资金:")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 0.08 || personalRate != 0.02 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费不缴费哦!")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}
	}
}

func Test_NewResolver_WorkInjuryRate_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.WorkInjuryRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("单位缴费:0.8%、1.6%、2.4")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 0.024 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("单位缴费:0.5%-2.5%")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 0.025 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("单位缴费:0.8%")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.008 || personalRate != 0.00 {
			t.Error("test failed")
		}

		if companyRate != 0.008 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:包含在医疗保险内，不单独缴纳生育保险")

		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:0.5%（2013年没调整，沿用2012年的标准）")

		if companyRate != 0.005 || personalRate != 0.00 {
			t.Log(companyRate)
			t.Error("test failed")
		}
	}
}

func Test_HousingFundRate_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.HousingFundRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("用人单位支付:6%~12%<br>个人支付:6%~12%<br>封顶－封底金额:封顶3515元、封底144元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.06 || personalRate != 0.06 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:5%、8%、10%<br>个人支付:5%、8%、10%<br>封顶－封底金额:封顶9501元、封底131元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.05 || personalRate != 0.05 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:12%<br>个人支付:12%<br>封顶－封底金额:封顶3385元、封底314元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.12 || personalRate != 0.12 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:5%-20%<br>个人支付:5%-20%<br>封顶－封底金额:上限7700元、下限131元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.05 || personalRate != 0.05 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("缴费基数:<br>缴费基数(上下限):最高不得超过市统计局公布的2008年度职工月平均工资1536元的3倍，即4608元；最低不能低于全市最低工资标准，即市区550元，两县450元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.00 || personalRate != 0.00 {
			t.Error("test failed")
		}
	}
}
