package protozqwl

type RegisterRequest struct {
	SN string `json:"sn"` // rtu sn号
	ID string `json:"id"` // rtu mac地址或者rtu imei
}
