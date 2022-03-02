package logic

import (
	"context"
	"time"

	"github.com/EAHITechnology/cage/dao"
	"github.com/EAHITechnology/cage/logic/manager"
	"github.com/EAHITechnology/raptor/elog"
)

func DemoLogic(ctx context.Context, info string) (int64, error) {
	fun := "DemoLogic -->"
	d := dao.DemoModel{
		Info:       info,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	// 调用 dao
	id, err := dao.InsertDemoTask(d)
	if err != nil {
		elog.Elog.Errorf("%s InsertDemoTask error:%v", fun, err)
		return 0, err
	}

	// 调用 manager 层进行第三方远程调用
	if err := manager.XxSdkManager.Call(id); err != nil {
		return 0, err
	}

	return id, nil
}
