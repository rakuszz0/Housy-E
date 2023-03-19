package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"housy/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/sign-up", h.SignUp)
	e.POST("/sign-in", h.SignIn) // add this code
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
