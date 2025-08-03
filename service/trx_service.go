package service

import (
	"errors"
	"fmt"
	"time"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateTransaction(userID uint, input dto.CreateTrxRequest) (*model.Trx, error) {
	if len(input.Items) == 0 {
		return nil, errors.New("transaksi tidak boleh kosong")
	}

	var totalHarga int
	var trxDetails []model.TrxDetail

	for _, item := range input.Items {
		var product model.Product
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			return nil, fmt.Errorf("produk %d tidak ditemukan", item.ProductID)
		}

		if product.Stok < item.Qty {
			return nil, fmt.Errorf("stok produk %s tidak cukup", product.NamaProduk)
		}

		// Buat log produk
		log := model.LogProduk{
			NamaProduk: product.NamaProduk,
			Slug:       product.Slug,
			Deskripsi:  product.Deskripsi,
			Harga:      product.Harga,
			Stok:       product.Stok,
			TokoID:     product.TokoID,
			CategoryID: product.CategoryID,
		}
		if err := config.DB.Create(&log).Error; err != nil {
			return nil, errors.New("gagal menyimpan log produk")
		}

		hargaTotal := product.Harga * item.Qty
		totalHarga += hargaTotal

		trxDetails = append(trxDetails, model.TrxDetail{
			LogProdukID: log.ID,
			TokoID:      product.TokoID,
			Qty:         item.Qty,
			HargaTotal:  hargaTotal,
		})

		// Kurangi stok produk
		product.Stok -= item.Qty
		config.DB.Save(&product)
	}

	var alamat model.Alamat
	if err := config.DB.First(&alamat, "id = ? AND user_id = ?", input.IDAlamat, userID).Error; err != nil {
		return nil, errors.New("alamat tidak ditemukan")
	}

	// Buat transaksi utama
	trx := model.Trx{
		UserID:           userID,
		KodeInvoice:      fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix()),
		AlamatID:         input.IDAlamat,
		MetodePembayaran: input.MetodePembayaran,
		HargaTotal:       totalHarga,
	}

	if err := config.DB.Create(&trx).Error; err != nil {
		return nil, errors.New("gagal membuat transaksi")
	}

	// Tambahkan ke detail_trx
	for i := range trxDetails {
		trxDetails[i].TrxID = trx.ID
	}
	if err := config.DB.Create(&trxDetails).Error; err != nil {
		return nil, errors.New("gagal menyimpan detail transaksi")
	}

	return &trx, nil
}

func GetUserTransactions(userID uint) ([]model.Trx, error) {
	var transaksi []model.Trx

	err := config.DB.
		Preload("Detail.LogProduk").
		Preload("Alamat").
		Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&transaksi).Error

	if err != nil {
		return nil, err
	}

	return transaksi, nil
}

func GetTransactionByID(userID uint, trxID string) (*model.Trx, error) {
	var trx model.Trx

	err := config.DB.
		Preload("Detail.LogProduk").
		Preload("Alamat").
		Preload("User").
		First(&trx, trxID).Error

	if err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}

	if trx.UserID != userID {
		return nil, errors.New("kamu tidak berhak mengakses transaksi ini")
	}

	return &trx, nil
}
