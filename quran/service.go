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
	GetSuraPage(suraNumber SuraNumber) int
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

func (s *service) GetSuraPage(suraNumber SuraNumber) int {
	var currentPage, currentSura, currentAya int = s.suras.GetSuraPage(suraNumber)

	var cSuraNumber SuraNumber = SuraNumber(currentSura)
	if cSuraNumber == suraNumber && currentAya > 1 {
		currentPage--
	}

	return  currentPage
}

func NewService(tr TranslationRepository, sr SuraRepository) Service {
	return &service {
		translations: tr,
		suras: sr,
	}
}

