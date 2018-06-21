package hueapi

import (
	"encoding/json"
	"net/http"

	"github.com/joostaarts/GolangHue/pkg/bridgediscovery"
)

func GetBridges(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bridgediscovery.GetBridges())

}

func GetLights(w http.ResponseWriter, r *http.Request) {

}
