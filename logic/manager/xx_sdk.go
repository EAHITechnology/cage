package manager

import (
	"github.com/EAHITechnology/raptor/elog"
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
	elog.Elog.Debugf("%s id:%d", fun, id)
	return nil
}
