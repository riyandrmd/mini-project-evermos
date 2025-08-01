package service

import (
	"errors"
	"fmt"
	"time"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"

	"gorm.io/gorm"
)

func CreateTrx(userID uint, req dto.CreateTrxRequest) error {
	db := config.DB

	var address model.Address
	if err := db.Where("id = ? AND user_id = ?", req.AddressID, userID).First(&address).Error; err != nil {
		return errors.New("invalid address")
	}

	//Buat kode invoice pake time
	invoiceCode := fmt.Sprintf("INV-%d", time.Now().UnixNano())

	trx := model.Trx{
		InvoiceCode: invoiceCode,
		Method:      req.Method,
		UserID:      userID,
		AddressID:   req.AddressID,
	}

	totalPrice := 0.0
	var detailTrxList []model.DetailTrx

	for _, item := range req.Items {
		var product model.Product
		if err := db.Where("id = ?", item.ProductID).Preload("Store").First(&product).Error; err != nil {
			return fmt.Errorf("product ID %d not found", item.ProductID)
		}

		logProduct := model.LogProduct{
			ProductID:     product.ID,
			Name:          product.Name,
			Slug:          product.Slug,
			Description:   product.Description,
			PriceReseller: product.PriceReseller,
			PriceCustomer: product.PriceCustomer,
			CategoryID:    product.CategoryID,
			StoreID:       product.StoreID,
		}
		if err := db.Create(&logProduct).Error; err != nil {
			return err
		}

		subTotal := product.PriceCustomer * float64(item.Qty)
		totalPrice += subTotal

		detailTrxList = append(detailTrxList, model.DetailTrx{
			LogProductID: logProduct.ID,
			StoreID:      product.StoreID,
			Qty:          item.Qty,
			TotalPrice:   subTotal,
		})
	}

	trx.TotalPrice = totalPrice
	if err := db.Create(&trx).Error; err != nil {
		return err
	}

	for i := range detailTrxList {
		detailTrxList[i].TrxID = trx.ID
	}
	if err := db.Create(&detailTrxList).Error; err != nil {
		return err
	}

	return nil
}

func GetMyTransactions(userID uint) ([]model.Trx, error) {
	db := config.DB

	var trxList []model.Trx
	err := db.Preload("DetailTrx").
		Preload("DetailTrx.LogProduct").
		Preload("Address").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&trxList).Error

	if err != nil {
		return nil, err
	}

	return trxList, nil
}

func GetTrxByID(userID uint, trxID uint) (*model.Trx, error) {
	db := config.DB

	var trx model.Trx
	err := db.Preload("DetailTrx").
		Preload("DetailTrx.LogProduct").
		Preload("Address").
		Where("id = ? AND user_id = ?", trxID, userID).
		First(&trx).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}

	return &trx, nil
}
