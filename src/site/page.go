package site

import (
	"fmt"
	"log"
	"mysite"
	"net/http"
	"strings"
)

type Page struct {
	Template    string
	Component   string
	ContentPath string
	Variables   map[string]string
	SubPages    map[string]Page
}

func (this *Page) Render(response http.ResponseWriter, request *http.Request, channel *Channel, requestVars map[string]string) {
	log.Printf("Component %s %v path %s", this.Component, requestVars, this.GetContentPath(requestVars))
	component_out := this.RenderComponent(response, request, channel, requestVars)
	response.Write([]byte(component_out))
}

func (this *Page) RenderComponent(response http.ResponseWriter, request *http.Request, channel *Channel, requestVars map[string]string) string {
	vars := make(map[string]interface{})
	for name, value := range channel.Variables {
		vars[name] = value
	}
	for name, value := range requestVars {
		vars[name] = value
	}
	for name, value := range this.Variables {
		vars[name] = value
	}
	if this.Component != "" {
		component_id, component_id_found := channel.Components[this.Component]
		if component_id_found {
			component, component_found := mysite.Components[component_id.ObjectName]
			if component_found {
				output := component.Render(response, request, vars)
				for name, value := range output {
					vars[name] = value
				}
				for name, subPage := range this.SubPages {
					vars[name] = subPage.RenderComponent(response, request, channel, requestVars)
				}
				templateFile := channel.Templates[this.Template].Filename
				return Render(templateFile, vars)

			} else {
				log.Printf("Component %s not found in mysite.", component_id.ObjectName)
			}
		} else {
			log.Printf("Component name %s not found in configuration.", this.Component)
		}
	}
	return ""
}

func (this *Page) GetContentPath(requestVars map[string]string) string {
	path := this.ContentPath
	for name, value := range requestVars {
		path = strings.Replace(path, fmt.Sprintf("{%s}", name), value, -1)
	}
	return path
}
