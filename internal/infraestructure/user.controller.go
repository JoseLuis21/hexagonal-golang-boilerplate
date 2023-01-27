package infraestructure

import (
	"errors"
	"hexagonal-go/internal/application"
	"hexagonal-go/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func UserCreateHandler(userCreateService application.UserCreateService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(createRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}

		var validate = validator.New()

		var errorsValidate []*ErrorResponse
		err := validate.Struct(*req)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var element ErrorResponse
				element.FailedField = err.StructNamespace()
				element.Tag = err.Tag()
				element.Value = err.Param()
				errorsValidate = append(errorsValidate, &element)
			}
			return c.Status(fiber.StatusBadRequest).JSON(errorsValidate)
		}

		err = userCreateService.CreateUser(c.Context(), req.ID, req.Name, req.Email, req.Password)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidUserID):
				c.Status(fiber.StatusBadRequest).JSON(err.Error())
				return err
			default:
				c.Status(fiber.StatusInternalServerError).JSON(err.Error())
				return err
			}
		}

		return c.Status(fiber.StatusCreated).JSON("OK Registro guardado")
	}
}
