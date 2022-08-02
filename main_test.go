package main

import (
	"os"
	"testing"
)

var (
	client IClient
)

func TestMain(m *testing.M) {
	config := Config{
		Token: "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMjkzNTQ0MjM2NTExMzIxMjMxNDQxIiwiaWF0IjoxNjU5NDA2MzA1LCJzdWIiOiJ7XCJ1c2VySWRcIjpcIjMyOTM1NDRcIixcImFjY291bnRJZFwiOlwiMzkwNjExMlwiLFwidXNlclR5cGVcIjowLFwiYXBwQ29kZVwiOlwiSFhHWUFQUFwiLFwiY2hhbm5lbENvZGVcIjpcIlBBVElFTlRfV0VDSEFUXCIsXCJkZXZpY2VudW1iZXJcIjpcIjIzNjUxMTMyMTIzMTQ0MVwiLFwiZGV2aWNlVHlwZVwiOlwiV1hfSDVcIixcImFjY291bnROb1wiOlwiMTgyODQ1NzA2MzNcIixcIm5hbWVcIjpudWxsLFwiZG9jdG9ySWRcIjpudWxsLFwib3JnYW5Db2RlXCI6XCJISUQwMTAxXCJ9IiwiZXhwIjoxNjYxOTk4MzA1fQ.k5oF5NVzKIsLp4JrkyhPuFzMKYdcXuhHo991x6iVKXc***HXGYAPP",
		Debug: true,
	}
	client = NewClient(config)

	exitCode := m.Run()
	os.Exit(exitCode)
}
