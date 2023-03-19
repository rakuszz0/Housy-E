package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
	// e.GET("/user/{id}", h.DeleteUser).Methods("DELETE")
	// e.GET("/user/{id}", middleware.UploadFile(h.UpdateUser)).Methods("PATCH")
	e.PATCH("/change-password", middleware.Auth(h.ChangePassword))
	e.PATCH("/change-image", middleware.Auth(middleware.UploadFile(h.ChangeImage)))
}
