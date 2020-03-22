// Package paging provides the use-case of paging.
package paging

import (
	"time"
	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) GetPage(page int, quran_text string, translation string) ([]*AyaModel, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetPage",
			"page", page,
			"quran_text", quran_text,
			"translation", translation,
			"took", time.Since(begin),
		)
	}(time.Now())
        return s.next.GetPage(page, quran_text, translation)
}

