package controllers

import (
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
	"strconv"
	"suntop/config"
	"suntop/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadProduct(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("error retrieving file", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error retrieving file",
		})
	}

	var products []models.Product

	if err := parseCSV(file, &products); err != nil {
		log.Println("error parsing CSV", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing CSV",
		})
	}

	for _, product := range products {
		config.DB.Create(&product)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "products added",
	})
}

func parseCSV(file *multipart.FileHeader, products *[]models.Product) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	// Skip the header row
	if _, err := reader.Read(); err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Convert quantity and code to appropriate types
		quantity, _ := strconv.Atoi(record[2])       // Assuming Quantity is the third column
		code, _ := strconv.ParseFloat(record[3], 64) // Assuming Code is the fourth column

		product := models.Product{
			ItemName:    record[0], // First column is ItemName
			Description: record[1], // Second column is Description
			Quantity:    quantity,
			Code:        code,
			Cartegory:   record[4], // Fifth column is Cartegory
			Caterlogue:  record[5], // Sixth column is Caterlogue
		}

		*products = append(*products, product)
	}

	return nil
}

func UploadSchedule(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("error retrieving file", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error retrieving file",
		})
	}

	var schedules []models.Schedule

	if err := parseScheduleCSV(file, &schedules); err != nil {
		log.Println("error parsing CSV", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing CSV",
		})
	}

	for _, schedule := range schedules {
		newSchedule := models.Schedule{
			Name:      schedule.Name,
			StartTime: schedule.StartTime,
			EndTime:   schedule.EndTime,
		}
		config.DB.Create(&newSchedule)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "schedules added",
	})
}

func parseScheduleCSV(file *multipart.FileHeader, schedules *[]models.Schedule) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	// Skip the header row
	if _, err := reader.Read(); err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Parse StartTime and EndTime
		startTime, err := time.Parse("2006-01-02 15:04", record[1]) // Adjust format as needed
		if err != nil {
			log.Println("error parsing StartTime:", err)
			return err
		}

		endTime, err := time.Parse("2006-01-02 15:04", record[2]) // Adjust format as needed
		if err != nil {
			log.Println("error parsing EndTime:", err)
			return err
		}

		schedule := models.Schedule{
			Name:      record[0],
			StartTime: startTime,
			EndTime:   endTime,
		}

		*schedules = append(*schedules, schedule)
	}

	return nil
}
