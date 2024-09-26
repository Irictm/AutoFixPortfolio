package tariffOperations

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITariffOperationsService interface {
	SaveTariffOperations(TariffOperations) (*TariffOperations, error)
	GetTariffOperationsById(uint32) (*TariffOperations, error)
	GetAllTariffOperations() ([]TariffOperations, error)
	UpdateTariffOperations(TariffOperations) error
	DeleteTariffOperationsById(uint32) error
}

type Controller struct {
	Service ITariffOperationsService
}

func (cntrl *Controller) postTariffOperations(c *gin.Context) {
	var tariffOperations TariffOperations
	if err := c.BindJSON(&tariffOperations); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffoperations: - %v", err)
		return
	}
	newTariffOperations, err := cntrl.Service.SaveTariffOperations(tariffOperations)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving tariffoperations: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newTariffOperations)
}

func (cntrl *Controller) getTariffOperationsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	tariffOperations, err := cntrl.Service.GetTariffOperationsById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting tariffoperations with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffOperations)
}

func (cntrl *Controller) getAllTariffOperations(c *gin.Context) {
	tariffOperations, err := cntrl.Service.GetAllTariffOperations()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all tariffoperationss: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffOperations)
}

func (cntrl *Controller) updateTariffOperations(c *gin.Context) {
	var tariffOperations TariffOperations
	if err := c.BindJSON(&tariffOperations); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to tariffoperations: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateTariffOperations(tariffOperations); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating tariffoperations: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffOperations)
}

func (cntrl *Controller) deleteTariffOperationsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to uint: - %v", err)
		return
	}

	err = cntrl.Service.DeleteTariffOperationsById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting tariffoperations with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/tariffs/operations", cntrl.postTariffOperations)
	rout.GET("/tariffs/operations/:id", cntrl.getTariffOperationsById)
	rout.GET("/tariffs/operations", cntrl.getAllTariffOperations)
	rout.PUT("/tariffs/operations", cntrl.updateTariffOperations)
	rout.DELETE("/tariffs/operations/:id", cntrl.deleteTariffOperationsById)
}
