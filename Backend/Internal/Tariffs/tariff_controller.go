package tariffs

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITariffService interface {
	SaveTariff(Tariff, string) (*Tariff, error)
	GetTariffById(uint32, string) (*Tariff, error)
	GetAllTariffs(string) ([]Tariff, error)
	UpdateTariff(Tariff, string) error
	DeleteTariffById(uint32, string) error
}

type TariffController struct {
	Service ITariffService
}

func (cntrl *TariffController) postTariff(c *gin.Context) {
	var tariff Tariff
	if err := c.BindJSON(&tariff); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to tariff - [%v]", err)
		return
	}
	newTariff, err := cntrl.Service.SaveTariff(tariff, c.Param("table"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving tariff - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newTariff)
}

func (cntrl *TariffController) getTariffById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	tariff, err := cntrl.Service.GetTariffById(uint32(id), c.Param("table"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting tariff with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariff)
}

func (cntrl *TariffController) getAllTariffs(c *gin.Context) {
	tariffs, err := cntrl.Service.GetAllTariffs(c.Param("table"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all tariffs - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariffs)
}

func (cntrl *TariffController) updateTariff(c *gin.Context) {
	var tariff Tariff
	if err := c.BindJSON(&tariff); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to tariff - [%v]", err)
		return
	}
	if err := cntrl.Service.UpdateTariff(tariff, c.Param("table")); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating tariff - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, tariff)
}

func (cntrl *TariffController) deleteTariffById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	err = cntrl.Service.DeleteTariffById(uint32(id), c.Param("table"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting tariff with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *TariffController) LinkPaths(rout *gin.Engine) {
	rout.POST("/tariffs/:table", cntrl.postTariff)
	rout.GET("/tariffs/:table/:id", cntrl.getTariffById)
	rout.GET("/tariffs/:table", cntrl.getAllTariffs)
	rout.PUT("/tariffs/:table", cntrl.updateTariff)
	rout.DELETE("/tariffs/:table/:id", cntrl.deleteTariffById)
}
