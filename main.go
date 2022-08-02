package main

import (
	"context"
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/jili0819/hx/base"
	"github.com/panjf2000/ants/v2"
	"runtime"
	"sync"
	"time"
)

var (
	susscess = make(map[string]bool, 0)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	return
	fmt.Println(time.Now())
	// base 初始化要抢科室信息
	// 1、获取指定科室，指定日期的所有可预约医生
	// 2、生成验证码
	// 3、抢号

	ctx := context.Background()
	token := "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzMjkzNTQ0MjM2NTExMzIxMjMxNDQxIiwiaWF0IjoxNjU5NDA2MzA1LCJzdWIiOiJ7XCJ1c2VySWRcIjpcIjMyOTM1NDRcIixcImFjY291bnRJZFwiOlwiMzkwNjExMlwiLFwidXNlclR5cGVcIjowLFwiYXBwQ29kZVwiOlwiSFhHWUFQUFwiLFwiY2hhbm5lbENvZGVcIjpcIlBBVElFTlRfV0VDSEFUXCIsXCJkZXZpY2VudW1iZXJcIjpcIjIzNjUxMTMyMTIzMTQ0MVwiLFwiZGV2aWNlVHlwZVwiOlwiV1hfSDVcIixcImFjY291bnROb1wiOlwiMTgyODQ1NzA2MzNcIixcIm5hbWVcIjpudWxsLFwiZG9jdG9ySWRcIjpudWxsLFwib3JnYW5Db2RlXCI6XCJISUQwMTAxXCJ9IiwiZXhwIjoxNjYxOTk4MzA1fQ.k5oF5NVzKIsLp4JrkyhPuFzMKYdcXuhHo991x6iVKXc***HXGYAPP"
	req := ReqInfo{
		AppCode:          "HXGYAPP",
		OrganCode:        "HID0101",
		ChannelCode:      "PATIENT_WECHAT",
		AppointmentType:  1,
		CardId:           "470972827781173248",
		HospitalCode:     "HID0101",
		HospitalAreaCode: "HID0101",
		SysScheduleId:    "463073395271462912",
		Type:             "WEB",
		ImageId:          "ghv4dggmvaj5ocuj3tpd7ra04jeib1c3",
	}
	// 随机10000个验证码，同时请求挂号
	infos := make([]ReqInfo, 10000)
	for _, value := range base.Code {
		req.VerifyCode = value
		infos = append(infos, req)
	}

	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(100, func(info interface{}) {
		SearchCompanyList(ctx, token, info)
		wg.Done()
	})
	if err != nil {
		return
	}
	defer p.Release()

	for _, value := range infos {
		wg.Add(1)
		_ = p.Invoke(value)
	}
	wg.Wait()

	if _, ok := susscess[req.CardId+req.SysScheduleId]; ok {
		fmt.Println("挂号成功")
	} else {
		fmt.Println("抢号失败")
	}
	fmt.Println(time.Now())
}

type ReqInfo struct {
	AppCode          string `json:"appCode"`
	OrganCode        string `json:"organCode"`
	ChannelCode      string `json:"channelCode"`
	AppointmentType  int    `json:"appointmentType"`
	CardId           string `json:"cardId"`
	HospitalCode     string `json:"hospitalCode"`
	HospitalAreaCode string `json:"hospitalAreaCode"`
	SysScheduleId    string `json:"sysScheduleId"`
	Type             string `json:"type"`
	VerifyCode       string `json:"verifyCode"`
	ImageId          string `json:"imageId"`
}

func SearchCompanyList(ctx context.Context, token string, req interface{}) error {
	header := GenerateHeader(token)
	resp := CommonRsp{}
	reqInfo, _ := req.(ReqInfo)
	info := gout.H{
		"appCode":          reqInfo.AppCode,
		"organCode":        reqInfo.OrganCode,
		"channelCode":      reqInfo.ChannelCode,
		"appointmentType":  reqInfo.AppointmentType,
		"cardId":           reqInfo.CardId,
		"hospitalCode":     reqInfo.HospitalCode,
		"hospitalAreaCode": reqInfo.HospitalAreaCode,
		"sysScheduleId":    reqInfo.SysScheduleId,
		"type":             reqInfo.Type,
		"verifyCode":       reqInfo.VerifyCode,
		"imageId":          reqInfo.ImageId,
	}
	if err := gout.POST("https://hytapiv2.cd120.com/cloud/hosplatcustomer/call/appointment/appointmentModel/sureAppointment").
		Debug(true).SetJSON(info).SetHeader(header).BindJSON(&resp).Do(); err != nil {
		return err
	}
	if resp.Code == "1" && resp.ErrCode == "0" {
		susscess[reqInfo.CardId+reqInfo.SysScheduleId] = true
	}
	return nil
}

// GenerateHeader header token&timespan
func GenerateHeader(token string) (header interface{}) {
	return gout.H{"Token": token, "accessToken": token}
}

type CommonRsp struct {
	Code    string   `json:"code"`
	ErrCode string   `json:"errCode"`
	Msg     string   `json:"msg"`
	Data    DataInfo `json:"data"`
}

type DataInfo struct {
	DoctorId          interface{} `json:"doctorId"`
	DocName           interface{} `json:"docName"`
	DocCode           interface{} `json:"docCode"`
	HospitalName      interface{} `json:"hospitalName"`
	AdmDate           interface{} `json:"admDate"`
	SysAppointmentId  string      `json:"sysAppointmentId"`
	AppointTime       interface{} `json:"appointTime"`
	AdmRange          interface{} `json:"admRange"`
	AdmTimeRange      string      `json:"admTimeRange"`
	RegFee            interface{} `json:"regFee"`
	AdmLocation       interface{} `json:"admLocation"`
	DeptName          string      `json:"deptName"`
	DeptCode          interface{} `json:"deptCode"`
	AppointmentNo     string      `json:"appointmentNo"`
	ChannelCode       interface{} `json:"channelCode"`
	ChannelName       string      `json:"channelName"`
	PatientId         interface{} `json:"patientId"`
	HospitalCode      interface{} `json:"hospitalCode"`
	BizSysSeq         interface{} `json:"bizSysSeq"`
	OrderCreateTime   interface{} `json:"orderCreateTime"`
	OrderStatus       interface{} `json:"orderStatus"`
	AppointmentResult string      `json:"appointmentResult"`
	AppointStatus     int         `json:"appointStatus"`
}
