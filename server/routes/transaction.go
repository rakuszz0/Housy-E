package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("", h.FindTransactions)
	e.GET("/:id", h.GetTransaction, middleware.Auth)
	e.POST("/transaction", h.CreateTransaction)
	// g.DELETE("/:id", h.DeleteTransaction, middleware.Auth)
	e.POST("/notification", h.Notification)
}
