package articleValidator

import (
	"github.com/saeidraei/go-realworld-clean/domain"
	"github.com/saeidraei/go-realworld-clean/uc"
)

type validator struct {
}

func New() uc.ArticleValidator {
	return validator{}
}

func (validator) BeforeCreationCheck(article *domain.Article) error { return nil }
func (validator) BeforeUpdateCheck(article *domain.Article) error   { return nil }
