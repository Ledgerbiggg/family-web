package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/services/interfaces"
	"github.com/gin-gonic/gin"
)

type TempController struct {
	c           *config.GConfig
	cm          *controllers.ControllerManager
	l           *log.ConsoleLogger
	homeService interfaces.IHomeService
}

func NewTempController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
	ls interfaces.IHomeService,
) *TempController {
	c := &TempController{
		c:           cf,
		cm:          cm,
		l:           l,
		homeService: ls,
	}
	c.RegisterController()
	return c
}

func (h *TempController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/temp", Handle: h.cards},
	}
}

func (h *TempController) RegisterController() {
	h.cm.AddController(h)
}

func (h *TempController) cards(context *gin.Context) {

}
