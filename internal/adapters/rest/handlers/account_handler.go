package handlers

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity/requests"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apiresponse"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apperror"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

type AccountHandler struct {
	serviceContainer *services.ServiceContainer
}

func NewAccountHandler(serviceContainer *services.ServiceContainer) *AccountHandler {
	return &AccountHandler{
		serviceContainer: serviceContainer,
	}
}

func (h *AccountHandler) GetAccountBalanceInfo(c *gin.Context) {
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

	account, err := h.serviceContainer.AccountService.GetAccountByUserID(c, userID)
	if err != nil {
		return
	}

	apiresponse.Success(c, account, "")
}

func (h *AccountHandler) Withdraw(c *gin.Context) {
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

	var req requests.WithdrawBalanceRequest
	if bindErr = c.ShouldBindJSON(&req); bindErr != nil {
		return
	}

	var userID int
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(int); ok {
			userID = v
		}
	}

	err = h.serviceContainer.AccountService.Withdraw(c, userID, req.Amount)
	if err != nil {
		return
	}

	apiresponse.Success(c, req.Amount, "Withdraw user balance successfully")
}

func (h *AccountHandler) Deposit(c *gin.Context) {
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

	var req requests.DepositBalanceRequest
	if bindErr = c.ShouldBindJSON(&req); bindErr != nil {
		return
	}

	var userID int
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(int); ok {
			userID = v
		}
	}

	err = h.serviceContainer.AccountService.Deposit(c, userID, req.Amount)
	if err != nil {
		return
	}

	apiresponse.Success(c, req.Amount, "Deposit user balance successfully")
}
