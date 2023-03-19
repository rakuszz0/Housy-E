package main

import (
	"fmt"
	"housy/database"
	"housy/pkg/mysql"
	"housy/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	mysql.DatabaseInit()
	database.RunMigration()
	var PORT = os.Getenv("PORT")
	fmt.Println("server running :" + PORT)

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	// fmt.Println("server running localhost:5000")
	e.Logger.Fatal(e.Start(":" + PORT))
}
