package storages

import (
	"time"
)

// Storages represent storage informastion.
type Storages struct {
	UUID         *string    `json:"uuid"`
	Comodity     *string    `json:"komoditas"`
	ProvinceArea *string    `json:"area_provinsi"`
	CityArea     *string    `json:"area_kota"`
	Size         *string    `json:"size"`
	Price        *string    `json:"price"`
	ParsedDate   *time.Time `json:"tgl_parsed"`
	TimeStamp    *string    `json:"timestamp"`
}

// AreaOption represent the available storage area options.
type AreaOption struct {
	Province *string `json:"province"`
	City     *string `json:"city"`
}

// SizeOption represent the available size option.
type SizeOption struct {
	Size *string `json:"size"`
}
