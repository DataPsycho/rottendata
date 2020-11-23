package app

import (
	"net/http"

	"github.com/DataPsycho/portfolio/src/controllers"
)

func mapUrls() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", controllers.CtrlHome)
	http.HandleFunc("/home", controllers.CtrlHome)
	http.HandleFunc("/ping", controllers.CtrlPing)
}
