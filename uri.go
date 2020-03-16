package gpass

import "github.com/sonda2208/gpass/walletobjects"

type URI struct {
	ID                   string
	URI                  string
	LocalizedDescription *LocalizedString
}

func (u *URI) toWO() *walletobjects.Uri {
	return &walletobjects.Uri{
		Id:                   u.ID,
		Kind:                 "walletobjects#uri",
		Uri:                  u.URI,
		LocalizedDescription: u.LocalizedDescription.toWO(),
	}
}

func listURIToWO(uris []*URI) []*walletobjects.Uri {
	res := make([]*walletobjects.Uri, len(uris))
	for i, s := range uris {
		res[i] = s.toWO()
	}

	return res
}

func woToUri(u *walletobjects.Uri) *URI {
	return &URI{
		ID:                   u.Id,
		URI:                  u.Uri,
		LocalizedDescription: woToLocalizedString(u.LocalizedDescription),
	}
}

type ImageUri struct {
	LocalizedDescription *LocalizedString
	URI                  string
}

func (iu *ImageUri) toWO() *walletobjects.ImageUri {
	return &walletobjects.ImageUri{
		Uri:                  iu.URI,
		LocalizedDescription: iu.LocalizedDescription.toWO(),
	}
}

func woToImageUri(u *walletobjects.ImageUri) *ImageUri {
	return &ImageUri{
		URI:                  u.Uri,
		LocalizedDescription: woToLocalizedString(u.LocalizedDescription),
	}
}
