package algo

import "github.com/thingio/go-algo/proto"

func Receive(msg interface{}) {
	switch req:= msg.(type) {
	case proto.CreateTaskReq:
	case proto.CancelTaskReq:
	}
}
