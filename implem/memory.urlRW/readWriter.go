package urlRw

import (
	"sync"

	"errors"

	"log"

	"github.com/saeidraei/go-realworld-clean/domain"
	"github.com/saeidraei/go-realworld-clean/uc"
)

type rw struct {
	store *sync.Map
}

func New() uc.UrlRW {
	return rw{
		store: &sync.Map{},
	}
}
func (rw rw) Create(url domain.Url) (*domain.Url, error) {
	if _, err := rw.GetByID(url.ID); err == nil {
		log.Println(err)
		return nil, uc.ErrAlreadyInUse
	}
	rw.store.Store(url.ID, url)

	return rw.GetByID(url.ID)
}

func (rw rw) Save(url domain.Url) (*domain.Url, error) {
	if _, err := rw.GetByID(url.ID); err != nil {
		return nil, uc.ErrNotFound
	}

	rw.store.Store(url.ID, url)

	return rw.GetByID(url.ID)
}

func (rw rw) GetByID(id string) (*domain.Url, error) {
	value, ok := rw.store.Load(id)
	if !ok {
		return nil, uc.ErrNotFound
	}

	url, ok := value.(domain.Url)
	if !ok {
		return nil, errors.New("not an url stored at key")
	}

	return &url, nil
}

func (rw rw) Delete(slug string) error {
	rw.store.Delete(slug)

	return nil
}
