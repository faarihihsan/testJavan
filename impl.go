package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func GetListKeluargaImpl() (res []Keluarga, err error) {
	row := dbClient.QueryRow("select array_to_json(array_agg(to_json(a))) from (\nselect k.id, k.nama, k.parent, jsonb_agg(jsonb_build_object('id', a.id, 'nama', a.nama, 'price', a.price)) as aset\nfrom keluarga k\nleft join keluarga_aset ka on k.id = ka.id_keluarga\nleft join aset a on a.id = ka.id_aset\ngroup by k.id\norder by k.id ) a;")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while query to db: %v", err.Error()))
	}

	var str []byte
	err = row.Scan(&str)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error scanning rows: %v", err.Error()))
	}

	err = json.Unmarshal(str, &res)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error scanning rows: %v", err.Error()))
	}

	for idx, val := range res {
		total := 0
		for _, i := range val.Aset {
			total += i.Price
		}
		val.TotalAset = total
		res[idx] = val
	}

	return res, nil

}

func AddKeluargaImpl(payload KeluargaPayload) error {
	var parent interface{}
	if payload.Parent == 0 {
		parent = nil
	} else {
		parent = payload.Parent
	}
	_, err := dbClient.Exec("INSERT INTO keluarga (nama, parent) VALUES ($1, $2)",
		payload.Nama, parent)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while writing to db: %v", err.Error()))
	}

	sendNotification()

	return nil
}

func UpdateKeluargaImpl(payload KeluargaPayload) error {
	_, err := dbClient.Exec("UPDATE keluarga SET nama = $1, parent = $2 WHERE id = $3",
		payload.Nama, payload.Parent, payload.Id)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while writing to db: %v", err.Error()))
	}

	sendNotification()

	return nil
}

func DeleteKeluargaImpl(payload KeluargaPayload) error {
	_, err := dbClient.Exec("DELETE FROM keluarga where id = $1", payload.Id)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while delete row from db: %v", err.Error()))
	}

	sendNotification()

	return nil
}

func AddAsetImpl(payload AsetPayload) error {
	price, err := getPriceAset(payload.Nama)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while getting price data: %v", err.Error()))
	}

	_, err = dbClient.Exec("insert into aset (nama, price) VALUES ($1, $2)", payload.Nama, price)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while writing to db: %v", err.Error()))
	}

	sendNotification()

	return err
}

func UpdateAsetImpl(payload AsetPayload) error {
	price, err := getPriceAset(payload.Nama)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while getting price data: %v", err.Error()))
	}

	_, err = dbClient.Exec("UPDATE aset SET nama = $1, price = $2 WHERE id = $3", payload.Nama, price, payload.Id)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while writing to db: %v", err.Error()))
	}

	sendNotification()

	return nil
}

func getPriceAset(nama string) (int, error) {
	strUrl := fmt.Sprintf("https://dummyjson.com/products/search?q=%v", url.QueryEscape(nama))
	response, err := http.Get(strUrl)
	if err != nil {
		return 0, err
	}

	var result ProductResult
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	if result.Total < 1 {
		return 0, errors.New("aset tidak ditemukan")
	}
	return result.Products[0].Price, nil

}

func DeleteAsetImpl(payload AsetPayload) error {
	_, err := dbClient.Exec("DELETE FROM aset where id = $1", payload.Id)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while delete row from db: %v", err.Error()))
	}

	sendNotification()

	return nil
}

func AddAsetKeluargaImpl(payload AsetKeluargaPayload) error {
	_, err := dbClient.Exec("insert into keluarga_aset (id_keluarga, id_aset) VALUES ($1, $2)",
		payload.IdKeluarga, payload.IdAset)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while writing to db: %v", err.Error()))
	}

	sendNotification()

	return nil
}

func DeleteAsetKeluargaImpl(payload AsetKeluargaPayload) error {
	_, err := dbClient.Exec("DELETE FROM keluarga_aset where id = $1", payload.Id)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while delete row from db: %v", err.Error()))
	}

	sendNotification()

	return nil
}
