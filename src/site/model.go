package site

import (
	"encoding/json"
	"io/ioutil"
)

type Template struct {
	Filename string
}

type Page struct {
	Template  string
	Component string
	SubPages  map[string] Page
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
	Components map[string] Component
	Templates map[string] Template
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

