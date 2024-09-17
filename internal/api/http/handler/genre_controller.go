package handler

import (
	"fmt"
	"go-hdflex/external/response"
	"go-hdflex/internal/api/service"
	"log/slog"
	"net/http"
)

type GenreHandler struct {
	log *slog.Logger
	s   service.GenreServiceInterface
}

func NewGenreHandler(
	log *slog.Logger,
	s service.GenreServiceInterface,
) *GenreHandler {
	return &GenreHandler{
		log: log,
		s:   s,
	}
}

func (h *GenreHandler) Get(w http.ResponseWriter, r *http.Request) {
	const op = "GenreHandler.Get() ->"

	genres, err := h.s.Get(r.Context())
	if err != nil {
		h.log.Error("Error getting genres", fmt.Errorf("%s %w", op, err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    genres,
	})
}
