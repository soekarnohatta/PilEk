package handler

import (
	"PilEk/router/middleware"
	"PilEk/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/login")
	guestUsers.POST("", h.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.PUT("/vote", h.Vote)
	user.POST("/update", h.Update)
	user.GET("/current", h.CurrentUser)

	kandidat := v1.Group("/kandidat")
	kandidat.GET("/get", h.GetOneKandidat)
	kandidat.GET("/getall", h.GetListKandidat)

	kelas := v1.Group("/kelas")
	kelas.GET("/get", h.GetOneKelas)
	kelas.GET("/getall", h.GetListKelas)
}
