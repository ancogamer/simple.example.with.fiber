package main

import (
	"github.com/gofiber/fiber/v2"
	route "github.com/ancogamer/simple.example.with.fiber/controller/handler"

)

func main(){
	app:=fiber.New()
	route.AllRoutes(app)
	app.Listen(":8081")
}
