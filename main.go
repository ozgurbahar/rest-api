// Rest API
//
//     Schemes: http
//     Host: localhost:8081
//     BasePath: /
//     Version: 0.1.0
//     Contact: test@gmail.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import (
	"net/http"
	. "rest-api/route"
)

func main() {
	router := AllRoute()
	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/")))
	router.PathPrefix("/swagger-ui/").Handler(sh)
	http.ListenAndServe(":8081", router)
}
