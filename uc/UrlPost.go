package uc

import (
	"github.com/saeidraei/go-realworld-clean/domain"
	"math/rand"
)

func (i interactor) UrlPost(url domain.Url) (*domain.Url, error) {

	//if err := i.articleValidator.BeforeCreationCheck(&article); err != nil {
	//	return nil, nil, err
	//}
	url.ID = randStringBytes(7)
	completeUrl, err := i.urlRW.Create(url)
	if err != nil {
		return nil, err
	}

	return completeUrl, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
