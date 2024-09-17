package handler

import (
	"fmt"
	"go-hdflex/external/response"
	"go-hdflex/internal/api/service"
	"log/slog"
	"net/http"
	"time"
)

type StorageHandler struct {
	log *slog.Logger

	s service.StorageServiceInterface
}

type StorageHandlerInterface interface {
}

func NewStorageHandler(
	log *slog.Logger,
	s service.StorageServiceInterface,
) *StorageHandler {
	return &StorageHandler{
		log: log,
		s:   s,
	}
}

func (h *StorageHandler) Resize(w http.ResponseWriter, r *http.Request) {
	const op = "StorageHandler.Resize() ->"

	data := h.s.GetStorageResizeData(r)

	file, path, err := h.s.Resize(data, false)
	if err != nil {
		h.log.Error("Error resizing image", fmt.Errorf("%s %w", op, err))
		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error resizing image",
			Data:    nil,
		})

		return
	}

	defer file.Close()

	http.ServeContent(w, r, path, time.Now(), file)
}

func (h *StorageHandler) ForceResize(w http.ResponseWriter, r *http.Request) {
	const op = "StorageHandler.ForceResize() ->"

	data := h.s.GetStorageResizeData(r)

	file, path, err := h.s.Resize(data, true)
	if err != nil {
		h.log.Error("Error resizing image", fmt.Errorf("%s %w", op, err))
		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error resizing image",
			Data:    nil,
		})

		return
	}

	defer file.Close()

	http.ServeContent(w, r, path, time.Now(), file)
}
