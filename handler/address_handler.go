package handler

import (
	"toko-api/model"
	"toko-api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input model.Alamat

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid", "error": err.Error()})
	}

	alamat, err := service.CreateAlamat(userID, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal buat alamat", "error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": true, "message": "Alamat dibuat", "data": alamat})
}

func GetAlamatSaya(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	alamat, err := service.GetAlamatByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal ambil alamat", "error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": true, "message": "Data alamat", "data": alamat})
}

func GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	alamat, err := service.GetAlamatByID(userID, id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Data alamat ditemukan", "data": alamat})
}

func UpdateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var input model.Alamat
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid", "error": err.Error()})
	}

	alamat, err := service.UpdateAlamat(userID, id, input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Alamat diperbarui", "data": alamat})
}

func DeleteAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	if err := service.DeleteAlamat(userID, id); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Alamat dihapus"})
}
