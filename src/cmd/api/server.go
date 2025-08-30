package main

import (
	services2 "Spy-Cat-Agency/src/internal/missions/application/services"
	repository2 "Spy-Cat-Agency/src/internal/missions/infrastructure/repository"
	handlers2 "Spy-Cat-Agency/src/internal/missions/interfaces/handlers"
	"Spy-Cat-Agency/src/internal/shared/db"
	"Spy-Cat-Agency/src/internal/shared/router"
	"Spy-Cat-Agency/src/internal/spycats/application/services"
	"Spy-Cat-Agency/src/internal/spycats/infrastructure/repository"
	"Spy-Cat-Agency/src/internal/spycats/interfaces/handlers"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// @title  SpyCats Service API
// @version 1.0
// @description API for managing spy cats and missions with targets.

// @schemes http
// @host localhost:8080
// @BasePath /api
func main() {

	if err := godotenv.Load(".env.example"); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	ctx := context.Background()

	newPool, err := db.ConnectDb(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	spyCatRepo := repository.NewSpyCatPgxRepository(newPool)

	spyCatService := services.NewSpyCatService(
		spyCatRepo,
		spyCatRepo,
		spyCatRepo,
		spyCatRepo,
	)

	spyCatHandler := &handlers.SpyCatHandler{
		Service: spyCatService,
	}

	missionRepo := repository2.NewMissionPgxRepository(newPool)

	missionService := services2.NewMissionService(
		missionRepo,
		missionRepo,
		missionRepo,
		missionRepo,
		spyCatRepo,
		newPool,
	)

	missionHandler := &handlers2.MissionHandler{
		Service: missionService,
	}

	mainRouter := router.SetUpRouter(spyCatHandler, missionHandler)

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println(time.Now())

	if err := mainRouter.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
