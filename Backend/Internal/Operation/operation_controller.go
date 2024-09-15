package operation

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IOperationService interface {
	SaveOperation(Operation) error
	GetOperationById(uint32) (*Operation, error)
	GetAllOperations() ([]Operation, error)
	UpdateOperation(Operation) error
	DeleteOperationById(uint32) error
}

type OperationController struct {
	Service IOperationService
}

func (cntrl *OperationController) postOperation(c *gin.Context) {
	var operation Operation
	if err := c.BindJSON(&operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to operation - [%v]", err)
		return
	}
	if err := cntrl.Service.SaveOperation(operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving operation - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operation)
}

func (cntrl *OperationController) getOperationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	operation, err := cntrl.Service.GetOperationById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting operation with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, operation)
}

func (cntrl *OperationController) getAllOperations(c *gin.Context) {
	operations, err := cntrl.Service.GetAllOperations()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all operations - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operations)
}

func (cntrl *OperationController) updateOperation(c *gin.Context) {
	var operation Operation
	if err := c.BindJSON(&operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to operation - [%v]", err)
		return
	}
	if err := cntrl.Service.UpdateOperation(operation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating operation - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, operation)
}

func (cntrl *OperationController) deleteOperationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	err = cntrl.Service.DeleteOperationById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting operation with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *OperationController) LinkPaths(rout *gin.Engine) {
	rout.POST("/operations", cntrl.postOperation)
	rout.GET("/operations/:id", cntrl.getOperationById)
	rout.GET("/operations", cntrl.getAllOperations)
	rout.PUT("/operations", cntrl.updateOperation)
	rout.DELETE("/operations/:id", cntrl.deleteOperationById)
}
