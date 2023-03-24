package protozqwl

type DefaultRequest struct {
	BaseStruct
}

func NewDefaultRequest() *DefaultRequest {
	r := &DefaultRequest{}
	r.RmtCmd = RmtCmdDefault
	return r
}

type DefaultResponse struct {
	BaseStruct
	Status Status `json:"status"`
}

func NewDefaultResponse(status Status) *DefaultResponse {
	r := &DefaultResponse{}
	r.RmtCmd = RmtCmdRetDefault
	r.Status = status
	return r
}
