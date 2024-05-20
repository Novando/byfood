package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/novando/byfood/be/pkg/postgresql"
	"os"
	"strconv"
)

func main() {
	// Init environment variable
	envPath := fmt.Sprint("./config/.env.local")
	err := godotenv.Load(envPath)
	if err != nil {
		envPath = fmt.Sprint("./config/.env")
		err = godotenv.Load(envPath)
	}
	if err != nil {
		panic(err)
	}

	// Init DB
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	pgxpool, query, err := postgresql.Init(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		dbPort,
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		panic(err)
	}
	defer pgxpool.Close()

	// init App
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})
	app.Use(cors.New())
	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		fmt.Println(err.Error())
	}
}
