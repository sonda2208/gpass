package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type TranslatedString struct {
	Language string
	Value    string
}

func (ts *TranslatedString) toWO() *walletobjects.TranslatedString {
	if ts == nil {
		return nil
	}

	return &walletobjects.TranslatedString{
		Kind:     "walletobjects#translatedString",
		Language: ts.Language,
		Value:    ts.Value,
	}
}

func listTranslatedStringToWO(ts []*TranslatedString) []*walletobjects.TranslatedString {
	res := make([]*walletobjects.TranslatedString, len(ts))
	for i, s := range ts {
		res[i] = s.toWO()
	}

	return res
}

func woToTranslatedString(s *walletobjects.TranslatedString) *TranslatedString {
	if s == nil {
		return nil
	}

	return &TranslatedString{
		Language: s.Language,
		Value:    s.Value,
	}
}
