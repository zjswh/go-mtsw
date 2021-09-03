package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"mtsw/global"
)

type Cron struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Level int `json:"level"`
	DependencyStatus int `json:"dependency_status"`
	Spec string `json:"spec"`
	Protocol int `json:"protocol"`
	HttpMethod string `json:"http_method"`
	Command string `json:"command"`
	Timeout int `json:"timeout"`
	Multi int `json:"multi"` // 是否单例 2是 1否
	NotifyStatus int `json:"notify_status"`
	NotifyType int `json:"notify_type"`
	RetryTimes int `json:"retry_times"`
	RetryInterval int `json:"retry_interval"`
}
type CronResult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data int `json:"data"`
}

func CreateCron(name string, date string, callback string,taskId int) (CronResult, error) {
	callback = url.QueryEscape(callback)
	cronUrl := global.GVA_CONFIG.Param.CronHost + "/api/v1/store"
	param := map[string]interface{} {
		"name" : name,
		"spec" : date,
		"command" : callback,
		"id" : taskId,
		"level" : 1,
		"dependency_status" : 1,
		"protocol" : 1,
		"timeout" : 10,
		"multi" : 2,
		"notify_status" : 1,
		"notify_type" : 2,
		"retry_times" : 0,
		"retry_interval" : 1,
		"http_method" : 1,
	}
	body, err := Request(cronUrl, param, map[string]interface{}{
		"Content-Type" : "multipart/form-data",
	}, "POST", "form")
	var cronResult CronResult
	json.Unmarshal(body, &cronResult)
	return cronResult, err
}

func DeleteCron(sTaskId int)  (CronResult, error){
	cronUrl := fmt.Sprintf("%s/%d", global.GVA_CONFIG.Param.CronHost + "/api/v1/remove", sTaskId)
	body, err := Request(cronUrl, map[string]interface{}{}, map[string]interface{}{}, "GET", "form")
	var cronResult CronResult
	json.Unmarshal(body, &cronResult)
	return cronResult, err
}

