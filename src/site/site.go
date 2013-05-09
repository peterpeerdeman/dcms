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
	"bytes"
	"regexp"
)

func Site() {
	go WatchConfiguration("mysite/configuration.json")
	http.HandleFunc("/", HandleAll)
	http.Handle("/assets/", http.StripPrefix("/mysite/assets", http.FileServer(http.Dir("./mysite/assets"))))
}

func HandleAll(response http.ResponseWriter, request *http.Request) {
	for channelName, channel := range SiteConfiguration.Channels {
		match, _ := matches(channelName, request.Host)
		if match {
			for sitemapURL, sitemapItem := range channel.Sitemap.Mapping {
				match, vars := matches(sitemapURL, request.RequestURI)
				if match {
					sitemapItem.Render(response, request, &channel, vars)
					return
				}
			}
		}
	}
	HttpError(errors.New("No page found"), response)
}

func matches(pattern, str string) (bool, map[string] string) {
	reg := regexp.MustCompile(pattern)
	match, vars := findStringSubmatchMap(reg, str)
	return match, vars
}

func findStringSubmatchMap(r *regexp.Regexp, s string) (bool, map[string] string) {
	captures := make(map[string] string)
	match := r.FindStringSubmatch(s)
	if match == nil {
		return false, captures
	}
	for i, name := range r.SubexpNames() {
		if i == 0 {
			continue
		}
		captures[name] = match[i]
	}
	return true, captures
}

func Render(templateFile string, vars map[string] interface{}) string {
	t := template.New("bunch")
	t.Funcs(template.FuncMap{"eq": reflect.DeepEqual})
	_, parseErr := t.ParseFiles(fmt.Sprintf("mysite/templates/%s", templateFile))
	if parseErr != nil {
		return fmt.Sprintf("%v", parseErr)
	}
	buffer := bytes.NewBufferString("")
	execErr := t.ExecuteTemplate(buffer, templateFile, vars)
	if execErr != nil {
		return fmt.Sprintf("%v", execErr)
	}
	return string(buffer.Bytes())
}

func HttpError(err error, response http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		response.Header().Set("Content-Type", "text/html; charset=utf-8")
		response.WriteHeader(500)
		tmplVars := make(map[string] interface{})
		tmplVars["Error"] = fmt.Sprintf("%v", err)
		tmplVars["Stack"] = string(debug.Stack())
		out := Render("error.tpl", tmplVars)
		response.Write([]byte(out))
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
