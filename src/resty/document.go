package resty

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
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
	id string
	value Document
	resp chan interface{}
}

type Document struct {
	Id string
	Path string
	Type string
	Name string
	Fields map[string] string
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
				documents[msg.id] = msg.value
			case opRead:
				msg.resp <- documents[msg.id]
			case opReadAll:
				msg.resp <- map2slice(documents)
			case opUpdate:
				doc := documents[msg.id]
				doc.Name = msg.value.Name
				doc.Fields = msg.value.Fields
				documents[msg.id] = doc
			case opDelete:
				delete(documents, msg.id)
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
	id := vars["id"]
	readChan := make(chan interface{})
	msg := message{op: opRead, id: id, resp: readChan}
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
	id := vars["id"]
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
	msg := message{op: opUpdate, id: id, value: content, resp: readChan}
	messageChannel <- msg
	response.Header().Set("Document-Type", "application/json")
	out, _ := json.Marshal(content)
	response.Write(out)
}

func PostDocument(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var content Document
	jsonErr := json.Unmarshal(bodyBytes, &content)
	if RestError(jsonErr, response) {
		return
	}
	content.Fields = make(map[string] string)
	id := sha1sum(content)
	content.Id = id
	readChan := make(chan interface{})
	msg := message{op: opCreate, id: id, value: content, resp: readChan}
	messageChannel <- msg
	response.Header().Set("Document-Type", "application/json")
	out, _ := json.Marshal(content)
	response.Write(out)
}

func DeleteDocument(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	msg := message{op: opDelete, id: id}
	messageChannel <- msg
}
