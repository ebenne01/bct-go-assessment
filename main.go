package main

import (
	"fmt"
	"log"

	"github.com/ebenne01/bct-go-assessment/controller"
	"github.com/ebenne01/bct-go-assessment/model"
	"github.com/labstack/echo/v4"
)

// Hard coding for the purpose of this demo.
// Info should be externalized
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "userdb"
)

func main() {

	// typically I would lookup the config
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	err := model.InitDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	defer model.CloseDB()

	e := echo.New()

	e.GET("/users", controller.GetAllUsers)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
