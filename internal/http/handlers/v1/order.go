package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	httpserver "top-selection-test/internal/http"
	"top-selection-test/internal/model"
)

// Usually we don't use repository direct in the handlers,
// but in current implementation abstraction between transport and DAL is not neccessary
type OrderRepository interface {
	Create(ctx context.Context, order model.Order) error
}

type Orders struct {
	orderRepository OrderRepository
}

func NewOrders(os OrderRepository) *Orders {
	return &Orders{
		orderRepository: os,
	}
}

func (h *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	var o model.Order
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&o); err != nil {
		return httpserver.ResponseError{
			Code:          http.StatusBadRequest,
			ResponseError: errors.New("invalid json format"),
			VerboseErr:    err,
		}
	}

	// TODO Truncate dates
	if err := h.orderRepository.Create(r.Context(), o); err != nil {
		// TODO Error handling for different repo errors
		return httpserver.ResponseError{
			Code:          http.StatusConflict,
			ResponseError: errors.New("room is not avaliable for given dates"),
			VerboseErr:    err,
		}
	}

	return nil
}
