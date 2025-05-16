package routes

import (
	"archive-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupArchiveRoutes(app fiber.Router) {
	app.Post("/journal", controllers.CreateJournalArchive)
	app.Post("/transaction", controllers.CreateTransactionArchive)
}
