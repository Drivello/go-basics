package main

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items      = make(map[int]Item)
	idxCounter = 1
	mutex      = &sync.Mutex{}
)

func main() {
	app := fiber.New()

	app.Get("/items", getItems)
	app.Get("/items/:id", getItem)
	app.Post("/items", createItem)
	app.Put("/items/:id", updateItem)
	app.Delete("/items/:id", deleteItem)

	app.Listen(":3000")
}

func getItems(c *fiber.Ctx) error {
	newItemList := make([]Item, 0, len(items))

	for _, item := range items {
		newItemList = append(newItemList, item)
	}
	return c.JSON(newItemList)
}

func getItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	item, exists := items[id]
	if !exists {
		return c.Status(404).SendString("Item not found")
	}
	return c.JSON(item)
}

func createItem(c *fiber.Ctx) error {
	var item Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	mutex.Lock()
	item.ID = idxCounter
	idxCounter++
	items[item.ID] = item
	mutex.Unlock()

	return c.Status(201).JSON(item)
}

func updateItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	var updated Item
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	mutex.Lock()
	defer mutex.Unlock()

	_, exists := items[id]
	if !exists {
		return c.Status(404).SendString("Item not found")
	}

	updated.ID = id
	items[id] = updated
	return c.JSON(updated)
}

func deleteItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	mutex.Lock()
	defer mutex.Unlock()

	_, exists := items[id]
	if !exists {
		return c.Status(404).SendString("Item not found")
	}

	delete(items, id)
	return c.SendStatus(204)
}
