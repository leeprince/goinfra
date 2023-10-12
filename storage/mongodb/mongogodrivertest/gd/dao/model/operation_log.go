package model

// 查询模型(mongDB官网库支持嵌入字段，但是必须要添加`bson:"bson:,inline`)：包含默认主键_id; 用于查询时能够返回默认主键_id
type QueryOperationLogModel struct {
	Id                      string `bson:"_id" json:"_id"` // 租户ID
	InsertOperationLogModel `bson:",inline"`
}

// 插入模型：不包含默认主键_id（MongoDB驱动程序mongDB默认会自动维护主键_id）
type InsertOperationLogModel struct {
	OrgId           int64      `bson:"org_id" json:"org_id"`                       // 租户ID
	UserId          int64      `bson:"user_id" json:"user_id"`                     // 用户ID
	UserName        string     `bson:"user_name" json:"user_name"`                 // 用户名
	UserNickname    string     `bson:"user_nickname" json:"user_nickname"`         // 用户昵称
	ClientIp        string     `bson:"client_ip" json:"client_ip"`                 // 发起请求的客户端IP
	ReqRouter       string     `bson:"req_router" json:"req_router"`               // 前端请求路径，包含网关前缀
	ReqRouterPrefix string     `bson:"req_router_prefix" json:"req_router_prefix"` // 前端请求路径的网关前缀
	OperateName     string     `bson:"operate_name" json:"operate_name"`           // 请求路由对应的操作名
	ReqHeader       string     `bson:"req_header" json:"req_header"`               // 请求头,json字符串
	ReqBody         string     `bson:"req_body" json:"req_body"`                   // 请求体,json字符串
	CostTime        int32      `bson:"cost_time" json:"cost_time"`                 // 请求处理完所花费时间，单位毫秒
	RespBody        RespBody   `bson:"resp_body" json:"resp_body"`                 // 响应体,json字符串,对应到 RespBody 结构体的json字符串
	LogId           string     `bson:"log_id" json:"log_id"`                       // 响应体的log_id
	RespStatus      RespStatus `bson:"resp_status" json:"resp_status"`             // 响应体的状态映射是否操作成功；1：成功；2：失败
	OperatedAt      int64      `bson:"operated_at" json:"operated_at"`             // 操作时间
	CreatedAt       int64      `bson:"created_at" json:"created_at"`               // 日志创建时间
}

var OperationLogField = struct {
	OrgId           string
	UserId          string
	UserName        string
	UserNickName    string
	ClientIp        string
	ReqRouter       string
	ReqRouterPrefix string
	OperateName     string
	ReqHeader       string
	ReqBody         string
	CostTime        string
	RespBody        string
	LogId           string
	RespStatus      string
	OperatedAt      string
	CreatedAt       string
}{
	OrgId:           "org_id",
	UserId:          "user_id",
	UserName:        "user_name",
	UserNickName:    "user_nickname",
	ClientIp:        "client_ip",
	ReqRouter:       "req_router",
	ReqRouterPrefix: "req_router_prefix",
	OperateName:     "operate_name",
	ReqHeader:       "req_header",
	ReqBody:         "req_body",
	CostTime:        "cost_time",
	RespBody:        "resp_body",
	LogId:           "log_id",
	RespStatus:      "resp_status",
	OperatedAt:      "operated_at",
	CreatedAt:       "created_at",
}

type RespBody struct {
	Code    int         `bson:"code" json:"code"` // 0 成功；非0失败
	Message string      `bson:"message" json:"message"`
	LogId   string      `bson:"logId" json:"log_id"`
	Data    interface{} `bson:"data" json:"data"`
}

type RespStatus int32

const (
	RespStatusSucc RespStatus = 1
	RespStatusFail RespStatus = 2
)

func (InsertOperationLogModel) DataBaseName() string {
	return "ptest-db"
}

func (InsertOperationLogModel) CollectionName() string {
	return "ptest-log"
}
