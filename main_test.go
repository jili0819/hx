package hx

import (
	"os"
	"testing"
)

var (
	client *Client
	idCard = "513023199306115619"
	token  = "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMjkzNTQ0MjM2NTExMzIxMjMxNDQxIiwiaWF0IjoxNjU5Nzc1NTk2LCJzdWIiOiJ7XCJ1c2VySWRcIjpcIjMyOTM1NDRcIixcImFjY291bnRJZFwiOlwiMzkwNjExMlwiLFwidXNlclR5cGVcIjowLFwiYXBwQ29kZVwiOlwiSFhHWUFQUFwiLFwiY2hhbm5lbENvZGVcIjpcIlBBVElFTlRfV0VDSEFUXCIsXCJkZXZpY2VudW1iZXJcIjpcIjIzNjUxMTMyMTIzMTQ0MVwiLFwiZGV2aWNlVHlwZVwiOlwiV1hfSDVcIixcImFjY291bnROb1wiOlwiMTgyODQ1NzA2MzNcIixcIm5hbWVcIjpudWxsLFwiZG9jdG9ySWRcIjpudWxsLFwib3JnYW5Db2RlXCI6XCJISUQwMTAxXCJ9IiwiZXhwIjoxNjYyMzY3NTk2fQ.tig00csxSDJXmej0xUwKAJZa68T8aW8zBqdjqje1w8A***HXGYAPP"
)

func TestMain(m *testing.M) {

	codeUrl, _ := os.Getwd()
	config := &Config{
		AppCode:        "HXGYAPP",
		OrganCode:      "HID0101",
		Debug:          true,
		ImageCachePath: codeUrl,
	}
	client = NewClient(config)
	client.AddCustomer(&CustomerInfo{
		Token:  token,
		IDCard: idCard,
	})
	exitCode := m.Run()
	os.Exit(exitCode)
}
