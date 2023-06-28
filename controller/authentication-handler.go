package controller

import (
	"net/http"
	"time"

	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/cyneptic/letsgo-authentication/internal/core/ports"
	"github.com/cyneptic/letsgo-authentication/internal/core/service"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)
type AuthenticationHandler struct {
	svc ports.UserServiceContract
}

func NewAuthenticationHandler() *AuthenticationHandler {
	svc := service.NewAuthenticationService()
	return &AuthenticationHandler{
		svc: svc,
	}
}
func AddAuthServiceRoutes(e echo.Echo)  {
	h := NewAuthenticationHandler()
	e.POST("/login", h.login)
	e.POST("/register", h.register)

}

func (h *AuthenticationHandler) login(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// err := validators.ValidateUserLogin(*user)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	token, err := h.svc.LoginHandler(*user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": "Invalid Email Or Password",
		})
	}

	return c.JSON(200, token)
}

func (h *AuthenticationHandler) register(c echo.Context) error {

	newUser := new(entities.User)

	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now()
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	// err := validators.ValidateUserRegister(*newUser)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	err := h.svc.AddUser(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"newUser": newUser,
	})
}