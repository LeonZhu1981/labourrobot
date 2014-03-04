package resolver

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	PensionBaseTypeId     = 17
	PensionRateTypeId     = 13
	JoblessBaseTypeId     = 18
	JoblessRateTypeId     = 8
	MedicalBaseTypeId     = 19
	MedicalRateTypeId     = 11
	WorkInjuryBaseTypeId  = 20
	WorkInjuryRateTypeId  = 6
	MaternityBaseTypeId   = 21
	MaternityRateTypeId   = 3
	HousingFundBaseTypeId = 22
	HousingFundRateTypeId = 14

	BRSplitSymblo             = "<br>"
	CommaSymblo               = "、"
	DashSymblo                = "-"
	DigitalRegMatch           = `[0-9.%]`
	WorkInjuryDigitalRegMatch = `[0-9.%、-]`
	PercentSymblo             = "%"
)

type Resolver interface {
	Resolve(tag string) (float32, float32, error)
}

func NewResolver(typeId int32) (Resolver, error) {
	var resolver Resolver
	switch typeId {
	case PensionBaseTypeId:
		resolver = &PensionBaseResolver{}
		return resolver, nil
	case PensionRateTypeId:
		resolver = &PensionRateResolver{}
		return resolver, nil
	case JoblessBaseTypeId:
		resolver = &JoblessBaseResolver{}
		return resolver, nil
	case JoblessRateTypeId:
		resolver = &JoblessRateResolver{}
		return resolver, nil
	case MedicalBaseTypeId:
		resolver = &MedicalBaseResolver{}
		return resolver, nil
	case MedicalRateTypeId:
		resolver = &MedicalRateResolver{}
		return resolver, nil
	case WorkInjuryBaseTypeId:
		resolver = &WorkInjuryBaseResolver{}
		return resolver, nil
	case WorkInjuryRateTypeId:
		resolver = &WorkInjuryRateResolver{}
		return resolver, nil
	case MaternityBaseTypeId:
		resolver = &MaternityBaseResolver{}
		return resolver, nil
	case MaternityRateTypeId:
		resolver = &MaternityRateResolver{}
		return resolver, nil
	case HousingFundBaseTypeId:
		resolver = &HousingFundBaseResolver{}
		return resolver, nil
	case HousingFundRateTypeId:
		resolver = &HousingFundRateResolver{}
		return resolver, nil
	default:
		err := errors.New(fmt.Sprintf("Can not found type id:%d for NewResolver function", typeId))
		return nil, err
	}
}

func printTagIsEmptyError(tag string) (float32, float32, error) {
	return 0, 0, errors.New(fmt.Sprintf("Can not resolve the target string: %s cause tag is empty.", tag))
}

func printNotFoundSplitSymbloError(tag string, splitSymblo string) (float32, float32, error) {
	return 0, 0, errors.New(fmt.Sprintf("Can not found split string: %s in the target string: %s", splitSymblo, tag))
}

func commonBaseResolve(tag string) (maxBase float32, minBase float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	if strings.IndexAny(tag, BRSplitSymblo) == -1 {
		return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
	}

	arr := strings.Split(tag, BRSplitSymblo)
	maxMinBaseString := arr[1]

	if strings.IndexAny(maxMinBaseString, CommaSymblo) == -1 {
		return printNotFoundSplitSymbloError(tag, CommaSymblo)
	}

	maxMinBaseArr := strings.Split(maxMinBaseString, CommaSymblo)

	re := regexp.MustCompile(DigitalRegMatch)

	maxBaseString := strings.Join(re.FindAllString(maxMinBaseArr[0], -1), "")
	minBaseString := strings.Join(re.FindAllString(maxMinBaseArr[1], -1), "")

	fmaxBase, err := strconv.ParseFloat(maxBaseString, 32)
	if err != nil {
		fmaxBase = 0.00
	}

	fminBase, err := strconv.ParseFloat(minBaseString, 32)
	if err != nil {
		fminBase = 0.00
	}

	maxBase = float32(fmaxBase)
	minBase = float32(fminBase)

	return maxBase, minBase, nil
}

