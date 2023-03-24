package protozqwl

type DevConfig struct {
	WM            *WMConfig      `json:"wm,omitempty"`
	SA            *SAConfig      `json:"sa,omitempty"`
	SB            *SBConfig      `json:"sb,omitempty"`
	MQTT          *MQTTConfig    `json:"mqtt,omitempty"`
	HB            *HBConfig      `json:"hb,omitempty"`
	Reg           *RegConfig     `json:"reg,omitempty"`
	HTTP          *HTTPConfig    `json:"http,omitempty"`
	RMT           *RMTConfig     `json:"rmt,omitempty"`
	UART          *UARTConfig    `json:"uart,omitempty"`
	NBT           *NBTConfig     `json:"nbt,omitempty"`
	AI            *AIConfig      `json:"ai,omitempty"`
	Poll          *PollConfig    `json:"poll,omitempty"`
	PollCmdConfig *PollCmdConfig `json:"pollcmd,omitempty"`
	RTUConfig     *RTUConfig     `json:"rtu,omitempty"`
	Info          *Info          `json:"info,omitempty"`
}

type WMConfig struct {
	Mode int `json:"mode"`
}

type SAConfig struct {
	Enable int    `json:"enable"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	UDP    int    `json:"udp"`
}

type SBConfig struct {
	Enable int    `json:"enable"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	UDP    int    `json:"udp"`
}

type MQTTConfig struct {
	Enable int    `json:"enable"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	UN     string `json:"un"`
	UP     string `json:"up"`
	ID     string `json:"id"`
	KA     int    `json:"ka"`
	Sub    string `json:"sub"`
	Pub1   string `json:"pub1"`
	Pub2   string `json:"pub2"`
	Pub3   string `json:"pub3"`
	Pub4   string `json:"pub4"`
	Cloud  int    `json:"cloud"`
	PK     string `json:"pk"`
	DN     string `json:"dn"`
	DS     string `json:"ds"`
	SSL    int    `json:"ssl"`
	QOS    int    `json:"qos"`
}

type HBConfig struct {
	Enable int    `json:"enable"`
	Time   int    `json:"time"`
	Data   string `json:"data"`
}

type RegConfig struct {
	Enable int    `json:"enable"`
	Mode   int    `json:"mode"`
	Data   string `json:"data"`
}

type HTTPConfig struct {
	Enable int    `json:"enable"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Post   int    `json:"post"`
	Header string `json:"header"`
}

type RMTConfig struct {
	Enable int    `json:"enable"`
	IP     string `json:"ip"`
	Code   string `json:"code"`
}

type UARTConfig struct {
	Rate  int `json:"rate"`
	Check int `json:"check"`
	Pack  int `json:"pack"`
}

type Info struct {
	RSSI   int        `json:"rssi"`
	Run    int        `json:"run"`
	GPS    string     `json:"gps"`
	Ver    string     `json:"ver"`
	Model  string     `json:"model"`
	IMEI   string     `json:"imei"`
	MAC    string     `json:"mac"`
	ICCID  string     `json:"iccid"`
	Type   int        `json:"type"`
	X      int        `json:"x"`
	Y      int        `json:"y"`
	ADC    int        `json:"adc"`
	Freem  int        `json:"freem"`
	Status InfoStatus `json:"status"`
}

type InfoStatus struct {
	SA    int `json:"sa"`
	SB    int `json:"sb"`
	HTTP  int `json:"http"`
	MQTT  int `json:"mqtt"`
	MQTCP int `json:"mqtcp"`
	MQSub int `json:"mqsub"`
}

type NBTConfig struct {
	Time int `json:"time"`
}

type AIConfig struct {
	Mode int `json:"mode"`
}

type PollConfig struct {
	Timeout      int    `json:"timeout"`
	Update       int    `json:"update"`
	Delay        int    `json:"delay"`
	UploadTime   int64  `json:"uploadtime"`
	UploadType   int    `json:"uploadtype"`
	Uploadchange int    `json:"uploadchange"`
	Timestamp    int64  `json:"timestamp"`
	TimeoutMode  int    `json:"timeoutmode"`
	Array        int    `json:"array"`
	YTrigger     int    `json:"y_trigger"`
	Prefix       string `json:"prefix"`
	Cmdcnt       int    `json:"cmdcnt"`
}

type PollCmdConfig struct {
	Cmd struct {
		Cnt   int     `json:"cnt"`
		Idx   int     `json:"idx"`
		Addr  int     `json:"addr"`
		Fun   int     `json:"fun"`
		Raddr int     `json:"raddr"`
		Rsize int     `json:"rsize"`
		DType int     `json:"dtype"`
		Float float64 `json:"float"`
		Key   string  `json:"key"`
		TV1   int     `json:"tv1"`
		TV2   int     `json:"tv2"`
		EM    int     `json:"em"`
	} `json:"cmd"`
}

type RTUConfig struct {
}
