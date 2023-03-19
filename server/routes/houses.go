package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func HouseRoutes(e *echo.Group) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)

	e.GET("/houses", h.FindHouses)
	e.GET("/houses-filter", h.FindHousesFilter)
	e.GET("/house/:id", h.GetHouse)
	e.POST("/house", middleware.UploadFile(h.CreateHouse))
	e.DELETE("/house/:id", h.DeleteHouse)
	e.PATCH("/house/:id", middleware.Auth(middleware.UploadFile(h.UpdateHouse)))
}
