package models

import "encoding/json"

func UnmarshalGroup(data []byte) (Group, error) {
	var r Group
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Group) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Group struct {
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
