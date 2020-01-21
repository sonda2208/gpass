package gpass

import (
	"context"
	"github.com/sonda2208/gpass/walletobjects"
)

type Issuer struct {
	IssuerID int64
	c *Client
}

func (c *Client) Issuer(id int64) *Issuer {
	return &Issuer{
		IssuerID: id,
		c:        c,
	}
}

func (c *Client) Issuers(ctx context.Context) ([]*Issuer, error) {
	res, err := c.wos.Issuer.List().Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	issuers := make([]*Issuer, len(res.Resources))
	for i, iss := range res.Resources {
		issuers[i] = toIssuer(iss, c)
	}
	
	return issuers, nil
}

func toIssuer(iss *walletobjects.Issuer, c *Client) *Issuer {
	return &Issuer{
		IssuerID: iss.IssuerId,
		c:        c,
	}
}

func (iss *Issuer) Create(ctx context.Context, im *IssuerMetadata) error {
	i, err := im.toWO()
	if err != nil {
		return err
	}

	_, err = iss.c.wos.Issuer.Insert(i).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (iss *Issuer) Metadata(ctx context.Context) (*IssuerMetadata, error) {
	i, err := iss.c.wos.Issuer.Get(iss.IssuerID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	meta, err := woToIssuerMetadata(i)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

type IssuerMetadata struct {
	ContactInfo *IssuerContactInfo
	HomepageURL string
	Name string
}

func (im *IssuerMetadata) toWO() (*walletobjects.Issuer, error) {
	i := &walletobjects.Issuer{}
	if im == nil {
		return i, nil
	}

	i.Name = im.Name
	i.HomepageUrl = im.HomepageURL
	i.ContactInfo = im.ContactInfo.toWO()
	return i, nil
}

func woToIssuerMetadata(i *walletobjects.Issuer) (*IssuerMetadata, error ) {
	meta := &IssuerMetadata{
		ContactInfo: woToIssuerContactInfo(i.ContactInfo),
		HomepageURL: i.HomepageUrl,
		Name:        i.Name,
	}
	return meta, nil
}

type IssuerContactInfo struct {
	AlertsEmails []string
	Email string
	Name string
	Phone string
}

func (ici *IssuerContactInfo) toWO() *walletobjects.IssuerContactInfo {
	i := &walletobjects.IssuerContactInfo{}
	if ici == nil {
		return nil
	}

	i.AlertsEmails = ici.AlertsEmails
	i.Email = ici.Email
	i.Name = ici.Name
	i.Phone = ici.Phone
	return i
}

func woToIssuerContactInfo(i *walletobjects.IssuerContactInfo) *IssuerContactInfo {
	return &IssuerContactInfo{
		AlertsEmails: i.AlertsEmails,
		Email:        i.Email,
		Name:         i.Name,
		Phone:        i.Phone,
	}
}