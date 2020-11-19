package verifier

import (
	"errors"
)

var ErrInvalidKey = errors.New("Invalid verification key")

type Verifier interface {
	Request(id, key string) error
	Verify(id, key string) (bool, error)
}

type Mock struct {
	data map[string]string
}

func (v *Mock) Request(id, key string) error {
	v.data[key] = id
	return nil
}

func (v *Mock) Verify(id, key string) (bool, error) {
	have, exists := v.data[key]
	if !exists {
		return false, ErrInvalidKey
	}
	if have != id {
		return false, ErrInvalidKey
	}
	// mark user as verified in DB
	return true, nil
}