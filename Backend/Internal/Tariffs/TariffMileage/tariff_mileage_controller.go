package tariffMileage

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITariffMileageService interface {
	SaveTariffMileage(TariffMileage) (*TariffMileage, error)
	GetTariffMileageById(uint32) (*TariffMileage, error)
	GetAllTariffMileage() ([]TariffMileage, error)
	UpdateTariffMileage(TariffMileage) error
	DeleteTariffMileageById(uint32) error
}

type Controller struct {
	Service ITariffMileageService
}

func (cntrl *Controller) postTariffMileage(c *gin.Context) {
	var tariffMileage TariffMileage
	if err := c.BindJSON(&tariffMileage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffmileage: - %v", err)
		return
	}
	newTariffMileage, err := cntrl.Service.SaveTariffMileage(tariffMileage)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving tariffmileage: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newTariffMileage)
}

func (cntrl *Controller) getTariffMileageById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	tariffMileage, err := cntrl.Service.GetTariffMileageById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting tariffmileage with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffMileage)
}

func (cntrl *Controller) getAllTariffMileage(c *gin.Context) {
	tariffMileages, err := cntrl.Service.GetAllTariffMileage()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all tariffMileage: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffMileages)
}

func (cntrl *Controller) updateTariffMileage(c *gin.Context) {
	var tariffMileage TariffMileage
	if err := c.BindJSON(&tariffMileage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffmileage: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateTariffMileage(tariffMileage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating tariffmileage: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffMileage)
}

func (cntrl *Controller) deleteTariffMileageById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	err = cntrl.Service.DeleteTariffMileageById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting tariffmileage with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/tariffs/mileage", cntrl.postTariffMileage)
	rout.GET("/tariffs/mileage/:id", cntrl.getTariffMileageById)
	rout.GET("/tariffs/mileage", cntrl.getAllTariffMileage)
	rout.PUT("/tariffs/mileage", cntrl.updateTariffMileage)
	rout.DELETE("/tariffs/mileage/:id", cntrl.deleteTariffMileageById)
}
