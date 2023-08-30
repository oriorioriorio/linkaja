package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/linkaja/go-app/helper"
	"github.com/marioheryanto/linkaja/go-app/library"
	"github.com/marioheryanto/linkaja/go-app/model"
)

type AccountController struct {
	lib library.AccountLibraryInterface
}

type AccountControllerInterface interface {
	CheckBalance(c *gin.Context)
	Transfer(c *gin.Context)
}

func NewAccountController(lib library.AccountLibraryInterface) AccountControllerInterface {
	return AccountController{
		lib: lib,
	}
}

func (ctrl AccountController) CheckBalance(c *gin.Context) {
	request := c.Param("number")
	response := model.Response{}

	if request == "" {
		response.Message = "account number is empty"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	acc, err := ctrl.lib.CheckBalance(context.Background(), request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = acc
	c.JSON(http.StatusOK, response)
}

func (ctrl AccountController) Transfer(c *gin.Context) {
	request := model.TransferParams{}
	response := model.Response{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	request.FromAccountNumber = c.Param("from_account_number")

	err = ctrl.lib.Transfer(context.Background(), request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	c.JSON(http.StatusCreated, nil)
}
