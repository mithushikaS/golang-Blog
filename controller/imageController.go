package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	filename := ""
	for _, file := range files {
		filename = randLetter(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}
	return c.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + filename,
	})
}