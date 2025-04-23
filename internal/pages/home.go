package pages

import "github.com/gofiber/fiber/v2"

var categories = []string {
	"Еда",
	"Животные",
	"Машины",
	"Спорт",
	"Музыка",
	"Технологии",
	"Прочее",
}

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}

	h.router.Get("/", h.home)
}

func (h HomeHandler) home(c *fiber.Ctx) error {
	return c.Render("home", categories)
}
