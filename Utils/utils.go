package Utils

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"mobile_storage_test_service/Logger"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func PreintForem(values url.Values) {
	for key, value := range values {
		zap.String(key, strings.Join(value, ","))
	}
}

func UTF8ToUnicode(s8 string) string {
	utf8Bytes := []byte(s8) // "你好"的UTF-8编码

	// 将UTF-8字节切片转换为Unicode字符串
	//unicodeString, err := utf8.DecodeString(string(utf8Bytes))
	//if err != nil {
	//	fmt.Println("解码错误:", err)
	//	return ""
	//}

	return string(utf8Bytes)
}

func TimestampToTime(s string) string {
	// 将时间戳转换为时间格式
	timestamp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		Logger.Logger.Error(err.Error())
	}
	tm := time.Unix(timestamp, 0)

	// 将时间格式化为 "2021-07-20 12:39:44" 的格式
	formattedTime := tm.Format("2006-01-02 15:04:05")
	return formattedTime
}

func PrintRequestBody(c *gin.Context) {
	// 获取请求的原始body
	body := c.Request.Body

	// 创建一个缓冲区来读取body内容
	buf := new(bytes.Buffer)

	// 将body内容读取到缓冲区中
	if _, err := buf.ReadFrom(body); err != nil {
		Logger.Logger.Error(err.Error())
		return
	}

	// 打印缓冲区中的内容
	Logger.Logger.Info("Request Body : " + buf.String())
}

func GetAuthValue(appKey string, appSecret string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		Logger.Logger.Error("无法获取执行文件路径:" + err.Error())
		return "", err
	}

	dir := filepath.Dir(exePath)
	timestamp := time.Now().Unix()
	//Logger.Logger.Info("获取当前时间:" + strconv.Itoa(int(timestamp)))
	//cmd := exec.Command(fmt.Sprintf("%s/app-signature", dir), fmt.Sprintf(" -timestamp=%s -nonce=1 -appkey=%s -appsecret=%s",
	//	fmt.Sprintf("%d", timestamp), appKey, appSecret))
	// 出现过chmod +x 然后，permission denied 问题
	exec.Command(fmt.Sprintf("chmod 777 %s/app-signature", dir))
	cmd := exec.Command(fmt.Sprintf("%s/app-signature", dir),
		"-timestamp", fmt.Sprintf("%d", timestamp),
		"-nonce", "1", "-appkey", appKey, "-appsecret", appSecret)
	Logger.Logger.Info("cmd:" + cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Logger.Logger.Fatal(err.Error())
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		Logger.Logger.Fatal(err.Error())
		return "", err
	}
	// 读取输出结果
	var authValue string
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		Logger.Logger.Fatal(err.Error())
		return "", err
	} else {
		authValue = string(opBytes)
		Logger.Logger.Info(authValue)
	}
	if err := cmd.Wait(); err != nil {
		Logger.Logger.Fatal(err.Error())
		return "", err
	}
	return authValue, nil
}
