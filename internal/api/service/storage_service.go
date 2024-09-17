package service

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/h2non/bimg"
	_struct "go-hdflex/internal/api/struct"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type StorageService struct {
	log *slog.Logger

	cachePath string
}

type StorageServiceInterface interface {
	Resize(_struct.StorageResize, bool) (*os.File, string, error)
	GetStorageResizeData(*http.Request) _struct.StorageResize
}

func NewStorageService(
	log *slog.Logger,
) *StorageService {
	return &StorageService{
		log:       log,
		cachePath: fmt.Sprintf("%s/.cache", os.Getenv("STORAGE_PATH")),
	}
}

func (s *StorageService) Resize(data _struct.StorageResize, force bool) (*os.File, string, error) {
	const op = "StorageService.Resize() ->"

	cachePath := fmt.Sprintf("%s/%dx%d/%s", s.cachePath, data.Width, data.Height, data.Path)
	if f, err := s.GetCachedFile(cachePath); err == nil {
		return f, cachePath, nil
	}

	oFilePath := fmt.Sprintf("%s/%s", os.Getenv("STORAGE_PATH"), data.Path)

	file, err := bimg.Read(oFilePath)
	if err != nil {
		return nil, "", fmt.Errorf("%s error reading file: %w", op, err)
	}

	img := bimg.NewImage(file)
	var rByteFile []byte

	if !force {
		rByteFile, err = img.Resize(data.Width, data.Height)
	} else {
		rByteFile, err = img.ForceResize(data.Width, data.Height)
	}
	if err != nil {
		return nil, "", fmt.Errorf("%s error resizing image: %w", op, err)
	}

	s._createCachePath(cachePath[:len(cachePath)-len(filepath.Base(data.Path))])

	if err := bimg.Write(cachePath, rByteFile); err != nil {
		return nil, "", fmt.Errorf("%s error writing file: %w", op, err)
	}

	rFile, err := os.Open(cachePath)
	if err != nil {
		return nil, "", fmt.Errorf("%s error opening file: %w", op, err)
	}

	return rFile, cachePath, nil
}

func (s *StorageService) GetCachedFile(path string) (*os.File, error) {
	const op = "StorageService.GetCachedFile() ->"

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s error opening file: %w", op, err)
	}

	return file, nil
}

func (s *StorageService) GetStorageResizeData(r *http.Request) _struct.StorageResize {
	width, _ := strconv.Atoi(chi.URLParam(r, "width"))
	height, _ := strconv.Atoi(chi.URLParam(r, "height"))

	length := len(fmt.Sprintf("%s/%s", chi.URLParam(r, "width"), chi.URLParam(r, "height")))

	return _struct.StorageResize{
		Width:  width,
		Height: height,
		Path:   r.URL.RequestURI()[13+length:],
	}
}

func (s *StorageService) _createCachePath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			panic(err)
		}
	}
}
