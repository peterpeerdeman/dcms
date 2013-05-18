package storage

import (
	"errors"
	"fmt"
)

type node struct {
	blobId   string
	parent   *node
	children map[string]node
}

func newNode() *node {
	var cur node
	cur.children = make(map[string]node)
	return &cur
}

func (this *node) mkdir(name string) (*node, error) {
	var cur node
	if this.Exists(name) {
		return nil, errors.New(fmt.Sprintf("A node with name %s aleady exists.", name))
	}
	cur.parent = this
	this.children[name] = cur
	return &cur, nil
}

func (this *node) create(name string, blobId string) (*node, error) {
	var cur node
	if this.Exists(name) {
		return nil, errors.New(fmt.Sprintf("A node with name %s aleady exists.", name))
	}
	cur.blobId = blobId
	cur.parent = this
	this.children[name] = cur
	return &cur, nil
}

func (this *node) Exists(name string) bool {
	_, exists := this.children[name]
	return exists
}
