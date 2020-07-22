package gpass

import (
	"time"

	"github.com/sonda2208/gpass/walletobjects"
)

type DateTime struct {
	Date time.Time
}

func (dt *DateTime) toWO() *walletobjects.DateTime {
	return &walletobjects.DateTime{
		Date: dt.Date.UTC().Format(time.RFC3339),
	}
}

func woToDateTime(dt *walletobjects.DateTime) *DateTime {
	if dt == nil {
		return nil
	}

	res, _ := time.Parse(time.RFC3339, dt.Date)
	return &DateTime{
		Date: res,
	}
}

type TimeInterval struct {
	Start *DateTime
	End   *DateTime
}

func (ti *TimeInterval) toWO() *walletobjects.TimeInterval {
	return &walletobjects.TimeInterval{
		Start: ti.Start.toWO(),
		End:   ti.End.toWO(),
		Kind:  "walletobjects#timeInterval",
	}
}

func woToTimeInterval(ti *walletobjects.TimeInterval) *TimeInterval {
	if ti == nil {
		return nil
	}

	return &TimeInterval{
		Start: woToDateTime(ti.Start),
		End:   woToDateTime(ti.End),
	}
}
