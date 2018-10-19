package walletobject

import "time"

type DateTime struct {
	Date time.Time `json:"date,omitempty"`
}
