package webserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joostaarts/GolangHue/api"
	"github.com/joostaarts/GolangHue/models"
)

func Startup() {
	router := mux.NewRouter()
	router.HandleFunc("/api/lights", hueapi.GetLights).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func TestSerialization() {
	b, _ := ioutil.ReadFile("C:\\Users\\310262118\\go\\src\\github.com\\joostaarts\\GolangHue\\models\\mocks\\scene.json")
	var scene models.Scene
	e := json.Unmarshal(b, &scene)
	if e != nil {
		log.Fatal(e)
	}

	b, _ = ioutil.ReadFile("C:\\Users\\310262118\\go\\src\\github.com\\joostaarts\\GolangHue\\models\\mocks\\light.json")
	var light models.Light
	e = json.Unmarshal(b, &light)
	if e != nil {
		log.Fatal(e)
	}
}
