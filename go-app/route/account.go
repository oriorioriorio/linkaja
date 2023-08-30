package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/linkaja/go-app/controller"
)

func AccountRoutes(app *gin.Engine, c controller.AccountControllerInterface) {
	account := app.Group("/account")
	account.GET("/:number", c.CheckBalance)
	account.POST("/:from_account_number/transfer", c.Transfer)

}
