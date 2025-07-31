package handler

import (
	"strconv"
	"toko-api/dto"
	"toko-api/service"
	"toko-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, 400, "Invalid request body")
	}
	category, err := service.CreateCategory(req)
	if err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}
	return utils.SuccessResponse(c, 201, "Category created", category)
}

func GetCategories(c *fiber.Ctx) error {
	categories, err := service.GetAllCategories()
	if err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}
	return utils.SuccessResponse(c, 200, "Categories retrieved", categories)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req dto.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, 400, "Invalid request body")
	}
	category, err := service.UpdateCategory(uint(id), req)
	if err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}
	return utils.SuccessResponse(c, 200, "Category updated", category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteCategory(uint(id)); err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}
	return utils.SuccessResponse(c, 200, "Category deleted", nil)
}
