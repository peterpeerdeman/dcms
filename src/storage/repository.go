package storage

type Repository struct {
}

func Init() *Repository {
	var repo Repository
	return &repo
}

func (this *Repository) Get(filename string) interface{} {
	return ""
}

func (this *Repository) List(filename string) []string {
	v := make([]string, 0)
	return v
}

func (this *Repository) Add(filename string, content interface{}) {

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
