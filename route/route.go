package route

import (
	"github.com/gorilla/mux"
	"net/http"
	. "rest-api/controller"
)

func AllRoute() *mux.Router {
	router:= mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/books", GetAllBooks).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books", AddBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", UpdateBook).Methods(http.MethodPatch)
	router.HandleFunc("/api/books/{id}", DeleteBook).Methods(http.MethodDelete)
	router.Path("/api/books/search/{id}").Queries("name", "{name}").HandlerFunc(DeleteBook).Name("DeleteBook").Methods(http.MethodGet)
	return router
}



