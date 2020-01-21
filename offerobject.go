package gpass

import (
	"context"

	"github.com/sonda2208/gpass/walletobjects"
)

type OfferObject struct {
	IssuerID      int64
	OfferClassID  string
	OfferObjectID string
	c             *Client
}

func (oc *OfferClass) OfferObject(id string) *OfferObject {
	return &OfferObject{
		IssuerID:      oc.IssuerID,
		OfferClassID:  oc.OfferClassID,
		OfferObjectID: id,
		c:             oc.c,
	}
}

func (oc *OfferClass) OfferObjects(ctx context.Context) ([]*OfferObject, error) {
	res, err := oc.c.wos.Offerobject.List().ClassId(oc.OfferClassID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	objects := make([]*OfferObject, len(res.Resources))
	for i, oo := range res.Resources {
		objects[i] = toOfferObject(oo, oc.c)
		objects[i].IssuerID = oc.IssuerID
	}

	return objects, nil
}

func toOfferObject(oo *walletobjects.OfferObject, c *Client) *OfferObject {
	return &OfferObject{
		OfferClassID:  oo.ClassId,
		OfferObjectID: oo.Id,
		c:             c,
	}
}

func (oo *OfferObject) Create(ctx context.Context, oom *OfferObjectMetadata) error {
	o, err := oom.toWO()
	if err != nil {
		return err
	}

	o.ClassId = oo.OfferClassID
	o.Id = oo.OfferObjectID
	_, err = oo.c.wos.Offerobject.Insert(o).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (oo *OfferObject) Metadata(ctx context.Context) (*OfferObjectMetadata, error) {
	o, err := oo.c.wos.Offerobject.Get(oo.OfferObjectID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	meta, err := woToOfferObjectMeta(o)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (oo *OfferObject) Update(ctx context.Context, oom *OfferObjectMetadataToUpdate) (*OfferObjectMetadata, error) {
	o, err := oom.toWO()
	if err != nil {
		return nil, err
	}

	res, err := oo.c.wos.Offerobject.Patch(oo.OfferObjectID, o).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return woToOfferObjectMeta(res)
}

func (oo *OfferObject) AddMessage(ctx context.Context, amr *AddMessageRequest) error {
	_, err := oo.c.wos.Offerobject.Addmessage(oo.OfferObjectID, amr.toWO()).Context(ctx).Do()
	if err != nil {
		return nil
	}

	return nil
}

type OfferObjectMetadata struct {
	State     string
	Locations []*LatLongPoint
	Barcode   *Barcode
}

func (oom *OfferObjectMetadata) toWO() (*walletobjects.OfferObject, error) {
	o := &walletobjects.OfferObject{
		State:   oom.State,
		Barcode: oom.Barcode.toWO(),
	}
	return o, nil
}

func woToOfferObjectMeta(oo *walletobjects.OfferObject) (*OfferObjectMetadata, error) {
	oom := &OfferObjectMetadata{
		State:   oo.State,
		Barcode: wotoBarcode(oo.Barcode),
	}

	oom.Locations = make([]*LatLongPoint, len(oo.Locations))
	for i, l := range oo.Locations {
		oom.Locations[i] = woToLatLongPoint(l)
	}

	return oom, nil
}

type OfferObjectMetadataToUpdate struct {
	State     string
	Locations []*LatLongPoint
	Barcode   *Barcode
}

func (oom *OfferObjectMetadataToUpdate) toWO() (*walletobjects.OfferObject, error) {
	o := &walletobjects.OfferObject{
		State:     oom.State,
		Locations: locationListToWO(oom.Locations),
		Barcode:   oom.Barcode.toWO(),
	}
	return o, nil
}
