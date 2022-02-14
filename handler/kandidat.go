package handler

import (
	"PilEk/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GetOneKandidat(c echo.Context) error {
	idkandidat, _ := strconv.Atoi(c.QueryParam("idkandidat"))
	k, err := h.ks.GetByIdKandidat(idkandidat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if k == nil {
		return c.JSON(http.StatusBadRequest, utils.NotFound())
	}

	newResp := newGetKandidatResponse(k)
	return c.JSON(http.StatusOK, newResp)
}

func (h *Handler) GetListKandidat(c echo.Context) error {
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	k, _, err := h.ks.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if k == nil {
		return c.JSON(http.StatusBadRequest, utils.NotFound())
	}

	newResp := newGetKandidatListResponse(k)
	return c.JSON(http.StatusOK, newResp)
}
