package handlers

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apiresponse"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apperror"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

type OrderHandler struct {
	serviceContainer *services.ServiceContainer
}

func NewOrderHandler(serviceContainer *services.ServiceContainer) *OrderHandler {
	return &OrderHandler{
		serviceContainer: serviceContainer,
	}
}

const (
	DefaultTopCustomerLimit = 5
)

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var (
		err     error
		bindErr error
	)
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, "There is an error", apperror.New(err.Error()))
		} else if bindErr != nil {
			apiresponse.Error(c, http.StatusBadRequest, "Validation failed", apperror.New(bindErr.Error()))
		}
	}()

	var userID int
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(int); ok {
			userID = v
		}
	}

	response, err := h.serviceContainer.OrderService.CreateOrder(c, userID)
	if err != nil {
		return
	}

	apiresponse.Success(c, response, "Order created successfully")
}

func (h *OrderHandler) GetTopCustomers(c *gin.Context) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, "Exception", apperror.New(err.Error()))
		}
	}()

	limit := DefaultTopCustomerLimit
	if v, ok := c.Get("limit"); ok {
		if v, ok := v.(int); ok {
			limit = v
		}
	}

	customers, err := h.serviceContainer.OrderService.GetTopCustomers(c, limit)
	if err != nil {
		return
	}

	apiresponse.Success(c, customers, "")
}
