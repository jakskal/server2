package storages

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Service implement method to work with storage.
type Service struct {
}

// NewService creates a new user service.
func NewService() Service {
	return Service{}
}

// StorageList fetch data storages from url
func (s *Service) StorageList() ([]Storages, error) {
	var dataStorages []Storages
	url := "​https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
	url = ensureProperURL(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bData, &dataStorages); err != nil {
		return nil, err
	}
	return dataStorages, nil
}

// ListAreaOption fetch data storages from url
func (s *Service) ListAreaOption() ([]AreaOption, error) {
	var dataAreaOption []AreaOption
	url := "​https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/option_area"
	url = ensureProperURL(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bData, &dataAreaOption); err != nil {
		return nil, err
	}
	return dataAreaOption, nil
}

// ListSizeOption fetch data storages from url
func (s *Service) ListSizeOption() ([]SizeOption, error) {
	var dataSizeOption []SizeOption
	url := "​https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/option_size"
	url = ensureProperURL(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bData, &dataSizeOption); err != nil {
		return nil, err
	}
	return dataSizeOption, nil
}

func ensureProperURL(url string) string {
	return strings.TrimPrefix(url, "\u200b")
}
