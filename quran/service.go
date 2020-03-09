// Package quran provides the use-case of quran.
package quran

import (
	//"fmt"
	//"errors"
	// "time"
)

// Service is the interface that provides quran methods.
type Service interface {
	GetTranslations() ([]*TranslationModel, error)
}

type service struct {
	translations		TranslationRepository
}

func (s *service) GetTranslations() ([]*TranslationModel, error) {
	var translationList []*TranslationModel
	translationList = s.translations.GetTranslationList()
	return translationList, nil
}

func NewService(tr TranslationRepository) Service {
	return &service {
		translations: tr,
	}
}

