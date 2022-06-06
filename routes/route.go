package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pharaohthecat/blog-application-go/blogbackend/controller"
	"github.com/pharaohthecat/blog-application-go/blogbackend/middleware"
)

// Configuração das rotas e métodos
func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)          // C- Criar post
	app.Get("/api/allpost", controller.AllPost)           // R - Listar todos os posts
	app.Get("/api/allpost/:id", controller.DetailPost)    // Detalhes do post
	app.Put("/api/updatepost/:id", controller.UpdatePost) // U - Atualizar post
	app.Get("/api/uniquepost", controller.UniquePost) // Post postados por um usuário identificado por id
	app.Delete("/api/deletepost/:id", controller.DeletePost) // D - Deletar um post
	app.Post("/api/upload-image", controller.Upload) // Upload da Imagem
	app.Static("/api/uploads", "./uploads") // Link gerado com os dados na pasta upload

}
