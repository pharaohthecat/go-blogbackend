package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pharaohthecat/blog-application-go/blogbackend/database"
	"github.com/pharaohthecat/blog-application-go/blogbackend/models"
	"github.com/pharaohthecat/blog-application-go/blogbackend/util"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error{
	var blogpost models.Blog

	if err := c.BodyParser(&blogpost); err != nil{
		fmt.Println("Não pude fazer o parse body")
	}

	if err := database.DB.Create(&blogpost).Error; err != nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" : "Payload inválido",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "Post está ativo!",
	})

}

func AllPost(c *fiber.Ctx) error{
	page,_ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page-1) * limit
	var total int64
	var getBlog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getBlog)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data" : getBlog,
		"meta" : fiber.Map{
			"total":total,
			"page":page,
			"last_page":math.Ceil(float64(int(total)/limit)),
		},
	})
}


func DetailPost(c *fiber.Ctx) error{
	id,_ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data":blogpost,
	})	
}


func UpdatePost(c *fiber.Ctx) error {
	id,_ :=strconv.Atoi(c.Params("id"))
	blog:=models.Blog{
		Id:uint(id),
	}
 
	if err:=c.BodyParser(&blog);err!=nil{
		fmt.Println("Não pude fazer o parse body")
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message":"Post atualizado com sucesso!",
	})	
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)

	return c.JSON(blog)
	
}


func DeletePost(c *fiber.Ctx) error {
	id,_ :=strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id:uint(id),
	}
	deleteQuery:=database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Não achei o registro a ser deletado",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Apagado com sucesso!",
	})
	

}


