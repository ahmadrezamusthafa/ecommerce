package handlers

import (
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apiresponse"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apperror"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strconv"
)

type ProductHandler struct {
	serviceContainer *services.ServiceContainer
}

func NewProductHandler(serviceContainer *services.ServiceContainer) *ProductHandler {
	return &ProductHandler{
		serviceContainer: serviceContainer,
	}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, err.Error(), apperror.New(err.Error()))
		}
	}()

	products, err := h.serviceContainer.ProductService.GetAllProducts(c)
	if err != nil {
		return
	}

	apiresponse.Success(c, products, "")
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, err.Error(), apperror.New(err.Error()))
		}
	}()

	productIDStr := c.Param("id")
	if productIDStr == "" {
		err = errors.New("product id is required")
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return
	}

	product, err := h.serviceContainer.ProductService.GetProductByID(c, productID)
	if err != nil {
		return
	}

	apiresponse.Success(c, product, "")
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			logger.Error(string(debug.Stack()))
			apiresponse.Error(c, http.StatusInternalServerError, "Unhandled exception", apperror.New("got panic exception"))
		} else if err != nil {
			apiresponse.Error(c, http.StatusInternalServerError, err.Error(), apperror.New(err.Error()))
		}
	}()

	query := c.Query("query")
	product, err := h.serviceContainer.ProductService.SearchProducts(c, query)
	if err != nil {
		return
	}

	apiresponse.Success(c, product, "")
}
