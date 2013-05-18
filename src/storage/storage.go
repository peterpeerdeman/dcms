package storage

import (
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
