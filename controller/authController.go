package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pharaohthecat/blog-application-go/blogbackend/database"
	"github.com/pharaohthecat/blog-application-go/blogbackend/models"
	"github.com/pharaohthecat/blog-application-go/blogbackend/util"
)

func validateEmail(email string) bool {
	//https://stackoverflow.com/questions/37203366/what-does-this-regex-mean-a-z0-9-can-you-help-me-by-answering-this
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Não pude fazer o parse body")
	}

	//Checar se a senha é menor que seis caracteres
	if len(data["password"].(string)) < 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "A senha deve conter mais de seis caracteres",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Endereço e-mail inválido!",
		})

	}

	// Checar se o e-mail já existe no banco de dados da aplicação
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "E-mail já existe!",
		})
	}

	// To-String Usuário
	user := models.User{
		Name:     data["first_name"].(string),
		LastName: data["last_name"].(string),
		Phone:    data["phone"].(string),
		Email:    strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)

	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Conta criada com sucesso!",
	})
}

//Login
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Não pude fazer o parse body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Conta não existente, crie uma!",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Senha incorreta!",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logado com sucesso!",
		"user":    user,
	})

}

type Claims struct {
	jwt.StandardClaims
}
