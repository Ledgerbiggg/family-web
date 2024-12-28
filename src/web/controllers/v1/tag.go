package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/dto/tag"
	"family-web-server/src/web/services/v1/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagController struct {
	c  *config.GConfig
	cm *controllers.ControllerManager
	l  *log.ConsoleLogger
	s  interfaces.ITagService
}

func NewTagController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
	s interfaces.ITagService,
) *TagController {
	c := &TagController{
		c:  cf,
		cm: cm,
		l:  l,
		s:  s,
	}
	c.RegisterController()
	return c
}
func (h *TagController) GetRoot() string {
	return "/tags"
}

func (h *TagController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: http.MethodGet, Path: "/:type", Handle: h.getTagsByType},     // 获取指定类型的标签
		{Method: http.MethodPost, Path: "/add", Handle: h.addTag},             // 添加标签
		{Method: http.MethodPut, Path: "/update", Handle: h.updateTag},        // 更新标签
		{Method: http.MethodDelete, Path: "/delete/:id", Handle: h.deleteTag}, // 删除标签
	}
}

func (h *TagController) RegisterController() {
	h.cm.AddController(h)
}

// getTagsByType godoc
// @Summary      获取指定类型的标签
// @Description  根据 URL 中的 type 参数获取指定类型的标签
// @Tags         tag
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        type   path      int     true  "标签类型,1:用户标签,2:相册标签,3:照片标签"
// @Success      200    {object}  common.Result
// @Router       /tags/{type} [get]
func (h *TagController) getTagsByType(context *gin.Context) {
	// 获取参数并尝试转换为 int
	tagTypeStr := context.Param("type")
	tagType, err := strconv.Atoi(tagTypeStr)
	if err != nil {
		context.Error(err)
		return
	}

	// 调用服务方法并传入转换后的 int 类型
	tags, err := h.s.GetTagsByType(tagType)
	if err != nil {
		h.l.Error("获取标签失败:" + err.Error())
		context.JSON(http.StatusOK, err)
		return
	}
	context.JSON(http.StatusOK, common.NewSuccessResultWithData(tags))
}

// addTag godoc
// @Summary      添加标签
// @Description  添加新的标签
// @Tags         tag
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  tag.TagDto  true  "标签信息"
// @Success      200    {object}  common.Result
// @Router       /tags/add [post]
func (h *TagController) addTag(context *gin.Context) {
	var v = &tag.TagDto{}
	if err := context.ShouldBindJSON(v); err != nil {
		h.l.Error("参数绑定失败:" + err.Error())
		context.Error(err)
		return
	}
	if err := h.s.AddTag(v); err != nil {
		h.l.Error("添加标签失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, common.NewSuccessResultWithData(nil))
}

// updateTag godoc
// @Summary      更新标签
// @Description  更新现有的标签
// @Tags         tag
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  tag.TagDto  true  "标签信息"
// @Success      200    {object}  common.Result
// @Router       /tags/update [put]
func (h *TagController) updateTag(context *gin.Context) {
	// 获取参数并尝试绑定
	tagDto := &tag.TagDto{}
	err := context.ShouldBindJSON(tagDto)
	if err != nil {
		h.l.Error("参数绑定失败:" + err.Error())
		context.Error(common.BadRequestError)
		return
	}
	// 更新标签
	if err = h.s.UpdateTag(tagDto); err != nil {
		h.l.Error("更新标签失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, common.NewSuccessResult())
}

// deleteTag godoc
// @Summary      删除标签
// @Description  根据 URL 中的 id 参数删除标签
// @Tags         tag
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int     true  "标签 ID"
// @Success      200  {object}  common.Result
// @Router       /tags/delete/{id} [delete]
func (h *TagController) deleteTag(context *gin.Context) {
	// 获取参数并尝试转换为 int
	tagIdStr := context.Param("id")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		context.Error(err)
		return
	}
	// 删除标签
	if err = h.s.DeleteTag(tagId); err != nil {
		h.l.Error("删除标签失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, common.NewSuccessResult())
}
