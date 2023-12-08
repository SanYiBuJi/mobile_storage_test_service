package Services

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mobile_storage_test_service/logger"
	"net/http"
)

type CaesarTestService struct {
}
type AcceptApplicationFormRequest struct {
	ID              string `json:"id"`        //表单ID
	AssetOID        string `json:"asset_oid"` //资产信息
	AssetID         string `json:"asset_id"`
	CategoryName    string `json:"category_name"`    //设备分类
	ClientID        string `json:"client_id"`        //终端ID
	ClientIP        string `json:"client_ip"`        //终端IP
	ClientName      string `json:"client_name"`      //终端名称
	DeviceID        string `json:"device_id"`        //设备硬件编号
	DeviceName      string `json:"device_name"`      //产品型号
	DeviceUID       string `json:"device_uid"`       //设备UID
	Intranet        int    `json:"intranet"`         //内网有效期
	Reason          string `json:"reason"`           //申请理由
	Remark          string `json:"remark"`           //管理员答复
	ResponsibleName string `json:"responsible_name"` //责任人名称
	StaffID         int    `json:"staff_id"`         //用户ID
	StaffName       string `json:"staff_name"`       //用户组织
	Type            int    `json:"type"`             //申请类型
	Status          int    `json:"status"`           //申请状态
	StartTime       string `json:"start_time"`       //发起时间
}

func AcceptApplicationForm(c *gin.Context) {
	var person AcceptApplicationFormRequest
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personString, err := json.Marshal(person)
	if err != nil {
		logger.Logger.Error("格式化JSON数据错误：" + err.Error())
	}
	logger.Logger.Info(string(personString))

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
