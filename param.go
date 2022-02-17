package atlib

const (
	SUCCESS = 200
	FAILED  = 300
)

const (
	QUERY_SUCCES  = "查询成功"
	SAVE_SUCCES   = "存储成功"
	SUBMIT_SUCCES = "反馈成功"
)

type PayStatus struct {
	VpnId            string
	Country          string
	Version          string
	DollarPrice      string
	Type             string
	ResultCode       string
	Level            string
	Code             string
	TimeStr          string
	Result           string
	VpnConnectStatus string `json:"vpnConnectStatus"`
	Status           string
}

type QueryBean struct {
	Count int
}

type LogFile struct {
	AccountId string
	Date      string
	Path      string
	FileName  string
}
type ProblemInfo struct {
	Id          int
	AccountId   string
	VpnId       string
	Problem     string
	ContactInfo string
	Date        string
}

type Event struct {
	Event      string `json:"event"`
	Timestamp  string
	TimePhone  string
	Uuid       string
	Androidid  string
	PhoneType  string
	Language   string
	Country    string
	AppVersion string
	OsVersion  string
	Segment    string
	Level      string
	SdkVersion string
}

type OtherParam struct {
	OrderId          string
	PackageName      string
	ProductId        string
	PurchaseTime     string
	PurchaseState    string
	DeveloperPayload string `json:"developerPayload"`
	PurchaseToken    string `json:"purchaseToken"`
}

type Country struct {
	Code string `json:"code"`
	En   string `json:"en"`
	Cn   string `json:"cn"`
}
