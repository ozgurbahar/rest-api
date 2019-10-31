package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"net/http"
	. "rest-api/model"
)
// swagger:operation GET /api/books users searchUser
// ---
// summary: Returns list of books by provided search parameters.
// description: HTTP status will be returned depending on first search term (a - 400, e - 403, rest - 200)
// parameters:
// - name: name
//   in: query
//   description: search params
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/usersResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "403":
//     "$ref": "#/responses/forbidden"


func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	book := &Book{}
	json.Unmarshal(reqBody, book)
	var db, errDb = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=1234 sslmode=disable")
	if errDb != nil {
		panic(errDb)
	}
	defer db.Close()
	//book := &Book{Name:"Test-1", Title:"Test title"}
	if dbc:=db.Create(book); dbc.Error != nil{
		fmt.Println(dbc.Error)
		panic(dbc.Error)
	}


	json.NewEncoder(w).Encode(book)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	name := r.FormValue("name")

	fmt.Println("id : ", id)
	fmt.Println("name : ", name)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)

}
