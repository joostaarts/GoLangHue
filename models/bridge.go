package models

type Root struct {
	SpecVersion SpecVersion `json:"specVersion" xml:"specVersion"`
	URLBase     string      `json:"URLBase" xml:"URLBase"`
	Device      Bridge      `json:"device" xml:"device"`
}

type Bridge struct {
	DeviceType       string   `json:"deviceType" xml:"deviceType"`
	FriendlyName     string   `json:"friendlyName" xml:"friendlyName"`
	Manufacturer     string   `json:"manufacturer" xml:"manufacturer"`
	ManufacturerURL  string   `json:"manufacturerURL" xml:"manufacturerURL"`
	ModelDescription string   `json:"modelDescription" xml:"modelDescription"`
	ModelName        string   `json:"modelName" xml:"modelName"`
	ModelNumber      string   `json:"modelNumber" xml:"modelNumber"`
	ModelURL         string   `json:"modelURL" xml:"modelURL"`
	SerialNumber     string   `json:"serialNumber" xml:"serialNumber"`
	Udn              string   `json:"UDN" xml:"UDN"`
	PresentationURL  string   `json:"presentationURL" xml:"presentationURL"`
	IconList         IconList `json:"iconList" xml:"iconList"`
}

type IconList struct {
	Icon Icon `json:"icon" xml:"icon"`
}

type Icon struct {
	Mimetype string `json:"mimetype" xml:"mimetype"`
	Height   string `json:"height" xml:"height"`
	Width    string `json:"width" xml:"width"`
	Depth    string `json:"depth" xml:"depth"`
	URL      string `json:"url" xml:"url"`
}

type SpecVersion struct {
	Major string `json:"major" xml:"major"`
	Minor string `json:"minor" xml:"minor"`
}
