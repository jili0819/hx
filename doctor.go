package main

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/jili0819/hx/base"
)

type (
	DoctorListReq struct {
		AppCode          string `json:"appCode"`
		OrganCode        string `json:"organCode"`
		ChannelCode      string `json:"channelCode"`
		AppointmentType  int    `json:"appointmentType"`
		DeptCategoryCode string `json:"deptCategoryCode"`
		HaveNo           int    `json:"haveNo"` // 是否有号
		TimeRange        int    `json:"timeRange"`
		HospitalAreaCode string `json:"hospitalAreaCode"`
		HospitalCode     string `json:"hospitalCode"`
		KeyWordEnumType  int    `json:"keyWordEnumType"`
		RegTitelCode     string `json:"regTitelCode"` // 全部职称
		ScheduleDate     string `json:"scheduleDate"` // 预约日期
	}
	DoctorListResp struct {
		Code    string `json:"code"`
		ErrCode string `json:"errCode"`
		Msg     string `json:"msg"`
		Data    []struct {
			DoctorId                  string `json:"doctorId"`
			DocName                   string `json:"docName"`
			DocCode                   string `json:"docCode"`
			DocNameInitial            string `json:"docNameInitial"`
			DocHeadImage              string `json:"docHeadImage"`
			DeptName                  string `json:"deptName"`
			DeptCode                  string `json:"deptCode"`
			DeptCategoryCode          string `json:"deptCategoryCode"`
			DeptCategoryName          string `json:"deptCategoryName"`
			ScheduleDeptCode          string `json:"scheduleDeptCode"`
			ScheduleDeptName          string `json:"scheduleDeptName"`
			DocExperience             string `json:"docExperience"`
			Introduction              string `json:"introduction"`
			HospitalCode              string `json:"hospitalCode"`
			HospitalAreaCode          string `json:"hospitalAreaCode"`
			HospitalName              string `json:"hospitalName"`
			BusinessTagItemRespVoList []struct {
				BusinessTagCode  string `json:"businessTagCode"`
				BusinessTagValue string `json:"businessTagValue"`
				BusinessId       string `json:"businessId"`
				BusinessStatus   int    `json:"businessStatus"`
			} `json:"businessTagItemRespVoList"`
			Status         int    `json:"status"`
			RegTitelName   string `json:"regTitelName"`
			RegTitelCode   string `json:"regTitelCode"`
			ScheduleDate   string `json:"scheduleDate"`
			SerialNumber   int    `json:"serialNumber"`
			AvailableCount int    `json:"availableCount"`
		} `json:"data"`
	}

	DoctorDetailReq struct {
		AppCode          string `json:"appCode"`
		OrganCode        string `json:"organCode"`
		ChannelCode      string `json:"channelCode"`
		AppointmentType  int    `json:"appointmentType"`
		DeptCode         string `json:"deptCode"`
		DeptCategoryCode string `json:"deptCategoryCode"`
		DoctorId         string `json:"doctorId"`
		HospitalAreaCode string `json:"hospitalAreaCode"`
		HospitalCode     string `json:"hospitalCode"`
		TabAreaCode      string `json:"tabAreaCode"`
	}
	DoctorDetailResp struct {
		Code    string `json:"code"`
		ErrCode string `json:"errCode"`
		Msg     string `json:"msg"`
		Data    struct {
			DoctorId                  string      `json:"doctorId"`
			DocCode                   string      `json:"docCode"`
			DocHeadImage              string      `json:"docHeadImage"`
			DocName                   string      `json:"docName"`
			HospitalCode              string      `json:"hospitalCode"`
			DeptName                  string      `json:"deptName"`
			ScheduleDeptName          string      `json:"scheduleDeptName"`
			RegTitelName              string      `json:"regTitelName"`
			HospitalName              string      `json:"hospitalName"`
			DocExperience             string      `json:"docExperience"`
			Introduction              string      `json:"introduction"`
			Tips                      string      `json:"tips"`
			HtmlTips                  interface{} `json:"htmlTips"`
			DrainageParamJson         interface{} `json:"drainageParamJson"`
			DrainageIconUrl           interface{} `json:"drainageIconUrl"`
			ScheduleDeptCode          string      `json:"scheduleDeptCode"`
			BusinessTagItemRespVoList []struct {
				BusinessTagCode  string `json:"businessTagCode"`
				BusinessTagValue string `json:"businessTagValue"`
				BusinessId       string `json:"businessId"`
				BusinessStatus   int    `json:"businessStatus"`
			} `json:"businessTagItemRespVoList"`
			SourceItemsRespVos []struct {
				SysScheduleId            string        `json:"sysScheduleId"`
				ReturnNo                 *int          `json:"returnNo"`
				ScheduleHisId            string        `json:"scheduleHisId"`
				ScheduleDate             string        `json:"scheduleDate"`
				ScheduleType             int           `json:"scheduleType"`
				ScheduleRange            int           `json:"scheduleRange"`
				ScheduleRangeOtherName   *string       `json:"scheduleRangeOtherName"`
				RangeName                interface{}   `json:"rangeName"`
				IsPrecise                int           `json:"isPrecise"`
				AvailableCount           int           `json:"availableCount"`
				RegFee                   float64       `json:"regFee"`
				ServiceFee               float64       `json:"serviceFee"`
				AdmLocation              string        `json:"admLocation"`
				StartNo                  int           `json:"startNo"`
				DeptName                 string        `json:"deptName"`
				DeptCode                 string        `json:"deptCode"`
				DeptCategoryCode         string        `json:"deptCategoryCode"`
				HospitalAreaCode         string        `json:"hospitalAreaCode"`
				HospitalAreaName         string        `json:"hospitalAreaName"`
				HospitalName             string        `json:"hospitalName"`
				DocName                  string        `json:"docName"`
				DoctorId                 string        `json:"doctorId"`
				Status                   int           `json:"status"`
				RegTitelName             string        `json:"regTitelName"`
				RegTitelCode             string        `json:"regTitelCode"`
				SerialNumber             int           `json:"serialNumber"`
				DelaySeconds             string        `json:"delaySeconds"`
				SourceItemsDetailRespVos []interface{} `json:"sourceItemsDetailRespVos"`
			} `json:"sourceItemsRespVos"`
			SourceItems []struct {
				HospitalCode       string `json:"hospitalCode"`
				AreaName           string `json:"areaName"`
				AreaCode           string `json:"areaCode"`
				SourceItemsRespVos []struct {
					SysScheduleId            string        `json:"sysScheduleId"`
					ReturnNo                 *int          `json:"returnNo"`
					ScheduleHisId            string        `json:"scheduleHisId"`
					ScheduleDate             string        `json:"scheduleDate"`
					ScheduleType             int           `json:"scheduleType"`
					ScheduleRange            int           `json:"scheduleRange"`
					ScheduleRangeOtherName   *string       `json:"scheduleRangeOtherName"`
					RangeName                interface{}   `json:"rangeName"`
					IsPrecise                int           `json:"isPrecise"`
					AvailableCount           int           `json:"availableCount"`
					RegFee                   float64       `json:"regFee"`
					ServiceFee               float64       `json:"serviceFee"`
					AdmLocation              string        `json:"admLocation"`
					StartNo                  int           `json:"startNo"`
					DeptName                 string        `json:"deptName"`
					DeptCode                 string        `json:"deptCode"`
					DeptCategoryCode         string        `json:"deptCategoryCode"`
					HospitalAreaCode         string        `json:"hospitalAreaCode"`
					HospitalAreaName         string        `json:"hospitalAreaName"`
					HospitalName             string        `json:"hospitalName"`
					DocName                  string        `json:"docName"`
					DoctorId                 string        `json:"doctorId"`
					Status                   int           `json:"status"`
					RegTitelName             string        `json:"regTitelName"`
					RegTitelCode             string        `json:"regTitelCode"`
					SerialNumber             int           `json:"serialNumber"`
					DelaySeconds             string        `json:"delaySeconds"`
					SourceItemsDetailRespVos []interface{} `json:"sourceItemsDetailRespVos"`
				} `json:"sourceItemsRespVos"`
			} `json:"sourceItems"`
		} `json:"data"`
	}
)

