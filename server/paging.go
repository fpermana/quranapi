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

	"github.com/fpermana/quranapi/paging"
)

type pagingHandler struct {
	s paging.Service

	logger kitlog.Logger
}

func (h *pagingHandler) router() chi.Router {
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
	r.Get("/{page}", h.getPage)

	//r.Method("GET", "/docs", http.StripPrefix("/booking/v1/docs", http.FileServer(http.Dir("booking/docs"))))

	return r
}

func (h *pagingHandler) getPage(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		return
	}
	quran_text := "quran_text_original"
	translation := "id_indonesian"
	u := r.URL
	if u.RawQuery != "" {
		m, err := url.ParseQuery(u.RawQuery)
		if err == nil {
			for k, v := range m {
				switch k {
				case "quran_text":
					quran_text = v[0]
				case "translation":
					translation = v[0]
				}
			}
		}
	}

	ayaList, _ := h.s.GetPage(page, quran_text, translation)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(ayaList); err != nil {
		h.logger.Log("error", err)
		//encodeError(ctx, err, w)
		return
	}
}

