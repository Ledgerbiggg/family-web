package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/eneity/login"
	"family-web-server/src/web/services/interfaces"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	c           *config.GConfig
	cm          *controllers.ControllerManager
	l           *log.ConsoleLogger
	homeService interfaces.IHomeService
}

func NewHomeController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
	ls interfaces.IHomeService,
) *HomeController {
	c := &HomeController{
		c:           cf,
		cm:          cm,
		l:           l,
		homeService: ls,
	}
	c.RegisterController()
	return c
}

func (h *HomeController) GetRoot() string {
	return "/home"
}

func (h *HomeController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/cards", Handle: h.cards},
	}
}

func (h *HomeController) RegisterController() {
	h.cm.AddController(h)
}

func (h *HomeController) cards(context *gin.Context) {
	value, exists := context.Get("role")
	if exists {
		role := value.(*login.Role)
		context.JSON(200, common.NewSuccessResult(h.homeService.GetHomeCardData(role)))
	} else {
		context.JSON(200, common.AdminRoleError)
	}
}
