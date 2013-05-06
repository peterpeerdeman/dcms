package site

import (
	"net/http"
	"html/template"
	"reflect"
	"log"
	"fmt"
	"runtime/debug"
)

func Site() {
	http.HandleFunc("/", HandleAll)
	http.Handle("/assets/", http.StripPrefix("/mysite/assets", http.FileServer(http.Dir("./mysite/assets"))))
}

func HandleAll(response http.ResponseWriter, request *http.Request) {
	var tmplVars struct {
		Menu string
	}
	tmplVars.Menu = "Home"
	tmplErr := Render(response, "home.tpl", tmplVars)
	if HttpError(tmplErr, response) {
		return
	}
}

func Render(response http.ResponseWriter, templateFile string, vars interface{}) error {
	t := template.New("bunch")
	t.Funcs(template.FuncMap{"eq": reflect.DeepEqual})
	_, parseErr := t.ParseFiles(fmt.Sprintf("mysite/templates/%s", templateFile))
	if parseErr != nil {
		return parseErr
	}
	execErr := t.ExecuteTemplate(response, templateFile, vars)
	if execErr != nil {
		return execErr
	}
	return nil
}

func HttpError(err error, response http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		response.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.WriteHeader(500)
		var tmplVars struct {
			Menu  string
			Error string
			Stack string
		}
		tmplVars.Error = fmt.Sprintf("%v", err)
		tmplVars.Stack = string(debug.Stack())
		renderErr := Render(response, "error.tpl", tmplVars)
		if renderErr != nil {
			log.Printf("%v", renderErr)
		}
		return true
	}
	return false
}
