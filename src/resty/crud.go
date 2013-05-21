package resty

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"storage"
)

func BuildQuery(Manager TypeManager, Path string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		docs, listErr := storage.Repo.List(fmt.Sprintf("/%s", Path))
		if listErr != nil {
			return
		}
		resp := Manager.MakeCollection()
		for _, file := range docs {
			data, getErr := storage.Repo.Get(fmt.Sprintf("/%s/%s", Path, file))
			if getErr == nil {
				element, err := Manager.Deserialize(data)
				if RestError(err, response) {
					return
				}
				resp.Append(element)
			} else {
				log.Printf("Unable to get %v", file)
			}
		}
		out, err := resp.Serialize()
		if RestError(err, response) {
			return
		}
		response.Header().Set("Document-Type", "application/json")
		response.Write(out)
	}
}

func BuildGet(Manager TypeManager, Path string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id := vars["id"]
		out, getErr := storage.Repo.Get(fmt.Sprintf("/%s/%s", Path, id))
		if RestError(getErr, response) {
			return
		}
		response.Header().Set("Document-Type", "application/json")
		response.Write(out)
	}
}

func BuildPut(Manager TypeManager, Path string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id := vars["id"]
		bodyBytes, readErr := ioutil.ReadAll(request.Body)
		if RestError(readErr, response) {
			return
		}
		addErr := storage.Repo.Add(fmt.Sprintf("/%s/%s", Path, id), bodyBytes)
		if RestError(addErr, response) {
			return
		}
		response.Header().Set("Document-Type", "application/json")
		response.Write(bodyBytes)
	}
}

func BuildPost(Manager TypeManager, Path string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		bodyBytes, readErr := ioutil.ReadAll(request.Body)
		if RestError(readErr, response) {
			return
		}
		entity := Manager.Create()
		jsonErr := json.Unmarshal(bodyBytes, &entity)
		if RestError(jsonErr, response) {
			return
		}

		log.Printf("----> %v", entity)

		entity = Manager.Bootstrap(entity)
		out, marsErr := Manager.Serialize(entity)
		if RestError(marsErr, response) {
			return
		}
		addErr := storage.Repo.Add(fmt.Sprintf("/%s/%s", Path, Manager.Identifier(entity)), out)
		if RestError(addErr, response) {
			return
		}
		response.Header().Set("Document-Type", "application/json")
		response.Write(out)
	}
}

func BuildDelete(Manager TypeManager, Path string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id := vars["id"]
		storage.Repo.Remove(fmt.Sprintf("/%s/%s", Path, id))
		response.Header().Set("Document-Type", "application/json")
	}
}
