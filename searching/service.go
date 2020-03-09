// Package searching provides the use-case of searching.
package searching

import (
	//"fmt"
	//"errors"
	// "time"
	"github.com/fpermana/quranapi/paging"
)

// Service is the interface that provides searching methods.
type Service interface {
	Search(keywords string, quran_text string, translation string, lastId int, limit int) ([]*paging.AyaModel, error)
}

type service struct {
	searching		SearchingRepository
}

func (s *service) Search(keywords string, quran_text string, translation string, lastId int, limit int) ([]*paging.AyaModel, error) {
	var ayaList []*paging.AyaModel
	ayaList = s.searching.Search(keywords, quran_text, translation, lastId, limit)
	return ayaList, nil
}

func NewService(sr SearchingRepository) Service {
	return &service {
		searching: sr,
	}
}

