package handler

import (
	"net/http"

	"github.com/RodrigoMarcelin/Authlume/src/internal/adapters/inbound/dto"
	"github.com/RodrigoMarcelin/Authlume/src/internal/application/usecase"
	"github.com/go-chi/render"
)

type UserHandler struct {
	CreateUserUseCase *usecase.CreateUserUseCase
}

func NewUserHandler(
	createUserUseCase *usecase.CreateUserUseCase,
) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUserUseCase,
	}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateUserRequest
	render.DecodeJSON(r.Body, &request)
	err := handler.CreateUserUseCase.Execute(r.Context(), request.Email, request.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
