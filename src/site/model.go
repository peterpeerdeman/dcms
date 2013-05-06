package site

import (
	"encoding/json"
	"io/ioutil"
)

type Template struct {
	Filename string
}

type Page struct {
	Template  Template
	Component Component
}

type Component struct {
	ObjectName string
}

type Sitemap struct {
	Mapping map[string] Page
}

type Channel struct {
	Sitemap  Sitemap
	Variables map[string] string
}

var SiteConfiguration struct {
	Channels map[string] Channel
}

func ReadConfiguration(configFile string) error {
	file, readErr := ioutil.ReadFile(configFile)
	if readErr != nil {
		return readErr
	}
	jsonErr := json.Unmarshal(file, &SiteConfiguration)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

