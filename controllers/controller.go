package controllers

import (
	"am_chat/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	var chat []models.Chat

	models.DB.Db.Find(&chat)

	return c.Status(fiber.StatusOK).JSON(chat)

}

func Create(c *fiber.Ctx) error {

	chat := new(models.Chat)

	if err := c.BodyParser(chat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&chat)

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func Show(c *fiber.Ctx) error {

	chat := &models.Chat{}
	id := c.Params("idrec")

	if err := models.DB.Db.First(chat, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(chat)
}

func Update(c *fiber.Ctx) error {

	chat := &models.Chat{}
	id := c.Params("idrec")

	if err := models.DB.Db.First(chat, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	if err := c.BodyParser(chat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Save(chat)

	return c.Status(fiber.StatusOK).JSON(chat)

}

func Delete(c *fiber.Ctx) error {

	chat := &models.Chat{}
	id := c.Params("idrec")

	if err := models.DB.Db.First(chat, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	models.DB.Db.Delete(chat, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Delete Data"})
}