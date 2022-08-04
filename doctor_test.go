package main

import (
	"github.com/jili0819/hx/base"
	"testing"
)

func TestClient_DoctorList(t *testing.T) {
	doctorListRep := DoctorListReq{
		AppCode:          base.AppCode,
		OrganCode:        base.OrganCode,
		ChannelCode:      "PATIENT_WECHAT",
		AppointmentType:  1,
		DeptCategoryCode: "7400-ZLXB",
		HaveNo:           1,
		TimeRange:        2,
		HospitalAreaCode: "HID0101",
		HospitalCode:     "HID0101",
		KeyWordEnumType:  0,
		ScheduleDate:     "2022-08-10",
	}
	doctorListResp, err := client.HxDoctorList(idCard, doctorListRep)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(doctorListResp)
}

func TestClient_DoctorDetail(t *testing.T) {
	doctorDetailRep := DoctorDetailReq{
		AppCode:          base.AppCode,
		OrganCode:        base.OrganCode,
		ChannelCode:      "PATIENT_WECHAT",
		AppointmentType:  1,
		DeptCode:         "1294",
		DeptCategoryCode: "7400-ZLXB",
		DoctorId:         "2c968082698199bf016a9ba6a00b43fe",
		HospitalAreaCode: "HID0101",
		HospitalCode:     "HID0101",
		TabAreaCode:      "HID0101",
	}
	doctorDetailResp, err := client.HxDoctorDetail(idCard, doctorDetailRep)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(doctorDetailResp)
}
