package controllers

import (
	"example.com/user/hello/databases"
	"example.com/user/hello/models"
	"github.com/gofiber/fiber/v2"
)

func New(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	product := models.Product{
		Name:   data["name"],
		Price:  data["price"],
		Weight: data["weight"],
	}
	databases.DB.Create(&product)
	return c.JSON(product)
}
func GetAll(c *fiber.Ctx) error {
	var product []models.Product
	res := databases.DB.Find(&product)
	if res.Error != nil {
		return res.Error
	}
	return c.JSON(product)
}
func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	res := databases.DB.Find(&product, id)
	if res.Error != nil {
		return c.SendString("err")
	}
	return c.JSON(product)
}
func Edit(c *fiber.Ctx) error {
	p := models.Product{}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	product := models.Product{
		Name:   data["name"],
		Price:  data["price"],
		Weight: data["weight"],
	}
	id := c.Params("id")

	databases.DB.Find(&p, id).Updates(
		map[string]interface{}{
			"Name":   product.Name,
			"Price":  product.Price,
			"weight": product.Weight,
		})

	// Using Find()
	//databases.DB.Find(&product, id).Update("Name", "San Diego")

	return c.SendString("Product update succesfully")
}
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	databases.DB.First(&product, id)
	databases.DB.Delete(&product)
	return c.SendString("Product succesfully delete from database")
}
