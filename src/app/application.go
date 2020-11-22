package app

import (
	"net/http"
)

// StartApplication is Responsible for Starting the Main Server
func StartApplication() {
	// get map of all the url from
	mapUrls()
	// attach to default server
	http.ListenAndServe(":8080", nil)
}
