package controllers

import (
	"strconv"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/kataras/iris/v12"
	"openwt.com/boat-app-backend/pkg/dto"
	"openwt.com/boat-app-backend/pkg/services"
)

type UsersController struct {
	UsersService services.UsersService
}

// @Summary      Get Me
// @Description  Get current User
// @Tags         auth,users
// @Accept       json
// @Produce      json
// @Security     JWT
// @Success      200  {object}  dto.UserDTO
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /users/me [get]
func (c *UsersController) GetMe(ctx iris.Context) *dto.UserDTO {
	ctxUser := ctx.User()
	id, err := ctxUser.GetID()
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	user, err := c.UsersService.GetUserByID(uint(idUint))
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	userDTO := dto.UserDTO{}
	err = dtoMapper.Map(&userDTO, user)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	return &userDTO
}
