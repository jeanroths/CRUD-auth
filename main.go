package main 

import ( 
	"github.com/gin-gonic/gin"
	godotenv "github.com/joho/godotenv"
	"github.com/jeanroths/CRUD-auth/src/controller/routes"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"log"
)
func main() {
	logger.Info("About to start application")
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err:= router.Run(":8080");	err != nil {
		log.Fatal(err)
	}
}