package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanroths/CRUD-auth/tree/main/src/configuration/rest_err"
)
func CreateUser(c *gin.Context){
	err := rest_err.NewBadRequestError("Rota chamada de maneira errada")
	c.JSON(err.Code, err)
	// code to create user
}
