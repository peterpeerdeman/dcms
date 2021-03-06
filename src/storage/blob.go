package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type blob []byte

func load_blob(shasum string) (blob, error) {
	content, readErr := ioutil.ReadFile(fmt.Sprintf("./data/%s", shasum))
	if readErr != nil {
		return content, readErr
	}
	return content, nil
}

func save_blob(content blob) (string, error) {
	blobId := sha1sum(string(content))
	filename := fmt.Sprintf("./data/%s", blobId)
	f, createErr := os.Create(filename)
	if createErr != nil {
		return blobId, createErr
	}
	defer f.Close()
	_, writeErr := f.Write(content)
	if writeErr != nil {
		return blobId, writeErr
	}
	return blobId, nil
}

func sha1sum(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
