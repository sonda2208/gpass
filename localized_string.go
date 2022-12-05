package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type LocalizedString struct {
	DefaultValue     *TranslatedString
	TranslatedValues []*TranslatedString
}

func (ls *LocalizedString) toWO() *walletobjects.LocalizedString {
	if ls == nil {
		return nil
	}

	res := walletobjects.LocalizedString{
		Kind:             "walletobjects#localizedString",
		TranslatedValues: listTranslatedStringToWO(ls.TranslatedValues),
	}

	if ls.DefaultValue != nil {
		res.DefaultValue = ls.DefaultValue.toWO()
	}

	return &res
}

func woToLocalizedString(s *walletobjects.LocalizedString) *LocalizedString {
	if s == nil {
		return nil
	}

	res := &LocalizedString{
		DefaultValue:     woToTranslatedString(s.DefaultValue),
		TranslatedValues: make([]*TranslatedString, len(s.TranslatedValues)),
	}

	for i, s := range s.TranslatedValues {
		res.TranslatedValues[i] = woToTranslatedString(s)
	}

	return res
}
