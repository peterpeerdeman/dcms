package storage

import (
	"testing"
)

func Test_store(t *testing.T) {
	cur := newTree()
	storeErr := cur.store("/a/b/c", "deadb33f")
	testErr(t, storeErr)
	blobId, fetchErr := cur.fetch("/a/b/c")
	testErr(t, fetchErr)
	if blobId != "deadb33f" {
		t.Error("The blobId was not deadb33f.")
	}
}
