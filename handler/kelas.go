package handler

import (
	"PilEk/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GetOneKelas(c echo.Context) error {
	idkelas, _ := strconv.Atoi(c.QueryParam("idkelas"))
	k, err := h.fs.GetByKelasId(uint(idkelas))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if k == nil {
		return c.JSON(http.StatusBadRequest, utils.NotFound())
	}

	newResp := newGetKelasResponse(k)
	return c.JSON(http.StatusOK, newResp)
}

func (h *Handler) GetListKelas(c echo.Context) error {
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	k, _, err := h.fs.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if k == nil {
		return c.JSON(http.StatusBadRequest, utils.NotFound())
	}

	newResp := newGetKelasListResponse(k)
	return c.JSON(http.StatusOK, newResp)
}
