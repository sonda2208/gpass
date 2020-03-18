package gpass

import (
	"context"
	"errors"

	"github.com/sonda2208/gpass/walletobjects"
)

type OfferClass struct {
	IssuerID     int64
	OfferClassID string
	c            *Client
}

func (iss *Issuer) OfferClass(id string) *OfferClass {
	return &OfferClass{
		c:            iss.c,
		IssuerID:     iss.IssuerID,
		OfferClassID: id,
	}
}

func (iss *Issuer) OfferClasses(ctx context.Context) ([]*OfferClass, error) {
	res, err := iss.c.wos.Offerclass.List().IssuerId(iss.IssuerID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	classes := make([]*OfferClass, len(res.Resources))
	for i, oc := range res.Resources {
		classes[i] = toOfferClass(oc, iss.c)
		classes[i].IssuerID = iss.IssuerID
	}

	return classes, nil
}

func toOfferClass(oc *walletobjects.OfferClass, c *Client) *OfferClass {
	if oc == nil {
		return nil
	}

	return &OfferClass{
		OfferClassID: oc.Id,
		c:            c,
	}
}

func (oc *OfferClass) Create(ctx context.Context, ocm *OfferClassMetadata) error {
	o, err := ocm.toWO()
	if err != nil {
		return err
	}

	o.Id = oc.OfferClassID
	_, err = oc.c.wos.Offerclass.Insert(o).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (oc *OfferClass) Metadata(ctx context.Context) (*OfferClassMetadata, error) {
	o, err := oc.c.wos.Offerclass.Get(oc.OfferClassID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	meta := woToOfferClassMeta(o)
	if meta == nil {
		return nil, errors.New("invalid metadata")
	}

	return meta, nil
}

func (oc *OfferClass) Update(ctx context.Context, ocm *OfferClassMetadataToUpdate) (*OfferClassMetadata, error) {
	o, err := ocm.toWO()
	if err != nil {
		return nil, err
	}

	res, err := oc.c.wos.Offerclass.Patch(oc.OfferClassID, o).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return woToOfferClassMeta(res), nil
}

func (oc *OfferClass) AddMessage(ctx context.Context, amr *AddMessageRequest) error {
	_, err := oc.c.wos.Offerclass.Addmessage(oc.OfferClassID, amr.toWO()).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

type OfferClassMetadata struct {
	IssuerName          string
	Provider            string
	RedemptionChannel   string
	ReviewStatus        string
	ShortTitle          string
	Title               string
	LocalizedShortTitle *LocalizedString
	LocalizedTitle      *LocalizedString
}

func (ocm *OfferClassMetadata) toWO() (*walletobjects.OfferClass, error) {
	if ocm == nil {
		return nil, nil
	}

	return &walletobjects.OfferClass{
		IssuerName:          ocm.IssuerName,
		Provider:            ocm.Provider,
		RedemptionChannel:   ocm.RedemptionChannel,
		ReviewStatus:        ocm.ReviewStatus,
		Title:               ocm.Title,
		ShortTitle:          ocm.ShortTitle,
		LocalizedShortTitle: ocm.LocalizedShortTitle.toWO(),
		LocalizedTitle:      ocm.LocalizedTitle.toWO(),
	}, nil
}

func woToOfferClassMeta(o *walletobjects.OfferClass) *OfferClassMetadata {
	if o == nil {
		return nil
	}

	ocm := &OfferClassMetadata{
		IssuerName:          o.IssuerName,
		Provider:            o.Provider,
		RedemptionChannel:   o.RedemptionChannel,
		ReviewStatus:        o.ReviewStatus,
		Title:               o.Title,
		ShortTitle:          o.ShortTitle,
		LocalizedShortTitle: woToLocalizedString(o.LocalizedShortTitle),
		LocalizedTitle:      woToLocalizedString(o.LocalizedTitle),
	}
	return ocm
}

type OfferClassMetadataToUpdate struct {
	ReviewStatus        string
	Title               string
	ShortTitle          string
	LocalizedShortTitle *LocalizedString
	LocalizedTitle      *LocalizedString
}

func (ocm *OfferClassMetadataToUpdate) toWO() (*walletobjects.OfferClass, error) {
	o := &walletobjects.OfferClass{
		ReviewStatus:        ocm.ReviewStatus,
		Title:               ocm.Title,
		ShortTitle:          ocm.ShortTitle,
		LocalizedShortTitle: ocm.LocalizedShortTitle.toWO(),
		LocalizedTitle:      ocm.LocalizedTitle.toWO(),
	}
	return o, nil
}
