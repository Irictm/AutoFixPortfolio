package vehicleController

import (
	"net/http"

	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/vehicle"
	"github.com/Irictm/AutoFixPortfolio/Backend/main/Services/vehicleService"
	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	Service vehicleService.VehicleService
}

func (cntrl *VehicleController) getAllVehicles(c *gin.Context) {
	vehicles, err := cntrl.Service.GetAllVehicles()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicles)
}

func (cntrl *VehicleController) postVehicle(c *gin.Context) {
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if err := cntrl.Service.SaveVehicle(vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func (cntrl *VehicleController) LinkPaths(rout *gin.Engine) {
	rout.POST("/vehicle", cntrl.postVehicle)
	rout.GET("/vehicle", cntrl.getAllVehicles)
}
