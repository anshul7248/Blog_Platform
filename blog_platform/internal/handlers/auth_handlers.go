package handlers

import (
	"blog_project/internal/services"
	"blog_project/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service *services.UserService
}

func NewAuthHandler(s *services.UserService) *AuthHandler { return &AuthHandler{Service: s} }

type SignupReq struct {
	UserName    string `json:"user_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var req SignupReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid payload")
	}
	u, err := h.Service.Register(req.UserName, req.Email, req.Password)
	if err != nil {
		return utils.Err(c, http.StatusBadRequest, err.Error())
	}
	return utils.JSON(c, http.StatusCreated, true, "user_created", u)
}

type loginReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req loginReq

	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid payload")
	}
	user, token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		return utils.Err(c, http.StatusBadRequest, "invalid credentials")
	}
	return utils.JSON(c, http.StatusOK, true, "login successfull", map[string]interface{}{"token": token, "user": user})
}
