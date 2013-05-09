package site

import (
	"encoding/json"
	"io/ioutil"
)

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
