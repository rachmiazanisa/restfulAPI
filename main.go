package main

import (
	"github.com/gin-gonic/gin"
	controllers "RestfulAPI/restfulAPI/controllers"
)

func main() {
	engine := gin.Default()

	//Read Data by id
	v3 := engine.Group("/db/home")
	{
		v3.GET("/", controllers.DBHome)
	}
	//Delete Data by id
	v4 := engine.Group("/db/delete")
	{
		v4.GET("/", controllers.DBDelete)
	}
	//Update Datab by id
	v5 := engine.Group("/db/update")
	{
		v5.GET("/", controllers.DBUpdate)
	}
	//Insert Data
	v6 := engine.Group("/db/insert")
	{
		v6.GET("/", controllers.DBInsert)
	}
	engine.Run(":8080")
}
