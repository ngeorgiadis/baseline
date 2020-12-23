package handlers

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("../../ui/dist/index.html")
	if err != nil {
		log.Error(err.Error())
	}
	w.Write(b)
}

type userFilter struct {
	Values    []interface{} `json:"values"`
	Opperator string        `json:"operator"`
	Type      string        `json:"type"`
}
