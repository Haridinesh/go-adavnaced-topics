package helper

import (
	"blogpost/models"

	"github.com/go-playground/validator"
)

func StructValidation(p models.Posts) error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}
