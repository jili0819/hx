package hx

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/jili0819/hx/base"
)

type (
	// HospitalReq request
	HospitalReq struct {
		AppCode         string `json:"appCode"`
		OrganCode       string `json:"organCode"`
		ChannelCode     string `json:"channelCode"`
		HospitalCode    string `json:"hospitalCode"` // 医院code，非院区
		AppointmentType int    `json:"appointmentType"`
	}
	// HospitalResp response
	HospitalResp struct {
		Code    string        `json:"code"`
		ErrCode string        `json:"errCode"`
		Msg     string        `json:"msg"`
		Data    *HospitalData `json:"data"`
	}
	HospitalData struct {
		DeptListTips                 string     `json:"deptListTips"`
		DeptListTipsNew              string     `json:"deptListTipsNew"`
		SelHospitalAreaRecordRespVos []Hospital `json:"selHospitalAreaRecordRespVos"`
	}
	// Hospital 院区
	Hospital struct {
		HospitalCode     string `json:"hospitalCode"`     // 医院code
		HospitalAreaCode string `json:"hospitalAreaCode"` // 院区code
		HospitalAreaName string `json:"hospitalAreaName"` // 院区名称
	}

	// DeptReq request
	DeptReq struct {
		AppCode         string `json:"appCode"`
		OrganCode       string `json:"organCode"`
		ChannelCode     string `json:"channelCode"`
		AppointmentType int    `json:"appointmentType"` // 1-可预约？
		HospitalArea    string `json:"hospitalArea"`    // 院区
		HospitalCode    string `json:"hospitalCode"`    // 医院code
	}
	// DeptResp response
	DeptResp struct {
		Code    string         `json:"code"`
		ErrCode string         `json:"errCode"`
		Msg     string         `json:"msg"`
		Data    []DeptCategory `json:"data"`
	}
	// DeptCategory 科室分类
	DeptCategory struct {
		HospitalCode            string      `json:"hospitalCode"`
		DeptCategoryCode        string      `json:"deptCategoryCode"`
		DeptCategoryName        string      `json:"deptCategoryName"`
		OrderNum                interface{} `json:"orderNum"`
		DeptCategoryNameInitial string      `json:"deptCategoryNameInitial"`
	}
)

// HxInitHospitalAreaRecord
//  @Description:获取所有院区
//  @receiver c
//  @param key
//  @param hospitalReq
//  @return err
func (c *Client) HxInitHospitalAreaRecord(
	key string,
	hospitalReq HospitalReq,
) (
	err error,
) {
	header := c.GenerateHeader(key)
	req := gout.H{
		"appCode":         hospitalReq.AppCode,
		"organCode":       hospitalReq.OrganCode,
		"channelCode":     hospitalReq.ChannelCode,
		"hospitalCode":    hospitalReq.HospitalCode,
		"appointmentType": hospitalReq.AppointmentType,
	}
	resp := HospitalResp{}
	if err = gout.POST(base.HxHost + base.HytHospitalArea).
		Debug(c.config.Debug).
		SetHeader(header).
		SetJSON(req).
		BindJSON(&resp).
		Do(); err != nil {
		return
	}
	if resp.Code == "0" {
		err = fmt.Errorf("url:%s,code:%s,errCode:%s,msg:%s", base.HxHost+base.HytHospitalArea, resp.Code, resp.ErrCode, resp.Msg)
		return
	}
	// 院区存入client
	if resp.Data != nil {
		if len(resp.Data.SelHospitalAreaRecordRespVos) > 0 {
			addHospitals(c, resp.Data.SelHospitalAreaRecordRespVos)
		}
	}
	return
}

// HxInitDept
//  @Description: 查询院区下科室
//  @receiver c
func (c *Client) HxInitDept(
	idCard string,
	deptReq DeptReq,
) (
	err error,
) {
	header := c.GenerateHeader(idCard)
	req := gout.H{
		"appCode":         deptReq.AppCode,
		"organCode":       deptReq.OrganCode,
		"channelCode":     deptReq.ChannelCode,
		"appointmentType": deptReq.AppointmentType,
		"hospitalArea":    deptReq.HospitalArea, // 院区code
		"hospitalCode":    deptReq.HospitalCode, // 医院code
	}
	resp := DeptResp{}
	if err = gout.POST(base.HxHost + base.HytDept).
		Debug(c.config.Debug).
		SetHeader(header).
		SetJSON(req).
		BindJSON(&resp).
		Do(); err != nil {
		return
	}
	if resp.Code == "0" {
		err = fmt.Errorf("url:%s,code:%s,errCode:%s,msg:%s", base.HxHost+base.HytDept, resp.Code, resp.ErrCode, resp.Msg)
		return
	}
	// 院区存入client
	if len(resp.Data) > 0 {
		addDeparts(c, deptReq.HospitalArea, resp.Data)
	}
	return
}

//
// addHospitals
//  @Description:添加院区到client
//  @receiver c
//  @param hospitals
//
func addHospitals(
	c IClient,
	hospitals []Hospital,
) {
	c.Client().Lk.Lock()
	defer c.Client().Lk.Unlock()
	for _, value := range hospitals {
		c.Client().Hospitals[value.HospitalAreaCode] = value
	}
}

// addDeparts
//  @Description: 添加院区科室
//
func addDeparts(
	c IClient,
	hospitalAreaCode string,
	departs []DeptCategory,
) {
	c.Client().Lk.Lock()
	defer c.Client().Lk.Unlock()
	c.Client().HospitalAreas[hospitalAreaCode] = departs
}
