package models

// Scene describes a scene that control the lights in a room
type Scene struct {
	Name        string   `json:"name"`
	Lights      []string `json:"lights"`
	Owner       string   `json:"owner"`
	Recycle     bool     `json:"recycle"`
	Locked      bool     `json:"locked"`
	AppData     AppData  `json:"appdata"`
	Picture     string   `json:"picture"`
	LastUpdated string   `json:"lastupdated"`
	Version     int      `json:"version"`
}

// AppData provides specific information about the scene
type AppData struct {
	Version int    `json:"version"`
	Data    string `json:"data"`
}
