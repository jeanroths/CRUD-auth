package main 

import ( 
	"github.com/gin-conic/gin"
	godotenv "github.com/joho/godotenv"
	"log"
)
func main() {
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