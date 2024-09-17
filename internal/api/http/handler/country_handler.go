package handler

import (
	"fmt"
	"go-hdflex/external/response"
	"go-hdflex/internal/api/service"
	"log/slog"
	"net/http"
)

type CountryHandler struct {
	log *slog.Logger
	s   service.CountryServiceInterface
}

func NewCountryHandler(
	log *slog.Logger,
	s service.CountryServiceInterface,
) *CountryHandler {
	return &CountryHandler{
		log: log,
		s:   s,
	}
}

func (h *CountryHandler) Get(w http.ResponseWriter, r *http.Request) {
	const op = "CountryHandler.Get() ->"

	countries, err := h.s.Get(r.Context())
	if err != nil {
		h.log.Error("Error getting countries", fmt.Errorf("%s %w", op, err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    countries,
	})
}
