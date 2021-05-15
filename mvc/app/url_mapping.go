package app
 
import (
	"github.com/kirankothule/golang-microservice/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}