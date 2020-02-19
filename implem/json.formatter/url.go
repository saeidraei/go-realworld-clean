package formatter

import (
	"github.com/saeidraei/go-realworld-clean/domain"
)

type Url struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

func NewUrlFromDomain(url domain.Url) Url {

	return Url{
		ID:      url.ID,
		Address: url.Address,
	}
}

func NewUrlsFromDomain(urls ...domain.Url) []Url {
	ret := []Url{} // return at least an empty array (not nil)

	for _, url := range urls {
		ret = append(ret, NewUrlFromDomain(url))
	}

	return ret
}
