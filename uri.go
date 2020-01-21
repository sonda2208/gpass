package gpass

import "github.com/sonda2208/gpass/walletobjects"

type Uri struct {
	Description string
	ID string
	URI string
}

func (u *Uri) toWO() *walletobjects.Uri {
	return &walletobjects.Uri{
		Description:          u.Description,
		Id:                   u.ID,
		Kind:                 "walletobjects#uri",
		Uri:                  u.URI,
	}
}

func woToUri(u *walletobjects.Uri) *Uri {
	return &Uri{
		Description: u.Description,
		ID:          u.Id,
		URI:         u.Uri,
	}
}

type ImageUri struct {
	Description string
	URI string
}

func (iu *ImageUri) toWO() *walletobjects.ImageUri {
	return &walletobjects.ImageUri{
		Description:          iu.Description,
		Uri:                  iu.URI,
	}
}

func woToImageUri(u *walletobjects.ImageUri) *ImageUri {
	return &ImageUri{
		Description: u.Description,
		URI:         u.Uri,
	}
}