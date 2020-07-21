package gpass

import (
	"context"
	"errors"

	"github.com/sonda2208/gpass/walletobjects"
)

type LoyaltyClass struct {
	IssuerID       int64
	LoyaltyClassID string
	c              *Client
}

func (iss *Issuer) LoyaltyClass(id string) *LoyaltyClass {
	return &LoyaltyClass{
		c:              iss.c,
		IssuerID:       iss.IssuerID,
		LoyaltyClassID: id,
	}
}

func (iss *Issuer) LoyaltyClasses(ctx context.Context) ([]*LoyaltyClass, error) {
	res, err := iss.c.wos.Loyaltyclass.List().IssuerId(iss.IssuerID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	classes := make([]*LoyaltyClass, len(res.Resources))
	for i, lc := range res.Resources {
		classes[i] = toLoyaltyClass(lc, iss.c)
		classes[i].IssuerID = iss.IssuerID
	}

	return classes, nil
}

func toLoyaltyClass(lc *walletobjects.LoyaltyClass, c *Client) *LoyaltyClass {
	if lc == nil {
		return nil
	}

	return &LoyaltyClass{
		LoyaltyClassID: lc.Id,
		c:              c,
	}
}

func (lc *LoyaltyClass) Create(ctx context.Context, lcm *LoyaltyClassMetadata) error {
	o, err := lcm.toWO()
	if err != nil {
		return err
	}

	o.Id = lc.LoyaltyClassID
	_, err = lc.c.wos.Loyaltyclass.Insert(o).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (lc *LoyaltyClass) Metadata(ctx context.Context) (*LoyaltyClassMetadata, error) {
	o, err := lc.c.wos.Loyaltyclass.Get(lc.LoyaltyClassID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	meta := woToLoyaltyClassMeta(o)
	if meta == nil {
		return nil, errors.New("invalid metadata")
	}

	return meta, nil
}

func (lc *LoyaltyClass) Update(ctx context.Context, ocm *LoyaltyClassMetadataToUpdate) (*LoyaltyClassMetadata, error) {
	o, err := ocm.toWO()
	if err != nil {
		return nil, err
	}

	res, err := lc.c.wos.Loyaltyclass.Patch(lc.LoyaltyClassID, o).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return woToLoyaltyClassMeta(res), nil
}

func (lc *LoyaltyClass) AddMessage(ctx context.Context, amr *AddMessageRequest) error {
	_, err := lc.c.wos.Loyaltyclass.Addmessage(lc.LoyaltyClassID, amr.toWO()).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

type LoyaltyClassMetadata struct {
	LocalizedAccountIdLabel   *LocalizedString
	LocalizedAccountNameLabel *LocalizedString
	HeroImage                 *Image
	HexBackgroundColor        string
	IssuerName                string
	ProgramName               string
	ReviewStatus              string
}

func (lcm *LoyaltyClassMetadata) toWO() (*walletobjects.LoyaltyClass, error) {
	if lcm == nil {
		return nil, nil
	}

	return &walletobjects.LoyaltyClass{
		LocalizedAccountIdLabel:   lcm.LocalizedAccountIdLabel.toWO(),
		LocalizedAccountNameLabel: lcm.LocalizedAccountNameLabel.toWO(),
		HeroImage:                 lcm.HeroImage.toWO(),
		HexBackgroundColor:        lcm.HexBackgroundColor,
		IssuerName:                lcm.IssuerName,
		ProgramName:               lcm.ProgramName,
		ReviewStatus:              lcm.ReviewStatus,
	}, nil
}

func woToLoyaltyClassMeta(o *walletobjects.LoyaltyClass) *LoyaltyClassMetadata {
	if o == nil {
		return nil
	}

	return &LoyaltyClassMetadata{
		LocalizedAccountIdLabel:   woToLocalizedString(o.LocalizedAccountIdLabel),
		LocalizedAccountNameLabel: woToLocalizedString(o.LocalizedAccountNameLabel),
		HeroImage:                 woToImage(o.HeroImage),
		HexBackgroundColor:        o.HexBackgroundColor,
		IssuerName:                o.IssuerName,
		ProgramName:               o.ProgramName,
		ReviewStatus:              o.ReviewStatus,
	}
}

type LoyaltyClassMetadataToUpdate struct {
	ReviewStatus string
	ProgramName  string
}

func (lcm *LoyaltyClassMetadataToUpdate) toWO() (*walletobjects.LoyaltyClass, error) {
	return &walletobjects.LoyaltyClass{
		ReviewStatus: lcm.ReviewStatus,
		ProgramName:  lcm.ProgramName,
	}, nil
}
