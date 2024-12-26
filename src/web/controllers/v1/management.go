package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"github.com/gin-gonic/gin"
)

type ManagementController struct {
	c  *config.GConfig
	cm *controllers.ControllerManager
	l  *log.ConsoleLogger
}

func NewManagementController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
) *ManagementController {
	c := &ManagementController{
		c:  cf,
		cm: cm,
		l:  l,
	}
	c.RegisterController()
	return c
}
func (h *ManagementController) GetRoot() string {
	return "/management"
}

func (h *ManagementController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/menus", Handle: h.menus},
	}
}

func (h *ManagementController) RegisterController() {
	h.cm.AddController(h)
}

// register godoc
// @Summary      获取管理页面的菜单
// @Description  根据当前的用户角色去获取菜单侧边栏
// @Tags         management
// @Produce      json
// @Success      200  {object}  common.Result
// @Router       /management/menus [get]
func (h *ManagementController) menus(context *gin.Context) {
	context.JSON(200, common.NewSuccessResultWithData(nil))
}
