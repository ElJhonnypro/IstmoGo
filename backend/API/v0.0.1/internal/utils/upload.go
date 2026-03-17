package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveImage(
	c *fiber.Ctx,
	file *multipart.FileHeader,
	prefix string, // car | rid | user
	basePath string, // ../data/uploads/
) (string, error) {

	if !IsvalidImage(file) {
		return "", ErrInvalidImage
	}

	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + "_" + prefix + ext
	fullPath := filepath.Join(basePath, filename)

	if err := c.SaveFile(file, fullPath); err != nil {
		return "", err
	}

	// Lo que se guarda en la DB
	return "uploads/" + filename, nil
}
