package server

import (
	//"context"
	//"encoding/json"
	"net/http"
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
	r.Get("/", h.getPage)

	//r.Method("GET", "/docs", http.StripPrefix("/booking/v1/docs", http.FileServer(http.Dir("booking/docs"))))

	return r
}

func (h *quranHandler) getPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
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

