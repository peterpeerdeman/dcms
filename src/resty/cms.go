package resty

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Cms() {

	serveSingle("/cms/", "assets/index.html")
	http.Handle("/cms/assets/", http.StripPrefix("/cms/assets", http.FileServer(http.Dir("./assets"))))

	Init()

	router := mux.NewRouter()
	router.HandleFunc("/rest/document", AllDocument).Methods("GET")
	router.HandleFunc("/rest/document", PostDocument).Methods("POST")
	router.HandleFunc("/rest/document/{id}", GetDocument).Methods("GET")
	router.HandleFunc("/rest/document/{id}", PutDocument).Methods("PUT")
	router.HandleFunc("/rest/document/{id}", DeleteDocument).Methods("DELETE")

	router.HandleFunc("/rest/template", AllTemplates).Methods("GET")
	router.HandleFunc("/rest/template", PostTemplate).Methods("POST")
	router.HandleFunc("/rest/template/{id}", GetTemplate).Methods("GET")
	router.HandleFunc("/rest/template/{id}", PutTemplate).Methods("PUT")
	router.HandleFunc("/rest/template/{id}", DeleteTemplate).Methods("DELETE")

	http.Handle("/rest/", router)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filename)
		})
}
