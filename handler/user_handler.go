package handler

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"

	"github.com/gofiber/fiber/v2"
)

func GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user model.User
	if err := config.DB.Preload("Toko").First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User tidak ditemukan",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil ambil profil",
		"data":    user,
	})
}

func UpdateMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal parsing request",
			"error":   err.Error(),
		})
	}

	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User tidak ditemukan",
			"error":   err.Error(),
		})
	}

	user.Nama = req.Nama
	user.Notelp = req.Notelp
	user.Email = req.Email
	user.TanggalLahir = req.TanggalLahir
	user.JenisKelamin = req.JenisKelamin
	user.Tentang = req.Tentang
	user.Pekerjaan = req.Pekerjaan
	user.IDProvinsi = req.IDProvinsi
	user.IDKota = req.IDKota

	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil update profil",
		"data":    user,
	})
}
