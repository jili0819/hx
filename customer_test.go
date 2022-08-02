package main

import (
	"github.com/jili0819/hx/base"
	"testing"
)

func TestClient_CardList(t *testing.T) {

	cardReq := CardReq{
		AppCode:     base.AppCode,
		OrganCode:   base.OrganCode,
		ChannelCode: "PATIENT_WECHAT",
		Guidance:    "0",
	}
	resp, err := client.CardList(cardReq)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
