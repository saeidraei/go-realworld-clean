package uc

import (
	"github.com/saeidraei/go-realworld-clean/domain"
)

func (i interactor) UrlPost(url domain.Url) (*domain.Url, error) {


	//if err := i.articleValidator.BeforeCreationCheck(&article); err != nil {
	//	return nil, nil, err
	//}

	completeUrl, err := i.urlRW.Create(url)
	if err != nil {
		return nil, err
	}


	return completeUrl, nil
}
