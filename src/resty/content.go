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
	id string
	value Content
	resp chan interface{}
}

type contentResponse struct {
	Id string
}

func ContentProcessor(messageChannel chan message) {
	var ContentMap map[string] Content
	ContentMap = make(map[string] Content)
	for {
		select {
		case msg, msg_ok := <-messageChannel:
			if !msg_ok {
				log.Println("ContentProcessor: Channel closed")
				return
			}
			log.Printf("%v", msg)
			switch msg.op {
			case opCreate:
				ContentMap[msg.id] = msg.value
			case opRead:
				msg.resp <- ContentMap[msg.id]
			case opReadAll:
				msg.resp <- ContentMap
			case opUpdate:
				ContentMap[msg.id] = msg.value
			case opDelete:
				delete(ContentMap, msg.id)
			}
		}
	}
}

type Content struct {
	Name string
	Fields []string
}

var messageChannel chan message

func Init() {
	messageChannel = make(chan message)
	go ContentProcessor(messageChannel)
}

func AllContent(response http.ResponseWriter, request *http.Request) {
	readChan := make(chan interface{})
	msg := message{op: opReadAll, resp: readChan}
	messageChannel <- msg
	resp := <- readChan
	out, jsonErr := json.Marshal(resp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(out)
}

func GetContent(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	log.Printf("GET content id = %s", id)
	readChan := make(chan interface{})
	msg := message{op: opRead, id: id, resp: readChan}
	messageChannel <- msg
	resp := <- readChan
	out, jsonErr := json.Marshal(resp)
	if RestError(jsonErr, response) {
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(out)
}

func PutContent(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var content Content
	jsonErr := json.Unmarshal(bodyBytes, &content)
	if RestError(jsonErr, response) {
		return
	}
	readChan := make(chan interface{})
	msg := message{op: opUpdate, id: id, value: content, resp: readChan}
	messageChannel <- msg
}

func PostContent(response http.ResponseWriter, request *http.Request) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if RestError(readErr, response) {
		return
	}
	var content Content
	jsonErr := json.Unmarshal(bodyBytes, &content)
	if RestError(jsonErr, response) {
		return
	}
	id := sha1sum(content)
	readChan := make(chan interface{})
	msg := message{op: opCreate, id: id, value: content, resp: readChan}
	messageChannel <- msg
	var resp struct{
		Id string
	}
	resp.Id = id
	response.Header().Set("Content-Type", "application/json")
	out, _ := json.Marshal(resp)
	response.Write(out)
}

func DeleteContent(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	msg := message{op: opDelete, id: id}
	messageChannel <- msg
}

func sha1sum(value interface{}) string {
	data, _ := json.Marshal(value)
	h := sha1.New()
	io.WriteString(h, string(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}
