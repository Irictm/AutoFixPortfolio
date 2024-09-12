package main

import (
	"database/sql"
	"log"

	"github.com/Irictm/AutoFixPortfolio/Backend/main/Controllers/vehicleController"
	"github.com/Irictm/AutoFixPortfolio/Backend/main/Repositories/vehicleRepository"
	"github.com/Irictm/AutoFixPortfolio/Backend/main/Services/vehicleService"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dataSourceName := "postgres://postgres:postgres@localhost:5432/autofix?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	vehicleRepository := vehicleRepository.VehicleRepository{DB: db}
	vehicleService := vehicleService.VehicleService{Repository: vehicleRepository}
	vehicleController := vehicleController.VehicleController{Service: vehicleService}

	router := gin.Default()
	vehicleController.LinkPaths(router)
	router.Run("localhost:8080")
}
