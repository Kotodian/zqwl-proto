package protozqwl

type GetRequest struct {
	BaseStruct
	Dev    []DevKey `json:"dev"`
	Cmdidx *int     `json:"cmdidx"`
}

func NewGetRequest(devKey ...DevKey) *GetRequest {
	r := &GetRequest{}
	r.RmtCmd = RmdCmdGet
	r.Dev = devKey
	return r
}

type GetResponse struct {
	BaseStruct
	Dev *DevConfig `json:"dev"`
}

func NewGetResponse(devConfig *DevConfig) *GetResponse {
	r := &GetResponse{}
	r.RmtCmd = RmdCmdRetGet
	r.Dev = devConfig
	return r
}
