package gpass

import "github.com/sonda2208/gpass/walletobjects"

type Image struct {
	SourceURI *ImageUri
}

func (img *Image) toWO() *walletobjects.Image {
	return &walletobjects.Image{
		Kind:      "walletobjects#image",
		SourceUri: img.SourceURI.toWO(),
	}
}

func woToImage(img *walletobjects.Image) *Image {
	if img == nil {
		return nil
	}

	return &Image{
		SourceURI: woToImageUri(img.SourceUri),
	}
}
