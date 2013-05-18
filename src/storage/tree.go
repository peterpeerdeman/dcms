package storage

import (
	"errors"
	"fmt"
)

type tree struct {
	Root *node
}

func newTree() *tree {
	var cur tree
	cur.Root = newNode()
	return &cur
}

func (this *tree) list(filename string) ([]string, error) {
	node, err := this.resolve(filename, false)
	if err != nil {
		return make([]string, 0), err
	}
	var out []string
	for name, subNode := range node.Children {
		if subNode.BlobId != "" {
			out = append(out, name)
		}
	}
	return out, nil
}

func (this *tree) fetch(filename string) (string, error) {
	node, err := this.resolve(filename, false)
	if err != nil {
		return "", err
	}
	return node.BlobId, nil
}

func (this *tree) store(filename string, blobId string) error {
	node, err := this.resolve(filename, true)
	if err != nil {
		return err
	}
	node.BlobId = blobId
	return nil
}

func (this *tree) resolve(filename string, create bool) (*node, error) {
	cur := this.Root
	path := split_path(filename)
	for index, part := range path {
		last := index == len(path)-1
		next, ok := cur.Children[part]
		if !ok {
			if create {
				if last {
					// create a file
					next, err := cur.create(part, "")
					if err != nil {
						return next, err
					}
					return next, nil
				} else {
					// create a directory
					newDir, err := cur.mkdir(part)
					if err != nil {
						return newDir, err
					}
					next = newDir
				}
			} else {
				return cur, errors.New(fmt.Sprintf("Couldn't resolve %s", part))
			}
		}
		cur = next
	}
	return cur, nil
}

func (this *tree) remove(filename string) error {
	cur := this.Root
	path := split_path(filename)
	if len(path) == 0 {
		return errors.New("Cannot remove the root node.")
	}
	for index, part := range path {
		last := index == len(path)-1
		if last {
			delete(cur.Children, part)
			return nil
		}
		next, ok := cur.Children[part]
		if !ok {
			return errors.New(fmt.Sprintf("Couldn't resolve %s", part))
		}
		cur = next
	}
	return errors.New("Cannot find the removable node.")
}
