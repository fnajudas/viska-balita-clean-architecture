package models

import "time"

// type Balita struct {
// 	ID                uint64 `gorm:"primaryKey;autoIncrement:true"`
// 	NIK               int
// 	Nama              string
// 	Alamat            string
// 	Birthday          *time.Time
// 	Image             string
// 	BeratBadan        string
// 	TinggiBadan       string
// 	LingkarLenganAtas string
// }

type RespGetDataBalita struct {
	Id                int        `json:"id"`
	CreatedOn         *time.Time `json:"created_on"`
	Nik               int        `json:"nik"`
	Nama              string     `json:"nama"`
	Alamat            string     `json:"alamat"`
	Birthday          time.Time  `json:"birthday"`
	Image             string     `json:"image"`
	BeratBadan        string     `json:"berat_badan"`
	TinggiBadan       string     `json:"tinggi_badan"`
	LingkarLenganAtas string     `json:"lingkar_lengan_atas"`
}

type ReqGetDataBalita struct {
	NIK  int    `json:"nik"`
	Nama string `json:"nama"`
}

type UpdateBalita struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type Balita struct {
	Id                int        `json:"id"`
	CreatedOn         *time.Time `json:"created_on"`
	Nik               int        `json:"nik"`
	Nama              string     `json:"nama"`
	Alamat            string     `json:"alamat"`
	Birthday          time.Time  `json:"birthday"`
	Image             string     `json:"image"`
	BeratBadan        string     `json:"berat_badan"`
	TinggiBadan       string     `json:"tinggi_badan"`
	LingkarLenganAtas string     `json:"lingkar_lengan_atas"`
}

type CreateDataBalita struct {
	Nik               int       `json:"nik"`
	Nama              string    `json:"nama"`
	Alamat            string    `json:"alamat"`
	Birthday          time.Time `json:"birthday"`
	Image             string    `json:"image"`
	BeratBadan        string    `json:"berat_badan"`
	TinggiBadan       string    `json:"tinggi_badan"`
	LingkarLenganAtas string    `json:"lingkar_lengan_atas"`
}

type DeteleBalita struct {
	Id int `json:"id"`
}
