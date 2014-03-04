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

	BRSplitSymblo   = "<br>"
	PercentSymblo   = "%"
	CommaSymblo     = "、"
	DigitalRegMatch = `\d`
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

type PensionBaseResolver struct {
}

func (r *PensionBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
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

	imaxBase, _ := strconv.Atoi(maxBaseString)
	iminBase, _ := strconv.Atoi(minBaseString)
	maxBase = float32(imaxBase)
	minBase = float32(iminBase)

	return maxBase, minBase, nil
}

type PensionRateResolver struct {
}

func (r *PensionRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	if len(tag) == 0 {
		return printTagIsEmptyError(tag)
	}

	tag = strings.Replace(tag, PercentSymblo, "", -1)
	if strings.IndexAny(tag, BRSplitSymblo) == -1 {
		return printNotFoundSplitSymbloError(tag, BRSplitSymblo)
	}

	arr := strings.Split(tag, BRSplitSymblo)
	companyRateString := arr[0]
	personalRateString := arr[1]
	re := regexp.MustCompile(DigitalRegMatch)

	companyRateString = strings.Join(re.FindAllString(companyRateString, -1), "")
	personalRateString = strings.Join(re.FindAllString(personalRateString, -1), "")

	icompanyRate, _ := strconv.Atoi(companyRateString)
	ipersonalRate, _ := strconv.Atoi(personalRateString)
	companyRate = float32(icompanyRate)
	personalRate = float32(ipersonalRate)
	return companyRate, personalRate, nil
}

type JoblessBaseResolver struct {
}

func (r *JoblessBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return 1.0, 1.0, nil
}

type JoblessRateResolver struct {
}

func (r *JoblessRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return 1.0, 1.0, nil
}

type MedicalBaseResolver struct {
}

func (r *MedicalBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return 1.0, 1.0, nil
}

type MedicalRateResolver struct {
}

func (r *MedicalRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return 1.0, 1.0, nil
}

type WorkInjuryBaseResolver struct {
}

func (r *WorkInjuryBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return 1.0, 1.0, nil
}

type WorkInjuryRateResolver struct {
}

func (r *WorkInjuryRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return 1.0, 1.0, nil
}

type MaternityBaseResolver struct {
}

func (r *MaternityBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return 1.0, 1.0, nil
}

type MaternityRateResolver struct {
}

func (r *MaternityRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return 1.0, 1.0, nil
}

type HousingFundBaseResolver struct {
}

func (r *HousingFundBaseResolver) Resolve(tag string) (maxBase float32, minBase float32, err error) {
	return 1.0, 1.0, nil
}

type HousingFundRateResolver struct {
}

func (r *HousingFundRateResolver) Resolve(tag string) (companyRate float32, personalRate float32, err error) {
	return 1.0, 1.0, nil
}
