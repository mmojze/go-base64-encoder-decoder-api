package main

import (
	"log"
	"os"

	router "./routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	/* Load env file */
	var err error
	err = godotenv.Load("config.env")

	/* Check if there was an error with the env file */
	if err != nil {
		log.Fatal("Error loading enviroment variables file")
	}

	/* Change gin mode to production and dettach debugger */
	if os.Getenv("MODE") == "production" {
		gin.SetMode("release")
	}

	/* Start server */
	server := gin.Default()
	server.Use(gin.Recovery())

	/* Common route to every endpoint */
	v1 := server.Group("/api/v1")

	/* Mount router */
	router.MountRoutes(v1)

	/* Run server */
	server.Run()

}
