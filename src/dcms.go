package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alleveenstra/godb"
	"github.com/gorilla/mux"
	"github.com/ziutek/mymysql/autorc"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"resty"
	"runtime/debug"
)

type Configuration struct {
	SqlHostname   string
	SqlUser       string
	SqlPassword   string
	SqlDatabase   string
	ListenAddress string
}

func main() {

	var config Configuration
	var configFile string

	flag.StringVar(&configFile, "config", "config.json", "The configuration file.")
	flag.Parse()

	config, configErr := ReadConfiguration(configFile)
	if configErr != nil {
		flag.Usage()
		os.Exit(1)
	}

	config.ListenAddress = "localhost:8080"

	if 1 == 2 {
		godb.Database = autorc.New("tcp", "", config.SqlHostname, config.SqlUser, config.SqlPassword, config.SqlDatabase)
		godb.Database.Register("SET NAMES utf8")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	resty.Init()
	router.HandleFunc("/rest/document", resty.AllDocument).Methods("GET")
	router.HandleFunc("/rest/document/{path:.*}", resty.PostDocument).Methods("POST")
	router.HandleFunc("/rest/document/{id}", resty.GetDocument).Methods("GET")
	router.HandleFunc("/rest/document/{id}", resty.PutDocument).Methods("PUT")
	router.HandleFunc("/rest/document/{id}", resty.DeleteDocument).Methods("DELETE")

	http.Handle("/", router)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(config.ListenAddress, nil)
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	var tmplVars struct {
		Menu string
	}
	tmplVars.Menu = "Home"
	tmplErr := Render(response, "home.tpl", tmplVars)
	if HttpError(tmplErr, response) {
		return
	}
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

func Render(response http.ResponseWriter, templateFile string, vars interface{}) error {
	t := template.New("bunch")
	t.Funcs(template.FuncMap{"eq": reflect.DeepEqual})
	_, parseErr := t.ParseFiles(fmt.Sprintf("templates/%s", templateFile), "templates/header.tpl", "templates/footer.tpl")
	if parseErr != nil {
		return parseErr
	}
	execErr := t.ExecuteTemplate(response, templateFile, vars)
	if execErr != nil {
		return execErr
	}
	return nil
}

func ReadConfiguration(configFile string) (Configuration, error) {
	var config Configuration
	file, readErr := ioutil.ReadFile(configFile)
	if readErr != nil {
		return config, readErr
	}
	jsonErr := json.Unmarshal(file, &config)
	Fatal(jsonErr)
	return config, nil
}

func Fatal(err error) {
	if err != nil {
		log.Fatalf("%v", err)
		debug.PrintStack()
		os.Exit(-1)
	}
}
