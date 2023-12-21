package middleware

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patrickmn/go-cache"
	"go-boilerplate/external/response"
	"go-boilerplate/internal/root/model"
	"log/slog"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	log *slog.Logger
	ch  *cache.Cache
}

func NewAuthMiddleware(
	log *slog.Logger,
	cache *cache.Cache,
) *AuthMiddleware {
	return &AuthMiddleware{
		log: log,
		ch:  cache,
	}
}

func (am *AuthMiddleware) New() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const op = "http.middleware.AuthMiddleware.New"

			log := am.log.With(
				slog.String("op", op),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			token := r.Header.Get("Authorization")
			if len(token) == 0 {
				log.Warn("token wasn't provided")

				response.Respond(w, response.Response{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})

				return
			}

			token = strings.Replace(token, "Bearer ", "", 1)

			user := _getUserFromCacheByToken(am.ch, token)
			if user == nil {
				log.Warn("user wasn't found")

				response.Respond(w, response.Response{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})

				return
			}

			user.Token = token

			ctx := context.WithValue(r.Context(), "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func _getUserFromCacheByToken(ch *cache.Cache, token string) *model.User {
	if v, ok := ch.Get(token); ok {
		if user, ok := v.(model.User); ok {
			return &user
		}

		slog.Error("failed to cast user from cache", slog.Any("user", v))
	}

	return nil
}
