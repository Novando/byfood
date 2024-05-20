package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/controller"
	"github.com/novando/byfood/be/internal/service"
	"github.com/novando/byfood/be/pkg/reposqlc"
)

func Init(app *fiber.App, db *reposqlc.Queries) {
	v1 := app.Group("/v1")

	sb := service.NewBookService(db)

	cb := controller.NewBookController(sb)

	book := v1.Group("/books")
	book.Delete("/:id", cb.Delete)
	book.Put("/:id", cb.Update)
	book.Post("/", cb.Create)
}
