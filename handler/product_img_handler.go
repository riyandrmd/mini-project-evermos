package handler

import (
	"strconv"
	"toko-api/config"
	"toko-api/model"
	"toko-api/service"
	"toko-api/utils"

	"github.com/gofiber/fiber/v2"
)

func UploadProductImage(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID := uint(userIDRaw.(float64)) // jwt.MapClaims menyimpan angka sebagai float64

	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	var user model.User
	if err := config.DB.Preload("Store").First(&user, userID).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "User not found")
	}

	var product model.Product
	if err := config.DB.Preload("Store").First(&product, productID).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
	}
	if product.Store == nil || product.Store.UserID != userID {
		return utils.ErrorResponse(c, fiber.StatusForbidden, "Unauthorized to upload image for this product")
	}

	file, err := c.FormFile("image")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Image file is required")
	}

	err = service.SaveProductImage(uint(productID), file)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Image uploaded successfully", nil)
}

func GetProductImages(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	images, err := service.GetProductImages(uint(productID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch product images")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product images fetched", images)
}

func DeleteProductImage(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	userID := uint(userIDRaw.(float64))

	imageID, err := strconv.Atoi(c.Params("imageId"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid image ID")
	}

	err = service.DeleteProductImage(userID, uint(imageID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product image deleted", nil)
}
