package handler

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaction(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.CreateTrxRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid", "error": err.Error()})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Validasi gagal", "error": err.Error()})
	}

	trx, err := service.CreateTransaction(userID, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	config.DB.Preload("Alamat").
		Preload("User").
		First(&trx)

	return c.Status(201).JSON(fiber.Map{"status": true, "message": "Transaksi berhasil", "data": trx})
}

func GetMyTransactions(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	data, err := service.GetUserTransactions(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal ambil transaksi", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "data": data})
}

func GetTransactionByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	trxID := c.Params("id")

	data, err := service.GetTransactionByID(userID, trxID)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "data": data})
}
