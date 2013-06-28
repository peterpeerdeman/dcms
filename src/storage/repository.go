package storage

import (
	"log"
)

type Repository struct {
	Master *commit
	Head   *tree
}

var Repo *Repository

func Init() {
	if repo_exists("head.json") {
		log.Printf("Loading existing repository.")
		repo, err := load_repo("head.json")
		if err != nil {
			log.Fatalf("Unable to parse head.json %v", err)
		}
		Repo = repo
		return
	}
	log.Printf("Creating new repository.")
	newRepository := new(Repository)
	newRepository.Head = newTree()
	save_repo("head.json", newRepository)
	Repo = newRepository
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

func (this *Repository) Remove(filename string) error {
	return this.Head.remove(filename)
}

func (this *Repository) Copy(from string, to string) {

}

func (this *Repository) Move(from string, to string) {

}

func (this *Repository) Merge(from string, to string) {

}

func (this *Repository) Commit(message string) error {
	commit := newCommit(this.Master, message, this.Head)
	this.Master = commit
	saveErr := save_repo("head.json", this)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func (this *Repository) Branch(name string) {

}

func (this *Repository) Checkout(branche_name string) {

}
