package storage

import (
	"errors"
	"fmt"
)

type node struct {
	BlobId   string
	Children map[string]*node
}

func newNode() *node {
	var cur node
	cur.Children = make(map[string]*node)
	return &cur
}

func (this *node) mkdir(name string) (*node, error) {
	if this.exists(name) {
		return nil, errors.New(fmt.Sprintf("A node with name %s aleady exists.", name))
	}
	cur := newNode()
	this.Children[name] = cur
	return cur, nil
}

func (this *node) create(name string, blobId string) (*node, error) {
	if this.exists(name) {
		return nil, errors.New(fmt.Sprintf("A node with name %s aleady exists.", name))
	}
	cur := newNode()
	cur.BlobId = blobId
	this.Children[name] = cur
	return cur, nil
}

func (this *node) exists(name string) bool {
	_, exists := this.Children[name]
	return exists
}
