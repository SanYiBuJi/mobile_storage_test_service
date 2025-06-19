package Services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mobile_storage_test_service/Databases"
	"mobile_storage_test_service/Logger"
	"mobile_storage_test_service/Models"
	"mobile_storage_test_service/Utils"
	"net/http"
	"strings"
)

// CreateApplicationForm 接受控制中心下发的表单数据
func CreateApplicationForm(c *gin.Context) {
	//Utils.PrintRequestBody(c)
	var person Models.AcceptApplicationFormRequest
	//data := c.Request.Body
	//Logger.Logger.Info(c.Request.)
	//reader := bufio.NewReader(data)
	//err := json.NewDecoder(reader).Decode(&person)
	//if err != nil {
	//	Logger.Logger.Error("Failed to decode JSON: %s" + err.Error())
	//	return
	//}
	if err := c.BindJSON(&person); err != nil {
		Logger.Logger.Error("解析数据异常" + err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 3, "message": err.Error()})
		return
	}
	//person.Type = Utils.UTF8ToUnicode(person.Type)
	//person.OutUseAuthMaxTime = Utils.UTF8ToUnicode(person.OutUseAuthMaxTime)
	//person.OutUseAuthValidDay = Utils.UTF8ToUnicode(person.OutUseAuthValidDay)
	//person.Status = Utils.UTF8ToUnicode(person.Status)
	personString, err := json.Marshal(person)
	if err != nil {
		Logger.Logger.Error("格式化JSON数据错误：" + err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 3, "message": err.Error()})
		return
	}
	//if person.Type == "\u0004" {
	//	Logger.Logger.Info("out_use_auth:" + person.OutUseAuth)
	//	Logger.Logger.Info("out_use_auth_max_time:" + person.OutUseAuthMaxTime)
	//	Logger.Logger.Info("out_use_auth_valid_day:" + person.OutUseAuthValidDay)
	//}

	Logger.Logger.Info(string(personString))
	// 持久话数据
	_, err = Databases.CreateApplicationFormData(&person)
	if err != nil {
		Logger.Logger.Error("插入数据库错误：" + err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 3, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
	return
}

func BatchGetApplicationForm(c *gin.Context) {
	var person Models.BatchGetApplicationFormRequest
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offset := (person.Page - 1) * person.Limit
	set, err := Databases.SelectApplicationFormDataList(person.Limit, offset)
	if err != nil {
		Logger.Logger.Error(err.Error())
	}
	if set != nil {
		//var body gin.H
		dataString, err := json.Marshal(set)
		Logger.Logger.Info(string(dataString))
		if err != nil {
			Logger.Logger.Error(err.Error())
		}
		//body["data"] = string(dataString)
		var response = Models.BatchGetApplicationFormResponse{Data: set}
		c.Status(http.StatusOK)
		c.JSON(http.StatusOK, response)
	}
}

func UpdateApplicationFormTestV1(c *gin.Context) {
	var person Models.UpdateApplicationFormTestV1Request
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if authValue, err := Utils.GetAuthValue(person.AppKey, person.AppSecret); err != nil {
		Logger.Logger.Error(err.Error())
	} else {
		authValue = strings.Split(authValue, "\n")[0]
		//请求Caesar
		caesarRequest := Models.CaesarMobileStoragePutCheck{Data: Models.CaesarMobileStoragePutCheckData{
			ID:       person.ID,
			AssetID:  person.AssetId,
			AssetOID: person.AssetOID,
			Reason:   person.Reason,
			Status:   person.Status,
		}}
		Logger.Logger.Info(caesarRequest.Data.ID)
		reqBytes, err := json.Marshal(caesarRequest)
		if err != nil {
			Logger.Logger.Error("请求数据转化失败:" + err.Error())
		}
		Logger.Logger.Info("Caesar请求数据:" + string(reqBytes))
		//	忽略证书
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("POST", fmt.Sprintf("https://%s:%s/caesar/v1/edge/mobile_storage_out_check", person.Addr, person.Port), bytes.NewBuffer(reqBytes))
		if err != nil {
			Logger.Logger.Error("创建请求失败:" + err.Error())
			return
		}
		headers := make(http.Header)                                // 创建一个空的请求头对象
		headers.Set("Authorization", "ZEUS-HMAC-SHA256 "+authValue) // 设置Authorization头，
		headers.Set("Content-Type", "application/json")
		req.Header = headers
		//req.Header.Set("authorization", authValue)
		// 发送请求并获取响应
		Logger.Logger.Info("authorization:" + req.Header.Get("authorization"))
		resp, err := client.Do(req)
		if err != nil {
			Logger.Logger.Error("发送请求失败:" + err.Error())
			return
		}
		defer resp.Body.Close()

		// 读取响应内容
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Logger.Logger.Error("读取响应失败:" + err.Error())
		}
		caesarRes := string(body)
		Logger.Logger.Info(caesarRes)
		c.JSON(http.StatusOK, gin.H{"message": caesarRes})
	}
}
