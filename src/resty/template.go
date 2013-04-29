package resty

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type templateField struct {
	Id string
	Name string
	Class string
	Type string
}

type template struct {
	Id string
	Name string
	Fields []templateField
}

var templates map[string] template

func AllTemplates(response http.ResponseWriter, request *http.Request) {
	out, jsonErr := json.Marshal(templates)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func GetTemplate(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, jsonErr := json.Marshal(templates[id])
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PostTemplate(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var tmp template
	jsonErr := json.Unmarshal(bodyBytes, &tmp)
	if RestError(jsonErr, response) {
		return
	}
	tmp.Id = sha1sum(tmp)
	templates[tmp.Id] = tmp
	out, jsonErr := json.Marshal(tmp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PutTemplate(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var tmp template
	jsonErr := json.Unmarshal(bodyBytes, &tmp)
	if RestError(jsonErr, response) {
		return
	}
	tmp.Id = id
	templates[id] = tmp
	out, jsonErr := json.Marshal(tmp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func DeleteTemplate(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	delete(templates, id)
}


