package patent

import (
	"net/http"
	"patent/controller/patent"

	"github.com/gin-gonic/gin"
)

var actionHandler = make(map[string]func(c *gin.Context))

func init() {
	initActionHandler()
}

func Load(rg *gin.RouterGroup) {

	rg.GET("/CarSchedule", func(c *gin.Context) {
		executeHttpHandlerFunc(c)
	})
	rg.POST("/CarSchedule", func(c *gin.Context) {
		executeHttpHandlerFunc(c)
	})
}
func executeHttpHandlerFunc(c *gin.Context) {
	action := c.Query("Action")
	if action == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Params {Action, systemKey} are required", "data": nil})
		return
	}

	handleFunc := actionHandler[action]
	if handleFunc == nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "No controller for action " + action, "data": nil})
		return
	}
	handleFunc(c)
}

func initActionHandler() {
	ctrl := &patent.Controller{}
	actionHandler["ParsingPdf"] = ctrl.ParsingPdf
}
