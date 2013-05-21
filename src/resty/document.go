package resty

import (
	"encoding/json"
	"github.com/extemporalgenome/slug"
)

type Document struct {
	Id     string
	Path   string
	Type   string
	Name   string
	Fields map[string]interface{}
}

type DocumentManager struct {
}

func (this *DocumentManager) Create() interface{} {
	doc := new(Document)
	doc.Fields = make(map[string]interface{})
	return doc
}

func (this *DocumentManager) MakeCollection() TypeCollection {
	collection := new(DocumentCollection)
	collection.documents = make([]Document, 0)
	return collection
}

func (this *DocumentManager) Bootstrap(in interface{}) interface{} {
	doc := in.(*Document)
	doc.Id = slug.Slug(doc.Name)
	return doc
}

func (this *DocumentManager) Serialize(in interface{}) ([]byte, error) {
	data, err := json.Marshal(in)
	return data, err
}

func (this *DocumentManager) Deserialize(data []byte) (interface{}, error) {
	var doc Document
	err := json.Unmarshal(data, &doc)
	return doc, err
}

func (this *DocumentManager) Identifier(in interface{}) string {
	return in.(*Document).Id
}

type DocumentCollection struct {
	documents []Document
}

func (this *DocumentCollection) Append(document interface{}) {
	this.documents = append(this.documents, document.(Document))
}

func (this *DocumentCollection) Serialize() ([]byte, error) {
	data, err := json.Marshal(this.documents)
	return data, err
}
