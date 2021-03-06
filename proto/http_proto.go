package proto

import "github.com/EAHITechnology/raptor/enet"

type DemoRequest struct {
	Info string `json:"info"`
}

type DemoRespData struct {
	Id int64 `json:"id"`
}

type DemoResp struct {
	*enet.CommonJsonResp
	Data DemoRespData `json:"data"`
}
