package routes 

import (
	"github.com/jeanroths/CRUD-auth/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.findUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.findUserByEmail)
	r.POST("/createUser", controller.createUser)
	r.PUT("/updateUser/:userId", controller.updateUser)
	r.DELETE("/deleteUser/:userId", controller.deleteUser)
}