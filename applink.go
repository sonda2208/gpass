package gpass

import "github.com/sonda2208/gpass/walletobjects"

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
	return &walletobjects.AppLinkDataAppLinkInfo{
		AppLogoImage: i.AppLogoImage.toWO(),
		AppTarget:    i.AppTarget.toWO(),
		Description:  i.Description.toWO(),
		Title:        i.Title.toWO(),
	}
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
	return &walletobjects.AppLinkData{
		AndroidAppLinkInfo: d.AndroidAppLinkInfo.toWO(),
		IosAppLinkInfo:     d.IOSAppLinkInfo.toWO(),
		WebAppLinkInfo:     d.WebAppLinkInfo.toWO(),
	}
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
