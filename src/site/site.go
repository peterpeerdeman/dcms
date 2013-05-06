package site

import (
	"net/http"
	"html/template"
	"reflect"
	"log"
	"fmt"
	"runtime/debug"
	"mysite"
	"os"
	"errors"
	"bytes"
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
	component_out := this.RenderComponent(response, request, channel)
	response.Write([]byte(component_out))
}

func (this *Page) RenderComponent(response http.ResponseWriter, request *http.Request, channel *Channel) string {
	vars := make(map[string] interface{})
	for name, value := range channel.Variables {
		vars[name] = value
	}
	if this.Component != "" {
		component_id, component_id_found := channel.Components[this.Component]
		if component_id_found {
			component, component_found := mysite.Components[component_id.ObjectName]
			if component_found {
				output := component.Render(response, request)
				for name, value := range output {
					vars[name] = value
				}
				for name, subPage := range this.SubPages {
					log.Printf("Rendering component %v", name)
					vars[name] = subPage.RenderComponent(response, request, channel)
				}
				templateFile := channel.Templates[this.Template].Filename
				t := template.New("bunch")
				t.Funcs(template.FuncMap{"eq": reflect.DeepEqual})
				_, parseErr := t.ParseFiles(fmt.Sprintf("mysite/templates/%s", templateFile))
				if parseErr != nil {
					 log.Fatalf("%v", parseErr)
				}
				buffer := bytes.NewBufferString("")
				execErr := t.ExecuteTemplate(buffer, templateFile, vars)
				if execErr != nil {
					log.Fatalf("%v", execErr)
				}
				return string(buffer.Bytes())

			} else {
				log.Printf("Component %s not found in mysite", component_id.ObjectName)
			}
		} else {
			log.Printf("Component name %s not found in configuration", this.Component)
		}
	}
	return ""
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
