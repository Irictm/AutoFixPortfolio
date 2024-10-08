package tariffAntiquity

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITariffAntiquityService interface {
	SaveAndParseTariffAntiquity(data map[string]interface{}) (map[string]interface{}, error)
	ReceiveTariffAntiquityCSV([][]string) error
	GetTariffAntiquityById(int64) (*TariffAntiquity, error)
	GetAllTariffAntiquity() ([]TariffAntiquity, error)
	UpdateTariffAntiquity(TariffAntiquity) error
	DeleteTariffAntiquityById(int64) error
}

type ICSVHandler interface {
	AttachCSV(*gin.Context) error
	ReceiveCSV(*gin.Context) ([][]string, error)
}

type Controller struct {
	Service    ITariffAntiquityService
	CsvHandler ICSVHandler
}

func (cntrl *Controller) postTariffAntiquity(c *gin.Context) {
	var err error
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed reading body data: - %v", err)
		return
	}
	bodyJson := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &bodyJson)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed marshaling body data: - %v", err)
		return
	}
	log.Println(bodyJson)

	tariffAntiquityMap, err := cntrl.Service.SaveAndParseTariffAntiquity(bodyJson)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing and saving tariff antiquity: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, tariffAntiquityMap)
}

func (cntrl *Controller) postTariffAntiquityCSV(c *gin.Context) {
	table, err := cntrl.CsvHandler.ReceiveCSV(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving tariffantiquity: - %v", err)
		return
	}
	err = cntrl.Service.ReceiveTariffAntiquityCSV(table)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving CSV tariff data: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, "CSV received and processed")
}

func (cntrl *Controller) getTariffAntiquityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	tariffAntiquity, err := cntrl.Service.GetTariffAntiquityById(int64(id))
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
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteTariffAntiquityById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting tariffantiquity with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/tariffs/antiquity", cntrl.postTariffAntiquity)
	rout.POST("/tariffs/antiquity/csv", cntrl.postTariffAntiquityCSV)
	rout.GET("/tariffs/antiquity/:id", cntrl.getTariffAntiquityById)
	rout.GET("/tariffs/antiquity", cntrl.getAllTariffAntiquity)
	rout.PUT("/tariffs/antiquity", cntrl.updateTariffAntiquity)
	rout.DELETE("/tariffs/antiquity/:id", cntrl.deleteTariffAntiquityById)
}
