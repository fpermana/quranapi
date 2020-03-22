// Package searching provides the use-case of searching.
package searching

import (
	"time"
	"github.com/go-kit/kit/log"
	"github.com/fpermana/quranapi/paging"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Search(keywords string, quran_text string, translation string, lastId int, limit int) ([]*paging.AyaModel, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Search",
			"keywords", keywords,
			"quran_text", quran_text,
			"translation", translation,
			"lastId", lastId,
			"limit", limit,
			"took", time.Since(begin),
		)
	}(time.Now())
        return s.next.Search(keywords, quran_text, translation, lastId, limit)
}

