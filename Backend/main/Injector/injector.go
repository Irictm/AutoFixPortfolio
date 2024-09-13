package injector

import (
	"context"
	"log"

	"github.com/Irictm/AutoFixPortfolio/Backend/main/Controllers/vehicleController"
	"github.com/Irictm/AutoFixPortfolio/Backend/main/Repositories/vehicleRepository"
	"github.com/Irictm/AutoFixPortfolio/Backend/main/Services/vehicleService"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func InjectDependencies(rout *gin.Engine) {
	db, err := ConnectPostgreSQL("postgres", "postgres", "localhost", "5432", "autofix")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	vehicleRepository := &vehicleRepository.VehicleRepository{DB: db}
	vehicleService := &vehicleService.VehicleService{Repository: vehicleRepository}
	vehicleController := vehicleController.VehicleController{Service: vehicleService}
	vehicleController.LinkPaths(rout)
}

func ConnectPostgreSQL(user string, pass string, host string, port string, dbname string) (*pgx.Conn, error) {
	dataSource := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/autofix?sslmode=disable"
	db, err := pgx.Connect(context.Background(), dataSource)
	return db, err
}
