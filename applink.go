package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type AppTarget struct {
	TargetURI *URI
}

func (t AppTarget) toWO() *walletobjects.AppLinkDataAppLinkInfoAppTarget {
	return &walletobjects.AppLinkDataAppLinkInfoAppTarget{
		TargetUri: t.TargetURI.toWO(),
	}
}

func woToAppTarget(t *walletobjects.AppLinkDataAppLinkInfoAppTarget) *AppTarget {
	if t == nil {
		return nil
	}

	return &AppTarget{
		TargetURI: woToUri(t.TargetUri),
	}
}

type AppLinkInfo struct {
	AppLogoImage *Image
	Title        *LocalizedString
	Description  *LocalizedString
	AppTarget    *AppTarget
}

func (i AppLinkInfo) toWO() *walletobjects.AppLinkDataAppLinkInfo {
	res := walletobjects.AppLinkDataAppLinkInfo{}
	if i.AppLogoImage != nil {
		res.AppLogoImage = i.AppLogoImage.toWO()
	}

	if i.AppTarget != nil {
		res.AppTarget = i.AppTarget.toWO()
	}

	if i.Description != nil {
		res.Description = i.Description.toWO()
	}

	if i.Title != nil {
		res.Title = i.Title.toWO()
	}

	return &res
}

func woToAppLinkInfo(i *walletobjects.AppLinkDataAppLinkInfo) *AppLinkInfo {
	if i == nil {
		return nil
	}

	return &AppLinkInfo{
		AppLogoImage: woToImage(i.AppLogoImage),
		Title:        woToLocalizedString(i.Title),
		Description:  woToLocalizedString(i.Description),
		AppTarget:    woToAppTarget(i.AppTarget),
	}
}

type AppLinkData struct {
	AndroidAppLinkInfo *AppLinkInfo
	IOSAppLinkInfo     *AppLinkInfo
	WebAppLinkInfo     *AppLinkInfo
}

func (d AppLinkData) toWo() *walletobjects.AppLinkData {
	res := walletobjects.AppLinkData{}
	if d.AndroidAppLinkInfo != nil {
		res.AndroidAppLinkInfo = d.AndroidAppLinkInfo.toWO()
	}

	if d.IOSAppLinkInfo != nil {
		res.IosAppLinkInfo = d.IOSAppLinkInfo.toWO()
	}

	if d.WebAppLinkInfo != nil {
		res.WebAppLinkInfo = d.WebAppLinkInfo.toWO()
	}

	return &res
}

func woToAppLinkData(d *walletobjects.AppLinkData) *AppLinkData {
	if d == nil {
		return nil
	}

	return &AppLinkData{
		AndroidAppLinkInfo: woToAppLinkInfo(d.AndroidAppLinkInfo),
		IOSAppLinkInfo:     woToAppLinkInfo(d.IosAppLinkInfo),
		WebAppLinkInfo:     woToAppLinkInfo(d.WebAppLinkInfo),
	}
}
