package gpass

import "github.com/sonda2208/gpass/walletobjects"

type ImageModuleData struct {
	ID        string
	MainImage *Image
}

func (d *ImageModuleData) toWO() *walletobjects.ImageModuleData {
	if d == nil {
		return nil
	}

	return &walletobjects.ImageModuleData{
		Id:        d.ID,
		MainImage: d.MainImage.toWO(),
	}
}

func listImageModuleDataToWO(d []*ImageModuleData) []*walletobjects.ImageModuleData {
	res := make([]*walletobjects.ImageModuleData, len(d))
	for i, s := range d {
		res[i] = s.toWO()
	}

	return res
}

func woToImageModuleData(d *walletobjects.ImageModuleData) *ImageModuleData {
	if d == nil {
		return nil
	}

	res := &ImageModuleData{
		ID:        d.Id,
		MainImage: woToImage(d.MainImage),
	}

	return res
}
