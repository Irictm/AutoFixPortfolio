package operation

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IOperationService interface {
	SaveOperation(Operation) (*Operation, error)
	GetOperationById(int64) (*Operation, error)
	GetAllOperations() ([]Operation, error)
	UpdateOperation(Operation) error
	DeleteOperationById(int64) error
}

type Controller struct {
	Service IOperationService
}

func (cntrl *Controller) postOperation(c *gin.Context) {
	var operation Operation
	if err := c.BindJSON(&operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to operation: - %v", err)
		return
	}
	newOperation, err := cntrl.Service.SaveOperation(operation)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving operation: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newOperation)
}

func (cntrl *Controller) getOperationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	operation, err := cntrl.Service.GetOperationById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting operation with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, operation)
}

func (cntrl *Controller) getAllOperations(c *gin.Context) {
	operations, err := cntrl.Service.GetAllOperations()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all operations: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operations)
}

func (cntrl *Controller) updateOperation(c *gin.Context) {
	var operation Operation
	if err := c.BindJSON(&operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to operation: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateOperation(operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating operation: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operation)
}

func (cntrl *Controller) deleteOperationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteOperationById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting operation with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/operations", cntrl.postOperation)
	rout.GET("/operations/:id", cntrl.getOperationById)
	rout.GET("/operations", cntrl.getAllOperations)
	rout.PUT("/operations", cntrl.updateOperation)
	rout.DELETE("/operations/:id", cntrl.deleteOperationById)
}
