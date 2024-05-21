package library

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/library/controller"
	"github.com/novando/byfood/be/internal/library/service"
	"github.com/novando/byfood/be/pkg/reposqlc"
)

func Init(r fiber.Router, db *reposqlc.Queries) {

	sb := service.NewBookService(db)

	cb := controller.NewBookController(sb)

	book := r.Group("/books")
	book.Delete("/:id", cb.Delete)
	book.Put("/:id", cb.Update)
	book.Get("/:id", cb.Detail)
	book.Get("/", cb.Read)
	book.Post("/", cb.Create)
}
