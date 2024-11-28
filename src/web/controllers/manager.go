package controllers

type ControllerManager struct {
	controllers []Base
}

func NewControllerManager() *ControllerManager {
	var controllers []Base
	return &ControllerManager{controllers: controllers}
}

func (cm *ControllerManager) GetControllers() []Base {
	return cm.controllers
}

func (cm *ControllerManager) AddController(controller Base) {
	cm.controllers = append(cm.controllers, controller)
}
