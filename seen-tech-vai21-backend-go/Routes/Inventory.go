package Routes

import (
	"SEEN-TECH-VAI21-BACKEND-GO/Controllers"

	"github.com/gofiber/fiber/v2"
)

func InventoryRoute(route fiber.Router) {
	route.Post("/new", Controllers.InventoryNew)
	route.Post("/get_all", Controllers.InventoryGetAll)
	route.Put("/set_status/:id/:new_status", Controllers.InventorySetStatus)
	route.Post("/modify", Controllers.InventoryModify)
}
