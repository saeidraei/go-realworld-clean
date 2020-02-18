package validator

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/saeidraei/go-realworld-clean/domain"
	"github.com/saeidraei/go-realworld-clean/uc"
)

type urlValidator struct{}

func New() uc.UrlValidator {
	return urlValidator{}
}

func (urlValidator) BeforeCreationCheck(url *domain.Url) error {
	if ok := govalidator.IsURL(url.Address); !ok {
		return errors.New("invalid url")
	}

	return nil
}
