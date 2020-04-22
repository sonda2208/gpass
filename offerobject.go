package gpass

import (
	"context"
	"errors"

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

	meta := woToOfferObjectMeta(o)
	if meta == nil {
		return nil, errors.New("invalid metadata")
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

	return woToOfferObjectMeta(res), nil
}

func (oo *OfferObject) AddMessage(ctx context.Context, amr *AddMessageRequest) error {
	_, err := oo.c.wos.Offerobject.Addmessage(oo.OfferObjectID, amr.toWO()).Context(ctx).Do()
	if err != nil {
		return nil
	}

	return nil
}

type OfferObjectMetadata struct {
	State            string
	Locations        []*LatLongPoint
	Barcode          *Barcode
	LinksModuleData  *LinksModuleData
	TextModulesData  []*TextModuleData
	ImageModulesData []*ImageModuleData
	AppLinkData      *AppLinkData
}

func (oom *OfferObjectMetadata) toWO() (*walletobjects.OfferObject, error) {
	o := &walletobjects.OfferObject{
		State:            oom.State,
		Locations:        locationListToWO(oom.Locations),
		Barcode:          oom.Barcode.toWO(),
		LinksModuleData:  oom.LinksModuleData.toWO(),
		TextModulesData:  listTextModuleDataToWO(oom.TextModulesData),
		ImageModulesData: listImageModuleDataToWO(oom.ImageModulesData),
	}
	return o, nil
}

func woToOfferObjectMeta(oo *walletobjects.OfferObject) *OfferObjectMetadata {
	if oo == nil {
		return nil
	}

	oom := &OfferObjectMetadata{
		State:            oo.State,
		Locations:        make([]*LatLongPoint, len(oo.Locations)),
		Barcode:          wotoBarcode(oo.Barcode),
		LinksModuleData:  woToLinksModuleData(oo.LinksModuleData),
		TextModulesData:  make([]*TextModuleData, len(oo.TextModulesData)),
		ImageModulesData: make([]*ImageModuleData, len(oo.ImageModulesData)),
	}

	for i, l := range oo.Locations {
		oom.Locations[i] = woToLatLongPoint(l)
	}

	for i, d := range oo.TextModulesData {
		oom.TextModulesData[i] = woToTextModuleData(d)
	}

	for i, d := range oo.ImageModulesData {
		oom.ImageModulesData[i] = woToImageModuleData(d)
	}

	return oom
}

type OfferObjectMetadataToUpdate struct {
	State            string
	Locations        []*LatLongPoint
	Barcode          *Barcode
	LinksModuleData  *LinksModuleData
	TextModulesData  []*TextModuleData
	ImageModulesData []*ImageModuleData
	AppLinkData      *AppLinkData
}

func (oom *OfferObjectMetadataToUpdate) toWO() (*walletobjects.OfferObject, error) {
	o := &walletobjects.OfferObject{
		State:            oom.State,
		Locations:        locationListToWO(oom.Locations),
		Barcode:          oom.Barcode.toWO(),
		LinksModuleData:  oom.LinksModuleData.toWO(),
		TextModulesData:  listTextModuleDataToWO(oom.TextModulesData),
		ImageModulesData: listImageModuleDataToWO(oom.ImageModulesData),
	}
	return o, nil
}
