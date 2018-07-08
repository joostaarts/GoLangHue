package hueapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joostaarts/GolangHue/pkg/bridgediscovery"
	"github.com/joostaarts/GolangHue/pkg/networking"
)

func GetBridges(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bridgediscovery.GetBridges())
}

func GetLights(w http.ResponseWriter, r *http.Request) {

}

func ConnectBridge(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// bridges := bridgediscovery.GetBridges()

	result, success := networking.CreateUser("http://192.168.2.2")
	log.Println(result)

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

	w.Write([]byte(result))
}
