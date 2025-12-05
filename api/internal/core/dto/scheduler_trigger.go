package dto

import "github.com/google/uuid"

type SchedulerTriggerReq struct {
	UID       uuid.UUID `json:"uuid,omitempty"`
	Email     string    `json:"email" binding:"required" example:"test001@gmail.com"`
	Message   string    `json:"message" binding:"required" example:"Teste de envio temporizado"`
	TriggerAt string    `json:"UTC_trigger_at" binding:"required" example:"2025-11-18T15:28:00Z"`
}

type SchedulerTriggerResp struct {
	Message string `json:"message" example:"Request Accepted"`
	UID     string `json:"uuid,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}
