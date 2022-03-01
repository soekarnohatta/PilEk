package main

import (
	"PilEk/db"
	"PilEk/handler"
	"PilEk/router"
	"PilEk/store"
	"github.com/labstack/echo/v4"
)

func main() {
	r := router.New()

	r.Static("/", "dist")
	r.HTTPErrorHandler = customHTTPErrorHandler
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewKandidatStore(d)
	fs := store.NewKelasStore(d)

	h := handler.NewHandler(*us, *as, *fs)
	h.Register(v1)
	_ = r.Start(":8080")
}

func customHTTPErrorHandler(_ error, c echo.Context) {
	errorPage := "dist/index.html"
	_ = c.File(errorPage)
}
