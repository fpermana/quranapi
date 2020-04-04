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
	GetSuraPage(suraNumber int) (int, error)
}

type service struct {
	translations		TranslationRepository
	suras			SuraRepository
}

func (s *service) GetTranslations() ([]*TranslationModel, error) {
	var translationList []*TranslationModel
	translationList = s.translations.GetTranslationList()
	return translationList, nil
}

func (s *service) GetSuraPage(suraNumber int) (int, error) {
	var currentPage int = s.suras.GetSuraPage(SuraNumber(suraNumber))

	return  currentPage, nil
}

func NewService(tr TranslationRepository, sr SuraRepository) Service {
	return &service {
		translations: tr,
		suras: sr,
	}
}

