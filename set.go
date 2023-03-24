package protozqwl

type SetRequest struct {
	BaseStruct
	Dev *DevConfig `json:"dev"`
}

func NewSetRequest(devConfig *DevConfig) *SetRequest {
	r := &SetRequest{}
	r.RmtCmd = RmdCmdSet
	r.Dev = devConfig
	return r
}

type SetResponse struct {
	BaseStruct
	Status Status `json:"status"`
}

func NewSetResponse(status Status) *SetResponse {
	r := &SetResponse{}
	r.RmtCmd = RmdCmdRetSet
	r.Status = status
	return r
}
