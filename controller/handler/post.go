package handler

import "github.com/gofiber/fiber/v2"

func Get(c*fiber.Ctx) error{

	return c.SendStatus(200).SendString("hello")
}