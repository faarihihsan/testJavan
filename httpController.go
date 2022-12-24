package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

func httpController(wg *sync.WaitGroup) {
	http.HandleFunc("/get-list-keluarga", GetListKeluargaHttp)
	http.HandleFunc("/add-keluarga", AddKeluargaHttp)
	http.HandleFunc("/update-keluarga", UpdateKeluargaHttp)
	http.HandleFunc("/delete-keluarga", DeleteKeluargaHttp)
	http.HandleFunc("/add-aset", AddAsetHttp)
	http.HandleFunc("/update-aset", UpdateAsetHttp)
	http.HandleFunc("/delete-aset", DeleteAsetHttp)
	http.HandleFunc("/add-aset-keluarga", AddAsetKeluargaHttp)
	http.HandleFunc("/delete-aset-keluarga", DeleteAsetKeluargaHttp)

	fmt.Println(fmt.Sprintf("Listening HTTP protocol on port %v", http_port))
	err := http.ListenAndServe(fmt.Sprintf(":%v", http_port), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wg.Done()
}

func GetListKeluargaHttp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := GetListKeluargaImpl()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while getting the data: %v", err.Error()), 400)
		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while marshalling response: %v", err.Error()), 400)
		return
	}

	_, err = w.Write(jsonRes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while writing response: %v", err.Error()), 400)
		return
	}
	return

}

func AddKeluargaHttp(w http.ResponseWriter, req *http.Request) {
	keluargaPayload, err := decodeKeluargaPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = AddKeluargaImpl(keluargaPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func UpdateKeluargaHttp(w http.ResponseWriter, req *http.Request) {
	keluargaPayload, err := decodeKeluargaPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = UpdateKeluargaImpl(keluargaPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func DeleteKeluargaHttp(w http.ResponseWriter, req *http.Request) {
	keluargaPayload, err := decodeKeluargaPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = DeleteKeluargaImpl(keluargaPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func AddAsetHttp(w http.ResponseWriter, req *http.Request) {
	asetPayload, err := decodeAsetPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = AddAsetImpl(asetPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func UpdateAsetHttp(w http.ResponseWriter, req *http.Request) {
	asetPayload, err := decodeAsetPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = UpdateAsetImpl(asetPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func DeleteAsetHttp(w http.ResponseWriter, req *http.Request) {
	asetPayload, err := decodeAsetPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = DeleteAsetImpl(asetPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func AddAsetKeluargaHttp(w http.ResponseWriter, req *http.Request) {
	asetKeluargaPayload, err := decodeAsetKeluargaPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = AddAsetKeluargaImpl(asetKeluargaPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func DeleteAsetKeluargaHttp(w http.ResponseWriter, req *http.Request) {
	asetKeluargaPayload, err := decodeAsetKeluargaPayload(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while decode request body: %v", err.Error()), 400)
		return
	}

	err = DeleteAsetKeluargaImpl(asetKeluargaPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while save the data: %v", err.Error()), 500)
		return
	}
}

func decodeKeluargaPayload(req *http.Request) (KeluargaPayload, error) {
	if req.Body == nil {
		return KeluargaPayload{}, errors.New("please send a request body")
	}

	var keluarga KeluargaPayload
	err := json.NewDecoder(req.Body).Decode(&keluarga)
	if err != nil {
		return KeluargaPayload{}, err
	}

	return keluarga, nil
}

func decodeAsetPayload(req *http.Request) (AsetPayload, error) {
	if req.Body == nil {
		return AsetPayload{}, errors.New("please send a request body")
	}

	var aset AsetPayload
	err := json.NewDecoder(req.Body).Decode(&aset)
	if err != nil {
		return AsetPayload{}, err
	}

	return aset, nil
}

func decodeAsetKeluargaPayload(req *http.Request) (AsetKeluargaPayload, error) {
	if req.Body == nil {
		return AsetKeluargaPayload{}, errors.New("please send a request body")
	}

	var asetKeluarga AsetKeluargaPayload
	err := json.NewDecoder(req.Body).Decode(&asetKeluarga)
	if err != nil {
		return AsetKeluargaPayload{}, err
	}

	return asetKeluarga, nil
}
