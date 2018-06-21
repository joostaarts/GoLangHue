package models

import "encoding/json"

func UnmarshalBridgeConfiguration(data []byte) (BridgeConfiguration, error) {
	var r BridgeConfiguration
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BridgeConfiguration) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BridgeConfiguration struct {
	Name             string               `json:"name"`
	ZigbeeChannel    int64                `json:"zigbeechannel"`
	BridgeID         string               `json:"bridgeid"`
	MAC              string               `json:"mac"`
	DHCP             bool                 `json:"dhcp"`
	IPAddress        string               `json:"ipaddress"`
	Netmask          string               `json:"netmask"`
	Gateway          string               `json:"gateway"`
	ProxyAddress     string               `json:"proxyaddress"`
	ProxyPort        int64                `json:"proxyport"`
	UTC              string               `json:"UTC"`
	LocalTime        string               `json:"localtime"`
	Timezone         string               `json:"timezone"`
	ModelID          string               `json:"modelid"`
	DataStoreVersion string               `json:"datastoreversion"`
	SwVersion        string               `json:"swversion"`
	APIVersion       string               `json:"apiversion"`
	SwUpdate         Swupdate             `json:"swupdate"`
	SwUpdate2        Swupdate2            `json:"swupdate2"`
	LinkButton       bool                 `json:"linkbutton"`
	PortalServices   bool                 `json:"portalservices"`
	PortalConnection string               `json:"portalconnection"`
	PortalState      Portalstate          `json:"portalstate"`
	InternetServices Internetservices     `json:"internetservices"`
	FactoryNew       bool                 `json:"factorynew"`
	ReplacesBridgeID string               `json:"replacesbridgeid"`
	Backup           Backup               `json:"backup"`
	StarterKitID     string               `json:"starterkitid"`
	Whitelist        map[string]Whitelist `json:"whitelist"`
}

type Backup struct {
	Status    string `json:"status"`
	ErrorCode int64  `json:"errorcode"`
}

type Internetservices struct {
	Internet     string `json:"internet"`
	RemoteAccess string `json:"remoteaccess"`
	Time         string `json:"time"`
	Swupdate     string `json:"swupdate"`
}

type Portalstate struct {
	Signedon      bool   `json:"signedon"`
	Incoming      bool   `json:"incoming"`
	Outgoing      bool   `json:"outgoing"`
	Communication string `json:"communication"`
}

type Swupdate struct {
	UpdateState    int64  `json:"updatestate"`
	CheckForUpdate bool   `json:"checkforupdate"`
	URL            string `json:"url"`
	Text           string `json:"text"`
	Notify         bool   `json:"notify"`
}

type Swupdate2 struct {
	CheckForUpdate bool         `json:"checkforupdate"`
	LastChange     string       `json:"lastchange"`
	Bridge         InstallState `json:"bridge"`
	State          string       `json:"state"`
	Autoinstall    Autoinstall  `json:"autoinstall"`
}

type Autoinstall struct {
	UpdateTime string `json:"updatetime"`
	On         bool   `json:"on"`
}

type InstallState struct {
	State       string `json:"state"`
	LastInstall string `json:"lastinstall"`
}

type Whitelist struct {
	LastUseDate string `json:"last use date"`
	CreateDate  string `json:"create date"`
	Name        string `json:"name"`
}
