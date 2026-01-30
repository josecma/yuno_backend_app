package handlers

import (
	usecases "yuno/src/modules/users/application/useCases"

	"github.com/gofiber/fiber/v2"
)

type CreateUserHandler struct {
	createUserUseCase *usecases.CreateUserUseCase
}

func NewCreateUserHandler(createUserUseCase *usecases.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *CreateUserHandler) Handle(c *fiber.Ctx) error {}
