package constants

// 交易状态（Success：付款成功；Ing：付款中；Fail：付款失败）
type TransactionStatus string

const (
	TransactionStatusSuccess TransactionStatus = "Success"
	TransactionStatusIng     TransactionStatus = "Ing"
	TransactionStatusFail    TransactionStatus = "Fail"
)

//事件类型（交易成功：TX_SUCCESS；交易失败：TX_FAIL）
type EventType string

const (
	EventTypeTXSUCCESS EventType = "TX_SUCCESS"
	EventTypeTXFAIL    EventType = "TX_FAIL"
)

var AllEventType = []EventType{
	EventTypeTXSUCCESS,
	EventTypeTXFAIL,
}

var EventTypeMapTransactionStatus = map[EventType]TransactionStatus{
	EventTypeTXSUCCESS: TransactionStatusSuccess,
	EventTypeTXFAIL:    TransactionStatusFail,
}
