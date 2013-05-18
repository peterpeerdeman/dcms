package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func split_path(path string) []string {
	if path == "/" {
		return make([]string, 0)
	}
	result := strings.Split(path, "/")
	out := make([]string, len(result))
	out_index := 0
	for _, part := range result[1:] {
		if part != "" {
			out[out_index] = part
			out_index++
		}
	}
	return out[:out_index]
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func repo_exists(name string) bool {
	if _, err := os.Stat(fmt.Sprintf("./data/%s", name)); err == nil {
		return true
	}
	return false
}

func load_repo(name string) (*Repository, error) {
	var repo Repository
	content, readErr := ioutil.ReadFile(fmt.Sprintf("./data/%s", name))
	if readErr != nil {
		return &repo, readErr
	}
	log.Printf("%v", string(content))
	jsonErr := json.Unmarshal(content, &repo)
	if jsonErr != nil {
		return &repo, jsonErr
	}
	return &repo, nil
}

func save_repo(name string, repo *Repository) error {
	content, jsonErr := json.Marshal(repo)
	if jsonErr != nil {
		return jsonErr
	}
	filename := fmt.Sprintf("./data/%s", name)
	writeErr := ioutil.WriteFile(filename, content, 0644)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
