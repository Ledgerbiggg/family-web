package v1

import (
	"family-web-server/src/web/controllers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"controllers",
	// 控制器管理者
	fx.Provide(controllers.NewControllerManager),
	// 注册控制器
	fx.Invoke(NewLoginController),  // 登录控制器
	fx.Invoke(NewHomeController),   // 主页控制器
	fx.Invoke(NewInviteController), // 邀请控制器
	fx.Invoke(NewAlbumController),  // 相册控制器
)
