package uc

import (
	"github.com/saeidraei/go-realworld-clean/domain"
	"time"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger           Logger
	urlRW            UrlRW
	cacheRW          CacheRW
	urlValidator     UrlValidator
}

// Logger : only used to log stuff
type Logger interface {
	Log(...interface{})
}


type UrlRW interface {
	Create(domain.Url) (*domain.Url, error)
	Save(domain.Url) (*domain.Url, error)
	GetByID(ID string) (*domain.Url, error)
	Delete(ID string) error
}

type CacheRW interface {
	Set(key string , value interface{},ttl time.Duration) error
	Get(key string) (interface{}, error)
}

type UrlValidator interface {
	BeforeCreationCheck(url *domain.Url) error
}