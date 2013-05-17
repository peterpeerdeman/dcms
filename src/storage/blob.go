package storage

type Blob struct {
	Type string
	Data interface{}
}

func NewBlob(dataType string, data interface{}) *Blob {
	var blob Blob
	blob.Type = dataType
	blob.Data = data
	return &blob
}
