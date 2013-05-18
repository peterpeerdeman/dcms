package mysite

import (
	"math/rand"
	"net/http"
)

type TestComponent struct {
}

func (this TestComponent) Render(response http.ResponseWriter, request *http.Request, vars map[string]interface{}) map[string]interface{} {
	var out = map[string]interface{}{"var3": "hoi"}
	out["rand"] = rand.Intn(100)
	return out
}
