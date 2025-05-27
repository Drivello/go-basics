package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"ent-orm-basic/ent"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var client *ent.Client

func main() {
	var err error

	// Conexión a PostgreSQL
	client, err = ent.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Error conectando a PostgreSQL: %v", err)
	}
	defer client.Close()

	// Crear tablas si no existen
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Error creando el esquema: %v", err)
	}

	// Iniciar Fiber
	app := fiber.New()

	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	log.Println("Servidor iniciado en http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}

func getUsers(c *fiber.Ctx) error {
	users, err := client.User.Query().All(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := client.User.Get(context.Background(), id)
	if err != nil {
		return c.Status(404).SendString("Usuario no encontrado")
	}
	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	type input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Age      int    `json:"age"`
	}

	var body input
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("JSON inválido")
	}

	user, err := client.User.
		Create().
		SetName(body.Name).
		SetEmail(body.Email).
		SetUsername(body.Username).
		SetAge(body.Age).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	type input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Age      int    `json:"age"`
	}

	var body input
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("JSON inválido")
	}

	user, err := client.User.
		UpdateOneID(id).
		SetName(body.Name).
		SetEmail(body.Email).
		SetUsername(body.Username).
		SetAge(body.Age).
		Save(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := client.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Usuario eliminado")
}
