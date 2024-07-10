package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	httpserver "top-selection-test/internal/http"
	"top-selection-test/internal/model"
	"top-selection-test/internal/repository"

	"github.com/go-playground/validator/v10"
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

	o.From = o.From.Truncate(24 * time.Hour)
	o.To = o.To.Truncate(24 * time.Hour)

	validate := validator.New()
	if err := validate.Struct(o); err != nil {
		return httpserver.ResponseError{
			Code:          http.StatusUnprocessableEntity,
			VerboseErr:    err,
			ResponseError: err,
		}
	}

	if err := h.orderRepository.Create(r.Context(), o); err != nil {
		return convertRepositoryErrorToHTTPError(err)
	}

	return nil
}

func convertRepositoryErrorToHTTPError(err error) error {
	var rErr error

	switch {
	case errors.Is(err, repository.ErrRoomDoesntExist):
		rErr = httpserver.ResponseError{
			Code:          http.StatusNotFound,
			ResponseError: errors.New("room not found"),
			VerboseErr:    err,
		}
	case errors.Is(err, repository.ErrRoomNotAvailable):
		rErr = httpserver.ResponseError{
			Code:          http.StatusConflict,
			ResponseError: errors.New("room is not available for given dates"),
			VerboseErr:    err,
		}
	default:
		rErr = err
	}

	return rErr
}
