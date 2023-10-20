package patent

import (
	"patent/controller"
	"patent/service/patent"
	err "patent/util"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// @Summary 查询车辆调度列表, Action: ParsingPdf
// @Description 查询车辆调度列表
// @Tags login (URL: {host}/api/login)
// @Accept  json
// @Produce json
// @Router /ParsingPdf [post]
// @Param Action  		 	 		query	  string				true    "ParsingPdf"
// @Param file  		 	 		body	  multipart.File		true    "patent pdf"
func (ctrl *Controller) ParsingPdf(c *gin.Context) {
	_, header, e := c.Request.FormFile("file")
	if e != nil {
		controller.SendResponse(c, err.InvalidParam, nil)
		return
	}
	s := patent.NewSystemService("")
	resp, er := s.ParsingPdf(header)
	if e != nil {
		controller.SendResponse(c, er, nil)
		return
	}
	controller.SendResponse(c, nil, resp)
}
