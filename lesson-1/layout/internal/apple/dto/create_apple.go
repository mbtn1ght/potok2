package dto

import (
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/google/uuid"
)

type CreateAppleOutput struct {
	ID uuid.UUID `json:"id"`
}

type CreateAppleInput struct {
	Name string `json:"name"`
}

func (i *CreateAppleInput) Validate() error {
	if i.Name == "" {
		return entity.ErrNameInvalid
	}

	return nil
}
