package handler

import (
	"fmt"
	"path/filepath"
	"toko-api/config"
	"toko-api/service"

	"github.com/gofiber/fiber/v2"
)

func GetMyToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	toko, err := service.GetTokoByUser(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Data toko", "data": toko})
}

func UpdateMyToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	toko, err := service.GetTokoByUser(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Toko tidak ditemukan"})
	}

	namaToko := c.FormValue("nama_toko")
	deskripsi := c.FormValue("deskripsi")

	file, err := c.FormFile("foto")
	if err == nil {
		filename := fmt.Sprintf("toko_%d_%s", userID, file.Filename)
		savePath := filepath.Join("uploads/toko", filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal menyimpan foto", "error": err.Error()})
		}

		toko.Foto = savePath
	}

	if namaToko != "" {
		toko.NamaToko = namaToko
	}
	if deskripsi != "" {
		toko.Deskripsi = deskripsi
	}

	if err := config.DB.Save(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal update toko", "error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Toko berhasil diperbarui",
		"data":    toko,
	})
}
