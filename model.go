package main

type Keluarga struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Parent    int    `json:"parent"`
	TotalAset int    `json:"totalAset"`
	Aset      []Aset `json:"aset"`
}

type Aset struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Price int    `json:"price"`
}

type KeluargaPayload struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Parent int    `json:"parent"`
}

type AsetPayload struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type AsetKeluargaPayload struct {
	Id         int `json:"id"`
	IdKeluarga int `json:"idKeluarga"`
	IdAset     int `json:"idAset"`
}

type ProductResult struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
}

type Product struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}
