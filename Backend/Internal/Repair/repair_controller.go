package repair

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IRepairService interface {
	SaveRepair(Repair) (*Repair, error)
	GetRepairById(int64) (*Repair, error)
	GetAllRepairs() ([]Repair, error)
	UpdateRepair(Repair) error
	DeleteRepairById(int64) error
}

type Controller struct {
	Service IRepairService
}

func (cntrl *Controller) postRepair(c *gin.Context) {
	var repair Repair
	if err := c.BindJSON(&repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to repair: - %v", err)
		return
	}
	newRepair, err := cntrl.Service.SaveRepair(repair)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving repair: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newRepair)
}

func (cntrl *Controller) getRepairById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	repair, err := cntrl.Service.GetRepairById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting repair with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, repair)
}

func (cntrl *Controller) getAllRepairs(c *gin.Context) {
	repairs, err := cntrl.Service.GetAllRepairs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all repairs: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, repairs)
}

func (cntrl *Controller) updateRepair(c *gin.Context) {
	var repair Repair
	if err := c.BindJSON(&repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to repair: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateRepair(repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating repair: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, repair)
}

func (cntrl *Controller) deleteRepairById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteRepairById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting repair with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/repairs", cntrl.postRepair)
	rout.GET("/repairs/:id", cntrl.getRepairById)
	rout.GET("/repairs", cntrl.getAllRepairs)
	rout.PUT("/repairs", cntrl.updateRepair)
	rout.DELETE("/repairs/:id", cntrl.deleteRepairById)
}
