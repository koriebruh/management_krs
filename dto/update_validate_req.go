package dto

type UpdateValidateReq struct {
	JobHost  string `json:"job_host" validate:"required"`
	JobAgent string `json:"job_agent" validate:"required"`
	TA       int    `json:"ta" validate:"required"`
}
