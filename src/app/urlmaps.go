package app

import (
	"net/http"

	"github.com/DataPsycho/portfolio/src/controllers"
)

func mapUrls() {
	// router.GET("/ping", controllers.Ping)
	http.HandleFunc("/home", controllers.CtrlHome)
	http.HandleFunc("/ping", controllers.CtrlPing)
}
