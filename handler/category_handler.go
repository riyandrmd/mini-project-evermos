package handler

import (
	"toko-api/dto"
	"toko-api/service"
	"toko-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	isAdmin := c.Locals("is_admin")
	if isAdmin != true {
		return utils.ErrorResponse(c, fiber.StatusForbidden, "Only admin can create category")
	}

	var input dto.CreateCategoryRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	category, err := service.CreateCategory(input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Category created", dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func GetAllCategories(c *fiber.Ctx) error {
	categories, err := service.GetAllCategories()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var response []dto.CategoryResponse
	for _, cat := range categories {
		response = append(response, dto.CategoryResponse{
			ID:   cat.ID,
			Name: cat.Name,
		})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "List of categories", response)
}

func UpdateCategory(c *fiber.Ctx) error {
	isAdmin := c.Locals("is_admin")
	if isAdmin != true {
		return utils.ErrorResponse(c, fiber.StatusForbidden, "Only admin can update category")
	}

	id := c.Params("id")
	var input dto.UpdateCategoryRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	category, err := service.UpdateCategory(id, input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Category updated", dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	isAdmin := c.Locals("is_admin")
	if isAdmin != true {
		return utils.ErrorResponse(c, fiber.StatusForbidden, "Only admin can delete category")
	}

	id := c.Params("id")
	if err := service.DeleteCategory(id); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Category deleted", nil)
}
