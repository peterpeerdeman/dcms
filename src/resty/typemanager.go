package resty

type TypeManager interface {
	Create() interface{}
	MakeCollection() TypeCollection
	Bootstrap(interface{}) interface{}
	Identifier(interface{}) string
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte) (interface{}, error)
}

type TypeCollection interface {
	Append(interface{})
	Serialize() ([]byte, error)
}
