package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Arraf18/go-sisko/app"
	"github.com/Arraf18/go-sisko/controller"
	"github.com/Arraf18/go-sisko/helper"
	"github.com/Arraf18/go-sisko/middleware"
	"github.com/Arraf18/go-sisko/model/domain"
	"github.com/Arraf18/go-sisko/repository"
	"github.com/Arraf18/go-sisko/service"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_sisko")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	siswaRepository := repository.NewSiswaRepository()
	siswaService := service.NewSiswaService(siswaRepository, db, validate)
	siswaController := controller.NewSiswaController(siswaService)
	router := app.NewRouter(siswaController)

	return middleware.NewAuthMiddleware(router)
}

func truncateSiswa(db *sql.DB) {
	db.Exec("TRUNCATE siswa")
}

func TestCreateSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/siswas", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/siswas", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)

	tx, _ := db.Begin()
	siswaRepository := repository.NewSiswaRepository()
	siswa := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/siswas/"+strconv.Itoa(siswa.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, siswa.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)

	tx, _ := db.Begin()
	siswaRepository := repository.NewSiswaRepository()
	siswa := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/siswas/"+strconv.Itoa(siswa.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestGetSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)

	tx, _ := db.Begin()
	siswaRepository := repository.NewSiswaRepository()
	siswa := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/siswas/"+strconv.Itoa(siswa.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, siswa.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, siswa.Nama, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, siswa.Alamat, responseBody["data"].(map[string]interface{})["alamat"])
	assert.Equal(t, siswa.TanggalLahir, responseBody["data"].(map[string]interface{})["tanggal_lahir"])
	assert.Equal(t, siswa.TempatLahir, responseBody["data"].(map[string]interface{})["tempat_lahir"])
	assert.Equal(t, siswa.JenisKelamin, responseBody["data"].(map[string]interface{})["jenis_kelamin"])
	assert.Equal(t, siswa.Agama, responseBody["data"].(map[string]interface{})["agama"])
	assert.Equal(t, siswa.GolonganDarah, responseBody["data"].(map[string]interface{})["golongan_darah"])
	assert.Equal(t, siswa.NoTelepon, responseBody["data"].(map[string]interface{})["no_telepon"])
}

func TestGetSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)

	tx, _ := db.Begin()
	siswaRepository := repository.NewSiswaRepository()
	siswa := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/siswas/"+strconv.Itoa(siswa.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteSiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)

	tx, _ := db.Begin()
	siswaRepository := repository.NewSiswaRepository()
	siswa := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/siswas/"+strconv.Itoa(siswa.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteSiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/siswas/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListSiswasSuccess(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)

	tx, _ := db.Begin()
	siswaRepository := repository.NewSiswaRepository()
	siswa1 := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Gadget",
	})
	siswa2 := siswaRepository.Save(context.Background(), tx, domain.Siswa{
		Nama: "Computer",
	})

	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/siswas", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var siswas = responseBody["data"].([]interface{})

	siswaResponse1 := siswas[0].(map[string]interface{})
	siswaResponse2 := siswas[1].(map[string]interface{})

	assert.Equal(t, siswa1.Id, int(siswaResponse1["id"].(float64)))
	assert.Equal(t, siswa1.Nama, siswaResponse1["name"])
	assert.Equal(t, siswa1.Alamat, siswaResponse1["alamat"])
	assert.Equal(t, siswa1.TanggalLahir, siswaResponse1["tanggal_lahir"])
	assert.Equal(t, siswa1.TempatLahir, siswaResponse1["tempat_lahir"])
	assert.Equal(t, siswa1.JenisKelamin, siswaResponse1["jenis_kelamin"])
	assert.Equal(t, siswa1.Agama, siswaResponse1["agama"])
	assert.Equal(t, siswa1.GolonganDarah, siswaResponse1["golongan_darah"])
	assert.Equal(t, siswa1.NoTelepon, siswaResponse1["no_telepon"])

	assert.Equal(t, siswa2.Id, int(siswaResponse2["id"].(float64)))
	assert.Equal(t, siswa2.Nama, siswaResponse2["name"])
	assert.Equal(t, siswa1.Alamat, siswaResponse2["alamat"])
	assert.Equal(t, siswa1.TanggalLahir, siswaResponse2["tanggal_lahir"])
	assert.Equal(t, siswa1.TempatLahir, siswaResponse2["tempat_lahir"])
	assert.Equal(t, siswa1.JenisKelamin, siswaResponse2["jenis_kelamin"])
	assert.Equal(t, siswa1.Agama, siswaResponse2["agama"])
	assert.Equal(t, siswa1.GolonganDarah, siswaResponse2["golongan_darah"])
	assert.Equal(t, siswa1.NoTelepon, siswaResponse2["no_telepon"])

}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateSiswa(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/siswas", nil)
	request.Header.Add("X-API_Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

}
