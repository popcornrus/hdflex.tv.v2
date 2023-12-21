package root

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-boilerplate/internal/root/http/handler"
	md "go-boilerplate/internal/root/http/middleware"
	"go-boilerplate/internal/root/http/middleware/logger"
	"log/slog"
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
	r.Use(logger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(ri chi.Router) {
		ri.Route("/users", func(ru chi.Router) {
			ru.Post("/sign-up", handlers.User.SignUp)
			ru.Post("/sign-in", handlers.User.SignIn)

			ru.Route("/me", func(rup chi.Router) {
				rup.Use(md.AuthMiddleware.New())

				rup.Get("/", handlers.User.Get)
				rup.Put("/update", handlers.User.Update)
			})
		})
	})

	return r
}
