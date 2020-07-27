package gpass

import (
	"context"
	"errors"

	"github.com/sonda2208/gpass/walletobjects"
)

type LoyaltyObject struct {
	IssuerID        int64
	LoyaltyClassID  string
	LoyaltyObjectID string
	c               *Client
}

func (lc *LoyaltyClass) LoyaltyObject(id string) *LoyaltyObject {
	return &LoyaltyObject{
		IssuerID:        lc.IssuerID,
		LoyaltyClassID:  lc.LoyaltyClassID,
		LoyaltyObjectID: id,
		c:               lc.c,
	}
}

func (lc *LoyaltyClass) LoyaltyObjects(ctx context.Context) ([]*LoyaltyObject, error) {
	res, err := lc.c.wos.Loyaltyobject.List().ClassId(lc.LoyaltyClassID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	objects := make([]*LoyaltyObject, len(res.Resources))
	for i, oo := range res.Resources {
		objects[i] = toLoyaltyObject(oo, lc.c)
		objects[i].IssuerID = lc.IssuerID
	}

	return objects, nil
}

func toLoyaltyObject(oo *walletobjects.LoyaltyObject, c *Client) *LoyaltyObject {
	return &LoyaltyObject{
		LoyaltyClassID:  oo.ClassId,
		LoyaltyObjectID: oo.Id,
		c:               c,
	}
}

func (lc *LoyaltyObject) Create(ctx context.Context, lom *LoyaltyObjectMetadata) error {
	o, err := lom.toWO()
	if err != nil {
		return err
	}

	o.ClassId = lc.LoyaltyClassID
	o.Id = lc.LoyaltyObjectID
	_, err = lc.c.wos.Loyaltyobject.Insert(o).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (lc *LoyaltyObject) Metadata(ctx context.Context) (*LoyaltyObjectMetadata, error) {
	o, err := lc.c.wos.Loyaltyobject.Get(lc.LoyaltyObjectID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	meta := woToLoyaltyObjectMetadata(o)
	if meta == nil {
		return nil, errors.New("invalid metadata")
	}

	return meta, nil
}

func (lc *LoyaltyObject) Update(ctx context.Context, lom *LoyaltyObjectMetadataUpdate) (*LoyaltyObjectMetadata, error) {
	o, err := lom.toWO()
	if err != nil {
		return nil, err
	}

	res, err := lc.c.wos.Loyaltyobject.Patch(lc.LoyaltyObjectID, o).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return woToLoyaltyObjectMetadata(res), nil
}

func (lc *LoyaltyObject) AddMessage(ctx context.Context, amr *AddMessageRequest) error {
	_, err := lc.c.wos.Offerobject.Addmessage(lc.LoyaltyObjectID, amr.toWO()).Context(ctx).Do()
	if err != nil {
		return nil
	}

	return nil
}

type LoyaltyObjectMetadata struct {
	AccountId              string
	AccountName            string
	AppLinkData            *AppLinkData
	Barcode                *Barcode
	ImageModulesData       []*ImageModuleData
	LinkedOfferIds         []string
	LinksModuleData        *LinksModuleData
	Locations              []*LatLongPoint
	LoyaltyPoints          *LoyaltyPoints
	SecondaryLoyaltyPoints *LoyaltyPoints
	State                  string
	TextModulesData        []*TextModuleData
	ValidTimeInterval      *TimeInterval
}

func (lom *LoyaltyObjectMetadata) toWO() (*walletobjects.LoyaltyObject, error) {
	lo := &walletobjects.LoyaltyObject{
		AccountId:        lom.AccountId,
		AccountName:      lom.AccountName,
		ImageModulesData: listImageModuleDataToWO(lom.ImageModulesData),
		LinkedOfferIds:   lom.LinkedOfferIds,
		Locations:        locationListToWO(lom.Locations),
		State:            lom.State,
		TextModulesData:  listTextModuleDataToWO(lom.TextModulesData),
	}

	if lom.AppLinkData != nil {
		lo.AppLinkData = lom.AppLinkData.toWo()
	}

	if lom.Barcode != nil {
		lo.Barcode = lom.Barcode.toWO()
	}

	if lom.LinksModuleData != nil {
		lo.LinksModuleData = lom.LinksModuleData.toWO()
	}

	if lom.LoyaltyPoints != nil {
		lo.LoyaltyPoints = lom.LoyaltyPoints.toWO()
	}

	if lom.SecondaryLoyaltyPoints != nil {
		lo.SecondaryLoyaltyPoints = lom.SecondaryLoyaltyPoints.toWO()
	}

	if lom.ValidTimeInterval != nil {
		lo.ValidTimeInterval = lom.ValidTimeInterval.toWO()
	}

	return lo, nil
}

func woToLoyaltyObjectMetadata(lo *walletobjects.LoyaltyObject) *LoyaltyObjectMetadata {
	if lo == nil {
		return nil
	}

	lom := &LoyaltyObjectMetadata{
		AccountId:              lo.AccountId,
		AccountName:            lo.AccountName,
		AppLinkData:            woToAppLinkData(lo.AppLinkData),
		Barcode:                wotoBarcode(lo.Barcode),
		ImageModulesData:       make([]*ImageModuleData, len(lo.ImageModulesData)),
		LinkedOfferIds:         lo.LinkedOfferIds,
		LinksModuleData:        woToLinksModuleData(lo.LinksModuleData),
		Locations:              make([]*LatLongPoint, len(lo.Locations)),
		LoyaltyPoints:          woToLoyaltyPoints(lo.LoyaltyPoints),
		SecondaryLoyaltyPoints: woToLoyaltyPoints(lo.SecondaryLoyaltyPoints),
		State:                  lo.State,
		TextModulesData:        make([]*TextModuleData, len(lo.TextModulesData)),
		ValidTimeInterval:      woToTimeInterval(lo.ValidTimeInterval),
	}

	for i, l := range lo.Locations {
		lom.Locations[i] = woToLatLongPoint(l)
	}

	for i, d := range lo.TextModulesData {
		lom.TextModulesData[i] = woToTextModuleData(d)
	}

	for i, d := range lo.ImageModulesData {
		lom.ImageModulesData[i] = woToImageModuleData(d)
	}

	return lom
}

type LoyaltyObjectMetadataUpdate struct {
	AccountId              string
	AccountName            string
	AppLinkData            *AppLinkData
	Barcode                *Barcode
	ImageModulesData       []*ImageModuleData
	LinkedOfferIds         []string
	LinksModuleData        *LinksModuleData
	Locations              []*LatLongPoint
	LoyaltyPoints          *LoyaltyPoints
	SecondaryLoyaltyPoints *LoyaltyPoints
	State                  string
	TextModulesData        []*TextModuleData
	ValidTimeInterval      *TimeInterval
}

func (lom *LoyaltyObjectMetadataUpdate) toWO() (*walletobjects.LoyaltyObject, error) {
	o := &walletobjects.LoyaltyObject{
		AccountId:              lom.AccountId,
		AccountName:            lom.AccountName,
		AppLinkData:            lom.AppLinkData.toWo(),
		Barcode:                lom.Barcode.toWO(),
		ImageModulesData:       listImageModuleDataToWO(lom.ImageModulesData),
		LinkedOfferIds:         lom.LinkedOfferIds,
		LinksModuleData:        lom.LinksModuleData.toWO(),
		Locations:              locationListToWO(lom.Locations),
		LoyaltyPoints:          lom.LoyaltyPoints.toWO(),
		SecondaryLoyaltyPoints: lom.SecondaryLoyaltyPoints.toWO(),
		State:                  lom.State,
		TextModulesData:        listTextModuleDataToWO(lom.TextModulesData),
		ValidTimeInterval:      lom.ValidTimeInterval.toWO(),
	}
	return o, nil
}
