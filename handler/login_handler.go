package handler

import (
	"genesis-alumni/database"
	"genesis-alumni/model/entity"
	"genesis-alumni/model/request"
	"genesis-alumni/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	log.Println(loginRequest)

	//Validate Request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//Check availablity alumnus name

	var alumnusName entity.AlumnusName

	err := database.DB.First(&alumnusName, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//Generate JWT
	claims := jwt.MapClaims{}
	claims["name"] = alumnusName.Name
	claims["email"] = alumnusName.Email
	claims["address"] = alumnusName.Address
	claims["phone"] = alumnusName.Phone
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	if alumnusName.Email == "indra@genesis.id" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGeneratedToken := utils.GenerateToken(&claims)
	if errGeneratedToken != nil {
		log.Println(errGeneratedToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})

	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}
