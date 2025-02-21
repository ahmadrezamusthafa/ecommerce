package handlers

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity/requests"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apiresponse"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apperror"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

type UserHandler struct {
	serviceContainer *services.ServiceContainer
}

func NewUserHandler(serviceContainer *services.ServiceContainer) *UserHandler {
	return &UserHandler{
		serviceContainer: serviceContainer,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
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

	var req requests.CreateUserRequest
	if bindErr = c.ShouldBindJSON(&req); bindErr != nil {
		return
	}

	user := entity.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	response, err := h.serviceContainer.UserService.Register(c, user)
	if err != nil {
		return
	}

	apiresponse.Success(c, response, "User registered successfully")
}

func (h *UserHandler) Update(c *gin.Context) {
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

	var req requests.UpdateUserRequest
	if bindErr = c.ShouldBindJSON(&req); bindErr != nil {
		return
	}

	var userID string
	if v, ok := c.Get("user_id"); ok {
		if v, ok := v.(string); ok {
			userID = v
		}
	}

	user := entity.User{
		ID:       userID,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	response, err := h.serviceContainer.UserService.Update(c, user)
	if err != nil {
		return
	}

	apiresponse.Success(c, response, "User updated successfully")
}

func (h *UserHandler) Login(c *gin.Context) {
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

	var req requests.UserLoginRequest
	if bindErr = c.ShouldBindJSON(&req); bindErr != nil {
		return
	}

	response, err := h.serviceContainer.UserService.Login(c, req.Email, req.Username, req.Password)
	if err != nil {
		return
	}

	apiresponse.Success(c, response, "User logged in successfully")
}
