package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type URI struct {
	ID                   string
	URI                  string
	LocalizedDescription *LocalizedString
}

func (u *URI) toWO() *walletobjects.Uri {
	res := walletobjects.Uri{
		Id:   u.ID,
		Kind: "walletobjects#uri",
		Uri:  u.URI,
	}

	if u.LocalizedDescription != nil {
		res.LocalizedDescription = u.LocalizedDescription.toWO()
	}

	return &res
}

func listURIToWO(uris []*URI) []*walletobjects.Uri {
	res := make([]*walletobjects.Uri, len(uris))
	for i, s := range uris {
		res[i] = s.toWO()
	}

	return res
}

func woToUri(u *walletobjects.Uri) *URI {
	if u == nil {
		return nil
	}

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
	if u == nil {
		return nil
	}

	return &ImageUri{
		URI:                  u.Uri,
		LocalizedDescription: woToLocalizedString(u.LocalizedDescription),
	}
}
