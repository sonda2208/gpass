package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type LinksModuleData struct {
	URIs []*URI
}

func (d *LinksModuleData) toWO() *walletobjects.LinksModuleData {
	if d == nil {
		return nil
	}

	return &walletobjects.LinksModuleData{
		Uris: listURIToWO(d.URIs),
	}
}

func woToLinksModuleData(d *walletobjects.LinksModuleData) *LinksModuleData {
	if d == nil {
		return nil
	}

	res := &LinksModuleData{
		URIs: make([]*URI, len(d.Uris)),
	}

	for i, u := range d.Uris {
		res.URIs[i] = woToUri(u)
	}

	return res
}
