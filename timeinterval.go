package gpass

import "github.com/sonda2208/gpass/walletobjects"

type TimeInterval struct {
	Start *walletobjects.DateTime
	End   *walletobjects.DateTime
}

func (ti *TimeInterval) toWO() *walletobjects.TimeInterval {
	return &walletobjects.TimeInterval{
		Start: ti.Start,
		End:   ti.End,
		Kind:  "walletobjects#timeInterval",
	}
}

func woToTimeInterval(ti *walletobjects.TimeInterval) *TimeInterval {
	if ti == nil {
		return nil
	}

	return &TimeInterval{
		Start: ti.Start,
		End:   ti.End,
	}
}
