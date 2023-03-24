package protozqwl

type BaseStruct struct {
	RmtCmd RmtCmd `json:"rmtcmd"`
	IMEI   string `json:"imei,omitempty"`
}

type RmtCmd string

const (
	RmdCmdGet        RmtCmd = "get"
	RmdCmdRetGet     RmtCmd = "retget"
	RmdCmdSet        RmtCmd = "set"
	RmdCmdRetSet     RmtCmd = "retset"
	RmtCmdRegister   RmtCmd = "register"
	RmtCmdHeartbeat  RmtCmd = "heartbeat"
	RmtCmdReset      RmtCmd = "reset"
	RmtCmdRetReset   RmtCmd = "retreset"
	RmtCmdDefault    RmtCmd = "default"
	RmtCmdRetDefault RmtCmd = "retdefault"
)

type DevKey string

const (
	DevWM   DevKey = "wm"
	DevSA   DevKey = "sa"
	DevSB   DevKey = "sb"
	DevMQTT DevKey = "mqtt"
	DevHB   DevKey = "hb"
	DevReg  DevKey = "reg"
	DevHTTP DevKey = "http"
	DevAPN  DevKey = "apn"
	DevRmt  DevKey = "rmt"
	DevPoll DevKey = "poll"
	DevRTU  DevKey = "rtu"
	DevUart DevKey = "uart"
	DevSMS  DevKey = "sms"
	DevInfo DevKey = "info"
	DevNbt  DevKey = "nbt"
	DevAi   DevKey = "ai"
)

type Enable uint8

const (
	Disabled Enable = 0
	Enabled  Enable = 1
)

type Status string

const (
	StatusFailed Status = "0"
	StatusOk     Status = "1"
)
