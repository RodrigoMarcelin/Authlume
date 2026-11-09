package rest

import (
	"github.com/RodrigoMarcelin/Authlume/src/internal/adapters/inbound/handler"
	"github.com/RodrigoMarcelin/Authlume/src/internal/application/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type userController struct {
	userHandler *handler.UserHandler
}

func NewUserController(
	createUserUseCase *usecase.CreateUserUseCase,
) *userController {
	return &userController{
		userHandler: handler.NewUserHandler(createUserUseCase),
	}
}

func (uc *userController) RoutesUser(r chi.Router) {
	r.Use(middleware.Logger)
	r.Post("/users", uc.userHandler.CreateUser)

}
