package handler

import (
	"net/http"

	"github.com/EAHITechnology/cage/logic"
	"github.com/EAHITechnology/cage/proto"
	"github.com/EAHITechnology/raptor/elog"
	"github.com/EAHITechnology/raptor/enet"
)

func TestPostHandler(c *enet.Context) {
	fun := "TestPostHandler --> "
	ctx := c.Request.Context()

	req := proto.DemoRequest{}
	resp := proto.DemoResp{
		Data:           proto.DemoRespData{},
		CommonJsonResp: &enet.CommonJsonResp{},
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

func TestGetHandler(c *enet.Context) {
	fun := "TestGetHandler --> "
	ctx := c.Request.Context()

	resp := proto.DemoResp{
		Data: proto.DemoRespData{},
	}

	info := c.Query("info")
	if info == "" {
		elog.Elog.ErrorfCtx(ctx, "%s info nil", fun)
		c.JSON(http.StatusOK, enet.CreateJsonResp(499, "参数错误"))
		return
	}

	id, err := logic.DemoLogic(ctx, info)
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
