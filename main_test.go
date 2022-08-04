package main

import (
	"os"
	"testing"
)

var (
	client IClient
	idCard = "513023199306115619"
	token  = "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMjkzNTQ0MjM2NTExMzIxMjMxNDQxIiwiaWF0IjoxNjU5NDA2MzA1LCJzdWIiOiJ7XCJ1c2VySWRcIjpcIjMyOTM1NDRcIixcImFjY291bnRJZFwiOlwiMzkwNjExMlwiLFwidXNlclR5cGVcIjowLFwiYXBwQ29kZVwiOlwiSFhHWUFQUFwiLFwiY2hhbm5lbENvZGVcIjpcIlBBVElFTlRfV0VDSEFUXCIsXCJkZXZpY2VudW1iZXJcIjpcIjIzNjUxMTMyMTIzMTQ0MVwiLFwiZGV2aWNlVHlwZVwiOlwiV1hfSDVcIixcImFjY291bnROb1wiOlwiMTgyODQ1NzA2MzNcIixcIm5hbWVcIjpudWxsLFwiZG9jdG9ySWRcIjpudWxsLFwib3JnYW5Db2RlXCI6XCJISUQwMTAxXCJ9IiwiZXhwIjoxNjYxOTk4MzA1fQ.k5oF5NVzKIsLp4JrkyhPuFzMKYdcXuhHo991x6iVKXc***HXGYAPP"
)

func TestMain(m *testing.M) {

	codeUrl, _ := os.Getwd()
	config := &Config{
		Debug:          true,
		ImageCachePath: codeUrl,
	}
	client = NewClient(config)
	client.AddCustomer(&CustomerInfo{
		Token:       token,
		OrganCode:   "HID0101",
		ChannelCode: "PATIENT_WECHAT",
		IDCard:      idCard,
	})
	exitCode := m.Run()
	os.Exit(exitCode)
}
