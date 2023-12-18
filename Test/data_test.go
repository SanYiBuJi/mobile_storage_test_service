package main

import (
	"mobile_storage_test_service/Databases"
	"mobile_storage_test_service/Models"
	"testing"
)

func TestData(t *testing.T) {
	_, err := Databases.CreateApplicationFormData(&Models.AcceptApplicationFormRequest{
		ID:              "A850421858166113384",
		AssetOID:        "2850421858166113384",
		AssetID:         "2850421858166113384",
		CategoryName:    "2850421858166113384",
		ClientID:        "2850421858166113384",
		ClientIP:        "2850421858166113384",
		ClientName:      "2850421858166113384",
		DeviceID:        "",
		DeviceName:      "",
		DeviceUID:       "",
		Intranet:        0,
		Reason:          "",
		Remark:          "",
		ResponsibleName: "",
		StaffID:         0,
		StaffName:       "",
		Type:            0,
		Status:          0,
		StartTime:       "",
	})
	if err != nil {
		println(err.Error())
	} else {
		return
	}
}
