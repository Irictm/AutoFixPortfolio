package tariffAntiquity

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITariffAntiquityService interface {
	SaveTariffAntiquity(TariffAntiquity) (*TariffAntiquity, error)
	GetTariffAntiquityById(uint32) (*TariffAntiquity, error)
	GetAllTariffAntiquity() ([]TariffAntiquity, error)
	UpdateTariffAntiquity(TariffAntiquity) error
	DeleteTariffAntiquityById(uint32) error
}

type Controller struct {
	Service ITariffAntiquityService
}

func (cntrl *Controller) postTariffAntiquity(c *gin.Context) {
	var tariffAntiquity TariffAntiquity
	if err := c.BindJSON(&tariffAntiquity); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffantiquity: - %v", err)
		return
	}
	newTariffAntiquity, err := cntrl.Service.SaveTariffAntiquity(tariffAntiquity)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving tariffantiquity: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newTariffAntiquity)
}

func (cntrl *Controller) getTariffAntiquityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	tariffAntiquity, err := cntrl.Service.GetTariffAntiquityById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting tariffantiquity with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffAntiquity)
}

func (cntrl *Controller) getAllTariffAntiquity(c *gin.Context) {
	tariffAntiquities, err := cntrl.Service.GetAllTariffAntiquity()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all tariffAntiquity: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffAntiquities)
}

func (cntrl *Controller) updateTariffAntiquity(c *gin.Context) {
	var tariffAntiquity TariffAntiquity
	if err := c.BindJSON(&tariffAntiquity); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffantiquity: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateTariffAntiquity(tariffAntiquity); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating tariffantiquity: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffAntiquity)
}

func (cntrl *Controller) deleteTariffAntiquityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	err = cntrl.Service.DeleteTariffAntiquityById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting tariffantiquity with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/tariffs/antiquity", cntrl.postTariffAntiquity)
	rout.GET("/tariffs/antiquity/:id", cntrl.getTariffAntiquityById)
	rout.GET("/tariffs/antiquity", cntrl.getAllTariffAntiquity)
	rout.PUT("/tariffs/antiquity", cntrl.updateTariffAntiquity)
	rout.DELETE("/tariffs/antiquity/:id", cntrl.deleteTariffAntiquityById)
}
