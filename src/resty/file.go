package resty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/extemporalgenome/slug"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"storage"
)

type File struct {
	Id       string
	Name     string
	MimeType string
}

type FileManager struct {
}

func (this *FileManager) Create() interface{} {
	var File File
	return File
}

func (this *FileManager) MakeCollection() TypeCollection {
	collection := new(FileCollection)
	collection.data = make([]File, 0)
	return collection
}

func (this *FileManager) Bootstrap(in interface{}) interface{} {
	doc := in.(File)
	doc.Id = slug.Slug(doc.Name)
	return doc
}

func (this *FileManager) Serialize(File interface{}) ([]byte, error) {
	data, err := json.Marshal(File)
	return data, err
}

func (this *FileManager) Deserialize(data []byte) (interface{}, error) {
	var doc File
	err := json.Unmarshal(data, &doc)
	return doc, err
}

func (this *FileManager) Identifier(in interface{}) string {
	return in.(File).Id
}

type FileCollection struct {
	data []File
}

func (this *FileCollection) Append(in interface{}) {
	this.data = append(this.data, in.(File))
}

func (this *FileCollection) Serialize() ([]byte, error) {
	data, err := json.Marshal(this.data)
	return data, err
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
