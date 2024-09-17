package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tokopedia/go-filter-parser"
	f "go-hdflex/external/filter"
	"go-hdflex/external/response"
	"go-hdflex/internal/api/service"
	"log/slog"
	"net/http"
	"strings"
)

type ContentHandler struct {
	log *slog.Logger

	s service.ContentServiceInterface
}

func NewContentHandler(
	log *slog.Logger,
	s service.ContentServiceInterface,
) *ContentHandler {
	return &ContentHandler{
		log: log,
		s:   s,
	}
}

func (h *ContentHandler) Get(w http.ResponseWriter, r *http.Request) {
	const op = "ContentHandler.Get() ->"

	log := h.log.With(
		"op", op,
	)

	paramsUrl := r.URL.RawQuery[:strings.LastIndex(r.URL.RawQuery, ";")]
	params := new(f.GetFilter)
	if err := filter.Parse(paramsUrl, params); err != nil {
		log.Error("Error parsing filter", slog.Any("error", err))
		response.Respond(w, response.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Error parsing filter: %v", err),
			Data:    nil,
		})
		return
	}

	ctx := context.WithValue(r.Context(), "filters", params)

	content, err := h.s.Get(ctx)
	if err != nil {
		log.Error("Error getting content", slog.Any("error", err))
		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error getting content",
			Data:    nil,
		})
		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    content.Items,
	})
}

func (h *ContentHandler) ShowBySlug(w http.ResponseWriter, r *http.Request) {
	const op = "ContentHandler.ShowBySlug() ->"

	slug := chi.URLParam(r, "slug")

	content, err := h.s.GetFirstBySlug(r.Context(), slug)
	if err != nil {
		h.log.Error("Error getting content by slug", fmt.Errorf("%s %w", op, err))

		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error getting content by slug",
			Data:    nil,
		})
		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    content,
	})
}

func (h *ContentHandler) Count(w http.ResponseWriter, r *http.Request) {
	const op = "ContentHandler.Count() ->"

	log := h.log.With(
		"op", op,
	)

	paramsUrl := r.URL.RawQuery[:strings.LastIndex(r.URL.RawQuery, ";")]
	filters := new(f.GetFilter)
	if err := filter.Parse(paramsUrl, filters); err != nil {
		log.Error("Error parsing filter", slog.Any("error", err))

		response.Respond(w, response.Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Error parsing filter: %v", err),
			Data:    nil,
		})
		return
	}

	ctx := context.WithValue(r.Context(), "filters", filters)

	count, err := h.s.Count(ctx)
	if err != nil {
		log.Error("Error getting content count", slog.Any("error", err))

		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error getting content count",
			Data:    nil,
		})
		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    count,
	})
}
