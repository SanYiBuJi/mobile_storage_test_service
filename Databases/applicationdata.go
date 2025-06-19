package Databases

import (
	"errors"
	"fmt"
	"mobile_storage_test_service/Logger"
	"mobile_storage_test_service/Models"
	"mobile_storage_test_service/Utils"
	"time"
)

func AcceptApplicationFormRequestTOData(apf *Models.AcceptApplicationFormRequest) *Models.ApplicationFormData {
	data := Models.ApplicationFormData{
		ID:                 0,
		FormID:             apf.ID,
		AssetOID:           apf.AssetOID,
		AssetID:            apf.AssetID,
		CategoryName:       apf.CategoryName,
		ClientID:           apf.ClientID,
		ClientIP:           apf.ClientIP,
		ClientName:         apf.ClientName,
		DeviceID:           apf.DeviceID,
		DeviceName:         apf.DeviceName,
		DeviceUID:          apf.DeviceUID,
		Intranet:           apf.Intranet,
		Reason:             apf.Reason,
		Remark:             apf.Remark,
		ResponsibleName:    apf.ResponsibleName,
		StaffID:            apf.StaffID,
		StaffName:          apf.StaffName,
		Type:               apf.ApplyType,
		Status:             apf.Status,
		StartTime:          apf.ApplyTime,
		CreateTime:         apf.CreateTime,
		UpdateTime:         apf.UpdateTime,
		OutUseAuth:         apf.OutUseAuth,
		OutUseAuthMaxTime:  apf.OutUseMaxTimes,
		OutUseAuthValidDay: Utils.TimestampToTime(fmt.Sprintf("%d", apf.OutUseValidDay)),
		//ApplyReason:         apf.ApplyReason,
		ApplyTime:           apf.ApplyTime,
		ApplyType:           apf.ApplyType,
		ApplyTypeDesc:       apf.ApplyTypeDesc,
		ClientNodeName:      apf.ClientNodeName,
		ClientNodePath:      apf.ClientNodePath,
		ClientType:          apf.ClientType,
		CustomSwapArea:      apf.CustomSwapArea,
		CustomSwapSize:      apf.CustomSwapSize,
		DeviceCapacity:      apf.DeviceCapacity,
		DevicePID:           apf.DevicePID,
		DeviceVID:           apf.DeviceVID,
		LastCategoryName:    apf.LastCategoryName,
		LastResponsibleName: apf.LastResponsibleName,
		RegisterType:        apf.RegisterType,
		StaffNodeName:       apf.StaffNodeName,
		StatusDesc:          apf.StatusDesc,
	}
	return &data
}

//	CreateApplicationFormData func (Models.ApplicationFormData) TableName() string {
//		return "Models.ApplicationFormData"
//	}
//
// 插入一条申请记录数据
func CreateApplicationFormData(apf *Models.AcceptApplicationFormRequest) (bool, error) {
	data := AcceptApplicationFormRequestTOData(apf)
	selectData, err := SelectApplicationFormDataOfFormID(data.FormID)
	if err != nil {
		return false, err
	}
	if selectData == nil && err == nil {
		// 无表单记录
		db, err := GetDBConnect()
		if err != nil {
			Logger.Logger.Error(err.Error())
			return false, err
		}
		if db != nil {
			//todo 设置时间戳
			// 未查询到记录，为新建设备申请记录，添加时间戳
			layout := "2006-01-02 15:04:05"
			data.CreateTime = time.Now().Format(layout)
			data.UpdateTime = data.CreateTime
			resDB := db.Create(data)
			if resDB.Error != nil {
				Logger.Logger.Error(resDB.Error.Error())
				return false, resDB.Error
			}
			//Logger.Logger.Info("rows affected: %d" + string(resDB.RowsAffected))
			Logger.Logger.Info("Create source:" + fmt.Sprintf("%s", data))
			return true, nil
		}
	} else {
		// 这里设计的是相同formID的内容不会再次插入，不确定后面的情况
		return false, nil
	}

	return false, err

}

// SelectApplicationFormDataOfFormID 查询一条申请记录 通过表单ID
func SelectApplicationFormDataOfFormID(formId string) (*Models.ApplicationFormData, error) {
	var formDatas []Models.ApplicationFormData
	db, err := GetDBConnect()
	if err != nil {
		Logger.Logger.Error(err.Error())
		return nil, err
	}
	if db != nil {
		resDBs := db.Where("form_id=?", formId).Find(&formDatas)
		if resDBs.Error != nil {
			Logger.Logger.Error(resDBs.Error.Error())
			return nil, resDBs.Error
		}
		if len(formDatas) != 1 {
			if len(formDatas) > 1 {
				Logger.Logger.Error("查询结果超过1条，出现重复，请联系管理员排查")
				return &formDatas[0], nil
			}
			if len(formDatas) == 0 {
				Logger.Logger.Info("查询结果为空")
				return nil, nil
			}
		} else {
			Logger.Logger.Debug("查询到一条记录：" + formDatas[0].FormID)
			return &formDatas[0], nil
		}
	} else {
		Logger.Logger.Error("Get db is null")
		return nil, errors.New("get db is null ")
	}
	return nil, nil
}

func UpdateApplicationForm(newDate *Models.ApplicationFormData) error {
	oldDate, err := SelectApplicationFormDataOfFormID(newDate.FormID)
	if err != nil {
		return err
	}
	//梳理，目前可知会进行数据修改的操作有：审批结果修改，管理员答复，更新时间
	if oldDate.Status == 4 {
		Logger.Logger.Info("已完成审批的记录，不允许再次修改 表单ID：" + oldDate.FormID)
		return errors.New("已完成审批的记录，不允许再次修改 表单ID：" + oldDate.FormID)
	} else if oldDate.Status == 1 && newDate.Status == 2 {

	}
	return err
}

func SelectApplicationFormDataList(limitNum int, length int) (*[]Models.ApplicationFormData, error) {
	db, err := GetDBConnect()
	if err != nil {
		return nil, err
	}
	var rows []Models.ApplicationFormData
	db.Limit(limitNum).Offset(length).Find(&rows)
	//Logger.Logger.Info("数据库查询内容长度:", len(rows))
	println(len(rows))
	return &rows, nil
}

// DeviceInfoData 结构体定义了设备信息的数据模型，包含设备的唯一标识符、ID、PID、VID、名称以及负责人的姓名和用户名。
type DeviceInfoData struct {
	ID                        uint   // 设备的唯一标识符
	DeviceUID                 string // 设备的唯一用户标识符
	DeviceID                  string // 设备的ID
	PID                       string // 设备的PID
	VID                       string // 设备的VID
	DeviceName                string // 设备的名称
	DeviceResponsibleName     string // 设备负责人的姓名
	DeviceResponsibleUserName string // 设备负责人的用户名
}
