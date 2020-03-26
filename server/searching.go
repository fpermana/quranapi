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

	"github.com/fpermana/quranapi/searching"
)

type searchingHandler struct {
	s searching.Service

	logger kitlog.Logger
}

func (h *searchingHandler) router() chi.Router {
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
	r.Get("/", h.search)

	//r.Method("GET", "/docs", http.StripPrefix("/booking/v1/docs", http.FileServer(http.Dir("booking/docs"))))

	return r
}

func (h *searchingHandler) search(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
	//Search(keywords string, quran_text string, translation string, lastId int, limit int) ([]*paging.AyaModel, error
	keyword := ""
	quran_text := "quran_text_original"
	translation := "id_indonesian"
	lastId := 1
	limit := 10
	u := r.URL
	if u.RawQuery != "" {
		m, err := url.ParseQuery(u.RawQuery)
		if err == nil {
			for k, v := range m {
				switch k {
				case "q":
					keyword = v[0]
				case "quran_text":
					quran_text = v[0]
				case "translation":
					translation = v[0]
				case "last_id":
					lastId, _ = strconv.Atoi(v[0])
				case "limit":
					limit, _ = strconv.Atoi(v[0])
				}
			}
		}
	}

	if keyword == "" {
		return
	}
	ayaList, _ := h.s.Search(keyword,quran_text,translation,lastId,limit)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(ayaList); err != nil {
		h.logger.Log("error", err)
		//encodeError(ctx, err, w)
		return
	}
	/*ctx := context.Background()

	trackingID := shipping.TrackingID(chi.URLParam(r, "trackingID"))

	c, err := h.s.LoadCargo(trackingID)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		Cargo booking.Cargo `json:"cargo"`
	}{
		Cargo: c,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}*/
}

