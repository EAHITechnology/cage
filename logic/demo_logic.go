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

	version, ok, err := dao.SetRedisLock("test", 10)
	if err != nil {
		elog.Elog.Errorf("%s SetRedisLock error:%v", fun, err)
		return 0, err
	}

	if !ok {
		return 0, nil
	}

	elog.Elog.Debugf("%s SetRedisLock version:%d", fun, version)

	// Call DAO
	id, err := dao.InsertDemoTask(d)
	if err != nil {
		elog.Elog.Errorf("%s InsertDemoTask error:%v", fun, err)
		return 0, err
	}

	if err := dao.Set(d.Info, d.Id); err != nil {
		elog.Elog.Errorf("%s Set error:%v", fun, err)
		return 0, err
	}

	val, err := dao.Get(d.Info)
	if err != nil {
		elog.Elog.Errorf("%s Get error:%v", fun, err)
		return 0, err
	}

	elog.Elog.Infof("%s Get val:%s", fun, val)

	if err := dao.DelRedisLock("test", version); err != nil {
		elog.Elog.Errorf("%s DelRedisLock error:%v", fun, err)
		return 0, err
	}

	// Call the manager layer to make third-party remote calls
	if err := manager.XxSdkManager.Call(id); err != nil {
		return 0, err
	}

	return id, nil
}
