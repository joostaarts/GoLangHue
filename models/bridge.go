package models

type Root struct {
	SpecVersion SpecVersion `json:"specVersion" xml:"specVersion"`
	URLBase     string      `json:"URLBase" xml:"URLBase"`
	Device      Bridge      `json:"device" xml:"device"`
}

type Bridge struct {
	DeviceType       string      `json:"devicetype" xml:"deviceType"`
	FriendlyName     string      `json:"friendlyname" xml:"friendlyName"`
	Name             string      `json:"name"`
	Manufacturer     string      `json:"manufacturer" xml:"manufacturer"`
	ManufacturerURL  string      `json:"manufacturerurl" xml:"manufacturerURL"`
	ModelDescription string      `json:"modeldescription" xml:"modelDescription"`
	ModelName        string      `json:"modelname" xml:"modelName"`
	ModelNumber      string      `json:"modelnumber" xml:"modelNumber"`
	ModelURL         string      `json:"modelurl" xml:"modelURL"`
	SerialNumber     string      `json:"serialnumber" xml:"serialNumber"`
	Udn              string      `json:"udn" xml:"UDN"`
	PresentationURL  string      `json:"presentationurl" xml:"presentationURL"`
	IconList         IconList    `json:"iconlist" xml:"iconList"`
	URLBase          string      `json:"urlbase" xml:"URLBase"`
	SmallConfig      SmallConfig `json:"smallconfig"`
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
