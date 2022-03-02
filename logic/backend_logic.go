package logic

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/EAHITechnology/cage/utils"
	"github.com/EAHITechnology/raptor/elog"

	"golang.org/x/net/context"
)

var stopFlag int32

const (
	STOP_FLAG = 1
)

func SetBackendStopFlag() {
	atomic.StoreInt32(&stopFlag, STOP_FLAG)
}

func GetBackendStopFlag() int32 {
	return atomic.LoadInt32(&stopFlag)
}

func InitBackend(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go InitDemoBackend(ctx, &wg)

	<-ctx.Done()
	SetBackendStopFlag()
	wg.Wait()
}

func InitDemoBackend(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	t := time.NewTimer(time.Second * 20)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if GetBackendStopFlag() == STOP_FLAG {
				break
			}
			elog.Elog.Debugf("config:%v", utils.Config.Test1)
			// do some thing
			t.Reset(time.Second * 20)
		case <-ctx.Done():
			return
		}

	}
}
