package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"toko-api/dto"
)

func GetAllProvinsi() ([]dto.Provinsi, error) {
	url := "https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("gagal fetch data provinsi")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("gagal baca response body")
	}

	var data []dto.Provinsi
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.New("gagal parse data JSON")
	}

	return data, nil
}
