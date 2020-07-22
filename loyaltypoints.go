package gpass

import "github.com/sonda2208/gpass/walletobjects"

type LoyaltyPoints struct {
	Balance        *walletobjects.LoyaltyPointsBalance
	Label          string
	LocalizedLabel *walletobjects.LocalizedString
}

func (lp *LoyaltyPoints) toWO() *walletobjects.LoyaltyPoints {
	return &walletobjects.LoyaltyPoints{
		Balance:        lp.Balance,
		Label:          lp.Label,
		LocalizedLabel: lp.LocalizedLabel,
	}
}

func woToLoyaltyPoints(lp *walletobjects.LoyaltyPoints) *LoyaltyPoints {
	if lp == nil {
		return nil
	}

	return &LoyaltyPoints{
		Balance:        lp.Balance,
		Label:          lp.Label,
		LocalizedLabel: lp.LocalizedLabel,
	}
}
