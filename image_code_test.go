package main

import (
	"testing"
)

func TestClient_GetImageCode(t *testing.T) {
	req := ImageCodeReq{
		AppCode:     "HXGYAPP",
		OrganCode:   "HID0101",
		ChannelCode: "PATIENT_WECHAT",
		Type:        "WEB",
	}
	err := client.Customer(idCard).HxGenerateImageCode(client, req)
	if err != nil {
		t.Error(err)
		return
	}
	code, err := client.Customer(idCard).HxGetImageCodeRand(client)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(code)
}
