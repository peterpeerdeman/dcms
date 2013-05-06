package mysite

import (
	"net/http"
	"math/rand"
)

type NewsComponent struct {
}

func (this NewsComponent) Render(response http.ResponseWriter, request *http.Request) map[string] interface{} {
	var out = map[string] interface{} {"var3": "halloooo"}
	out["rand"] = rand.Intn(100)
	return out
}
