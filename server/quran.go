package server

import (
	//"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	//"time"
	//"fmt"

	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"

	"github.com/fpermana/quranapi/quran"
)

type quranHandler struct {
	s quran.Service

	logger kitlog.Logger
}

func (h *quranHandler) router() chi.Router {
	r := chi.NewRouter()

	/*r.Route("/paging", func(r chi.Router) {
		//r.Post("/", h.bookCargo)
		r.Get("/", h.getPage)
		r.Route("/{trackingID}", func(r chi.Router) {
			r.Get("/", h.loadCargo)
			r.Get("/request_routes", h.requestRoutes)
			r.Post("/assign_to_route", h.assignToRoute)
			r.Post("/change_destination", h.changeDestination)
		})

	})*/
	//r.Get("/locations", h.listLocations)
	r.Get("/translations", h.getTranslations)
	r.Get("/sura-page", h.getSuraPage)

	//r.Method("GET", "/docs", http.StripPrefix("/booking/v1/docs", http.FileServer(http.Dir("booking/docs"))))

	return r
}

func (h *quranHandler) getTranslations(w http.ResponseWriter, r *http.Request) {

	translationList, _ := h.s.GetTranslations()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(translationList); err != nil {
		h.logger.Log("error", err)
		//encodeError(ctx, err, w)
		return
	}
}

func (h *quranHandler) getSuraPage(w http.ResponseWriter, r *http.Request) {
	id := 0
	u := r.URL
	if u.RawQuery != "" {
		m, err := url.ParseQuery(u.RawQuery)
		if err == nil {
			for k, v := range m {
				switch k {
				case "id":
					id, _ = strconv.Atoi(v[0])
				}
			}
		}
	}

	if id <= 0 {
		return
	}
	page, _ := h.s.GetSuraPage(id)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(page); err != nil {
		h.logger.Log("error", err)
		//encodeError(ctx, err, w)
		return
	}
}

