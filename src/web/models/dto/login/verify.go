package login

type VerifyDto struct {
	Username string `json:"username" binding:"required"`
	ReaName  string `json:"reaName" binding:"required"`
}
