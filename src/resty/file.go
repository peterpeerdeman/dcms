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
	"storage"
)

type File struct {
	Id       string
	Name     string
	MimeType string
}

func AllFile(response http.ResponseWriter, request *http.Request) {
	docs, listErr := storage.Repo.List("/files")
	if listErr != nil {
		return
	}
	var resp []File
	for _, file := range docs {
		data, getErr := storage.Repo.Get(fmt.Sprintf("/files/%s", file))
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
	out, getErr := storage.Repo.Get(fmt.Sprintf("/content/%s", id))
	if RestError(getErr, response) {
		return
	}
	out2, getErr2 := storage.Repo.Get(fmt.Sprintf("/files/%s", id))
	if RestError(getErr2, response) {
		return
	}
	var file File
	jsonErr := json.Unmarshal(out2, &file)
	if RestError(jsonErr, response) {
		return
	}
	//response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	response.Header().Set("Content-Type", file.MimeType)
	response.Write(out)
}

func GetFile(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	out, getErr := storage.Repo.Get(fmt.Sprintf("/files/%s", id))
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
	addErr := storage.Repo.Add(fmt.Sprintf("/files/%s", id), bodyBytes)
	if RestError(addErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(bodyBytes)
}

func PostFile(response http.ResponseWriter, request *http.Request) {
	uploaded, tweede, formErr := request.FormFile("file")
	if RestError(formErr, response) {
		return
	}
	var doc File
	doc.Name = tweede.Filename
	doc.MimeType = tweede.Header.Get("Content-Type")
	doc.Id = uuid()
	var buf bytes.Buffer
	io.Copy(&buf, uploaded)
	out, marsErr := json.Marshal(doc)
	log.Printf("%v", tweede.Header)
	if RestError(marsErr, response) {
		return
	}
	addErr := storage.Repo.Add(fmt.Sprintf("/files/%s", doc.Id), out)
	if RestError(addErr, response) {
		return
	}
	fileErr := storage.Repo.Add(fmt.Sprintf("/content/%s", doc.Id), buf.Bytes())
	if RestError(fileErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func DeleteFile(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	storage.Repo.Remove(fmt.Sprintf("/files/%s", id))
	response.Header().Set("Document-Type", "application/json")
}
