package proto

type OrderActionType int

const (
	Buy OrderActionType = iota
	Sell
	SellToday
	SellYesterday
	Short
	Cover
	CoverToday
	CoverYesterday
	Auto
	AutoLong
	AutoShort
)

type EntrustOrder struct {
	Security string
	Action   OrderActionType
	Price    int64
	Seq      float64
	Size     int32
	RefNo    int64
}

type AlgoOptions struct {
	Name   string
	Params map[string]string
}

type CreateTaskReq struct {
	Req
	Orders []EntrustOrder
	Algo   AlgoOptions
}

type CreateTaskRsp struct {
	Rsp
	TaskId string
}

type CancelTaskReq struct {
	Req
	TaskId string
}

type CancelTaskRsp struct {
	Rsp
}

type Req struct {
	ReqId string
}

type Rsp struct {
	ReqId string
}