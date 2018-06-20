package models

// http://10.0.1.2/api/xxxxx/lights

// Light describes a light
type Light struct {
	LightType        string       `json:"lighttype"`
	Name             string       `json:"name"`
	ModelID          string       `json:"modelid"`
	ManufacturerName string       `json:"manufacturername"`
	ProductName      string       `json:"productname"`
	UniqueID         string       `json:"uniqueid"`
	SwVersion        string       `json:"swversion"`
	State            LightState   `json:"state"`
	SwUpdate         SwUpdate     `json:"swupdate"`
	Capabilities     Capabilities `json:"capabilities"`
}

// LightState describes the state of a light
type LightState struct {
	On        bool         `json:"on"`
	Bri       int          `json:"bri"`
	Hue       int          `json:"hue"`
	Sat       int          `json:"sat"`
	Effect    string       `json:"effect"`
	Xy        []Coordinate `json:"xy"`
	Ct        int          `json:"ct"`
	Alert     string       `json:"alert"`
	ColorMode string       `json:"colormode"`
	Mode      string       `json:"mode"`
	Reachable bool         `json:"reachable"`
}

// SwUpdate describes software update state of a light
type SwUpdate struct {
	State       string `json:"state"`
	LastInstall string `json:"lastinstall"`
}

// Config describes the configuration of a light
type Config struct {
	Archetype string `json:"archetype"`
	Function  string `json:"function"`
	Direction string `json:"direction"`
}

// Capabilities describes the capabilities of a light
type Capabilities struct {
	Streaming StreamState `json:"streaming"`
	Control   Control     `json:"control"`
	Certified bool        `json:"certified"`
}

// Control describes light characteristics
type Control struct {
	MinDimLevel    int          `json:"mindimlevel"`
	MaxLumen       int          `json:"maxlumen"`
	ColorGamutType string       `json:"colorgamuttype"`
	ColorGamut     []Coordinate `json:"colorgamut"`
	Ct             MinMax       `json:"ct"`
}

// MinMax defines a min-max range
type MinMax struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// StreamState describes the streaming state of the light
type StreamState struct {
	Renderer bool `json:"renderer"`
	Proxy    bool `json:"proxy"`
}

/*
"state": {
            "on": false,
            "bri": 1,
            "hue": 10778,
            "sat": 251,
            "effect": "none",
            "xy": [
                0.5609,
                0.4042
            ],
            "ct": 500,
            "alert": "none",
            "colormode": "xy",
            "mode": "homeautomation",
            "reachable": true
        },
        "swupdate": {
            "state": "noupdates",
            "lastinstall": null
        },
        "type": "Extended color light",
        "name": "Hue color lamp 1",
        "modelid": "LCT007",
        "manufacturername": "Philips",
        "productname": "Hue color lamp",
        "capabilities": {
            "certified": true,
            "control": {
                "mindimlevel": 2000,
                "maxlumen": 800,
                "colorgamuttype": "B",
                "colorgamut": [
                    [
                        0.675,
                        0.322
                    ],
                    [
                        0.409,
                        0.518
                    ],
                    [
                        0.167,
                        0.04
                    ]
                ],
                "ct": {
                    "min": 153,
                    "max": 500
                }
            },
            "streaming": {
                "renderer": true,
                "proxy": true
            }
        },
        "config": {
            "archetype": "sultanbulb",
            "function": "mixed",
            "direction": "omnidirectional"
        },
        "uniqueid": "00:17:88:01:00:f7:b6:4b-0b",
        "swversion": "5.105.0.21536"
    }
*/
