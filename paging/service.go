// Package paging provides the use-case of paging.
package paging

import (
	// "errors"
	// "time"
)

// Service is the interface that provides booking methods.
type Service interface {
	// GetPage(page int, qurant_text string, translation string) ([]*AyaModel, error)
}

type service struct {
	ayas			AyaRepository
}

/*func (s *service) GetPage(page int, qurant_text string, translation string) ([]*AyaModel, error) {

}*/