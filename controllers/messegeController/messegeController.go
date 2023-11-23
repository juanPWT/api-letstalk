package messegecontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juanPWT/api-letstalk/models"
)

func Index(c *fiber.Ctx) error {

	var messeges []models.Messege

	models.DB.Find(&messeges)

	return c.Status(fiber.StatusOK).JSON(messeges)
}

func Create(c *fiber.Ctx) error {
	var messege models.Messege

	if err := c.BodyParser(&messege); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messege": err.Error(),
		})
	}

	if err := models.DB.Create(&messege).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messege": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"messege": "success created messsge",
		"from":    messege.From,
		"content": messege.Content,
	})
}
