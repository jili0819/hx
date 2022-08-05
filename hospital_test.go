package hx

import (
	"testing"
)

func TestClient_InitDept(t *testing.T) {
	if err := client.HxInitDept(idCard, DeptReq{
		AppCode:         "HXGYAPP",
		OrganCode:       "HID0101",
		ChannelCode:     "PATIENT_WECHAT",
		HospitalCode:    "HID0101",
		HospitalArea:    "HID0101",
		AppointmentType: 1,
	}); err != nil {
		t.Error(err)
	}
	t.Log(client.Client().Hospitals)
}

func TestClient_InitHospitalAreaRecord(t *testing.T) {
	if err := client.HxInitHospitalAreaRecord(idCard, HospitalReq{
		AppCode:         "HXGYAPP",
		OrganCode:       "HID0101",
		ChannelCode:     "PATIENT_WECHAT",
		HospitalCode:    "HID0101",
		AppointmentType: 1,
	}); err != nil {
		t.Error(err)
	}
	t.Log(client.Client().Hospitals)
}
