package receipt

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IReceiptService interface {
	SaveReceipt(Receipt) (*Receipt, error)
	GetReceiptById(uint32) (*Receipt, error)
	GetAllReceipts() ([]Receipt, error)
	UpdateReceipt(Receipt) error
	DeleteReceiptById(uint32) error
}

type ReceiptController struct {
	Service IReceiptService
}

func (cntrl *ReceiptController) postReceipt(c *gin.Context) {
	var receipt Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to receipt - [%v]", err)
		return
	}
	newReceipt, err := cntrl.Service.SaveReceipt(receipt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving receipt - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newReceipt)
}

func (cntrl *ReceiptController) getReceiptById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	receipt, err := cntrl.Service.GetReceiptById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting receipt with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, receipt)
}

func (cntrl *ReceiptController) getAllReceipts(c *gin.Context) {
	receipts, err := cntrl.Service.GetAllReceipts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all receipts - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, receipts)
}

func (cntrl *ReceiptController) updateReceipt(c *gin.Context) {
	var receipt Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to receipt - [%v]", err)
		return
	}
	if err := cntrl.Service.UpdateReceipt(receipt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating receipt - [%v]", err)
		return
	}
	c.IndentedJSON(http.StatusOK, receipt)
}

func (cntrl *ReceiptController) deleteReceiptById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to uint - [%v]", err)
		return
	}

	err = cntrl.Service.DeleteReceiptById(uint32(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting receipt with id %d - [%v]", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *ReceiptController) LinkPaths(rout *gin.Engine) {
	rout.POST("/receipts", cntrl.postReceipt)
	rout.GET("/receipts/:id", cntrl.getReceiptById)
	rout.GET("/receipts", cntrl.getAllReceipts)
	rout.PUT("/receipts", cntrl.updateReceipt)
	rout.DELETE("/receipts/:id", cntrl.deleteReceiptById)
}
