package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/services/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	c           *config.GConfig
	cm          *controllers.ControllerManager
	l           *log.ConsoleLogger
	homeService interfaces.IHomeService
}

func NewAlbumController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
) *AlbumController {
	c := &AlbumController{
		c:  cf,
		cm: cm,
		l:  l,
	}
	c.RegisterController()
	return c
}

func (h *AlbumController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/album/photo", Handle: h.photos},
	}
}

func (h *AlbumController) RegisterController() {
	h.cm.AddController(h)
}

func (h *AlbumController) photos(context *gin.Context) {
	// 获取 URL 参数 id
	pic := context.Query("pic")
	if pic == "" {
		context.Error(common.BadRequestError)
		return
	}

	// 假设根目录下有 src/static/img 目录存储图片
	imagePath := fmt.Sprintf("./src/static/img/%s", pic)
	// 设置 Content-Type
	context.Header("Content-Type", "image/jpeg")
	// 返回图片文件
	context.File(imagePath)
}
