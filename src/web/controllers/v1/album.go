package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/eneity/login"
	"family-web-server/src/web/services/v1/interfaces"
	"family-web-server/src/web/utils"
	"github.com/gin-gonic/gin"
	"strings"
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
		{Method: "GET", Path: "/:category/photos", Handle: h.photosByCategory},
		{Method: "GET", Path: "/photo", Handle: h.photoByPid},
		{Method: "GET", Path: "/fresh-photo", Handle: h.freshPhoto},
	}
}

func (h *AlbumController) RegisterController() {
	h.cm.AddController(h)
}

// categories godoc
// @Summary      获取所有的相册分类
// @Description  获取所有的相册分类详情
// @Tags         album
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  common.Result
// @Router       /album/category-list [get]
func (h *AlbumController) categories(context *gin.Context) {
	value, exists := context.Get("role")
	if exists {
		role := value.(*login.Role)
		context.JSON(200, common.NewSuccessResultWithData(h.albumService.GetCategoryList(role)))
		return
	}
	context.JSON(200, common.AdminRoleError)
}

// photosByCategory godoc
// @Summary      相册分类照片获取
// @Description  通过相册分类id获取这个分类下面的所有照片信息
// @Tags         album
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category   path   string  true  "分类的id"
// @Success      200  {object}  common.Result
// @Router       /album/{category}/photos [get]
func (h *AlbumController) photosByCategory(context *gin.Context) {
	category := context.Param("category")
	var role *login.Role
	if value, exists := context.Get("role"); exists {
		role = value.(*login.Role)
	}
	if category == "" {
		h.l.Error("category is empty")
		context.Error(common.BadRequestError)
		return
	}
	context.JSON(200, common.NewSuccessResultWithData(h.albumService.GetCategoryPhotos(category, role)))
}

// photoByPid godoc
// @Summary      获取照片数据
// @Description  获取某个相册的照片字节数据
// @Tags         album
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        pid   query   string  true  "图片的ID"
// @Param        categoryId query   string  false "图片所属类别ID"
// @Success      200  {object}  common.Result
// @Router       /album/photo [get]
func (h *AlbumController) photoByPid(context *gin.Context) {
	// 获取 URL 参数 id
	categoryId := context.Query("categoryId")
	pid := context.Query("pid")
	if pid == "" {
		context.Error(common.BadRequestError)
		return
	}
	imageBytes, err := h.albumService.GetImageBytesByName(categoryId, pid)
	if err != nil {
		context.Error(err)
		return
	}
	// 设置 Content-Type
	contentType := "image/jpeg"
	if strings.HasSuffix(pid, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(pid, ".gif") {
		contentType = "image/gif"
	}
	context.Header("Content-Type", contentType)
	// 返回图片文件
	context.Writer.Write(imageBytes)
}

func (h *AlbumController) freshPhoto(context *gin.Context) {
	go func() {
		utils.ReadPathAllDir("./images",
			h.albumService.SaveCategoryByCategoryName,
			h.albumService.SavePhotoByCategoryIdAndPhotoName)
	}()
	context.JSON(200, common.NewSuccessResult())
}
