package api

import (
	"os"

	"github.com/4babushkin/threads_rest_api/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {

	//Load environmenatal variables
	// err :=
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	port := os.Getenv("APP_PORT")
	route := gin.Default()
	// models.ConnectDB() // new

	// route.GET("/", func(context *gin.Context) {
	// 	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	// })
	route.GET("/instanses", controllers.GetAllInstanses)
	route.GET("/instanse/:instans", controllers.Get)
	route.Run(port)
}
