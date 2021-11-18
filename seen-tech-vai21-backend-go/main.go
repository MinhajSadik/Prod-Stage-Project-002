package main

import (
	"SEEN-TECH-VAI21-BACKEND-GO/DBManager"
	"SEEN-TECH-VAI21-BACKEND-GO/Routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetupRoutes(app *fiber.App) {
	Routes.MaterialRoute(app.Group("/material"))
	Routes.InventoryRoute(app.Group("/inventory"))
	Routes.UnitsOfMeasurementRoute(app.Group("/units_of_measurement"))
	Routes.ProductRoute(app.Group("/product"))
	Routes.ContactRoute(app.Group("/contact"))
	Routes.PriceListRoute(app.Group("/price_list"))
	Routes.SettingRoute(app.Group("/setting"))
	Routes.ProdStagesRoute(app.Group("/prodstages"))
	Routes.SalesRoute(app.Group("/sales"))
}

func main() {

	fmt.Println(("Hello SEEN-TECH-CHIR"))
	fmt.Print("Initializing Database Connection ... ")
	initState := DBManager.InitCollections()

	if initState {
		fmt.Println("[OK]")
	} else {
		fmt.Println("[FAILED]")
		return
	}

	fmt.Print("Initializing the server ... ")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(pprof.New())
	SetupRoutes(app)
	app.Static("/", "./Public")
	fmt.Println("[OK]")
	app.Listen(":8080")

}
