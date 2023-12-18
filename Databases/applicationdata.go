package Databases

import (
	"errors"
	"fmt"
	"mobile_storage_test_service/Models"
	"mobile_storage_test_service/logger"
	"time"
)

func AcceptApplicationFormRequestTOData(apf *Models.AcceptApplicationFormRequest) *Models.ApplicationFormData {
	data := Models.ApplicationFormData{
		ID:              0,
		FormID:          apf.ID,
		AssetOID:        apf.AssetOID,
		AssetID:         apf.AssetID,
		CategoryName:    apf.CategoryName,
		ClientID:        apf.ClientID,
		ClientIP:        apf.ClientIP,
		ClientName:      apf.ClientName,
		DeviceID:        apf.DeviceID,
		DeviceName:      apf.DeviceName,
		DeviceUID:       apf.DeviceUID,
		Intranet:        apf.Intranet,
		Reason:          apf.Reason,
		Remark:          apf.Remark,
		ResponsibleName: apf.ResponsibleName,
		StaffID:         apf.StaffID,
		StaffName:       apf.StaffName,
		Type:            apf.Type,
		Status:          apf.Status,
		StartTime:       "2023-07-29 15:30:00",
		//CreateTime:      time.Now().Unix(),
		//UpgradeTime:     time.Now().Unix(),
	}
	return &data
}

//func (Models.ApplicationFormData) TableName() string {
//	return "Models.ApplicationFormData"
//}

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
			logger.Logger.Error(err.Error())
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
				logger.Logger.Error(resDB.Error.Error())
				return false, resDB.Error
			}
			//logger.Logger.Info("rows affected: %d" + string(resDB.RowsAffected))
			logger.Logger.Info("Create source:" + fmt.Sprintf("%s", data))
			return true, nil
		}
	} else {
		// 这里设计的是相同formID的内容不会再次插入，不确定后面的情况
		return false, nil
	}

	return false, err

}

func SelectApplicationFormDataOfFormID(formId string) (*Models.ApplicationFormData, error) {
	var formDatas []Models.ApplicationFormData
	db, err := GetDBConnect()
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}
	if db != nil {
		resDBs := db.Where("form_id=?", formId).Find(&formDatas)
		if resDBs.Error != nil {
			logger.Logger.Error(resDBs.Error.Error())
			return nil, resDBs.Error
		}
		if len(formDatas) != 1 {
			if len(formDatas) > 1 {
				logger.Logger.Error("查询结果超过1条，出现重复，请联系管理员排查")
				return &formDatas[0], nil
			}
			if len(formDatas) == 0 {
				logger.Logger.Info("查询结果为空")
				return nil, nil
			}
		} else {
			logger.Logger.Debug("查询到一条记录：" + formDatas[0].FormID)
			return &formDatas[0], nil
		}
	} else {
		logger.Logger.Error("Get db is null")
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
		logger.Logger.Info("已完成审批的记录，不允许再次修改 表单ID：" + oldDate.FormID)
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
	//logger.Logger.Info("数据库查询内容长度:", len(rows))
	println(len(rows))
	return &rows, nil
}

type DeviceInfoData struct {
	ID                        uint
	DeviceUID                 string
	DeviceID                  string
	PID                       string
	VID                       string
	DeviceName                string
	DeviceResponsibleName     string
	DeviceResponsibleUserName string
}
