package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marioheryanto/linkaja/go-app/controller"
	"github.com/marioheryanto/linkaja/go-app/database"
	"github.com/marioheryanto/linkaja/go-app/helper"
	"github.com/marioheryanto/linkaja/go-app/library"
	"github.com/marioheryanto/linkaja/go-app/repository"
	"github.com/marioheryanto/linkaja/go-app/route"
)

func init() {
	godotenv.Load()
}

func main() {
	// clients
	dbClient := database.ConnectDB()
	validator := helper.NewValidator()

	// repo
	accountRepo := repository.NewAccountRepository(dbClient)

	// library
	accountLib := library.NewAccountLibrary(accountRepo, validator)

	// controller
	accountCtrl := controller.NewAccountController(accountLib)

	router := gin.Default()

	// config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	// config.AllowHeaders = []string{}

	// ----- Router -----
	router.Use(cors.New(config))

	route.AccountRoutes(router, accountCtrl)

	// router.Run(fmt.Sprintf("http://localhost:%v", os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))

}
