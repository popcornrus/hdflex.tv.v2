package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-hdflex/internal/api/http/handler"
	md "go-hdflex/internal/api/http/middleware"
	"log/slog"
	"net/http"
	"time"
)

func NewRouter(
	log *slog.Logger,
	handlers *handler.Handlers,
	md *md.Middleware,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	//r.Use(logger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors())

	r.Route("/storage", func(r chi.Router) {
		r.Get("/r/{width}/{height}/*", handlers.Storage.Resize)
		r.Get("/fr/{width}/{height}/*", handlers.Storage.ForceResize)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/content", func(r chi.Router) {
			r.Get("/", handlers.Content.Get)
			r.Get("/{slug}", handlers.Content.ShowBySlug)
			r.Get("/count", handlers.Content.Count)
		})

		r.Route("/genres", func(r chi.Router) {
			r.Get("/", handlers.Genre.Get)
		})

		r.Route("/countries", func(r chi.Router) {
			r.Get("/", handlers.Country.Get)
		})
	})

	return r
}

func cors() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.WriteHeader(http.StatusOK)
				return
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}
}
