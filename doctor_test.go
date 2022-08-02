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
		ScheduleDate:     "2022-08-05",
	}
	doctorListResp, err := client.DoctorList(doctorListRep)
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
		DeptCode:         "331",
		DeptCategoryCode: "7400-ZLXB",
		DoctorId:         "2c9480827028f503017074da72f54a35",
		HospitalAreaCode: "HID0101",
		HospitalCode:     "HID0101",
		TabAreaCode:      "HID0101",
	}
	doctorDetailResp, err := client.DoctorDetail(doctorDetailRep)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(doctorDetailResp)
}
