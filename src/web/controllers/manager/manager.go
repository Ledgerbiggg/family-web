package manager

import (
	"family-web-server/src/web/controllers/base"
)

type ControllerManager struct {
	controllers []base.ControllerBase
}

func NewControllerManager() *ControllerManager {
	var controllers []base.ControllerBase
	return &ControllerManager{controllers: controllers}
}

func (cm *ControllerManager) GetControllers() []base.ControllerBase {
	return cm.controllers
}

func (cm *ControllerManager) AddController(controller base.ControllerBase) {
	cm.controllers = append(cm.controllers, controller)
}
