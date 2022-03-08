package handler

import (
	"genesis-alumni/database"
	"genesis-alumni/model/entity"
	"genesis-alumni/model/request"
	"genesis-alumni/utils"
	"log"

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

	hashedPass, err := utils.HashingPassword(alumnus.Password)
	if err != nil {
		log.Println("error while hash password", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newAlumnus.Password = (hashedPass)

	//Save data alumnus to database
	errCreateAlumnus := database.DB.Create(&newAlumnus).Error
	if errCreateAlumnus != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newAlumnus,
	})
}

func AlumnusHandlerGetAll(ctx *fiber.Ctx) error {

	userInfo := ctx.Locals("userInfo")
	log.Println("user info :", userInfo)

	var alumnusNames []entity.AlumnusName
	result := database.DB.Debug().Find(&alumnusNames)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(alumnusNames)
}