func commonRateResolve(tag string) (companyRate float32, personalRate float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	if strings.IndexAny(tag, BRSplitSymblo) == -1 {
		return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
	}

	arr := strings.Split(tag, BRSplitSymblo)
	companyRateString := arr[0]
	personalRateString := arr[1]
	re := regexp.MustCompile(DigitalRegMatch)

	companyRateString = strings.Join(re.FindAllString(companyRateString, -1), "")
	personalRateString = strings.Join(re.FindAllString(personalRateString, -1), "")

	if firstIndex := strings.Index(companyRateString, PercentSymblo); firstIndex > -1 {
		companyRateString = companyRateString[0:firstIndex]
	}

	fcompanyRate, err := strconv.ParseFloat(companyRateString, 32)

	if err != nil {
		fcompanyRate = 0.00
	}

	if firstIndex := strings.Index(personalRateString, PercentSymblo); firstIndex > -1 {
		personalRateString = personalRateString[0:firstIndex]
	}

	fpersonalRate, err := strconv.ParseFloat(personalRateString, 32)
	if err != nil {
		fpersonalRate = 0.00
	}

	companyRate = float32(fcompanyRate) / 100
	personalRate = float32(fpersonalRate) / 100
	return companyRate, personalRate, nil
}

type PensionBaseResolver struct {
}

func (r *PensionBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type PensionRateResolver struct {
}

func (r *PensionRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return commonRateResolve(tag)
}

type JoblessBaseResolver struct {
}

func (r *JoblessBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type JoblessRateResolver struct {
}

func (r *JoblessRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return commonRateResolve(tag)
}

type MedicalBaseResolver struct {
}

func (r *MedicalBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type MedicalRateResolver struct {
}

func (r *MedicalRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	if strings.IndexAny(tag, BRSplitSymblo) == -1 {
		return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
	}

	arr := strings.Split(tag, BRSplitSymblo)
	companyRateString := arr[0]
	personalRateString := arr[2]
	re := regexp.MustCompile(DigitalRegMatch)

	companyRateString = strings.Join(re.FindAllString(companyRateString, -1), "")
	if firstIndex := strings.Index(companyRateString, PercentSymblo); firstIndex > -1 {
		companyRateString = companyRateString[0:firstIndex]
	}

	personalRateString = strings.Join(re.FindAllString(personalRateString, -1), "")
	if firstIndex := strings.Index(personalRateString, PercentSymblo); firstIndex > -1 {
		personalRateString = personalRateString[0:firstIndex]
	}

	fcompanyRate, err := strconv.ParseFloat(companyRateString, 32)
	if err != nil {
		fcompanyRate = 0.00
	}

	fpersonalRate, err := strconv.ParseFloat(personalRateString, 32)
	if err != nil {
		fpersonalRate = 0.00
	}

	companyRate = float32(fcompanyRate) / 100
	personalRate = float32(fpersonalRate) / 100
	return companyRate, personalRate, nil
}

type WorkInjuryBaseResolver struct {
}

func (r *WorkInjuryBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type WorkInjuryRateResolver struct {
}

func (r *WorkInjuryRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}
	re := regexp.MustCompile(WorkInjuryDigitalRegMatch)
	resultString := strings.Join(re.FindAllString(tag, -1), "")

	if strings.IndexAny(resultString, DashSymblo) > -1 {
		arr := strings.Split(resultString, DashSymblo)
		companyRateString := arr[len(arr)-1]
		if firstIndex := strings.Index(companyRateString, PercentSymblo); firstIndex > -1 {
			companyRateString = companyRateString[0:firstIndex]
		}

		fcompanyRate, err := strconv.ParseFloat(companyRateString, 32)
		if err != nil {
			fcompanyRate = 0.00
		}

		companyRate = float32(fcompanyRate) / 100
		return companyRate, 0.00, nil
	} else if strings.IndexAny(resultString, CommaSymblo) > -1 {
		arr := strings.Split(tag, CommaSymblo)
		companyRateString := arr[len(arr)-1]
		if firstIndex := strings.Index(companyRateString, PercentSymblo); firstIndex > -1 {
			companyRateString = companyRateString[0:firstIndex]
		}

		fcompanyRate, err := strconv.ParseFloat(companyRateString, 32)
		if err != nil {
			fcompanyRate = 0.00
		}

		companyRate = float32(fcompanyRate) / 100
		return companyRate, 0.00, nil
	} else {
		if firstIndex := strings.Index(resultString, PercentSymblo); firstIndex > -1 {
			resultString = resultString[0:firstIndex]
		}
		fcompanyRate, err := strconv.ParseFloat(resultString, 32)
		if err != nil {
			fcompanyRate = 0.00
		}

		companyRate = float32(fcompanyRate) / 100
		return companyRate, 0.00, nil
	}
}

type MaternityBaseResolver struct {
}

func (r *MaternityBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type MaternityRateResolver struct {
}

func (r *MaternityRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return commonRateResolve(tag)
}

type HousingFundBaseResolver struct {
}

func (r *HousingFundBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type HousingFundRateResolver struct {
}

func (r *HousingFundRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return 1.0, 1.0, nil
}
