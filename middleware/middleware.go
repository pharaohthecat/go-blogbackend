package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pharaohthecat/blog-application-go/blogbackend/util"
)

func IsAuthenticate(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")

	if _,err := util.ParseJwt(cookie); err != nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"Não autenticado",
		})
	}

	return c.Next()
}