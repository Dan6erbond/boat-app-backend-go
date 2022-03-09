package controllers

import (
	"errors"

	"github.com/kataras/iris/v12"
	"openwt.com/boat-app-backend/pkg/dto"
	"openwt.com/boat-app-backend/pkg/services"
)

type AuthController struct {
	AuthService services.AuthService
}

// @Summary      Login
// @Description  Logs user into the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.LoginRequest  true  "Login"
// @Success      200   {object}  dto.LoginRequest
// @Failure      400   {string}  string
// @Failure      404   {string}  string
// @Failure      500   {string}  string
// @Router       /auth/login [post]
func (c *AuthController) PostLogin(ctx iris.Context) *dto.LoginResponse {
	var loginRequest dto.LoginRequest
	err := ctx.ReadJSON(&loginRequest)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid request"))
		return nil
	}
	loginResponse, err := c.AuthService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	return loginResponse
}

// @Summary      Sign up
// @Description  Sign up a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.SignUpRequest  true  "Sign up"
// @Success      200   {object}  dto.SignUpResponse
// @Failure      400   {string}  string
// @Failure      404   {string}  string
// @Failure      500   {string}  string
// @Router       /auth/register [post]
func (c *AuthController) PostRegister(ctx iris.Context) *dto.SignUpResponse {
	var signUpRequest dto.SignUpRequest
	err := ctx.ReadJSON(&signUpRequest)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid request"))
		return nil
	}
	signUpResponse, err := c.AuthService.SignUp(signUpRequest.Username, signUpRequest.Password, signUpRequest.FirstName, signUpRequest.LastName)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	return signUpResponse
}
