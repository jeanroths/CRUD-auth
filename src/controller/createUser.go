package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanroths/CRUD-auth/src/controller/model/request"
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"fmt"
)
func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil	{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Invalid fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
	// code to create user
}
