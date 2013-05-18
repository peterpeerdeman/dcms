package resty

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

func sha1sum(value interface{}) string {
	data, _ := json.Marshal(value)
	h := sha1.New()
	io.WriteString(h, string(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func uuid() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "err"
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid)
}
