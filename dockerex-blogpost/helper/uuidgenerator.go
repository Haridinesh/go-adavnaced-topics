package helper

import "github.com/google/uuid"

func Uuidgen() string {
	return uuid.NewString()
}
