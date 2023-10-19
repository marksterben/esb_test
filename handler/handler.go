package handler

import (
	"context"
	"encoding/json"
	"esb/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Usecase interface {
	GetCustomers(ctx context.Context) (*domain.ApiResponse[[]domain.CustomerEntity], error)
	GetItems(ctx context.Context) (*domain.ApiResponse[[]domain.ItemEntity], error)
	CreateInvoice(ctx context.Context, request domain.PostInvoiceRequest) (*domain.ApiResponse[domain.InvoiceEntity], error)
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

func (h *Handler) GetItems(e echo.Context) error {
	ctx := e.Request().Context()
	resp, err := h.usecase.GetItems(ctx)
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

func (h *Handler) CreateInvoice(e echo.Context) error {
	var request domain.PostInvoiceRequest

	err := e.Bind(&request)
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			ResponseError{Message: domain.ErrBadParamInput.Error()},
		)
	}

	ctx := e.Request().Context()
	resp, err := h.usecase.CreateInvoice(ctx, request)
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			ResponseError{Message: domain.ErrInternalServerError.Error()},
		)
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	e.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(e.Response()).Encode(resp)
}
