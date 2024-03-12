package routes 

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.findUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.findUserByEmail)
	r.POST("/createUser", controller.createUser)
	r.PUT("/updateUser/:userId", controller.updateUser)
	r.DELETE("/deleteUser/:userId", controller.deleteUser)
}