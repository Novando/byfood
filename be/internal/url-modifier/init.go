package urlModifier

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/url-modifier/controller"
	"github.com/novando/byfood/be/pkg/reposqlc"
)

func Init(r fiber.Router, db *reposqlc.Queries) {
	r.Post("/url-cleanup", controller.ProcessUrl)
}
