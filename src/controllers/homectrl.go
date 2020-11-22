package controllers

import (
	"log"
	"net/http"

	"github.com/DataPsycho/portfolio/src/templates"
)

// CtrlHome is controler for Home
func CtrlHome(w http.ResponseWriter, req *http.Request) {
	data := "DataPsycho"
	err := templates.TPL.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
