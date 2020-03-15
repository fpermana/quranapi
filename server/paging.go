package server

import (
	//"context"
	//"encoding/json"
	"net/http"
	//"time"

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

	r.Route("/paging", func(r chi.Router) {
		//r.Post("/", h.bookCargo)
		r.Get("/", h.getPage)
		/*r.Route("/{trackingID}", func(r chi.Router) {
			r.Get("/", h.loadCargo)
			r.Get("/request_routes", h.requestRoutes)
			r.Post("/assign_to_route", h.assignToRoute)
			r.Post("/change_destination", h.changeDestination)
		})*/

	})
	//r.Get("/locations", h.listLocations)

	//r.Method("GET", "/docs", http.StripPrefix("/booking/v1/docs", http.FileServer(http.Dir("booking/docs"))))

	return r
}

func (h *pagingHandler) getPage(w http.ResponseWriter, r *http.Request) {
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

/*func (h *bookingHandler) listCargos(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	cs := h.s.Cargos()

	var response = struct {
		Cargos []booking.Cargo `json:"cargos"`
	}{
		Cargos: cs,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
*/
