package handler

import (
	"toko-api/model"
	"toko-api/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	data, err := service.GetAllCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal ambil kategori", "error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": true, "data": data})
}

func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := service.GetCategoryByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Kategori tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"status": true, "data": data})
}

func CreateCategory(c *fiber.Ctx) error {
	var input model.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input invalid", "error": err.Error()})
	}

	data, err := service.CreateCategory(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal tambah kategori", "error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"status": true, "message": "Kategori dibuat", "data": data})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var input model.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input invalid", "error": err.Error()})
	}
	data, err := service.UpdateCategory(id, input)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Kategori tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"status": true, "message": "Kategori diperbarui", "data": data})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteCategory(id); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Kategori tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"status": true, "message": "Kategori dihapus"})
}
