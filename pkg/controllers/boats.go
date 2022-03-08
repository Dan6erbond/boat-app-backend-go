package controllers

import (
	"errors"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/kataras/iris/v12"
	"openwt.com/boat-app-backend/pkg/dto"
	"openwt.com/boat-app-backend/pkg/repositories"
	"openwt.com/boat-app-backend/pkg/services"
)

type BoatsController struct {
	BoatsService services.BoatsService
}

// @Summary      Create new boat
// @Description  Create new boat
// @Tags         boats
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateUpdateBoatDTO  true  "Create new boat"
// @Success      200   {object}  dto.BoatDTO
// @Failure      400   {string}  string
// @Failure      404   {string}  string
// @Failure      500   {string}  string
// @Router       /boats [post]
func (c *BoatsController) Post(ctx iris.Context) *dto.BoatDTO {
	var createDTO dto.CreateUpdateBoatDTO
	err := ctx.ReadJSON(&createDTO)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid request"))
		return nil
	}
	boat, err := c.BoatsService.CreateBoat(&createDTO)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	var boatDTO dto.BoatDTO
	err = dtoMapper.Map(&boatDTO, boat)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}
	return &boatDTO
}

// @Summary      Get all boats
// @Description  Get all boats
// @Tags         boats
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.BoatDTO
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /boats [get]
func (c *BoatsController) Get(ctx iris.Context) []dto.BoatDTO {
	var boatDTOs []dto.BoatDTO
	err := dtoMapper.Map(&boatDTOs, c.BoatsService.GetBoats())

	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}

	return boatDTOs
}

// @Summary      Get boat by id
// @Description  Get boat by id
// @Tags         boats
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Boat id"
// @Success      200  {object}  dto.BoatDTO
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /boats/{id} [get]
func (c *BoatsController) GetBy(id uint, ctx iris.Context) *dto.BoatDTO {
	boat, err := c.BoatsService.GetBoatByID(id)

	if errors.Is(err, repositories.ErrBoatNotFound) {
		ctx.StopWithError(iris.StatusNotFound, err)
		return nil
	}

	boatDTO := dto.BoatDTO{}
	err = dtoMapper.Map(&boatDTO, boat)

	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}

	return &boatDTO
}

// @Summary      Update boat
// @Description  Update boat
// @Tags         boats
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateUpdateBoatDTO  true  "Update boat"
// @Success      200   {object}  dto.BoatDTO
// @Failure      400   {string}  string
// @Failure      404   {string}  string
// @Failure      500   {string}  string
// @Router       /boats [put]
func (c *BoatsController) PutBy(id uint, ctx iris.Context) *dto.BoatDTO {
	var updateDTO dto.CreateUpdateBoatDTO
	err := ctx.ReadJSON(&updateDTO)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid request"))
		return nil
	}

	boat, err := c.BoatsService.UpdateBoat(id, &updateDTO)
	if errors.Is(err, repositories.ErrBoatNotFound) {
		ctx.StopWithError(iris.StatusNotFound, err)
		return nil
	}

	boatDTO := dto.BoatDTO{}
	err = dtoMapper.Map(&boatDTO, boat)

	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return nil
	}

	return &boatDTO
}

// @Summary      Delete boat
// @Description  Delete boat
// @Tags         boats
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Boat id"
// @Success      200  {object}  dto.BoatDTO
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /boats/{id} [delete]
func (c *BoatsController) DeleteBy(id uint, ctx iris.Context) *dto.BoatDTO {
	err := c.BoatsService.DeleteBoatByID(id)
	if errors.Is(err, repositories.ErrBoatNotFound) {
		ctx.StopWithError(iris.StatusNotFound, err)
		return nil
	}

	return nil
}
