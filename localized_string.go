package gpass

import "github.com/sonda2208/gpass/walletobjects"

type LocalizedString struct {
	DefaultValue     *TranslatedString
	TranslatedValues []*TranslatedString
}

func (ls *LocalizedString) toWO() *walletobjects.LocalizedString {
	if ls == nil {
		return nil
	}

	return &walletobjects.LocalizedString{
		Kind:             "walletobjects#localizedString",
		DefaultValue:     ls.DefaultValue.toWO(),
		TranslatedValues: listTranslatedStringToWO(ls.TranslatedValues),
	}
}

func woToLocalizedString(s *walletobjects.LocalizedString) *LocalizedString {
	res := &LocalizedString{
		DefaultValue: woToTranslatedString(s.DefaultValue),
	}

	res.TranslatedValues = make([]*TranslatedString, len(s.TranslatedValues))
	for i, s := range s.TranslatedValues {
		res.TranslatedValues[i] = woToTranslatedString(s)
	}

	return res
}