// HxDoctorList 获取预约当天医生排班
func (c *Client) HxDoctorList(
	idCard string,
	doctorListReq DoctorListReq,
) (
	doctorListResp DoctorListResp,
	err error,
) {
	header := c.GenerateHeader(idCard)
	req := gout.H{
		"appCode":          doctorListReq.AppCode,
		"organCode":        doctorListReq.OrganCode,
		"channelCode":      doctorListReq.ChannelCode,
		"appointmentType":  doctorListReq.AppointmentType,
		"deptCategoryCode": doctorListReq.DeptCategoryCode,
		"haveNo":           doctorListReq.HaveNo,
		"timeRange":        doctorListReq.TimeRange,
		"hospitalAreaCode": doctorListReq.HospitalAreaCode,
		"hospitalCode":     doctorListReq.HospitalCode,
		"keyWordEnumType":  doctorListReq.KeyWordEnumType,
	}
	if doctorListReq.RegTitelCode != "" {
		req["regTitelCode"] = doctorListReq.RegTitelCode
	}
	if doctorListReq.ScheduleDate != "" {
		req["scheduleDate"] = doctorListReq.ScheduleDate
	}
	doctorListResp = DoctorListResp{}
	if err = gout.POST(base.HytHost + base.HytDoctorList).
		Debug(c.Config().Debug).
		SetHeader(header).
		SetJSON(req).
		BindJSON(&doctorListResp).
		Do(); err != nil {
		return
	}
	if doctorListResp.Code == "0" {
		err = fmt.Errorf("url:%s,code:%s,errCode:%s,msg:%s", base.HytHost+base.HytDoctorList, doctorListResp.Code, doctorListResp.ErrCode, doctorListResp.Msg)
		return
	}
	return
}

