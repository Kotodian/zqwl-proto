package protozqwl

type ResetRequest struct {
	BaseStruct
}

func NewResetRequest() *ResetRequest {
	r := &ResetRequest{}
	r.RmtCmd = RmtCmdReset
	return r
}

type ResetResponse struct {
	BaseStruct
	Status Status `json:"status"`
}

func NewResetResponse(status Status) *ResetResponse {
	r := &ResetResponse{}
	r.RmtCmd = RmtCmdRetReset
	r.Status = status
	return r
}
