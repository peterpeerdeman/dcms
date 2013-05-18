package storage

import (
	"testing"
)

func Test_split_path_full(t *testing.T) {
	parts := split_path("/this/is/a/test")
	if len(parts) != 4 {
		t.Error("Parts has incorrect size.")
	}
	if parts[0] != "this" {
		t.Error("Parts does not have this.")
	}
	if parts[1] != "is" {
		t.Error("Parts does not have is.")
	}
	if parts[2] != "a" {
		t.Error("Parts does not have a.")
	}
	if parts[3] != "test" {
		t.Error("Parts does not have test.")
	}
}

func Test_split_path_one(t *testing.T) {
	parts := split_path("/this")
	if len(parts) != 1 {
		t.Error("Parts has incorrect size.")
	}
	if parts[0] != "this" {
		t.Error("Parts does not have this.")
	}
}

func Test_split_path_root(t *testing.T) {
	parts := split_path("/")
	if len(parts) != 0 {
		t.Error("Parts has incorrect size.")
	}
}

func Test_split_path_slashes(t *testing.T) {
	parts := split_path("/this//is")
	if len(parts) != 2 {
		t.Error("Parts has incorrect size.")
	}
	if parts[0] != "this" {
		t.Error("Parts does not have this.")
	}
	if parts[1] != "is" {
		t.Error("Parts does not have is.")
	}
}
