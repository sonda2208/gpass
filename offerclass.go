package gpass

import (
	"context"

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

	meta, err := woToOfferClassMeta(o)
	if err != nil {
		return nil, err
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

	return woToOfferClassMeta(res)
}

func (oc *OfferClass) AddMessage(ctx context.Context, amr *AddMessageRequest) error {
	_, err := oc.c.wos.Offerclass.Addmessage(oc.OfferClassID, amr.toWO()).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

type OfferClassMetadata struct {
	IssuerName        string
	Provider          string
	RedemptionChannel string
	ReviewStatus      string
	Title             string
}

func (ocm *OfferClassMetadata) toWO() (*walletobjects.OfferClass, error) {
	return &walletobjects.OfferClass{
		IssuerName:        ocm.IssuerName,
		Provider:          ocm.Provider,
		RedemptionChannel: ocm.RedemptionChannel,
		ReviewStatus:      ocm.ReviewStatus,
		Title:             ocm.Title,
	}, nil
}

func woToOfferClassMeta(o *walletobjects.OfferClass) (*OfferClassMetadata, error) {
	ocm := &OfferClassMetadata{
		IssuerName:        o.IssuerName,
		Provider:          o.Provider,
		RedemptionChannel: o.RedemptionChannel,
		ReviewStatus:      o.ReviewStatus,
		Title:             o.Title,
	}
	return ocm, nil
}

type OfferClassMetadataToUpdate struct {
	ReviewStatus string
	Title        string
}

func (ocm *OfferClassMetadataToUpdate) toWO() (*walletobjects.OfferClass, error) {
	o := &walletobjects.OfferClass{
		ReviewStatus: ocm.ReviewStatus,
		Title:        ocm.Title,
	}
	return o, nil
}
