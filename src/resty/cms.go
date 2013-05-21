package resty

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Cms() {

	serveSingle("/cms/", "assets/index.html")
	http.Handle("/cms/assets/", http.StripPrefix("/cms/assets", http.FileServer(http.Dir("./assets"))))

	router := mux.NewRouter()

	documentManager := new(DocumentManager)
	router.HandleFunc("/rest/document", BuildQuery(documentManager, "documents")).Methods("GET")
	router.HandleFunc("/rest/document", BuildPost(documentManager, "documents")).Methods("POST")
	router.HandleFunc("/rest/document/{id}", BuildGet(documentManager, "documents")).Methods("GET")
	router.HandleFunc("/rest/document/{id}", BuildPut(documentManager, "documents")).Methods("PUT")
	router.HandleFunc("/rest/document/{id}", BuildDelete(documentManager, "documents")).Methods("DELETE")

	documentTypeManager := new(DocumentTypeManager)
	router.HandleFunc("/rest/document-type", BuildQuery(documentTypeManager, "document-types")).Methods("GET")
	router.HandleFunc("/rest/document-type", BuildPost(documentTypeManager, "document-types")).Methods("POST")
	router.HandleFunc("/rest/document-type/{id}", BuildGet(documentTypeManager, "document-types")).Methods("GET")
	router.HandleFunc("/rest/document-type/{id}", BuildPut(documentTypeManager, "document-types")).Methods("PUT")
	router.HandleFunc("/rest/document-type/{id}", BuildDelete(documentTypeManager, "document-types")).Methods("DELETE")

	fileManager := new(FileManager)
	router.HandleFunc("/rest/file", BuildQuery(fileManager, "files")).Methods("GET")
	router.HandleFunc("/rest/file", PostFile).Methods("POST")
	router.HandleFunc("/rest/file/{id}", BuildGet(fileManager, "files")).Methods("GET")
	router.HandleFunc("/rest/file/{id}", BuildPut(fileManager, "files")).Methods("PUT")
	router.HandleFunc("/rest/file/{id}", BuildDelete(fileManager, "files")).Methods("DELETE")

	http.Handle("/rest/", router)

	contentRouter := mux.NewRouter()
	contentRouter.HandleFunc("/content/{id}", GetContent).Methods("GET")
	http.Handle("/content/", contentRouter)

}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
