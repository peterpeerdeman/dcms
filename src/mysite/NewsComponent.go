package mysite

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"resty"
	"storage"
)

type NewsComponent struct {
}

func (this NewsComponent) Render(response http.ResponseWriter, request *http.Request, vars map[string]interface{}) map[string]interface{} {
	var out = map[string]interface{}{"var3": "halloooo"}
	out["rand"] = rand.Intn(100)
	docData, getErr := storage.Repo.Get(fmt.Sprintf("/documents/%s", vars["documentName"]))
	if getErr == nil {
		var doc resty.Document
		json.Unmarshal(docData, &doc)
		out["document"] = doc
	}
	return out
}
