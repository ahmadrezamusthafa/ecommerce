package handlers

import (
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity/requests"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apiresponse"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apperror"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strconv"
)

type CartHandler struct {
	serviceContainer *services.ServiceContainer
}

func NewCartHandler(serviceContainer *services.ServiceContainer) *CartHandler {
	return &CartHandler{
		serviceContainer: serviceContainer,
	}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, "Exception", apperror.New(err.Error()))
		}
	}()

	var userID int
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(int); ok {
			userID = v
		}
	}

	cart, err := h.serviceContainer.CartService.GetCart(c, userID)
	if err != nil {
		return
	}

	apiresponse.Success(c, cart, "")
}

func (h *CartHandler) AddItemToCart(c *gin.Context) {
	var (
		err     error
		bindErr error
	)
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, err.Error(), apperror.New(err.Error()))
		} else if bindErr != nil {
			apiresponse.Error(c, http.StatusBadRequest, err.Error(), apperror.New(bindErr.Error()))
		}
	}()

	var req requests.AddToCartRequest
	if bindErr = c.ShouldBindJSON(&req); bindErr != nil {
		return
	}

	var userID int
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(int); ok {
			userID = v
		}
	}

	cartItem := entity.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	err = h.serviceContainer.CartService.AddItemToCart(c, userID, cartItem)
	if err != nil {
		return
	}

	apiresponse.Success(c, req.ProductID, "Item added successfully")
}

func (h *CartHandler) RemoveItemFromCart(c *gin.Context) {
	var (
		err     error
		bindErr error
	)
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, err.Error(), apperror.New(err.Error()))
		} else if bindErr != nil {
			apiresponse.Error(c, http.StatusBadRequest, err.Error(), apperror.New(bindErr.Error()))
		}
	}()

	var userID int
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(int); ok {
			userID = v
		}
	}

	productIDStr := c.Param("id")
	if productIDStr == "" {
		err = errors.New("product id is required")
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return
	}

	err = h.serviceContainer.CartService.RemoveItemFromCart(c, userID, productID)
	if err != nil {
		return
	}

	apiresponse.Success(c, productID, "Item removed successfully")
}
