package Services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mobile_storage_test_service/Databases"
	"mobile_storage_test_service/Models"
	"mobile_storage_test_service/Utils"
	"mobile_storage_test_service/logger"
	"net/http"
	"strings"
)

func AcceptApplicationForm(c *gin.Context) {
	var person Models.AcceptApplicationFormRequest
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personString, err := json.Marshal(person)
	if err != nil {
		logger.Logger.Error("格式化JSON数据错误：" + err.Error())
	}
	logger.Logger.Info(string(personString))
	_, err = Databases.CreateApplicationFormData(&person)
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
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
		logger.Logger.Error(err.Error())
	}
	if set != nil {
		//var body gin.H
		dataString, err := json.Marshal(set)
		logger.Logger.Info(string(dataString))
		if err != nil {
			logger.Logger.Error(err.Error())
		}
		//body["data"] = string(dataString)
		var response = Models.BatchGetApplicationFormResponse{Data: set}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateApplicationFormTestV1(c *gin.Context) {
	var person Models.UpdateApplicationFormTestV1Request
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if authValue, err := Utils.GetAuthValue(person.Appkey, person.Appsecret); err != nil {
		logger.Logger.Error(err.Error())
	} else {
		authValue = strings.Split(authValue, "\n")[0]
		//请求Caesar
		caesarRequest := Models.CaesarMobileStoragePutCheck{
			ID:       person.ID,
			AssetID:  person.AssetId,
			AssetOID: person.AssetOID,
			Reason:   person.Reason,
			Status:   person.Status,
		}
		logger.Logger.Info(caesarRequest.ID)
		reqBytes, err := json.Marshal(caesarRequest)
		if err != nil {
			logger.Logger.Error("请求数据转化失败:" + err.Error())
		}
		logger.Logger.Info("Caesar请求数据:" + string(reqBytes))
		//	忽略证书
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("POST", fmt.Sprintf("https://%s:%s/caesar/v1/edge/mobile_storage_out_check", person.Addr, person.Port), bytes.NewBuffer(reqBytes))
		if err != nil {
			logger.Logger.Error("创建请求失败:" + err.Error())
			return
		}
		headers := make(http.Header)                                // 创建一个空的请求头对象
		headers.Set("Authorization", "ZEUS-HMAC-SHA256 "+authValue) // 设置Authorization头，
		headers.Set("Content-Type", "application/json")
		req.Header = headers
		//req.Header.Set("authorization", authValue)
		// 发送请求并获取响应
		logger.Logger.Info("authorization:" + req.Header.Get("authorization"))
		resp, err := client.Do(req)
		if err != nil {
			logger.Logger.Error("发送请求失败:" + err.Error())
			return
		}
		defer resp.Body.Close()

		// 读取响应内容
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Logger.Error("读取响应失败:" + err.Error())
		}
		caesarRes := string(body)
		logger.Logger.Info(caesarRes)
		c.JSON(http.StatusOK, gin.H{"message": caesarRes})
	}
}
