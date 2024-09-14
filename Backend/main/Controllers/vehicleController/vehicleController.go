package vehicleController

import (
	"log"
	"net/http"
	"strconv"

	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/vehicle"
	"github.com/gin-gonic/gin"
)

type IVehicleService interface {
	SaveVehicle(Vehicle) error
	GetVehicleById(uint32) (*Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	UpdateVehicle(Vehicle) error
	DeleteVehicleById(uint32) error
}

type VehicleController struct {
	Service IVehicleService
}

func (cntrl *VehicleController) postVehicle(c *gin.Context) {
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to vehicle - [%v]", err)
		return
	}
	if err := cntrl.Service.SaveVehicle(vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving vehicle - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func (cntrl *VehicleController) getVehicleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	vehicle, err := cntrl.Service.GetVehicleById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting vehicle with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func (cntrl *VehicleController) getAllVehicles(c *gin.Context) {
	vehicles, err := cntrl.Service.GetAllVehicles()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all vehicles - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicles)
}

func (cntrl *VehicleController) updateVehicle(c *gin.Context) {
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to vehicle - [%v]", err)
		return
	}
	if err := cntrl.Service.UpdateVehicle(vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating vehicle - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func (cntrl *VehicleController) deleteVehicleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	err = cntrl.Service.DeleteVehicleById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting vehicle with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *VehicleController) LinkPaths(rout *gin.Engine) {
	rout.POST("/vehicles", cntrl.postVehicle)
	rout.GET("/vehicles/:id", cntrl.getVehicleById)
	rout.GET("/vehicles", cntrl.getAllVehicles)
	rout.PUT("/vehicles", cntrl.updateVehicle)
	rout.DELETE("/vehicles/:id", cntrl.deleteVehicleById)
}
