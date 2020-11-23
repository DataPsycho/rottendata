package controllers

import (
	"log"
	"net/http"

	"github.com/DataPsycho/portfolio/src/templates"
	uuid "github.com/satori/go.uuid"
)

// CtrlHome is controler for Home
func CtrlHome(w http.ResponseWriter, req *http.Request) {
	data := "DataPsycho"
	err := templates.TPL.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func getCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "_psychosession",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
}
