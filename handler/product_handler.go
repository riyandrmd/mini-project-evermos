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
	userIDRaw := c.Locals("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	var input dto.CreateProductRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.New().Struct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	product, err := service.CreateProduct(userID, input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res := dto.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Product created", res)
}

func GetAllProducts(c *fiber.Ctx) error {
	products, err := service.GetAllProducts()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var result []dto.ProductResponse
	for _, p := range products {
		result = append(result, dto.ProductResponse{
			ID:         p.ID,
			Name:       p.Name,
			Price:      p.Price,
			Stock:      p.Stock,
			StoreID:    p.StoreID,
			CategoryID: p.CategoryID,
		})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Products fetched", result)
}

func GetProductByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	product, err := service.GetProductByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
	}

	res := dto.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product fetched", res)
}

func UpdateProduct(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID := uint(userIDRaw.(float64))

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	var input dto.CreateProductRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.New().Struct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	product, err := service.UpdateProductByID(userID, uint(id), input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res := dto.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product updated", res)
}

func DeleteProduct(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID := uint(userIDRaw.(float64))

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	if err := service.DeleteProductByID(userID, uint(id)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product deleted", nil)
}

func GetMyProducts(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	products, err := service.GetProductsByUser(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var result []dto.ProductResponse
	for _, p := range products {
		result = append(result, dto.ProductResponse{
			ID:         p.ID,
			Name:       p.Name,
			Price:      p.Price,
			Stock:      p.Stock,
			StoreID:    p.StoreID,
			CategoryID: p.CategoryID,
		})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "My products fetched", result)
}
