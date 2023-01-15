package routes

import (
	"Sample-APIs/pkg/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func GetRoutes() {
	app := fiber.New()
	app.Get("/", controllers.HandleError)
	app.Get("api/number/:id", controllers.GetDataById)
	app.Get("api/name/:name", controllers.GetDataByName)
	app.Listen(":8080")

	limiterConfig := limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1:8080"
		},
		Max:        5,
		Expiration: 30 * time.Second,
		//Storage:    myCustomStorage{},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON("Limit reached")
		},
	}

	app.Use(limiter.New(limiterConfig))

}
