package handler

import (
	"net/http"

	"github.com/EAHITechnology/cage/logic"
	"github.com/EAHITechnology/cage/proto"
	"github.com/EAHITechnology/raptor/elog"
	"github.com/EAHITechnology/raptor/enet"
)

func TestHandler(c *enet.Context) {
	fun := "TestHandler --> "
	ctx := c.Request.Context()

	req := proto.DemoRequest{}
	resp := proto.DemoResp{
		Data: proto.DemoRespData{},
	}

	if err := c.ShouldBind(&req); err != nil {
		elog.Elog.ErrorfCtx(ctx, "%s ShouldBind err:%v", fun, err)
		c.JSON(http.StatusOK, enet.CreateJsonResp(400, "参数错误"))
		return
	}

	id, err := logic.DemoLogic(ctx, req.Info)
	if err != nil {
		elog.Elog.ErrorfCtx(ctx, "%s DemoLogic err:%v", fun, err)
		c.JSON(http.StatusOK, enet.CreateJsonResp(500, err.Error()))
		return
	}

	resp.Data.Id = id
	resp.Code = 200
	resp.Msg = "success"

	c.JSON(http.StatusOK, resp)
}
