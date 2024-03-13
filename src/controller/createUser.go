package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanroths/CRUD-auth/src/controller/model/request"
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"go.uber.org/zap"
	"net/http"
	"github.com/jeanroths/CRUD-auth/src/controller/model/response"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context){
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil	{
		logger.Error("Error trying to validate user info", err,
			zap.String("journey","createUser"))
		restErr := rest_err.NewBadRequestError("Some fields are incorrect")

		c.JSON(restErr.Code, restErr)
		return
	}
	
	response := response.UserResponse{
		Id:"test",
		Name: userRequest.Name,
		Email: userRequest.Email,
	}
	logger.Info("User created succesfully",
			zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, response)
	// code to create user
}
