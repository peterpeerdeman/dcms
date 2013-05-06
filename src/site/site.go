package site

import (
	"net/http"
	"html/template"
	"reflect"
	"log"
	"fmt"
	"runtime/debug"
	"os"
	"errors"
)

func Site() {
	readErr := ReadConfiguration("mysite/configuration.json")
	Fatal(readErr)
	log.Printf("%v", SiteConfiguration)
	http.HandleFunc("/", HandleAll)
	http.Handle("/assets/", http.StripPrefix("/mysite/assets", http.FileServer(http.Dir("./mysite/assets"))))
}

func HandleAll(response http.ResponseWriter, request *http.Request) {

	log.Printf("%v", request.RequestURI)

	for channelName, channel := range SiteConfiguration.Channels {
		if matches(channelName, request.Host) {
			for sitemapURL, sitemapItem := range channel.Sitemap.Mapping {
				if matches(request.RequestURI, sitemapURL) {
					sitemapItem.Render(response, request, &channel)
					return
				}
			}
		}
	}

	HttpError(errors.New("No page found"), response)

}

func (this *Page) Render(response http.ResponseWriter, request *http.Request, channel *Channel) {
	var tmplVars struct {
		Channel map[string] string
	}
	tmplVars.Channel = channel.Variables
	tmplErr := Render(response, this.Template.Filename, tmplVars)
	if HttpError(tmplErr, response) {
		return
	}
}

func matches(pattern, str string) bool {
	return pattern == str
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


func Fatal(err error) {
	if err != nil {
		log.Fatalf("%v", err)
		debug.PrintStack()
		os.Exit(-1)
	}
}
