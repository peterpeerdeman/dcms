package resty

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Document struct {
	Id     string
	Path   string
	Type   string
	Name   string
	Fields map[string]string
}

func AllDocument(response http.ResponseWriter, request *http.Request) {
	docs, listErr := Repo.List("/documents")
	if listErr != nil {
		return
	}
	var resp []Document
	for _, file := range docs {
		data, getErr := Repo.Get(fmt.Sprintf("/documents/%s", file))
		if getErr == nil {
			var doc Document
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

func GetDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, getErr := Repo.Get(fmt.Sprintf("/documents/%s", id))
	if RestError(getErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PutDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	addErr := Repo.Add(fmt.Sprintf("/documents/%s", id), bodyBytes)
	if RestError(addErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(bodyBytes)
}

func PostDocument(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var doc Document
	jsonErr := json.Unmarshal(bodyBytes, &doc)
	if RestError(jsonErr, response) {
		return
	}
	doc.Id = uuid()
	doc.Fields = make(map[string]string)
	out, marsErr := json.Marshal(doc)
	if RestError(marsErr, response) {
		return
	}
	addErr := Repo.Add(fmt.Sprintf("/documents/%s", doc.Id), out)
	if RestError(addErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func DeleteDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	Repo.Remove(fmt.Sprintf("/documents/%s", id))
	response.Header().Set("Document-Type", "application/json")
}
