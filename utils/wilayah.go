package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FetchProvinces() ([]Province, error) {
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var provinces []Province
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	if len(provinces) == 0 {
		return nil, errors.New("no province data found")
	}

	return provinces, nil
}
