package server

import (
	"net/http"

	"github.com/EAHITechnology/raptor/elog"
	"github.com/EAHITechnology/raptor/enet"
)

const (
	USER_HEADER = "user"
)

func authMiddle(c *enet.Context) {
	fun := "AuthMiddle -->"
	user := c.GetHeader(USER_HEADER)
	if user == "" {
		c.Abort()
		elog.Elog.Infof("%s AuthMiddle fail", fun)
		c.JSON(http.StatusOK, enet.CreateJsonResp(500, "auth fail"))
	}
}
