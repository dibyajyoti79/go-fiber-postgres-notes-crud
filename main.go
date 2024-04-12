package main

import (
	"fmt"
	"go-fiber-crud/config"
	"go-fiber-crud/controller"
	"go-fiber-crud/model"
	"go-fiber-crud/repository"
	"go-fiber-crud/router"
	"go-fiber-crud/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Print("Run service ...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	// Init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	// Init Controller
	noteController := controller.NewNoteController(noteService)

	// Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8080"))
}
