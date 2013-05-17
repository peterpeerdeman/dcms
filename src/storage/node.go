package storage

type Node struct {
	BlobId   string
	Parent   *Node
	Children map[string]Node
}

func NewNode() *Node {
	var node Node
	node.Children = make(map[string]Node)
	return &node
}

func (this *Node) Mkdir(name string) *Node {
	var node Node
	node.Parent = this
	this.Children[name] = node
	return &node
}

func (this *Node) Create(name string, blobId string) *Node {
	var node Node
	node.BlobId = blobId
	node.Parent = this
	this.Children[name] = node
	return &node
}
