package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/trigo-William-HOANG/biweb/models"
	"github.com/trigo-William-HOANG/biweb/storage"
	"gorm.io/gorm"
)

type App struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Link        string `json:"link"`
}
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateApp(c *fiber.Ctx) error {
	app := App{}
	err := c.BodyParser(&app)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&app).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to create app"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "app created successfully"})
	return nil

}

func (r *Repository) GetApps(c *fiber.Ctx) error {
	appModels := &[]models.Apps{}
	err := r.DB.Find(appModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to get apps"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "apps fetched successfully",
			"data":    appModels,
		})
	return nil
}

func (r *Repository) GetAppsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	appModel := &models.Apps{}
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id is required"})
		return nil
	}

	fmt.Println("The ID is ", id)

	err := r.DB.Where("id = ?", id).First(appModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to get app"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "app fetched successfully", "data": appModel})
	return nil
}

func (r *Repository) DeleteApp(c *fiber.Ctx) error {
	appModel := &models.Apps{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id is required"})
		return nil
	}

	err := r.DB.Delete(appModel, id)

	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to delete app"})
		return err.Error
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "app deleted successfully"})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/add_applications", r.CreateApp)
	api.Delete("delete_app/:id", r.DeleteApp)
	api.Get("/get_apps/:id", r.GetAppsByID)
	api.Get("/get_apps", r.GetApps)
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database ")
	}

	err = models.MigrateApps(db)

	if err != nil {
		log.Fatal("Could not migrate the database")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":3000")
}
