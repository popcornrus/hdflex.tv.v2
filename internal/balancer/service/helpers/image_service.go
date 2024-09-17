package helpers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"go-hdflex/internal/balancer/_struct"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ImageService struct {
}

type ImageServiceInterface interface {
	Download(_struct.Image) (string, error)
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (i *ImageService) Download(im _struct.Image) (string, error) {
	const op = "ImageService.Download() ->"

	request, err := http.NewRequest("GET", im.Url.String(), nil)
	if err != nil {
		return "", fmt.Errorf("%s %w", op, err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	request.Header.Set("Referer", "https://hdflex.tv")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("%s %w", op, err)
	}

	if strings.Contains(response.Request.URL.String(), "no-poster.gif") {
		return "", fmt.Errorf("%s %w", op, errors.New("no poster"))
	}

	if im.Filename == "" {
		h := sha256.New()
		h.Write([]byte(filepath.Base(im.Url.String())))
		im.Filename = fmt.Sprintf("%x.%s", h.Sum(nil), "webp")
	}

	defer response.Body.Close()

	path, err := i.saveInStorage(im.Filename, response.Body)
	if err != nil {
		return "", fmt.Errorf("%s %w", op, err)
	}

	return path, nil
}

func (i *ImageService) saveInStorage(title string, body io.ReadCloser) (string, error) {
	const op = "ImageService.saveInStorage() ->"

	defer body.Close()

	path := fmt.Sprintf("%s/%s/%s", os.Getenv("STORAGE_PATH"), title[0:2], title[2:4])
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return "", fmt.Errorf("%s %w", op, err)
		}
	}

	filepath := fmt.Sprintf("%s/%s", path, title)
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("%s %w", op, err)
	}

	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return "", fmt.Errorf("%s %w", op, err)
	}

	return fmt.Sprintf("%s/%s/%s", title[0:2], title[2:4], title), nil
}
