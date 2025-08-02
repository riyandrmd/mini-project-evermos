package handler

import (
	"strconv"
	"toko-api/dto"
	"toko-api/service"
	"toko-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var req dto.CreateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	userIDFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	if err := service.CreateProduct(userID, req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Product created", nil)
}

func GetAllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	name := c.Query("name")
	categoryID, _ := strconv.Atoi(c.Query("category_id", "0"))

	products, total, err := service.GetAllProducts(page, limit, name, uint(categoryID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch products")
	}

	response := fiber.Map{
		"products": products,
		"total":    total,
		"page":     page,
		"limit":    limit,
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Products fetched", response)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	userIDFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	if err := service.UpdateProduct(userID, id, req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product updated", nil)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	userIDFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	if err := service.DeleteProduct(userID, id); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product deleted", nil)
}

func GetMyProducts(c *fiber.Ctx) error {
	userIDFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	products, err := service.GetMyProducts(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Products fetched", products)
}
