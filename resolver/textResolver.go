package resolver

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	PensionBaseTypeId     = "17"
	PensionRateTypeId     = "13"
	JoblessBaseTypeId     = "18"
	JoblessRateTypeId     = "8"
	MedicalBaseTypeId     = "19"
	MedicalRateTypeId     = "11"
	WorkInjuryBaseTypeId  = "20"
	WorkInjuryRateTypeId  = "6"
	MaternityBaseTypeId   = "21"
	MaternityRateTypeId   = "3"
	HousingFundBaseTypeId = "22"
	HousingFundRateTypeId = "14"
	BRSplitSymblo         = "<br>"
	CloseBRSplitSymblo    = "<br/>"
	CommaSymblo           = ","
	FullWidthCommaSymblo  = "，"
	ChineseCommaSymblo    = "、"
	DashSymblo            = "-"
	PercentSymblo         = "%"
	SwungDashSymblo       = "~"
	CurrencySymblo        = "元"
	MaxLimitSymblo        = "上限"
	DigitalRegMatch       = `[0-9.%、，,~-]`
)

type TextResolver interface {
	Resolve(tag string) (float32, float32, error)
}

func NewResolver(typeId string) (TextResolver, error) {
	var resolver TextResolver
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

func resolveParamerterFromRegexp(tag1 string, tag2 string, regExpression string) (retTag1 float32, retTag2 float32, err error) {
	if firstIndex := strings.Index(tag1, CurrencySymblo); firstIndex > -1 {
		tag1 = tag1[0:firstIndex]
	}
	if firstIndex := strings.Index(tag2, CurrencySymblo); firstIndex > -1 {
		tag2 = tag2[0:firstIndex]
	}

	re := regexp.MustCompile(regExpression)
	tag1 = strings.Join(re.FindAllString(tag1, -1), "")
	tag2 = strings.Join(re.FindAllString(tag2, -1), "")

	fTag1, err := strconv.ParseFloat(tag1, 32)
	if err != nil {
		fTag1 = 0.00
	}

	fTag2, err := strconv.ParseFloat(tag2, 32)
	if err != nil {
		fTag2 = 0.00
	}

	retTag1 = float32(fTag1)
	retTag2 = float32(fTag2)

	if retTag1 > retTag2 {
		return retTag1, retTag2, nil
	} else {
		return retTag2, retTag1, nil
	}
}

func commonBaseResolve(tag string) (maxBase float32, minBase float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	if strings.Index(tag, MaxLimitSymblo) > -1 && strings.Index(tag, CloseBRSplitSymblo) == -1 {
		arr := strings.Split(tag, MaxLimitSymblo)
		return resolveParamerterFromRegexp(arr[1], arr[0], DigitalRegMatch)
	} else {
		if strings.Index(tag, CloseBRSplitSymblo) == -1 {
			if strings.Index(tag, SwungDashSymblo) > -1 {
				arr := strings.Split(tag, SwungDashSymblo)
				return resolveParamerterFromRegexp(arr[1], arr[0], DigitalRegMatch)
			} else if strings.Index(tag, FullWidthCommaSymblo) > -1 {
				arr := strings.Split(tag, FullWidthCommaSymblo)
				return resolveParamerterFromRegexp(arr[0], arr[1], DigitalRegMatch)
			} else if strings.Index(tag, CommaSymblo) > -1 {
				arr := strings.Split(tag, CommaSymblo)
				return resolveParamerterFromRegexp(arr[0], arr[1], DigitalRegMatch)
			}
		} else {
			var arr []string
			arr = strings.Split(tag, BRSplitSymblo)
			if len(arr) < 2 {
				arr = strings.Split(tag, CloseBRSplitSymblo)
			}

			maxMinBaseString := arr[1]
			if strings.Index(maxMinBaseString, ChineseCommaSymblo) > -1 {
				maxMinBaseArr := strings.Split(maxMinBaseString, ChineseCommaSymblo)
				return resolveParamerterFromRegexp(maxMinBaseArr[0], maxMinBaseArr[1], DigitalRegMatch)
			} else if strings.Index(tag, MaxLimitSymblo) > -1 && strings.Index(tag, "元 ") == -1 {
				if strings.Index(tag, "缴费基数(上下限):上限") > -1 {
					arr = strings.Split(tag, "下限")
					return resolveParamerterFromRegexp(arr[0], arr[1], DigitalRegMatch)
				}
				arr = strings.Split(tag, MaxLimitSymblo)
				return resolveParamerterFromRegexp(arr[1], arr[0], DigitalRegMatch)
			} else if strings.Index(tag, "元 ") > -1 {
				arr = strings.Split(tag, "元 ")
				return resolveParamerterFromRegexp(arr[0], arr[1], DigitalRegMatch)
			} else {
				if strings.Index(arr[1], "（年）") > -1 {
					maxBase, minBase, err = resolveParamerterFromRegexp(arr[1], arr[1], DigitalRegMatch)
					tmp := fmt.Sprintf("%.2f", (maxBase / 12))
					tmpDig, _ := strconv.ParseFloat(tmp, 32)
					maxBase = float32(tmpDig)
					minBase = maxBase
					return maxBase, minBase, err
				}
				return resolveParamerterFromRegexp(arr[1], arr[1], DigitalRegMatch)
			}
		}
	}
	return 0.00, 0.00, nil
}

func commonRateResolve(tag string) (companyRate float32, personalRate float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	var arr []string

	arr = strings.Split(tag, BRSplitSymblo)
	if len(arr) < 2 {
		arr = strings.Split(tag, CloseBRSplitSymblo)
		if len(arr) < 2 {
			return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
		}
	}

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

	var arr []string
	arr = strings.Split(tag, BRSplitSymblo)
	if len(arr) < 2 {
		arr = strings.Split(tag, CloseBRSplitSymblo)
		if len(arr) < 2 {
			return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
		}
	}

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
	re := regexp.MustCompile(DigitalRegMatch)
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
	} else if strings.IndexAny(resultString, ChineseCommaSymblo) > -1 {
		arr := strings.Split(tag, ChineseCommaSymblo)
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
	if firstIndex := strings.Index(tag, PercentSymblo); firstIndex > -1 {
		tag = tag[0:firstIndex]
	}

	re := regexp.MustCompile(DigitalRegMatch)
	resultString := strings.Join(re.FindAllString(tag, -1), "")

	fcompanyRate, err := strconv.ParseFloat(resultString, 32)
	if err != nil {
		fcompanyRate = 0.00
	}

	companyRate = float32(fcompanyRate) / 100
	return companyRate, 0.00, nil
}

type HousingFundBaseResolver struct {
}

func (r *HousingFundBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return commonBaseResolve(tag)
}

type HousingFundRateResolver struct {
}

func (r *HousingFundRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	var arr []string
	arr = strings.Split(tag, BRSplitSymblo)
	if len(arr) < 2 {
		arr = strings.Split(tag, CloseBRSplitSymblo)
		if len(arr) < 2 {
			return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
		}
	}

	var processor = func(arg string) (rate float32, err error) {
		re := regexp.MustCompile(DigitalRegMatch)
		resultString := strings.Join(re.FindAllString(arg, -1), "")

		var handler = func(paramStr string, splitSymblo string) (rate float32, err error) {
			arr := strings.Split(paramStr, splitSymblo)
			rateString := arr[0]
			if firstIndex := strings.Index(rateString, PercentSymblo); firstIndex > -1 {
				rateString = rateString[0:firstIndex]
			} else {
				rateString = ""
			}

			frate, err := strconv.ParseFloat(rateString, 32)
			if err != nil {
				frate = 0.00
			}

			rate = float32(frate) / 100
			return rate, nil
		}

		if strings.IndexAny(resultString, DashSymblo) > -1 {
			return handler(resultString, DashSymblo)
		} else if strings.IndexAny(resultString, ChineseCommaSymblo) > -1 {
			return handler(resultString, ChineseCommaSymblo)
		} else if strings.IndexAny(resultString, CommaSymblo) > -1 {
			return handler(resultString, CommaSymblo)
		} else if strings.IndexAny(resultString, SwungDashSymblo) > -1 {
			return handler(resultString, SwungDashSymblo)
		} else if strings.IndexAny(resultString, FullWidthCommaSymblo) > -1 {
			return handler(resultString, FullWidthCommaSymblo)
		} else {
			if firstIndex := strings.Index(resultString, PercentSymblo); firstIndex > -1 {
				resultString = resultString[0:firstIndex]
			} else {
				resultString = ""
			}

			frate, err := strconv.ParseFloat(resultString, 32)
			if err != nil {
				frate = 0.00
			}

			rate = float32(frate) / 100
			return rate, nil
		}
	}

	companyRate, _ = processor(arr[0])
	personalRate, _ = processor(arr[1])

	return companyRate, personalRate, nil
}
