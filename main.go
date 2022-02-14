package main

import (
	"PilEk/db"
	"PilEk/handler"
	"PilEk/router"
	"PilEk/store"

	_ "PilEk/docs"                               // docs is generated by Swag CLI, you have to import it.
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Swagger Example API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @host 127.0.0.1:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.New()

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewKandidatStore(d)
	fs := store.NewKelasStore(d)

	h := handler.NewHandler(*us, *as, *fs)
	h.Register(v1)
	r.Logger.Error(r.Start("192.168.1.5:8080"))
}