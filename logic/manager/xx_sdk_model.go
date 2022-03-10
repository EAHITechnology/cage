package manager

import "github.com/EAHITechnology/raptor/enet"

type DemoRespData struct {
	Id int64 `json:"id"`
}

type DemoResp struct {
	*enet.CommonJsonResp
	Data DemoRespData `json:"data"`
}
