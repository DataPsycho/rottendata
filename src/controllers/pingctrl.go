package controllers

import (
	"log"
	"net/http"

	"github.com/DataPsycho/portfolio/src/templates"
)

// CtrlPing is Pinging The Server
func CtrlPing(w http.ResponseWriter, req *http.Request) {
	err := templates.TPL.ExecuteTemplate(w, "pong.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
