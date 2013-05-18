package resty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type File struct {
	Id       string
	Name     string
	MimeType string
}

func AllFile(response http.ResponseWriter, request *http.Request) {
	docs, listErr := Repo.List("/files")
	if listErr != nil {
		return
	}
	var resp []File
	for _, file := range docs {
		data, getErr := Repo.Get(fmt.Sprintf("/files/%s", file))
		if getErr == nil {
			var doc File
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

func GetContent(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, getErr := Repo.Get(fmt.Sprintf("/content/%s", id))
	if RestError(getErr, response) {
		return
	}
	//response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func GetFile(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, getErr := Repo.Get(fmt.Sprintf("/files/%s", id))
	if RestError(getErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PutFile(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	addErr := Repo.Add(fmt.Sprintf("/files/%s", id), bodyBytes)
	if RestError(addErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(bodyBytes)
}

func PostFile(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var doc File
	jsonErr := json.Unmarshal(bodyBytes, &doc)
	if RestError(jsonErr, response) {
		return
	}
	uploaded, tweede, formErr := request.FormFile("file")
	if RestError(formErr, response) {
		return
	}
	doc.Name = tweede.Filename
	doc.MimeType = tweede.Header.Get("Document-Type")
	doc.Id = uuid()
	var buf bytes.Buffer
	io.Copy(&buf, uploaded)
	out, marsErr := json.Marshal(doc)
	if RestError(marsErr, response) {
		return
	}
	addErr := Repo.Add(fmt.Sprintf("/files/%s", doc.Id), out)
	if RestError(addErr, response) {
		return
	}
	fileErr := Repo.Add(fmt.Sprintf("/content/%s", doc.Id), buf.Bytes())
	if RestError(fileErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func DeleteFile(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	Repo.Remove(fmt.Sprintf("/files/%s", id))
	response.Header().Set("Document-Type", "application/json")
}
