package handler

import (
	"context"
	"encoding/json"
	"esb/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Usecase interface {
	GetCustomers(ctx context.Context) (*domain.MultipleResponse[domain.CustomerEntity], error)
}

type ResponseError struct {
	Message string `json:"message"`
}

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase}
}

func (h *Handler) GetCustomers(e echo.Context) error {
	ctx := e.Request().Context()
	resp, err := h.usecase.GetCustomers(ctx)
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			ResponseError{Message: domain.ErrInternalServerError.Error()},
		)
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	e.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(e.Response()).Encode(resp)
}
