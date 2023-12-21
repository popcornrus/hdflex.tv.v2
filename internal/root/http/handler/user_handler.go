package handler

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"go-boilerplate/external/response"
	request "go-boilerplate/internal/root/http/request/users"
	"go-boilerplate/internal/root/model"
	"go-boilerplate/internal/root/service/user_service"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	log       *slog.Logger
	validator *validator.Validate

	um user_service.UserServiceInterface
}

func NewUserHandler(
	log *slog.Logger,
	um user_service.UserServiceInterface,
) *UserHandler {
	return &UserHandler{
		log:       log,
		validator: validator.New(),
		um:        um,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    r.Context().Value("user").(*model.User),
	})

	return
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	const op = "UserHandler.Update"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req request.UpdateRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Error("failed to decode request", err)

		response.Respond(w, response.Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Data:    nil,
		})

		return
	}

	if err := h.validator.Struct(req); err != nil {
		log.Error("failed to validate request", err)

		response.Respond(w, response.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "unprocessable entity",
			Data:    nil,
		})

		return
	}

	user := r.Context().Value("user").(*model.User)
	err := h.um.Update(r.Context(), user, req)
	if err != nil {
		log.Error("failed to update user", err)

		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		})

		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    nil,
	})

	return
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	const op = "UserHandler.SignUp"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req request.SignUpRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Error("failed to decode request", err)

		response.Respond(w, response.Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Data:    nil,
		})

		return
	}

	if err := h.validator.Struct(req); err != nil {
		log.Error("failed to validate request", err)

		response.Respond(w, response.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "unprocessable entity",
			Data:    nil,
		})

		return
	}

	if _, err := h.um.FindUserByEmail(r.Context(), req.Email); err == nil {
		log.Error("email already exists")

		response.Respond(w, response.Response{
			Status:  http.StatusConflict,
			Message: "email already exists",
			Data:    nil,
		})

		return
	}

	token, err := h.um.SignUp(r.Context(), req)
	if err != nil {
		log.Error("failed to sign up", err)

		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data: struct {
				Token string `json:"token"`
			}{
				Token: *token,
			},
		})
		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: struct {
			Token *string `json:"token"`
		}{
			Token: token,
		},
	})

	return
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	const op = "UserHandler.SignIn"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req request.SignInRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Error("failed to decode request", err)

		response.Respond(w, response.Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
		})

		return
	}

	if err := h.validator.Struct(req); err != nil {
		log.Error("failed to validate request", err)

		response.Respond(w, response.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "unprocessable entity",
		})

		return
	}

	token, err := h.um.SignIn(r.Context(), req)
	if err != nil {
		log.Error("failed to sign in", err)

		response.Respond(w, response.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})

		return
	}

	response.Respond(w, response.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: struct {
			Token *string `json:"token"`
		}{
			Token: token,
		},
	})

	return
}
