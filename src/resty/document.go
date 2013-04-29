package resty

import (
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
	"crypto/sha1"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

type opCode int

const (
	opCreate opCode = iota
	opRead
	opReadAll
    opUpdate
	opDelete
)

type message struct {
	op opCode
	path string
	value Document
	resp chan interface{}
}

func DocumentProcessor(messageChannel chan message) {
	documents := make(map[string] Document)
	for {
		select {
		case msg, msg_ok := <-messageChannel:
			if !msg_ok {
				log.Println("DocumentProcessor: Channel closed")
				return
			}
			log.Printf("%v", msg)
			switch msg.op {
			case opCreate:
				documents[msg.path] = msg.value
			case opRead:
				msg.resp <- documents[msg.path]
			case opReadAll:
				msg.resp <- map2slice(documents)
			case opUpdate:
				documents[msg.path] = msg.value
			case opDelete:
				delete(documents, msg.path)
			}
		}
	}
}

func map2slice(m map[string] Document) []Document {
	v := make([]Document, len(m))
	idx := 0
	for _, value := range m {
		v[idx] = value
		idx++
	}
	return v
}

type Document struct {
	Id string
	Name string
	Fields map[string] string
}

func AllDocument(response http.ResponseWriter, request *http.Request) {
	readChan := make(chan interface{})
	msg := message{op: opReadAll, resp: readChan}
	messageChannel <- msg
	resp := <- readChan
	out, jsonErr := json.Marshal(resp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func GetDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := vars["path"]
	readChan := make(chan interface{})
	msg := message{op: opRead, path: path, resp: readChan}
	messageChannel <- msg
	resp := <- readChan
	out, jsonErr := json.Marshal(resp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Document-Type", "application/json")
	response.Write(out)
}

func PutDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := vars["path"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var content Document
	jsonErr := json.Unmarshal(bodyBytes, &content)
	if RestError(jsonErr, response) {
		return
	}
	readChan := make(chan interface{})
	msg := message{op: opUpdate, path: path, value: content, resp: readChan}
	messageChannel <- msg
	response.Header().Set("Document-Type", "application/json")
	out, _ := json.Marshal(content)
	response.Write(out)
}

func PostDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := vars["path"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var content Document
	jsonErr := json.Unmarshal(bodyBytes, &content)
	if RestError(jsonErr, response) {
		return
	}
	id := sha1sum(content)
	content.Id = id
	readChan := make(chan interface{})
	msg := message{op: opCreate, path: path, value: content, resp: readChan}
	messageChannel <- msg
	response.Header().Set("Document-Type", "application/json")
	out, _ := json.Marshal(content)
	response.Write(out)
}

func DeleteDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := vars["path"]
	msg := message{op: opDelete, path: path}
	messageChannel <- msg
}

func sha1sum(value interface{}) string {
	data, _ := json.Marshal(value)
	h := sha1.New()
	io.WriteString(h, string(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}
