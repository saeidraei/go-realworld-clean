package uc

import "github.com/saeidraei/go-realworld-clean/domain"

func (i interactor) UrlGet(id string) (*domain.Url, error) {

	url, err := i.urlRW.GetByID(id)
	if err != nil {
		return nil, err
	}

	return url, nil
}
