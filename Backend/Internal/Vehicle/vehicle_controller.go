package vehicle

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IVehicleService interface {
	SaveVehicle(Vehicle) (*Vehicle, error)
	GetVehicleById(uint32) (*Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	UpdateVehicle(Vehicle) error
	DeleteVehicleById(uint32) error
}

type Controller struct {
	Service IVehicleService
}

func (cntrl *Controller) postVehicle(c *gin.Context) {
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to vehicle: - %v", err)
		return
	}
	newVehicle, err := cntrl.Service.SaveVehicle(vehicle)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving vehicle: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newVehicle)
}

func (cntrl *Controller) getVehicleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	vehicle, err := cntrl.Service.GetVehicleById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting vehicle with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func (cntrl *Controller) getAllVehicles(c *gin.Context) {
	vehicles, err := cntrl.Service.GetAllVehicles()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all vehicles: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicles)
}

func (cntrl *Controller) updateVehicle(c *gin.Context) {
	var vehicle Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to vehicle: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateVehicle(vehicle); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating vehicle: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func (cntrl *Controller) deleteVehicleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	err = cntrl.Service.DeleteVehicleById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting vehicle with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/vehicles", cntrl.postVehicle)
	rout.GET("/vehicles/:id", cntrl.getVehicleById)
	rout.GET("/vehicles", cntrl.getAllVehicles)
	rout.PUT("/vehicles", cntrl.updateVehicle)
	rout.DELETE("/vehicles/:id", cntrl.deleteVehicleById)
}
