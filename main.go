package main

import (
	"github.com/Arraf18/go-sisko/app"
	"github.com/Arraf18/go-sisko/controller"
	"github.com/Arraf18/go-sisko/helper"
	"github.com/Arraf18/go-sisko/middleware"
	"github.com/Arraf18/go-sisko/repository"
	"github.com/Arraf18/go-sisko/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	siswaRepository := repository.NewSiswaRepository()
	siswaService := service.NewSiswaService(siswaRepository, db, validate)
	siswaController := controller.NewSiswaController(siswaService)
	router := app.NewRouter(siswaController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
