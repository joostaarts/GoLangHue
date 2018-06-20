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
	Name    string   `json:"name"`
	Lights  []string `json:"lights"`
	Type    string   `json:"type"`
	State   State    `json:"state"`
	Recycle bool     `json:"recycle"`
	Class   string   `json:"class"`
	Action  Action   `json:"action"`
}

type Action struct {
	On        bool      `json:"on"`
	Bri       int64     `json:"bri"`
	Hue       int64     `json:"hue"`
	Sat       int64     `json:"sat"`
	Effect    string    `json:"effect"`
	Xy        []float64 `json:"xy"`
	CT        int64     `json:"ct"`
	Alert     string    `json:"alert"`
	ColorMode string    `json:"colormode"`
}

type State struct {
	AllOn bool `json:"all_on"`
	AnyOn bool `json:"any_on"`
}
