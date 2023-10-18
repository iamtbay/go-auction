package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/database"
	"github.com/iamtbay/go-auction/helpers"
	"github.com/iamtbay/go-auction/models"
)

type Auth struct{}

var authDB = database.AuthDBInit()

// @Summary Login
// @Description Login For User
// @ID auth-login
// @Accept json
// @Produce json
// @Param userInfo body models.LoginModel true "User login data"
// @Success 200 {object} map[string]interface{} "Succesful register"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /auth/login [post]
// @Tags auth
func (x *Auth) Login(c *fiber.Ctx) error {
	//get user info
	var userInfo *models.LoginModel
	err := c.BodyParser(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//send values to db
	infoDB, err := authDB.Login(userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//create jwt
	token, err := helpers.CreateJWT(infoDB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//create cookie
	err = helpers.CreateCookie(c, token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//return
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Hello %v", infoDB.Username),
	})
}

// / @Summary Register For New User
// @Description Register For New User
// @ID auth-register
// @Accept json
// @Produce json
// @Param userInfo body models.RegisterModel true "User register data"
// @Success 200 {object} map[string]interface{} "Succesful register"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /auth/register [post]
// @Tags auth
func (x *Auth) Register(c *fiber.Ctx) error {
	//get user info
	var userInfo models.RegisterModel
	err := c.BodyParser(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	}
	//send values to db
	err = authDB.Register(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//send positive json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succesfully registered",
	})

}

// @Summary Update User
// @Description Update Infos For User
// @ID update-user
// @Accept json
// @Produce json
// @Param userInfo body models.RegisterModel true "User update data"
// @Success 200 {object} map[string]interface{} "User updated succesfully"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /auth/update [patch]
// @Tags auth
func (x *Auth) Update(c *fiber.Ctx) error {
	//get userinfo
	var userInfo models.RegisterModel
	err := c.BodyParser(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//get userinfo from jwt
	tokenString := c.Cookies("accessToken")
	claims, err := helpers.ParseJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//db ops
	err = authDB.Update(claims.ID, &userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//send json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succesfulLy updated",
	})

}

// @Summary Logout User
// @Description User Logout Funciton
// @ID logout-user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Succesful register"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /auth/logout [post]
// @Tags auth
func (x *Auth) Logout(c *fiber.Ctx) error {
	//delete cookie if exist
	err := helpers.DeleteCookie(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//send positive json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logout succesfully",
	})

}
