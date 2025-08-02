package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"toko-api/config"
	"toko-api/model"
	"toko-api/utils"
)

func SaveProductImage(productID uint, file *multipart.FileHeader) error {

	var product model.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		return errors.New("product not found")
	}

	timestamp := time.Now().Unix()
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("product_%d_%d%s", productID, timestamp, ext)
	path := filepath.Join("uploads/products", filename)

	if err := os.MkdirAll("uploads/products", os.ModePerm); err != nil {
		return err
	}
	if err := utils.SaveFile(file, path); err != nil {
		return err
	}

	image := model.ProductImage{
		ProductID: productID,
		ImageURL:  filename,
	}
	return config.DB.Create(&image).Error
}

func GetProductImages(productID uint) ([]model.ProductImage, error) {
	var images []model.ProductImage
	if err := config.DB.Where("product_id = ?", productID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func DeleteProductImage(userID, imageID uint) error {
	var image model.ProductImage
	if err := config.DB.Preload("Product.Store").First(&image, imageID).Error; err != nil {
		return errors.New("image not found")
	}

	if image.Product.Store == nil || image.Product.Store.UserID != userID {
		return errors.New("you are not the owner of this product image")
	}

	filePath := fmt.Sprintf("./uploads/products/%s", image.ImageURL)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return errors.New("failed to delete image file")
	}

	if err := config.DB.Delete(&image).Error; err != nil {
		return errors.New("failed to delete image from database")
	}

	return nil
}
