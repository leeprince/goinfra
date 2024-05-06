package params

type SendMsgToUserReq struct {
	WsToken string `json:"wstoken"`
	Data    string `json:"data"`
}
