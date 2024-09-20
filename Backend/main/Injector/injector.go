package injector

import (
	"context"
	"log"

	bonus "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Bonus"
	operation "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Operation"
	receipt "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Receipt"
	repair "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Repair"
	vehicle "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Vehicle"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func InjectDependencies(rout *gin.Engine) {
	db, err := ConnectPostgreSQL("postgres", "postgres", "localhost", "5432", "autofix")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	vehicleRepository := &vehicle.VehicleRepository{DB: db}
	vehicleService := &vehicle.VehicleService{Repository: vehicleRepository}
	vehicleController := vehicle.VehicleController{Service: vehicleService}
	vehicleController.LinkPaths(rout)

	operationRepository := &operation.OperationRepository{DB: db}
	operationService := &operation.OperationService{Repository: operationRepository}
	operationController := operation.OperationController{Service: operationService}
	operationController.LinkPaths(rout)

	repairRepository := &repair.RepairRepository{DB: db}
	repairService := &repair.RepairService{Repository: repairRepository}
	repairController := repair.RepairController{Service: repairService}
	repairController.LinkPaths(rout)

	receiptRepository := &receipt.ReceiptRepository{DB: db}
	receiptService := &receipt.ReceiptService{Repository: receiptRepository}
	receiptController := receipt.ReceiptController{Service: receiptService}
	receiptController.LinkPaths(rout)

	bonusRepository := &bonus.BonusRepository{DB: db}
	bonusService := &bonus.BonusService{Repository: bonusRepository}
	bonusController := bonus.BonusController{Service: bonusService}
	bonusController.LinkPaths(rout)

}

func ConnectPostgreSQL(user string, pass string, host string, port string, dbname string) (*pgx.Conn, error) {
	dataSource := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/autofix?sslmode=disable"

	db, err := pgx.Connect(context.Background(), dataSource)
	return db, err
}
