// Package paging provides the use-case of paging.
package paging

import (
	//"fmt"
	"errors"
	// "time"
)

// Service is the interface that provides booking methods.
type Service interface {
	GetPage(page int, quran_text string, translation string) ([]*AyaModel, error)
	GetTotalPage() (int, error)
}

type service struct {
	ayas			AyaRepository
	pages			PageRepository
}

func (s *service) GetPage(page int, quran_text string, translation string) ([]*AyaModel, error) {
	var ayaList []*AyaModel

	var i,j int = s.pages.GetSuraAyaStart(page)
	//fmt.Println(i,j)
	if i == -1 && j == -1 {
		return ayaList, errors.New("Sura and Aya not found")
	}
	var k,l int = s.pages.GetSuraAyaStart(page+1)
	//fmt.Println(k,l)
	if k == -1 && l == -1 {
		var ayaNumberStart AyaNumber = s.ayas.GetNumberBySuraAya(quran_text, i, j)
		ayaList = s.ayas.GetAyaListFrom(quran_text, translation, ayaNumberStart)
	} else {
		var ayaNumberStart AyaNumber = s.ayas.GetNumberBySuraAya(quran_text, i, j)
		var ayaNumberEnd AyaNumber = s.ayas.GetNumberBySuraAya(quran_text, k, l)
		ayaList = s.ayas.GetAyaListBetween(quran_text, translation, ayaNumberStart, ayaNumberEnd, false)
	}

	return ayaList, nil
}

func (s *service) GetTotalPage() (int, error) {
	var totalPage int = 604

	totalPage = s.pages.GetTotalPage()

	return totalPage, nil
}

func NewService(ar AyaRepository, pr PageRepository) Service {
	return &service {
		ayas: ar,
		pages: pr,
	}
}

