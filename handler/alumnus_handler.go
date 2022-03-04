package handler

import (
	"genesis-alumni/model/entity"
	"genesis-alumni/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AlumnusHandlerCreate(ctx *fiber.Ctx) error {
	alumnus := new(request.AlumnusNameRequest)
	if errRequest := ctx.BodyParser(alumnus); errRequest != nil {
		return errRequest
	}

	//Validation Request
	validate := validator.New()
	errValidate := validate.Struct(alumnus)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newAlumnus := entity.AlumnusName{
		Name:    alumnus.Name,
		Email:   alumnus.Email,
		Address: alumnus.Address,
		Phone:   alumnus.Phone,
	}
}
