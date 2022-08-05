package http

import (
	"net/http"
	"strconv"

	"github.com/arnaz06/deposit"
	"github.com/arnaz06/deposit/customerror"
	"github.com/arnaz06/deposit/service"
	"github.com/labstack/echo"
)

type depositHandler struct {
	service service.DepositService
}

func AddDepositHandler(e *echo.Echo, service service.DepositService) {
	if service == nil {
		panic("service is nil")
	}

	handler := &depositHandler{service}

	e.GET("/deposit/:walletID", handler.get)
	e.POST("/deposit", handler.deposit)
}

func (h *depositHandler) get(c echo.Context) error {
	ctx := c.Request().Context()

	walletIDString := c.Param("walletID")

	walletID, err := strconv.ParseInt(walletIDString, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, customerror.ValidationError(err.Error()))
	}

	deposit, err := h.service.Get(ctx, walletID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, deposit)
}

func (h *depositHandler) deposit(c echo.Context) error {
	ctx := c.Request().Context()

	var input deposit.Deposit
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, customerror.ValidationError(err.Error()))
	}

	deposit, err := h.service.Deposit(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, deposit)
}
