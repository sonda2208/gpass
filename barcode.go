package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type Barcode struct {
	AlternativeText string
	Type            string
	Value           string
	RenderEncoding  string
}

func (b *Barcode) toWO() *walletobjects.Barcode {
	if b == nil {
		return nil
	}

	return &walletobjects.Barcode{
		AlternateText:  b.AlternativeText,
		RenderEncoding: b.RenderEncoding,
		Type:           b.Type,
		Value:          b.Value,
	}
}

func wotoBarcode(b *walletobjects.Barcode) *Barcode {
	if b == nil {
		return nil
	}

	return &Barcode{
		AlternativeText: b.AlternateText,
		Type:            b.Type,
		Value:           b.Value,
		RenderEncoding:  b.RenderEncoding,
	}
}
