package resty

import (
	"fmt"
	"crypto/sha1"
	"io"
	"encoding/json"
)

func sha1sum(value interface{}) string {
	data, _ := json.Marshal(value)
	h := sha1.New()
	io.WriteString(h, string(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}
