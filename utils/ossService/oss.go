package ossService

import (
	"encoding/json"
	"errors"
	"fmt"
	"mtsw/utils"
	"strings"
)

const (
	uploadUrl     = "v1/file/upload"
	copyUrl       = "v1/file/copy"
	getOssInfoUrl = "v1/file/getOssInfo"
	deleteUrl     = "v1/file/delete"
	AuthKey = "e22488067969afd2c63f722d5727321"
	OssDomain = "http://114.55.64.105/"
	StaticDomain = "https://static-pro.guangdianyun.tv"
)

func Copy(tempUrl string, uin int, module string, isComplete bool) (string, bool, error)  {
	tempData := strings.Split(tempUrl, "?")
	tempUrl = tempData[0]

	arr := strings.Split(tempUrl, "/")
	if len(arr) < 4 {
		return tempUrl, false, nil
	}

	tempUrl = strings.Join(arr[len(arr) - 4:], "/")
	if strings.Contains(tempUrl, fmt.Sprintf("/%d/%s", uin, module)) {
		return tempUrl, false, nil
	}

	urlx := OssDomain + copyUrl + "?authKey=" + AuthKey
	res, _ := utils.Request(urlx, map[string]interface{}{
		"tempUrl" : tempUrl,
		"uin" : uin,
		"module" : module,
	}, map[string]interface{}{}, "POST", "json")
	var result struct{
		Code int `json:"code"`
		ErrorCode int `json:"errorCode"`
		ErrorMessage string `json:"errorMessage"`
		Data struct{
			ShortPath string `json:"shortPath"`
			CompletePath string `json:"completePath"`
		} `json:"data"`
	}

	json.Unmarshal(res, &result)
	if result.Code != 200 || result.ErrorCode != 0 {
		return "", true, errors.New(result.ErrorMessage)
	}
	path := result.Data.ShortPath
	if isComplete == true {
		path = result.Data.CompletePath
	}
	return path, true, nil
}

func Get(tempUrl string, uin int, module string) string {
	if strings.Contains(tempUrl, "http") {
		return tempUrl
	}
	return fmt.Sprintf("%s/%d/%s/%s", StaticDomain, uin, module, tempUrl)
}

func Delete(tempUrl string, uin int, module string) {
	urlx := OssDomain + copyUrl + "?authKey=" + AuthKey
	arr := strings.Split(tempUrl, "/")
	if len(arr) > 4 {
		tempUrl = strings.Join(arr[len(arr) - 4:], "/")
	}
	utils.Request(urlx, map[string]interface{}{
		"tempUrl" : tempUrl,
		"uin" : uin,
		"module" : module,
	}, map[string]interface{}{}, "POST", "json")
}
