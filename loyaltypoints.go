package gpass

import "github.com/sonda2208/gpass/walletobjects"

type Money struct {
	CurrencyCode string
	Micros       int64
}

func (m *Money) toWO() *walletobjects.Money {
	return &walletobjects.Money{
		CurrencyCode: m.CurrencyCode,
		Kind:         "walletobjects#money",
		Micros:       m.Micros,
	}
}

func woToMoney(m *walletobjects.Money) *Money {
	if m != nil {
		return nil
	}

	return &Money{
		CurrencyCode: m.CurrencyCode,
		Micros:       m.Micros,
	}
}

type LoyaltyPointsBalance struct {
	Double float64
	Int    int64
	Money  *Money
	String string
}

func (lpb *LoyaltyPointsBalance) toWO() *walletobjects.LoyaltyPointsBalance {
	return &walletobjects.LoyaltyPointsBalance{
		Double: lpb.Double,
		Int:    lpb.Int,
		String: lpb.String,
		Money:  lpb.Money.toWO(),
	}
}

func woToLoyaltyPointsBalance(lpb *walletobjects.LoyaltyPointsBalance) *LoyaltyPointsBalance {
	if lpb != nil {
		return nil
	}

	return &LoyaltyPointsBalance{
		Double: lpb.Double,
		Int:    lpb.Int,
		String: lpb.String,
		Money:  woToMoney(lpb.Money),
	}
}

type LoyaltyPoints struct {
	Balance        *LoyaltyPointsBalance
	Label          string
	LocalizedLabel *LocalizedString
}

func (lp *LoyaltyPoints) toWO() *walletobjects.LoyaltyPoints {
	return &walletobjects.LoyaltyPoints{
		Balance:        lp.Balance.toWO(),
		Label:          lp.Label,
		LocalizedLabel: lp.LocalizedLabel.toWO(),
	}
}

func woToLoyaltyPoints(lp *walletobjects.LoyaltyPoints) *LoyaltyPoints {
	if lp == nil {
		return nil
	}

	return &LoyaltyPoints{
		Balance:        woToLoyaltyPointsBalance(lp.Balance),
		Label:          lp.Label,
		LocalizedLabel: woToLocalizedString(lp.LocalizedLabel),
	}
}
