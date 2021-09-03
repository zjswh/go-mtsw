package aodianyun

import "encoding/json"

func ReplayOp(uin int, appId string, stream string, stype string) (ApiResult, error) {
	param := map[string]interface{}{
		"uin" : uin,
		"appid" : appId,
		"stream" : stream,
		"type" : stype,
	}
	url := "LSS.ReplayOp"
	res, err := OpenApi(url, uin, param)
	var apiResult ApiResult
	json.Unmarshal(res, &apiResult)
	if err != nil {
		return apiResult, err
	}
	return apiResult, err
}

func EnableLiveReview(uin int, appId string, stream string) (interface{}, error) {
	param := map[string]interface{}{
		"uin" : uin,
		"app" : appId,
		"stream" : stream,
	}
	url := "LSS.EnableLiveReview"
	res, err := OpenApi(url, uin, param)
	return res, err
}

func RemoveLiveReview(uin int, appId string, stream string) (interface{}, error) {
	param := map[string]interface{}{
		"uin" : uin,
		"app" : appId,
		"stream" : stream,
	}
	url := "LSS.RemoveLiveReview"
	res, err := OpenApi(url, uin, param)
	return res, err
}
