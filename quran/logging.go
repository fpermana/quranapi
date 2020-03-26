// Package quran provides the use-case of quran.
package quran

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

func (s *loggingService) GetTranslations() ([]*TranslationModel, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetTranslations",
			"took", time.Since(begin),
		)
	}(time.Now())
        return s.next.GetTranslations()
}

func (s *loggingService) GetSuraPage(suraNumber int) (int, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetSuraPage",
			"suraNumber", suraNumber,
			"took", time.Since(begin),
		)
	}(time.Now())
        return s.next.GetSuraPage(suraNumber)
}

