package tariffRepairNumber

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITariffRepairNumberService interface {
	SaveTariffRepairNumber(TariffRepairNumber) (*TariffRepairNumber, error)
	GetTariffRepairNumberById(int64) (*TariffRepairNumber, error)
	GetAllTariffRepairNumber() ([]TariffRepairNumber, error)
	UpdateTariffRepairNumber(TariffRepairNumber) error
	DeleteTariffRepairNumberById(int64) error
}

type Controller struct {
	Service ITariffRepairNumberService
}

func (cntrl *Controller) postTariffRepairNumber(c *gin.Context) {
	var tariffRepairNumber TariffRepairNumber
	if err := c.BindJSON(&tariffRepairNumber); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffrepairnumber: - %v", err)
		return
	}
	newTariffRepairNumber, err := cntrl.Service.SaveTariffRepairNumber(tariffRepairNumber)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving tariffrepairnumber: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newTariffRepairNumber)
}

func (cntrl *Controller) getTariffRepairNumberById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	tariffRepairNumber, err := cntrl.Service.GetTariffRepairNumberById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting tariffrepairnumber with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffRepairNumber)
}

func (cntrl *Controller) getAllTariffRepairNumber(c *gin.Context) {
	tariffRepairNumbers, err := cntrl.Service.GetAllTariffRepairNumber()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all tariffRepairNumber: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffRepairNumbers)
}

func (cntrl *Controller) updateTariffRepairNumber(c *gin.Context) {
	var tariffRepairNumber TariffRepairNumber
	if err := c.BindJSON(&tariffRepairNumber); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffrepairnumber: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateTariffRepairNumber(tariffRepairNumber); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating tariffrepairnumber: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffRepairNumber)
}

func (cntrl *Controller) deleteTariffRepairNumberById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteTariffRepairNumberById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting tariffrepairnumber with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/tariffs/repairNumber", cntrl.postTariffRepairNumber)
	rout.GET("/tariffs/repairNumber/:id", cntrl.getTariffRepairNumberById)
	rout.GET("/tariffs/repairNumber", cntrl.getAllTariffRepairNumber)
	rout.PUT("/tariffs/repairNumber", cntrl.updateTariffRepairNumber)
	rout.DELETE("/tariffs/repairNumber/:id", cntrl.deleteTariffRepairNumberById)
}
