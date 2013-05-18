package storage

type Repository struct {
	head tree
}

func Init() *Repository {
	var repo Repository
	return &repo
}

func (this *Repository) Get(filename string) ([]byte, error) {
	return []byte(""), nil
}

func (this *Repository) List(filename string) ([]string, error) {
	v := make([]string, 0)
	return v, nil
}

func (this *Repository) Add(filename string, content []byte) error {
	return nil
}

func (this *Repository) Remove(filename string) {

}

func (this *Repository) Copy(from string, to string) {

}

func (this *Repository) Move(from string, to string) {

}

func (this *Repository) Commit(message string) {

}

func (this *Repository) Branch(name string) {

}

func (this *Repository) Checkout(branche_name string) {

}
