package bonus

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IBonusService interface {
	SaveBonus(Bonus) (*Bonus, error)
	GetBonusById(int64) (*Bonus, error)
	GetAllBonuses() ([]Bonus, error)
	UpdateBonus(Bonus) error
	DeleteBonusById(int64) error
}

type Controller struct {
	Service IBonusService
}

func (cntrl *Controller) postBonus(c *gin.Context) {
	var bonus Bonus
	if err := c.BindJSON(&bonus); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to bonus: - %v", err)
		return
	}
	newBonus, err := cntrl.Service.SaveBonus(bonus)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed saving bonus: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, newBonus)
}

func (cntrl *Controller) getBonusById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	bonus, err := cntrl.Service.GetBonusById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting bonus with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, bonus)
}

func (cntrl *Controller) getAllBonuses(c *gin.Context) {
	bonuss, err := cntrl.Service.GetAllBonuses()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed getting all bonuses: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, bonuss)
}

func (cntrl *Controller) updateBonus(c *gin.Context) {
	var bonus Bonus
	if err := c.BindJSON(&bonus); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing JSON to bonus: - %v", err)
		return
	}
	if err := cntrl.Service.UpdateBonus(bonus); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed updating bonus: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, bonus)
}

func (cntrl *Controller) deleteBonusById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed parsing id to int: - %v", err)
		return
	}

	err = cntrl.Service.DeleteBonusById(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("failed deleting bonus with id %d: - %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/bonuses", cntrl.postBonus)
	rout.GET("/bonuses/:id", cntrl.getBonusById)
	rout.GET("/bonuses", cntrl.getAllBonuses)
	rout.PUT("/bonuses", cntrl.updateBonus)
	rout.DELETE("/bonuses/:id", cntrl.deleteBonusById)
}
