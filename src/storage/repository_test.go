package storage

import (
	"testing"
)

func Test_Init(t *testing.T) {
	repo := Init()
	if len(repo.List("/")) > 0 {
		t.Error("Empty repo contains files!")
	}
}
