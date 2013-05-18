package storage

import (
	"testing"
)

func Test_Init(t *testing.T) {
	repo := Init()
	listing, listingErr := repo.List("/")
	testErr(t, listingErr)
	if len(listing) > 0 {
		t.Error("Empty repo contains files!")
	}
}

func Test_Add(t *testing.T) {
	repo := Init()
	addErr := repo.Add("/test123", []byte("Test 123"))
	testErr(t, addErr)
	listing, listingErr := repo.List("/")
	testErr(t, listingErr)
	if !contains(listing, "test123") {
		t.Error("File test123 not found.")
	}
	t.Logf("%v", listing)
	get, getErr := repo.Get("/test123")
	testErr(t, getErr)
	t.Logf("%v", get)
	if string(get) != "Test 123" {
		t.Error("Content not identical.")
	}
	t.Logf("%v", get)
}

func Test_AddUnkownDir(t *testing.T) {
	repo := Init()
	addErr := repo.Add("/documents/test123", []byte("Test 123"))
	testErr(t, addErr)
	listing, listingErr := repo.List("/documents")
	testErr(t, listingErr)
	if !contains(listing, "test123") {
		t.Error("File test123 not found.")
	}
	t.Logf("%v", listing)
	get, getErr := repo.Get("/documents/test123")
	testErr(t, getErr)
	t.Logf("%v", get)
	if string(get) != "Test 123" {
		t.Error("Content not identical.")
	}
	t.Logf("%v", get)
}

func testErr(t *testing.T, e error) {
	if e != nil {
		t.Errorf("Error %v", e)
	}
}
