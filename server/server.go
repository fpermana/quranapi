package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	//"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/fpermana/quranapi/paging"
	"github.com/fpermana/quranapi/quran"
	"github.com/fpermana/quranapi/searching"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Paging		paging.Service
	Quran		quran.Service
	Searching	searching.Service

	Logger kitlog.Logger

	router chi.Router
}

// New returns a new HTTP server.
func New(ps paging.Service, qs quran.Service, ss searching.Service, logger kitlog.Logger) *Server {
	s := &Server{
		Paging:  ps,
		Quran: qs,
		Searching: ss,
		Logger:   logger,
	}

	r := chi.NewRouter()

	r.Use(accessControl)

	r.Route("/paging", func(r chi.Router) {
		h := pagingHandler{s.Paging, s.Logger}
		r.Mount("/v1", h.router())
	})
	r.Route("/quran", func(r chi.Router) {
		h := quranHandler{s.Quran, s.Logger}
		r.Mount("/v1", h.router())
	})
	r.Route("/searching", func(r chi.Router) {
		h := searchingHandler{s.Searching, s.Logger}
		r.Mount("/v1", h.router())
	})

	//r.Method("GET", "/metrics", promhttp.Handler())*/

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	/*case shipping.ErrUnknownCargo:
		w.WriteHeader(http.StatusNotFound)
	case tracking.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)*/
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

