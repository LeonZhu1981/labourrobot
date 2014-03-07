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

		companyRate, personalRate, _ = r1.Resolve("单位缴费:12%12313<br/>个人缴费:8.1%4.512")
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

		companyRate, personalRate, _ = r1.Resolve("单位缴费:18%<br/>个人缴费:8%")

		if companyRate != 0.18 || personalRate != 0.08 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:20%（2013年没调整，沿用2012年的标准）<br/>个人缴费:8%（2013年没调整，沿用2012年的标准）")

		if companyRate != 0.20 || personalRate != 0.08 {
			t.Error("test failed")
		}
	}
}

func Test_NewResolver_CommonBase_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.PensionBaseTypeId)
	if err == nil {
		maxRate, minRate, err := r1.Resolve("缴费基数:2530元<br/>缴费基数(上下限):上限12780元、下限2530元")
		if err != nil {
			t.Error(err.Error())
		}

		if maxRate != 12780 || minRate != 2530 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1200元（2013年没调整，沿用2012年的标准）<br/>缴费基数(上下限):上限10119元、下限1200元（2013年没调整，沿用2012年的标准）")

		if maxRate != 10119 || minRate != 1200 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1050元<br/>缴费基数(上下限):")

		if maxRate != 0.00 || minRate != 0.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br/>缴费基数(上下限):下限1957元上限11289元")

		if maxRate != 11289 || minRate != 1957 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br/>缴费基数(上下限):2603")

		if maxRate != 2603 || minRate != 2603 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br/>缴费基数(上下限):26723（年）")

		if maxRate != 2226.92 || minRate != 2226.92 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br/>缴费基数(上下限):上限8932元 下限1787元")
		if maxRate != 8932 || minRate != 1787 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1600<br>缴费基数(上下限):下限1600上限4260")
		if maxRate != 4260 || minRate != 1600 {
			t.Log(maxRate)
			t.Log(minRate)
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:2080元（固定值）<br/>缴费基数(上下限):无上下限")
		if maxRate != 0 || minRate != 0 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1532元<br/>缴费基数(上下限):")
		if maxRate != 0 || minRate != 0 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:2327.4<br/>缴费基数(上下限):上限11637下限2327.4")
		if maxRate != 11637 || minRate != 2327.4 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:上限8195元、下限800元<br/>缴费基数(上下限):1967元")
		if maxRate != 8195 || minRate != 800 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:2600<br/>缴费基数(上下限):下限2600")
		if maxRate != 2600 || minRate != 2600 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1480元、1280元<br/>缴费基数(上下限):上限13940元、下限1480元、1280元")
		if maxRate != 13940 || minRate != 1480 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1200元<br/>缴费基数(上下限):无上限、下限1200元")
		if maxRate != 1200 || minRate != 1200 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:840元（2013年没调整，沿用2012年的标准）<br/>缴费基数(上下限):上限13499元、下限840元（2013年没调整，沿用2012年的标准）")
		if maxRate != 13499 || minRate != 840 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br/>缴费基数(上下限):下限800上限无")
		if maxRate != 800 || minRate != 800 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:以上半年工资平均基数为准<br/>缴费基数(上下限):")
		if maxRate != 0 || minRate != 0 {
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
		if maxRate != 0.00 || minRate != 0.00 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:2138元<br/>缴费基数(上下限):固定值2138元")
		if maxRate != 2138 || minRate != 2138 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br>缴费基数(上下限):最高标准将由2009年的1512元上调至1840元")
		if maxRate != 1840 || minRate != 1840 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br>缴费基数(上下限):单位和职工住房公积金月缴存额下限各为40.50元；市区单位和职工住房公积金月缴存额上限各为2392元，开平市、台山市、鹤山市及恩平市单位和职工住房公积金月缴存额上限各为2025元")
		if maxRate != 2392 || minRate != 40.50 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br>缴费基数(上下限):无最高上限，最低下限为1581元")
		if maxRate != 1581 || minRate != 1581 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:<br>缴费基数(上下限):下限800上限无")
		if maxRate != 800 || minRate != 800 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1294 <br>缴费基数(上下限):1413.6元-9090元 ")
		if maxRate != 9090 || minRate != 1413.6 {
			t.Error("test failed")
		}

		maxRate, minRate, _ = r1.Resolve("缴费基数:1083 <br/>缴费基数(上下限):1083~8277 ")

		if maxRate != 8277 || minRate != 1083 {
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

		companyRate, personalRate, _ = r1.Resolve("单位缴费:1%<br/>个人缴费:农村：不缴费 城镇：0.2%")

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

		companyRate, personalRate, _ = r1.Resolve("单位缴费:2%<br/>个人缴费:1%")

		if companyRate != 0.02 || personalRate != 0.01 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:1%<br/>个人缴费:")

		if companyRate != 0.01 || personalRate != 0.00 {
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:2%(2013年没调整，沿用2012年的标准)<br/>个人缴费:1%(2013年没调整，沿用2012年的标准)")

		if companyRate != 0.02 || personalRate != 0.01 {
			t.Error("test failed")
		}

	}
}

func Test_NewResolver_MedicalRate_Resolve(t *testing.T) {
	r1, err := resolver.NewResolver(resolver.MedicalRateTypeId)
	if err == nil {
		companyRate, personalRate, err := r1.Resolve("单位缴费:8%<br/>单位大额医疗费用互助资金:<br/>个人缴费:2%<br/>个人大额医疗费用互助资金:")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 0.08 || personalRate != 0.02 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("单位缴费:6%（2013年没调整，沿用2012年的标准）<br/>单位大额医疗费用互助资金:<br/>个人缴费:2%（2013年没调整，沿用2012年的标准）<br/>个人大额医疗费用互助资金:")
		if err != nil {
			t.Error(err.Error())
		}

		if companyRate != 0.06 || personalRate != 0.02 {
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

func Test_NewResolver_MaternityRate_Resolve(t *testing.T) {
	r1, _ := resolver.NewResolver(resolver.MaternityRateTypeId)
	companyRate, personalRate, _ := r1.Resolve("单位缴费:0.8%")
	if companyRate != 0.008 || personalRate != 0.00 {
		t.Error("test failed")
	}

	companyRate, personalRate, _ = r1.Resolve("单位缴费:0.7%")
	if companyRate != 0.007 || personalRate != 0.00 {
		t.Error("test failed")
	}

	companyRate, personalRate, _ = r1.Resolve("单位缴费:包含在医疗保险内，不单独缴纳生育保险")
	if companyRate != 0.00 || personalRate != 0.00 {
		t.Error("test failed")
	}

	companyRate, personalRate, _ = r1.Resolve("单位缴费:0.5%（2013年没调整，沿用2012年的标准）")
	if companyRate != 0.005 || personalRate != 0.00 {
		t.Error("test failed")
	}

	companyRate, personalRate, _ = r1.Resolve("单位缴费:05%")
	if companyRate != 0.05 || personalRate != 0.00 {
		t.Error("test failed")
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
			t.Error("test failed")
		}

		companyRate, personalRate, _ = r1.Resolve("单位缴费:0.25%、0.5%、0.75%（2013年没调整，沿用2012年标准）")

		if companyRate != 0.0075 || personalRate != 0.00 {
			t.Error("test failed")
		}
		companyRate, personalRate, _ = r1.Resolve("单位缴费:1.3%、2.1%、2.9%")

		if companyRate != 0.029 || personalRate != 0.00 {
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

		companyRate, personalRate, err = r1.Resolve("用人单位支付:6%~12%<br/>个人支付:6%~12%<br/>封顶－封底金额:封顶3515元、封底144元")
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

		companyRate, personalRate, err = r1.Resolve("用人单位支付:5%-12%<br/>个人支付:5%-12%<br/>封顶－封底金额:")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.05 || personalRate != 0.05 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:8%-12%<br/>个人支付:8%-12% <br/>封顶－封底金额:封顶2516元、封底321元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.08 || personalRate != 0.08 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:12%<br/>个人支付:10%<br/>封顶－封底金额:")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.12 || personalRate != 0.10 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:5%/8%/10%/12%<br/>个人支付:5%/8%/10%/12%<br/>封顶－封底金额:10014")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.05 || personalRate != 0.05 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:5%、8%、10%<br/>个人支付:5%、8%、10%<br/>封顶－封底金额:封顶9501元、封底131元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.05 || personalRate != 0.05 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:6%~12%<br/>个人支付:6%~12%<br/>封顶－封底金额:封顶3515元、封底144元")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.06 || personalRate != 0.06 {
			t.Error("test failed")
		}

		companyRate, personalRate, err = r1.Resolve("用人单位支付:12%（2013年没调整，沿用2012年的标准）<br/>个人支付:12%（2013年没调整，沿用2012年的标准）<br/>封顶－封底金额:")
		if err != nil {
			t.Error(err.Error())
		}
		if companyRate != 0.12 || personalRate != 0.12 {
			t.Error("test failed")
		}
	}
}
