package mysite

import (
	"net/http"
	"math/rand"
)

type TestComponent struct {
}

func (this TestComponent) Render(response http.ResponseWriter, request *http.Request) map[string] interface{} {
	var out = map[string] interface{} {"var3": "hoi"}
	out["rand"] = rand.Intn(100)
	return out
}
