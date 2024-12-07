package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/services/interfaces"
	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	c            *config.GConfig
	cm           *controllers.ControllerManager
	l            *log.ConsoleLogger
	albumService interfaces.IAlbumService
}

func NewAlbumController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
	albumService interfaces.IAlbumService,
) *AlbumController {
	c := &AlbumController{
		c:            cf,
		cm:           cm,
		l:            l,
		albumService: albumService,
	}
	c.RegisterController()
	return c
}

func (h *AlbumController) GetRoot() string {
	return "/album"
}

func (h *AlbumController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/category-list", Handle: h.categories},
		{Method: "GET", Path: "/:category/photos", Handle: h.categoryPhotos},
		{Method: "GET", Path: "/photo", Handle: h.photo},
	}
}

func (h *AlbumController) RegisterController() {
	h.cm.AddController(h)
}

func (h *AlbumController) categories(context *gin.Context) {
	context.JSON(200, common.NewSuccessResult(h.albumService.GetCategoryList()))
}

func (h *AlbumController) categoryPhotos(context *gin.Context) {
	category := context.Param("category")
	if category == "" {
		h.l.Error("category is empty")
		context.Error(common.BadRequestError)
		return
	}
	context.JSON(200, common.NewSuccessResult(h.albumService.GetCategoryPhotos(category)))
}

func (h *AlbumController) photo(context *gin.Context) {
	// 获取 URL 参数 id
	picName := context.Query("pic")
	if picName == "" {
		context.Error(common.BadRequestError)
		return
	}
	imageBytes, err := h.albumService.GetImageBytesByName(picName)
	if err != nil {
		context.Error(err)
		return
	}
	// 设置 Content-Type
	context.Header("Content-Type", "image/jpeg")
	// 返回图片文件
	context.Writer.Write(imageBytes)
}
