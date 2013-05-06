package mysite

import "net/http"

var Components = map[string] Component {
	"mysite.TestComponent": TestComponent{},
	"mysite.NewsComponent": NewsComponent{}}

type Component interface {
	Render(response http.ResponseWriter, request *http.Request) map[string] interface{}
}
