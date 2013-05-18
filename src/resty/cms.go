package resty

import (
	"github.com/gorilla/mux"
	"net/http"
	"storage"
)

var Repo *storage.Repository

func Cms() {

	serveSingle("/cms/", "assets/index.html")
	http.Handle("/cms/assets/", http.StripPrefix("/cms/assets", http.FileServer(http.Dir("./assets"))))

	Init()

	Repo = storage.Init()

	router := mux.NewRouter()
	router.HandleFunc("/rest/document", AllDocument).Methods("GET")
	router.HandleFunc("/rest/document", PostDocument).Methods("POST")
	router.HandleFunc("/rest/document/{id}", GetDocument).Methods("GET")
	router.HandleFunc("/rest/document/{id}", PutDocument).Methods("PUT")
	router.HandleFunc("/rest/document/{id}", DeleteDocument).Methods("DELETE")

	router.HandleFunc("/rest/document-type", AllDocumentType).Methods("GET")
	router.HandleFunc("/rest/document-type", PostDocumentType).Methods("POST")
	router.HandleFunc("/rest/document-type/{id}", GetDocumentType).Methods("GET")
	router.HandleFunc("/rest/document-type/{id}", PutDocumentType).Methods("PUT")
	router.HandleFunc("/rest/document-type/{id}", DeleteDocumentType).Methods("DELETE")

	http.Handle("/rest/", router)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