// HxDoctorDetail 获取医生的所有排班
func (c *Client) HxDoctorDetail(
	idCard string,
	doctorDetailReq DoctorDetailReq,
) (
	doctorDetailResp DoctorDetailResp,
	err error,
) {
	header := c.GenerateHeader(idCard)
	req := gout.H{
		"appCode":          doctorDetailReq.AppCode,
		"organCode":        doctorDetailReq.OrganCode,
		"channelCode":      doctorDetailReq.ChannelCode,
		"appointmentType":  doctorDetailReq.AppointmentType,
		"deptCode":         doctorDetailReq.DeptCode,
		"deptCategoryCode": doctorDetailReq.DeptCategoryCode,
		"doctorId":         doctorDetailReq.DoctorId,
		"hospitalAreaCode": doctorDetailReq.HospitalAreaCode,
		"hospitalCode":     doctorDetailReq.HospitalCode,
		"tabAreaCode":      doctorDetailReq.TabAreaCode,
	}
	doctorDetailResp = DoctorDetailResp{}
	if err = gout.POST(base.HytHost + base.HytDoctorDetail).
		Debug(c.Config().Debug).
		SetHeader(header).
		SetJSON(req).
		BindJSON(&doctorDetailResp).
		Do(); err != nil {
		return
	}
	if doctorDetailResp.Code == "0" {
		err = fmt.Errorf("url:%s,code:%s,errCode:%s,msg:%s", base.HytHost+base.HytDoctorDetail, doctorDetailResp.Code, doctorDetailResp.ErrCode, doctorDetailResp.Msg)
		return
	}
	return
}
