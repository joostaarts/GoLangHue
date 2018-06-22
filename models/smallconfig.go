package models

import "encoding/json"

func UnmarshalSmallConfig(data []byte) (SmallConfig, error) {
	var r SmallConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SmallConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SmallConfig struct {
	Name             string `json:"name"`
	DatastoreVersion string `json:"datastoreversion"`
	SwVersion        string `json:"swversion"`
	APIVersion       string `json:"apiversion"`
	MAC              string `json:"mac"`
	BridgeID         string `json:"bridgeid"`
	FactoryNew       bool   `json:"factorynew"`
	ModelID          string `json:"modelid"`
	StarterKitID     string `json:"starterkitid"`
}
