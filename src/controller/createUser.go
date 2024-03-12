package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanroths/CRUD-auth/src/controller/model/request"
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"fmt"
	"log"
)
func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil	{
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		restErr := rest_err.NewBadRequestError("Some fields are incorrect")

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
	// code to create user
}
