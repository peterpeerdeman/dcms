package resty

import (
	"encoding/json"
	"fmt"
	"github.com/extemporalgenome/slug"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"storage"
)

type documentTypeField struct {
	Name string
	Type string
	Max  int
	Min  int
}

type documentType struct {
	Id     string
	Name   string
	Fields []documentTypeField
}

func AllDocumentType(response http.ResponseWriter, request *http.Request) {
	docs, listErr := storage.Repo.List("/document-types")
	if listErr != nil {
		return
	}
	var resp []documentType
	for _, file := range docs {
		data, getErr := storage.Repo.Get(fmt.Sprintf("/document-types/%s", file))
		if getErr == nil {
			var doc documentType
			err := json.Unmarshal(data, &doc)
			if err == nil {
				resp = append(resp, doc)
			}
		} else {
			log.Printf("Unable to get %v", file)
		}
	}
	out, jsonErr := json.Marshal(resp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func GetDocumentType(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, getErr := storage.Repo.Get(fmt.Sprintf("/document-types/%s", id))
	if RestError(getErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PutDocumentType(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	addErr := storage.Repo.Add(fmt.Sprintf("/document-types/%s", id), bodyBytes)
	if RestError(addErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(bodyBytes)
}

func PostDocumentType(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var doc documentType
	jsonErr := json.Unmarshal(bodyBytes, &doc)
	if RestError(jsonErr, response) {
		return
	}
	doc.Id = slug.Slug(doc.Name)
	out, marsErr := json.Marshal(doc)
	if RestError(marsErr, response) {
		return
	}
	addErr := storage.Repo.Add(fmt.Sprintf("/document-types/%s", doc.Id), out)
	if RestError(addErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func DeleteDocumentType(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	storage.Repo.Remove(fmt.Sprintf("/document-types/%s", id))
	response.Header().Set("Document-Type", "application/json")
}
