package dto

type LivenessResp struct {
	Message string `json:"message" binding:"required" example:"OK"`
	Sumary  string `json:"sumary" binding:"required" example:"go-scheduler-trigger-api:8080 in TagVersion: 0.0.0 on Envoriment:dev responds OK"`
}
