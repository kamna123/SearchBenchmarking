package datamodel

type Pmp struct{
	Adsize string
	AdFormats string
	Gps string
	Ifa string
	Integration string
	Clickable string
	Instream string
	Incentivized string
	Autoplay string
	Sound string
	Skippable string
	Viewability string
	Moat string
	AdsizeDesk string
	AdunitDesk string
	Apporwebtargetting string
	Categories string
	CategoriesInclude string
	DomainBundle string
	DomainBundleInclude string
	AdultContent string
	DeviceOSinclude string
	DeviceOS string
	DeviceID string
	DeviceCategory string
	DeviceMake string
	DeviceModel string
	DeviceOSVersion string
	CarrierWhitelist string
	Ispwhitelist string
	WifiTargeting string
	Country string
	Postalcode string
	Dma string
	DpExpression string
	EnabledOnDesktop string
	EnabledOnMobile string
	EnabledOnTablet string
	DesktopSelected string
	MobSelected string
	TabSelected string
	DeviceTypeError string
	States string
	Cities string
	DesignId string
	Site string
}

func GetMetaDataMapping() map[string]string {

	pmpMapping := make(map[string]string)
	pmpMapping["Adsize"] = "adsize"
	pmpMapping["Gps"] = "gps"
	pmpMapping["AdunitDesk"] = "adunitdesk"
	pmpMapping["AdFormats"] = "adformats"
	pmpMapping["Country"] = "country"
	pmpMapping["DomainBundle"] = "domainbundle"
	pmpMapping["Ifa"] = "ifa"
	pmpMapping["Integration"] = "integration"
	pmpMapping["Clickable"] = "clickable"
	pmpMapping["Instream"] = "instream"
	pmpMapping["Incentivized"] = "incentivized"
	pmpMapping["Autoplay"] = "autoplay"
	pmpMapping["Sound"] = "sound"
	pmpMapping["Skippable"] = "skippable"
	pmpMapping["Viewability"] = "viewability"
	pmpMapping["Moat"] = "moat"
	pmpMapping["AdsizeDesk"] = "adsizedesk"
	pmpMapping["Adunitdesk"] = "adunitdesk"
	pmpMapping["Apporwebtargetting"] = "apporwebtargetting"
	pmpMapping["Categories"] = "categories"
	pmpMapping["CategoriesInclude"] = "categoriesinclude"
	pmpMapping["DomainBundle"] = "domainbundle"
	pmpMapping["DomainBundleInclude"] = "domainbundleinclude"
	pmpMapping["AdultContent"] = "adultcontent"
	pmpMapping["DeviceOSinclude"] = "deviceOSinclude"
	pmpMapping["DeviceOS"] = "deviceOS"
	pmpMapping["DeviceID"] = "deviceID"
	pmpMapping["DeviceCategory"] = "deviceCategory"
	pmpMapping["DeviceMake"] = "deviceMake"
	pmpMapping["DeviceModel"] = "devicemodel"
	pmpMapping["DeviceOSVersion"] = "deviceosversion"
	pmpMapping["CarrierWhitelist"] = "carrierwhitelist"
	pmpMapping["Ispwhitelist"] = "ispwhitelist"
	pmpMapping["WifiTargeting"] = "wifitargeting"
	pmpMapping["Country"] = "country"
	pmpMapping["States"] = "states"
	pmpMapping["Postalcode"] = "postalcode"
	pmpMapping["Dma"] = "dma"
	pmpMapping["DpExpression"] = "dp_expression"
	pmpMapping["EnabledOnDesktop"] = "enabled_on_desktop"
	pmpMapping["EnabledOnMobile"] = "enabled_on_mobile"
	pmpMapping["EnabledOnTablet"] = "enabled_on_tablet"
	pmpMapping["DesktopSelected"] = "desktopSelected"
	pmpMapping["MobSelected"] = "MobSelected"
	pmpMapping["TabSelected"] = "tabSelected"
	pmpMapping["DeviceTypeError"] = "DeviceTypeError"
	pmpMapping["Cities"] = "cities"
	pmpMapping["DesignId"] = "designid"
	pmpMapping["Site"] = "site"


	return pmpMapping
}