package routes

import (
	"github.com/harshtricon/notes-backend-main/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/notes", controllers.GetNotes)
	app.Get("/api/note/:NoteID", controllers.GetNoteById)
	app.Put("/api/update-note/:NoteID", controllers.UpdateNoteById)
	app.Post("/api/create-note", controllers.CreateNote)
	app.Delete("/api/delete-note/:NoteID", controllers.DeleteNote)

}
