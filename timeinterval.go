package gpass

import (
	"time"

	"github.com/Hutchison-Technologies/gpass/walletobjects"
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
	res := walletobjects.TimeInterval{
		Kind: "walletobjects#timeInterval",
	}

	if ti.Start != nil {
		res.Start = ti.Start.toWO()
	}

	if ti.End != nil {
		res.End = ti.End.toWO()
	}

	return &res
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
