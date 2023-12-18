package Utils

import (
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"mobile_storage_test_service/logger"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func PreintForem(values url.Values) {
	for key, value := range values {
		zap.String(key, strings.Join(value, ","))
	}
}

func GetAuthValue(appKey string, appSecret string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		logger.Logger.Error("无法获取执行文件路径:" + err.Error())
		return "", err
	}

	dir := filepath.Dir(exePath)
	timestamp := time.Now().Unix()
	//logger.Logger.Info("获取当前时间:" + strconv.Itoa(int(timestamp)))
	//cmd := exec.Command(fmt.Sprintf("%s/app-signature", dir), fmt.Sprintf(" -timestamp=%s -nonce=1 -appkey=%s -appsecret=%s",
	//	fmt.Sprintf("%d", timestamp), appKey, appSecret))
	cmd := exec.Command(fmt.Sprintf("%s/app-signature", dir),
		"-timestamp", fmt.Sprintf("%d", timestamp),
		"-nonce", "1", "-appkey", appKey, "-appsecret", appSecret)
	logger.Logger.Info("cmd:" + cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		logger.Logger.Fatal(err.Error())
		return "", err
	}
	// 读取输出结果
	var authValue string
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		logger.Logger.Fatal(err.Error())
		return "", err
	} else {
		authValue = string(opBytes)
		logger.Logger.Info(authValue)
	}
	if err := cmd.Wait(); err != nil {
		logger.Logger.Fatal(err.Error())
		return "", err
	}
	return authValue, nil
}
