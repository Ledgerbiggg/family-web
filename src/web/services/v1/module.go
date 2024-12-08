package v1

import (
	impls2 "family-web-server/src/web/services/v1/impls"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	// 注册服务
	fx.Provide(impls2.NewLoginService),  // 登录服务
	fx.Provide(impls2.NewHomeService),   // 登录服务
	fx.Provide(impls2.NewInviteService), // 邀请服务
	fx.Provide(impls2.NewAlbumService),  // 相册服务
)
