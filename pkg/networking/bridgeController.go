package networking

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/joostaarts/GolangHue/models/hueapimodels"
)

func CreateUser(bridgeBaseUrl string) (result string, success bool) {
	success = true
	body := bytes.NewBuffer([]byte(`{"devicetype":"Vue Web App#browser"}`))
	resp, err := http.Post(bridgeBaseUrl+"/api", "application/json", body)

	if err != nil {
		log.Println(err)
		return
	}

	b := make([]byte, 1024)
	n, err := resp.Body.Read(b)
	if err != nil {
		log.Println(err)
		return
	}

	b = b[:n]

	var res HueApiModels.CreateUserResults
	json.Unmarshal(b, &res)

	if res[0].Error.Type != 0 {
		return res[0].Error.Description, false
	}

	return res[0].Success.Username, true
}

func FullConfig(bridgeBaseUrl string) {
}
