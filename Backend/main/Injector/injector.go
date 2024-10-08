package injector

import (
	"context"
	"log"

	bonus "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Bonus"
	operation "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Operation"
	receipt "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Receipt"
	repair "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Repair"
	tariff "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Tariffs"
	tariffAntiquity "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Tariffs/TariffAntiquity"
	tariffMileage "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Tariffs/TariffMileage"
	tariffOperations "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Tariffs/TariffOperations"
	tariffRepairNumber "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Tariffs/TariffRepairNumber"
	vehicle "github.com/Irictm/AutoFixPortfolio/Backend/Internal/Vehicle"
	csvHandler "github.com/Irictm/AutoFixPortfolio/Backend/Utils/CSVHandler"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func InjectDependencies(rout *gin.Engine) {
	db, err := ConnectPostgreSQL("postgres", "postgres", "localhost", "5432", "autofix")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	csvHandler := csvHandler.CsvHandler{}

	vehicleRepository := &vehicle.Repository{DB: db}
	vehicleService := &vehicle.Service{Repository: vehicleRepository}
	vehicleController := vehicle.Controller{Service: vehicleService}
	vehicleController.LinkPaths(rout)

	tariffAntiquityRepository := tariffAntiquity.Repository{DB: db}
	tariffAntiquityService := tariffAntiquity.Service{Repository: &tariffAntiquityRepository}
	tariffAntiquityController := tariffAntiquity.Controller{Service: &tariffAntiquityService, CsvHandler: &csvHandler}
	tariffAntiquityController.LinkPaths(rout)

	tariffMileageRepository := tariffMileage.Repository{DB: db}
	tariffMileageService := tariffMileage.Service{Repository: &tariffMileageRepository}
	tariffMileageController := tariffMileage.Controller{Service: &tariffMileageService}
	tariffMileageController.LinkPaths(rout)

	tariffOperationsRepository := tariffOperations.Repository{DB: db}
	tariffOperationsService := tariffOperations.Service{Repository: &tariffOperationsRepository}
	tariffOperationsController := tariffOperations.Controller{Service: &tariffOperationsService}
	tariffOperationsController.LinkPaths(rout)

	tariffRepairNumberRepository := tariffRepairNumber.Repository{DB: db}
	tariffRepairNumberService := tariffRepairNumber.Service{Repository: &tariffRepairNumberRepository}
	tariffRepairNumberController := tariffRepairNumber.Controller{Service: &tariffRepairNumberService}
	tariffRepairNumberController.LinkPaths(rout)

	tariffService := &tariff.TariffService{TariffAntiquity: &tariffAntiquityService, TariffMileage: &tariffMileageService,
		TariffOperations: &tariffOperationsService, TariffRepairNumber: &tariffRepairNumberService}

	operationRepository := &operation.Repository{DB: db}
	operationService := &operation.Service{Repository: operationRepository, TarService: tariffService}
	operationController := operation.Controller{Service: operationService}
	operationController.LinkPaths(rout)

	repairRepository := &repair.Repository{DB: db}
	repairService := &repair.Service{Repository: repairRepository}
	repairController := repair.Controller{Service: repairService}
	repairController.LinkPaths(rout)

	bonusRepository := &bonus.Repository{DB: db}
	bonusService := &bonus.Service{Repository: bonusRepository}
	bonusController := bonus.Controller{Service: bonusService}
	bonusController.LinkPaths(rout)

	receiptRepository := &receipt.Repository{DB: db}
	receiptService := &receipt.Service{Repository: receiptRepository, BonusService: bonusService,
		RepairService: repairService, OperationService: operationService, TariffService: tariffService,
		VehicleService: vehicleService}
	receiptController := receipt.Controller{Service: receiptService}
	receiptController.LinkPaths(rout)
}

func ConnectPostgreSQL(user string, pass string, host string, port string, dbname string) (*pgx.Conn, error) {
	dataSource := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/autofix?sslmode=disable"

	db, err := pgx.Connect(context.Background(), dataSource)
	return db, err
}
