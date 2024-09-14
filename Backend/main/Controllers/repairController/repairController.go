package repairController

import (
	"log"
	"net/http"
	"strconv"

	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/repair"
	"github.com/gin-gonic/gin"
)

type IRepairService interface {
	SaveRepair(Repair) error
	GetRepairById(uint32) (*Repair, error)
	GetAllRepairs() ([]Repair, error)
	UpdateRepair(Repair) error
	DeleteRepairById(uint32) error
}

type RepairController struct {
	Service IRepairService
}

func (cntrl *RepairController) postRepair(c *gin.Context) {
	var repair Repair
	if err := c.BindJSON(&repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to repair - [%v]", err)
		return
	}
	if err := cntrl.Service.SaveRepair(repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving repair - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, repair)
}

func (cntrl *RepairController) getRepairById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	repair, err := cntrl.Service.GetRepairById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting repair with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, repair)
}

func (cntrl *RepairController) getAllRepairs(c *gin.Context) {
	repairs, err := cntrl.Service.GetAllRepairs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all repairs - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, repairs)
}

func (cntrl *RepairController) updateRepair(c *gin.Context) {
	var repair Repair
	if err := c.BindJSON(&repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to repair - [%v]", err)
		return
	}
	if err := cntrl.Service.UpdateRepair(repair); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating repair - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, repair)
}

func (cntrl *RepairController) deleteRepairById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	err = cntrl.Service.DeleteRepairById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting repair with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *RepairController) LinkPaths(rout *gin.Engine) {
	rout.POST("/repairs", cntrl.postRepair)
	rout.GET("/repairs/:id", cntrl.getRepairById)
	rout.GET("/repairs", cntrl.getAllRepairs)
	rout.PUT("/repairs", cntrl.updateRepair)
	rout.DELETE("/repairs/:id", cntrl.deleteRepairById)
}
