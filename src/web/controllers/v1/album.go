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
	"net/http"
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
		{Method: http.MethodGet, Path: "/category-list", Handle: h.categories},              // 获取所有的相册
		{Method: http.MethodPost, Path: "/category-add", Handle: h.categoryAdd},             // 添加相册
		{Method: http.MethodPut, Path: "/category-update", Handle: h.categoryUpdate},        // 添加相册
		{Method: http.MethodDelete, Path: "/category-delete/:id", Handle: h.categoryDelete}, // 删除相册
		{Method: http.MethodGet, Path: "/:category/photos", Handle: h.photosByCategory},     // 获取相册下的照片
		{Method: http.MethodGet, Path: "/photo", Handle: h.photoByPid},                      // 获取照片字节
		{Method: http.MethodGet, Path: "/fresh-photo", Handle: h.freshPhoto},                // 同步照片
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
		context.JSON(http.StatusOK, common.NewSuccessResultWithData(h.albumService.GetCategoryList(role)))
		return
	}
	context.JSON(http.StatusOK, common.AdminRoleError)
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
	context.JSON(http.StatusOK, common.NewSuccessResultWithData(h.albumService.GetCategoryPhotos(category, role)))
}

// photoByPid godoc
// @Summary      获取照片数据
// @Description  获取某个相册的照片字节数据
// @Tags         album
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        pid   query   string  true  "图片的ID"
// @Success      200  {object}  common.Result
// @Router       /album/photo [get]
func (h *AlbumController) photoByPid(context *gin.Context) {
	// 获取 URL 参数 id
	pid := context.Query("pid")
	if pid == "" {
		context.Error(common.BadRequestError)
		return
	}
	imageBytes, err := h.albumService.GetImageBytesByCategoryIdAndPid(pid)
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

// photoByPid godoc
// @Summary      同步照片
// @Description  将数据库的照片同步到本地数据库存储
// @Tags         album
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  common.Result
// @Router       /album/fresh-photo [get]
func (h *AlbumController) freshPhoto(context *gin.Context) {
	go func() {
		utils.ReadPathAllDir(h.c.Static.Dir+"/img/",
			h.albumService.SaveCategoryByCategoryName,
			h.albumService.SavePhotoByCategoryIdAndPhotoName)
	}()
	context.JSON(http.StatusOK, common.NewSuccessResult())
}

func (h *AlbumController) categoryAdd(context *gin.Context) {

}

func (h *AlbumController) categoryUpdate(context *gin.Context) {

}

func (h *AlbumController) categoryDelete(context *gin.Context) {

}
