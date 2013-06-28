package storage

type commit struct {
	Tree    tree
	Parent  *commit
	Message string
}

func newCommit(parent *commit, message string, tree *tree) *commit {
	cur := new(commit)
	cur.Parent = parent
	cur.Message = message
	cur.Tree = *tree
	return cur
}
