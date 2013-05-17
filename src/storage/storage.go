package storage

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func load(shasum string, value *interface{}) error {
	file, readErr := ioutil.ReadFile(fmt.Sprintf("data/%s", shasum))
	if readErr != nil {
		return readErr
	}
	jsonErr := json.Unmarshal(file, &value)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func save(value *interface{}) error {
	data, jsonErr := json.Marshal(value)
	if jsonErr != nil {
		return jsonErr
	}
	filename := fmt.Sprintf("data/%s", sha1sum(string(data)))
	writeErr := ioutil.WriteFile(filename, data, 0755)
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func sha1sum(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
