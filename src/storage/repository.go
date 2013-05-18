package storage

import (
	"log"
)

type Repository struct {
	Head *tree
}

func Init() *Repository {
	if repo_exists("head.json") {
		repo, err := load_repo("head.json")
		if err != nil {
			log.Fatalf("Unable to parse head.json %v", err)
		}
		return repo
	}
	var repo Repository
	repo.Head = newTree()
	save_repo("head.json", &repo)
	return &repo
}

func (this *Repository) Get(filename string) ([]byte, error) {
	blobId, fetchErr := this.Head.fetch(filename)
	if fetchErr != nil {
		return make([]byte, 0), fetchErr
	}
	data, loadErr := load_blob(blobId)
	if loadErr != nil {
		return data, loadErr
	}
	return data, nil
}

func (this *Repository) List(filename string) ([]string, error) {
	listing, err := this.Head.list(filename)
	if err != nil {
		return listing, err
	}
	return listing, nil
}

func (this *Repository) Add(filename string, content []byte) error {
	blobId, err := save_blob(content)
	if err != nil {
		return err
	}
	storeErr := this.Head.store(filename, blobId)
	if storeErr != nil {
		return storeErr
	}
	saveErr := save_repo("head.json", this)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func (this *Repository) Remove(filename string) {

}

func (this *Repository) Copy(from string, to string) {

}

func (this *Repository) Move(from string, to string) {

}

func (this *Repository) Merge(from string, to string) {

}

func (this *Repository) Commit(message string) {

}

func (this *Repository) Branch(name string) {

}

func (this *Repository) Checkout(branche_name string) {

}
