package handler

import (
	"fmt"
	"path/filepath"
	"strconv"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	namaProduk := c.FormValue("nama_produk")
	deskripsi := c.FormValue("deskripsi")
	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))
	harga, _ := strconv.Atoi(c.FormValue("harga"))
	stok, _ := strconv.Atoi(c.FormValue("stok"))

	if namaProduk == "" || categoryID == 0 || harga == 0 || stok == 0 {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Field wajib tidak boleh kosong"})
	}

	product, err := service.CreateProduct(userID, dto.CreateProductRequest{
		NamaProduk: namaProduk,
		Deskripsi:  deskripsi,
		CategoryID: uint(categoryID),
		Harga:      harga,
		Stok:       stok,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	form, err := c.MultipartForm()
	if err == nil && form != nil {
		files := form.File["photos"]
		for _, file := range files {
			// Buat nama file unik
			filename := fmt.Sprintf("produk_%d_%s", product.ID, file.Filename)
			savePath := filepath.Join("uploads/products", filename)

			if err := c.SaveFile(file, savePath); err != nil {
				continue
			}

			service.SaveProductImage(product.ID, savePath)
		}
	}

	config.DB.Preload("Images").
		Preload("Toko").
		Preload("Category").
		First(&product)

	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "Produk berhasil dibuat",
		"data":    product,
	})
}

func GetAllProducts(c *fiber.Ctx) error {
	filters := map[string]string{
		"nama_produk": c.Query("nama_produk"),
		"id_category": c.Query("id_category"),
		"id_toko":     c.Query("id_toko"),
		"harga_min":   c.Query("harga_min"),
		"harga_max":   c.Query("harga_max"),
		"page":        c.Query("page"),
		"limit":       c.Query("limit"),
	}

	products, err := service.GetAllProducts(filters)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal mengambil produk", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "data": products})
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := service.GetProductByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Produk tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"status": true, "data": product})
}
