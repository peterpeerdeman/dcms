package resty

import (
	"encoding/json"
	"github.com/extemporalgenome/slug"
	"log"
)

type DocumentTypeField struct {
	Name string
	Type string
	Max  int
	Min  int
}

type DocumentType struct {
	Id     string
	Name   string
	Fields []DocumentTypeField
}

type DocumentTypeManager struct {
}

func (this *DocumentTypeManager) Create() interface{} {
	data := new(DocumentType)
	data.Fields = make([]DocumentTypeField, 0)
	return data
}

func (this *DocumentTypeManager) MakeCollection() TypeCollection {
	collection := new(DocumentTypeCollection)
	collection.data = make([]DocumentType, 0)
	return collection
}

func (this *DocumentTypeManager) Bootstrap(in interface{}) interface{} {
	doc := in.(*DocumentType)
	doc.Id = slug.Slug(doc.Name)
	return doc
}

func (this *DocumentTypeManager) Serialize(in interface{}) ([]byte, error) {
	data, err := json.Marshal(in)
	return data, err
}

func (this *DocumentTypeManager) Deserialize(data []byte) (interface{}, error) {
	var doc DocumentType
	log.Printf("%v", string(data))
	err := json.Unmarshal(data, &doc)
	return doc, err
}

func (this *DocumentTypeManager) Identifier(in interface{}) string {
	return in.(*DocumentType).Id
}

type DocumentTypeCollection struct {
	data []DocumentType
}

func (this *DocumentTypeCollection) Append(in interface{}) {
	this.data = append(this.data, in.(DocumentType))
}

func (this *DocumentTypeCollection) Serialize() ([]byte, error) {
	data, err := json.Marshal(this.data)
	return data, err
}
