package gpass

import "github.com/sonda2208/gpass/walletobjects"

type TextModuleData struct {
	Body            string
	Header          string
	ID              string
	LocalizedBody   *LocalizedString
	LocalizedHeader *LocalizedString
}

func (d *TextModuleData) toWO() *walletobjects.TextModuleData {
	if d == nil {
		return nil
	}

	return &walletobjects.TextModuleData{
		Body:            d.Body,
		Header:          d.Header,
		Id:              d.ID,
		LocalizedBody:   d.LocalizedBody.toWO(),
		LocalizedHeader: d.LocalizedHeader.toWO(),
	}
}

func listTextModuleDataToWO(d []*TextModuleData) []*walletobjects.TextModuleData {
	res := make([]*walletobjects.TextModuleData, len(d))
	for i, s := range d {
		res[i] = s.toWO()
	}

	return res
}

func woToTextModuleData(d *walletobjects.TextModuleData) *TextModuleData {
	return &TextModuleData{
		Body:            d.Body,
		Header:          d.Header,
		ID:              d.Id,
		LocalizedBody:   woToLocalizedString(d.LocalizedBody),
		LocalizedHeader: woToLocalizedString(d.LocalizedHeader),
	}
}
