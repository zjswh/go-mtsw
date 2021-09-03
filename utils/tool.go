package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"math/rand"
	"mtsw/global"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func DefaultIntParam(str string, defaultValue int, c *gin.Context) int {
	param := c.Query(str)
	if param == "" {
		return defaultValue
	}
	intParam, _ := strconv.Atoi(param)
	return intParam
}

func DefaultIntFormValue(str string, defaultValue int, c *gin.Context) int {
	param := c.Request.FormValue(str)
	if param == "" {
		return defaultValue
	}
	intParam, _ := strconv.Atoi(param)
	return intParam
}

func Request(url string, data map[string]interface{}, header map[string]interface{}, method string, stype string) (body []byte, err error) {
	url = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(url, "\n", ""), " ", ""), "\r", "")
	param := []byte("")
	if stype == "json" {
		param, _ = json.Marshal(data)
		header["Content-Type"] = "application/json"
	} else {
		s := ""
		for k, v := range data {
			s += fmt.Sprintf("%s=%v&", k, v)
		}
		header["Content-Type"] = "application/x-www-form-urlencoded"
		param = []byte(s)
	}

	//获取环境变量
	if global.GVA_CONFIG.Param.XCaStage != "" {
		header["x-ca-stage"] = global.GVA_CONFIG.Param.XCaStage
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(param))
	if err != nil {
		err = fmt.Errorf("new request fail: %s", err.Error())
		return
	}

	for k, v := range header {
		req.Header.Add(k, fmt.Sprintf("%s", v))
	}

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("do request fail: %s", err.Error())
		return
	}

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("read res body fail: %s", err.Error())
		return
	}
	return
}

func RequestGet(url string) (body []byte, err error) {
	return Request(url, map[string]interface{}{}, map[string]interface{}{}, "GET", "")
}

func GetCronTime(t time.Time) string {
	h := fmt.Sprintf("%d %d %d %d %d *", t.Second(), t.Minute(), t.Hour(), t.Day(), t.Month())
	return h
}

func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

func Export(c *gin.Context, filename string, file *xlsx.File)  {
	if filename == "" {
		time := strconv.Itoa(int(time.Now().Unix()))
		filename = time
	}
	filename = filename + ".xlsx"
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Transfer-Encoding", "binary")
	file.Write(c.Writer)
}

func SendDms(topic string, cmd string, uin int, data interface{}, withUin bool)  {
	extra, _ := json.Marshal(data)
	url := global.GVA_CONFIG.Param.MessageHost + "/v1/message/Index/send"
	if !withUin {
		url = global.GVA_CONFIG.Param.MessageHost + "/v1/message/Index/sendDms"
	}
	res, err := Request(url, map[string]interface{}{
		"topic": topic,
		"cmd" : cmd,
		"uin" : uin,
		"extra": string(extra),
	}, map[string]interface{}{}, "POST", "form")
	if err != nil {
		fmt.Println(string(res))
	}
}

func StringContains(arr []string, str string) bool {
	flag := false
	for _, v := range arr {
		if v == str {
			flag = true
			break
		}
	}
	return flag
}

func RandomCoupon() string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	var length = 16 // 2个-
	rand.Seed(time.Now().UnixNano())

	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		if i > 0 && i%4 == 0 {
			result = append(result, "-")
		}
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			// result = append(result, string(rand.Intn(26)+65))
			result = append(result, string(rune(rand.Intn(26)+65)))
		} else {
			// result = append(result, string(rand.Intn(26)+97))
			result = append(result, string(rune(rand.Intn(26)+97)))
		}
	}
	return strings.ToUpper(strings.Join(result, ""))
}

func StructToMap(data interface{}) map[string]interface{} {
	dataMap := map[string]interface{}{}
	dataBytes, _ := json.Marshal(data)

	//防止底层在输出的时候会进行格式化防止出现科学计数法
	d := json.NewDecoder(bytes.NewReader(dataBytes))
	d.UseNumber()
	_ = d.Decode(&dataMap)
	return dataMap
}

func MapToJson(om map[string]interface{}, order ...string) string {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	l := len(order)
	for i, k := range order {
		fmt.Fprintf(buf, "\"%s\":\"%v\"", k, om[k])
		if i < l-1 {
			buf.WriteByte(',')
		}
	}
	buf.Write([]byte{'}'})
	return buf.String()
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
