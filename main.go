package main

import (
	"github.com/gofiber/fiber"
	route "github.com/ancogamer/simple.example.with.fiber/controller/handler"

)

func main(){
	app:=fiber.New(fiber.Settings{
		ErrorHandler:              nil,
		ServerHeader:              "",
		StrictRouting:             false,
		CaseSensitive:             false,
		Immutable:                 false,
		UnescapePath:              false,
		ETag:                      false,
		Prefork:                   false,
		BodyLimit:                 10*1024*1024, //setting the body limit of all requests to 10 megabytes
		Concurrency:               0,
		DisableKeepalive:          false,
		DisableDefaultDate:        false,
		DisableDefaultContentType: false,
		DisableHeaderNormalizing:  false,
		DisableStartupMessage:     false,
		Views:                     nil,
		ReadTimeout:               0,
		WriteTimeout:              0,
		IdleTimeout:               0,
		ReadBufferSize:            0,
		WriteBufferSize:           0,
		CompressedFileSuffix:      "",
	})
	route.AllRoutes(app)
	app.Listen(":8081")
}
