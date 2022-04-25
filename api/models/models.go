package models

import "fmt"

type Model interface {
	TODO
}

func NotFound[T Model](m T, id interface{}) error {
	return fmt.Errorf("%T not found: %v", m, id)
}
