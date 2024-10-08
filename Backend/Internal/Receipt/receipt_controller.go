package receipt

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IReceiptService interface {
	SaveReceipt(Receipt) (*Receipt, error)
	GetReceiptById(int64) (*Receipt, error)
	GetAllReceipts() ([]Receipt, error)
	UpdateReceipt(Receipt) error
	CalcTotalAmount(int64) (*Receipt, error)
	DeleteReceiptById(int64) error
}

type Controller struct {
	Service IReceiptService
}

func (cntrl *Controller) postReceipt(c *gin.Context) {
	var receipt Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to receipt: - %v", err)
		return
	}
	newReceipt, err := cntrl.Service.SaveReceipt(receipt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving receipt: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newReceipt)
}

func (cntrl *Controller) getReceiptById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	receipt, err := cntrl.Service.GetReceiptById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting receipt with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, receipt)
}

func (cntrl *Controller) getAllReceipts(c *gin.Context) {
	receipts, err := cntrl.Service.GetAllReceipts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all receipts: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, receipts)
}

func (cntrl *Controller) updateReceipt(c *gin.Context) {
	var receipt Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to receipt: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateReceipt(receipt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating receipt: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, receipt)
}

func (cntrl *Controller) calcTotalCostReceipt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}
	newReceipt, err := cntrl.Service.CalcTotalAmount(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed calculating total value receipt: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newReceipt)
}

func (cntrl *Controller) deleteReceiptById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteReceiptById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting receipt with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/receipts", cntrl.postReceipt)
	rout.GET("/receipts/:id", cntrl.getReceiptById)
	rout.GET("/receipts", cntrl.getAllReceipts)
	rout.PUT("/receipts", cntrl.updateReceipt)
	rout.PUT("/receipts/calculate/:id", cntrl.calcTotalCostReceipt)
	rout.DELETE("/receipts/:id", cntrl.deleteReceiptById)
}
