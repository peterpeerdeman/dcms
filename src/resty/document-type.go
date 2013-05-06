package resty

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
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

var documentTypes map[string] documentType

func AllDocumentType(response http.ResponseWriter, request *http.Request) {
	out, jsonErr := json.Marshal(documentTypes)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func GetDocumentType(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, jsonErr := json.Marshal(documentTypes[id])
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PostDocumentType(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var tmp documentType
	jsonErr := json.Unmarshal(bodyBytes, &tmp)
	if RestError(jsonErr, response) {
		return
	}
	tmp.Fields = make([]documentTypeField, 0)
	tmp.Id = sha1sum(tmp)
	documentTypes[tmp.Id] = tmp
	out, jsonErr := json.Marshal(tmp)
	if RestError(jsonErr, response) {
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
	var tmp documentType
	jsonErr := json.Unmarshal(bodyBytes, &tmp)
	if RestError(jsonErr, response) {
		return
	}
	tmp.Id = id
	documentTypes[id] = tmp
	out, jsonErr := json.Marshal(tmp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func DeleteDocumentType(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	delete(documentTypes, id)
}


