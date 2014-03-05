package model

var (
	LookupTypeIdList = []string{"17", "13", "18", "8", "19", "11", "20", "6", "21", "3", "22", "14"}
)

type TaxParameter struct {
	TaxRate                float32
	MaxWholeMonthTaxAmount float32
	MinWholeMonthTaxAmount float32
	DeductAmout            float32
}

type InsuranceParameter struct {
	ProvinceName                string
	ProvinceShortName           string
	CityName                    string
	CityShortName               string
	PensionRateByCompany        float32
	PensionRateByIndividual     float32
	PensionMaxBase              float32
	PensionMinBase              float32
	JoblessRateByCompany        float32
	JoblessRateByIndividual     float32
	JoblessMaxBase              float32
	JoblessMinBase              float32
	MedicalRateByCompany        float32
	MedicalRateByIndividual     float32
	MedicalMaxBase              float32
	MedicalMinBase              float32
	WorkInjuryRateByCompany     float32
	WorkInjuryRateByIndividual  float32
	WorkInjuryMaxBase           float32
	WorkInjuryMinBase           float32
	MaternityRateByCompany      float32
	MaternityRateByIndividual   float32
	MaternityMaxBase            float32
	MaternityMinBase            float32
	HousingFundRateByCompany    float32
	HousingFundRateByIndividual float32
	HousingFundMaxBase          float32
	HousingFundMinBase          float32
}

type TaxInsurance struct {
	ChinaTaxThreshold      float32
	ForeignTaxThreshold    float32
	InsuranceParameterList []InsuranceParameter
	TaxParameterList       []TaxParameter
}

type TaxInsuranceMetaData struct {
	CityName   string
	TypeID     string
	DataBody   string
	Data1      float32
	Data2      float32
	EffectTime string
}