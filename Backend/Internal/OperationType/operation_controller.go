package operationType

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IOperationTypeService interface {
	SaveOperationType(OperationType) (*OperationType, error)
	GetOperationTypeById(int64) (*OperationType, error)
	GetAllOperationTypes() ([]OperationType, error)
	UpdateOperationType(OperationType) error
	DeleteOperationTypeById(int64) error
}

type Controller struct {
	Service IOperationTypeService
}

func (cntrl *Controller) postOperationType(c *gin.Context) {
	var operationType OperationType
	if err := c.BindJSON(&operationType); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to operationType: - %v", err)
		return
	}
	newOperationType, err := cntrl.Service.SaveOperationType(operationType)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving operationType: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newOperationType)
}

func (cntrl *Controller) getOperationTypeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	operationType, err := cntrl.Service.GetOperationTypeById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting operationType with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, operationType)
}

func (cntrl *Controller) getAllOperationTypes(c *gin.Context) {
	operationTypes, err := cntrl.Service.GetAllOperationTypes()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all operationTypes: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operationTypes)
}

func (cntrl *Controller) updateOperationType(c *gin.Context) {
	var operationType OperationType
	if err := c.BindJSON(&operationType); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to operationType: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateOperationType(operationType); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating operationType: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operationType)
}

func (cntrl *Controller) deleteOperationTypeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteOperationTypeById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting operationType with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/operationTypes", cntrl.postOperationType)
	rout.GET("/operationTypes/:id", cntrl.getOperationTypeById)
	rout.GET("/operationTypes", cntrl.getAllOperationTypes)
	rout.PUT("/operationTypes", cntrl.updateOperationType)
	rout.DELETE("/operationTypes/:id", cntrl.deleteOperationTypeById)
}
