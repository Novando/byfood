package urlModifier

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/url-modifier/controller"
)

func Init(r fiber.Router) {
	r.Post("/url-modifier", controller.ProcessUrl)
}
