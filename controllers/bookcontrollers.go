package controllers

import (
	"net/http"

	"github.com/PriantikoNap/go-fiber-book.git/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var books []database.Book

	database.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(books)
}

func Create(c *fiber.Ctx) error {
	var books database.Book

	if err := c.BodyParser(&books); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&books).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(books)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var book database.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(book)
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var book database.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if database.DB.Where("id=?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak Dapat Mengupdate Data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var book database.Book

	if database.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat Menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Data berhasil dihapus",
	})
}
