package controllers

import (
	"amrank/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var rank []models.Rank

	models.DB.Db.Order("rseult desc").Find(&rank)

	return c.Status(fiber.StatusOK).JSON(rank)

}

func Create(c *fiber.Ctx) error {

	rank := new(models.Rank)

	if err := c.BodyParser(rank); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&rank)

	return c.Status(fiber.StatusCreated).JSON(rank)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var rank models.Rank

	result := models.DB.Db.Where("id = ?", id).First(&rank)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "Quiz not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.JSON(rank)
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	var updatedData models.Rank

	// Parse the body to get the updated user data
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if models.DB.Db.Where("id = ?", id).Updates(&updatedData).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})

}