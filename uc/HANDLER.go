package uc

import (
	"log"

	"github.com/saeidraei/go-realworld-clean/domain"
)

type Handler interface {
	UrlLogic
}

type UrlLogic interface {
	UrlPost(url domain.Url) (*domain.Url, error)
	UrlGet(id string) (*domain.Url, error)
}

type HandlerConstructor struct {
	Logger       Logger
	UrlRW        UrlRW
	CacheRW      CacheRW
	UrlValidator UrlValidator
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.UrlRW == nil {
		log.Fatal("missing UrlRW")
	}

	return interactor{
		logger:       c.Logger,
		urlRW:        c.UrlRW,
		cacheRW:      c.CacheRW,
		urlValidator: c.UrlValidator,
	}
}
