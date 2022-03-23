package manager

import (
	"encoding/json"

	"github.com/EAHITechnology/raptor/elog"
	"github.com/EAHITechnology/raptor/erpc"
	"golang.org/x/net/context"
)

type xxSdkManager struct{}

var XxSdkManager *xxSdkManager

func init() {
	XxSdkManager = NewXxSdkManager()
}

func NewXxSdkManager() *xxSdkManager {
	return &xxSdkManager{}
}

func (x *xxSdkManager) Call(id int64) error {
	fun := "Call -->"

	client, err := erpc.HttpManager.GetClient("cage_downstream")
	if err != nil {
		return err
	}

	resp, err := client.Send(context.Background(), erpc.GET, nil, "/test", nil, nil, nil)
	if err != nil {
		return err
	}

	r := DemoResp{}
	if err := json.Unmarshal(resp, &r); err != nil {
		return err
	}

	elog.Elog.Debugf("%s code:%d msg:%s", fun, r.Code, r.Msg)

	return nil
}
